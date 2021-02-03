package fund

import (
	"encoding/json"
	"fund/db/dbconnect"
	"fund/model/fund"
	"fund/util/ginResponseUtil"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
)

//	保存数据
func AddFunds(c *gin.Context)  {
	response := ginResponseUtil.Gin{c}
	var funds []model.Fund
	data, _ := ioutil.ReadAll(c.Request.Body)
	var mymap map[string][]model.Fund
	json.Unmarshal(data,&mymap)
	funds = mymap["data"]
	if err :=dbconnect.PgGormDb.Create(&funds).Error;err != nil {
		log.Println("批量保存数据失败：",err)
		response.Error()
		return
	}
	response.Ok()
}

type returnData struct {
	ExpectGrowth 			string `json:"expectGrowth"`			//	当前基金单位净值估算日涨幅,单位为百分比
	ExpectWorthDate 		string `json:"expectWorthDate"`			//	净值估算更新日期,,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
}
//	查询数据
func GetFundDetail(c *gin.Context)  {
	response := ginResponseUtil.Gin{c}
	code := c.Query("code")
	var returnDatas []returnData
	dbconnect.PgGormDb.Model(&model.Fund{}).Select("expect_growth,expect_worth_date").Where(" code = ? and expect_worth_date like ?",code,time.Now().Format("2006-01-02")+"%").Scan(&returnDatas)
	var expectWorthDate []string
	var expectGrowth []string
	for _, v := range returnDatas {
		expectGrowth = append(expectGrowth, v.ExpectGrowth)
		expectWorthDate = append(expectWorthDate, v.ExpectWorthDate)
	}
	var myMap = make(map[string]interface{})
	myMap["expectWorthDate"] = expectWorthDate
	myMap["expectGrowth"] = expectGrowth
	response.OkData(myMap)
}