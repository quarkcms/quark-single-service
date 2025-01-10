package task

import (
	"log"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

type BillTask struct{}

// 账单任务
func RunBillTask(cron *gocron.Scheduler) {
	billTask := &BillTask{}
	// 生成昨日的日账单
	cron.Every(1).Day().At("00:00").Do(billTask.MakeBill, "yestoday")

	// 生成上周的周账单
	cron.Every(1).Monday().At("00:00").Do(billTask.MakeBill, "week")

	// 生成上月的月账单
	cron.Every(1).Month(1).At("00:00").Do(billTask.MakeBill, "month")
}

// 生成昨日的日账单
func (p *BillTask) MakeBill(period string) {
	var (
		billRecordTitle, billRecordDay string
		billRecordType                 int8
		startDatetime, endDatetime     datetime.Datetime
	)
	switch period {
	case "yestoday":
		yestoday := datetime.Now().AddDate(0, 0, -1).Format("2006-01-02")
		billRecordTitle = "日账单"
		billRecordType = 1
		// 账单起始日期
		startDatetime, _ = datetime.ParseDatetime(yestoday + " 00:00:00")
		// 账单截止日期
		endDatetime, _ = datetime.ParseDatetime(yestoday + " 23:59:59")
		billRecordDay = yestoday
	case "week":
		billRecordTitle = "周账单"
		billRecordType = 2
		// 账单起止日期
		startDatetime, endDatetime = getLastWeekStartDatetimeAndEndDatetime()
		_, week := endDatetime.ISOWeek()
		billRecordDay = "第" + strconv.Itoa(week) + "周（" + endDatetime.Format("01月") + "）"
	case "month":
		billRecordTitle = "月账单"
		billRecordType = 3
		// 账单起止日期
		startDatetime, endDatetime = getLastMonthStartDatetimeAndEndDatetime()
		billRecordDay = startDatetime.Format("2006-01")
	default:
		return
	}

	// 初始化（收入、支出、入账）金额
	var entryPrice, expPrice, incomePrice float64

	bills := service.NewBillService().GetListByPeriod(startDatetime, endDatetime)
	for _, bill := range bills {
		// PM：0-支出；1-获得
		switch bill.PM {
		case 0:
			expPrice = expPrice + bill.Number
		case 1:
			entryPrice = entryPrice + bill.Number
		}
	}
	incomePrice = entryPrice - expPrice

	// 创建账单
	billRecord := model.BillRecord{
		Title:         billRecordTitle,
		Day:           billRecordDay,
		Type:          billRecordType,
		EntryPrice:    entryPrice,
		ExpPrice:      expPrice,
		IncomePrice:   incomePrice,
		StartDatetime: startDatetime,
		EndDatetime:   endDatetime,
	}
	if err := service.NewBillRecordService().CreateBillRecord(billRecord); err != nil {
		log.Println("创建账单失败，错误信息：", err, " 账单信息为：", billRecord)
	}
}

// 获取上周起止日期
func getLastWeekStartDatetimeAndEndDatetime() (startDatetime, endDatetime datetime.Datetime) {
	// 获取当前时间
	now := datetime.Now()
	// 获取当前年月日
	year, month, day := now.Date()
	// 获取当前日期是星期几
	weekDay := datetime.Now().Weekday()
	// 计算上周开始日期
	startOfLastWeek := time.Date(year, month, day-int(weekDay+6), 0, 0, 0, 0, now.Location())
	// 计算上周结束日期
	endOfLastWeek := startOfLastWeek.AddDate(0, 0, 6)

	startDatetime, _ = datetime.ParseDatetime(startOfLastWeek.Format("2006-01-02") + " 00:00:00")
	endDatetime, _ = datetime.ParseDatetime(endOfLastWeek.Format("2006-01-02") + " 23:59:59")

	return startDatetime, endDatetime
}

// 获取上个月起止日期
func getLastMonthStartDatetimeAndEndDatetime() (startDatetime, endDatetime datetime.Datetime) {
	// 获取当前日期
	now := datetime.Now()
	// 获取上个月第一天
	firstDayOfLastMonth := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location())
	// 获取上个月最后一天
	lastDayOfLastMonth := firstDayOfLastMonth.AddDate(0, 1, -1)

	startDatetime, _ = datetime.ParseDatetime(firstDayOfLastMonth.Format("2006-01-02") + " 00:00:00")
	endDatetime, _ = datetime.ParseDatetime(lastDayOfLastMonth.Format("2006-01-02") + " 23:59:59")

	return startDatetime, endDatetime
}
