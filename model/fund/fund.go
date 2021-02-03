package model

import (
	"fund/db/dbconnect"
	"gorm.io/gorm"
)

type Fund struct {
	gorm.Model
	Code 					string `json:"code"`					//	基金代码
	Name 					string `json:"name"`					//	基金名称
	NetWorth 				float64 `json:"netWorth"`				//	当前基金单位净值
	ExpectWorth 			float64 `json:"expectWorth"`			//	当前基金单位净值估算
	ExpectGrowth 			string `json:"expectGrowth"`			//	当前基金单位净值估算日涨幅,单位为百分比
	DayGrowth				string `json:"dayGrowth"`				//	单位净值日涨幅,单位为百分比
	LastWeekGrowth 			string `json:"lastWeekGrowth"`			//	单位净值周涨幅,单位为百分比
	LastMonthGrowth 		string `json:"lastMonthGrowth"`			//	单位净值月涨幅,单位为百分比
	LastThreeMonthsGrowth 	string `json:"lastThreeMonthsGrowth"`	//	单位净值三月涨幅,单位为百分比
	LastSixMonthsGrowth 	string `json:"lastSixMonthsGrowth"`		//	单位净值六月涨幅,单位为百分比
	LastYearGrowth 			string `json:"lastYearGrowth"`			//	单位净值年涨幅,单位为百分比
	NetWorthDate 			string `json:"netWorthDate"`			//	净值更新日期,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
	ExpectWorthDate 		string `json:"expectWorthDate"`			//	净值估算更新日期,,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
}

func init()  {
	//	判断数据库中是否存在该表，不存在则新建
	h := dbconnect.PgGormDb.Migrator().HasTable(&Fund{})
	if !h {
		dbconnect.PgGormDb.Migrator().CreateTable(&Fund{})
	}
}