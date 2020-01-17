package models

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego"
)

type OProvince struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Citys []OCity `json:"cityList"`
}

type OCity struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Areas []OArea `json:"areaList"`
}

type OArea struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ProvinceCity struct {
	ID               int       `gorm:"column:Id"`
	Name             string    `gorm:"column:Name"`
	Level            int       `gorm:"column:Level"`
	Code             string    `gorm:"column:Code"`
	ParentId         int       `gorm:"column:ParentId"`
	CreationDate     time.Time `gorm:"column:Create_Date"`
	ModificationDate time.Time `gorm:"column:Modify_Date"`
}

func addProvinceData() {
	tx := GetDb().Begin()
	var test []OProvince
	paperBytes, _ := ioutil.ReadFile("/home/dashuai/go/src/PaperUtils/知识点文件/官方省份数据.txt")
	index := 1
	err := json.Unmarshal(paperBytes, &test)
	if err != nil {
		beego.Debug(err)
		return
	}

	for t := range test {
		province := test[t]
		var data1 ProvinceCity
		data1.ID = index
		data1.Name = province.Name
		data1.Code = province.Code
		data1.Level = 1
		data1.ParentId = 0
		data1.CreationDate = time.Now()
		data1.ModificationDate = time.Now()
		index = index + 1
		err := tx.Table("ba_province_city").Create(&data1).Error
		beego.Debug("a")
		if err != nil {
			beego.Debug(err)
			tx.Rollback()

			return
		}

		for c := range province.Citys {
			city := province.Citys[c]
			var data2 ProvinceCity
			data2.ID = index
			data2.Name = city.Name
			data2.Code = city.Code
			data2.Level = 2
			data2.ParentId = data1.ID
			data2.CreationDate = time.Now()
			data2.ModificationDate = time.Now()
			index = index + 1
			err = tx.Table("ba_province_city").Create(&data2).Error

			if err != nil {
				beego.Debug(err)
				tx.Rollback()
				return
			}

			for a := range city.Areas {
				area := city.Areas[a]
				var data3 ProvinceCity
				data3.ID = index
				data3.Name = area.Name
				data3.Code = area.Code
				data3.Level = 3
				data3.ParentId = data2.ID
				data3.CreationDate = time.Now()
				data3.ModificationDate = time.Now()
				index = index + 1

				err = tx.Table("ba_province_city").Create(&data3).Error

				if err != nil {
					beego.Debug(err)
					tx.Rollback()
					return
				}
			}
		}
	}

	tx.Commit()
}
