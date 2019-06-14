package models

import "time"

type OPaperSimple struct {
	ID               int       `gorm:"column:id"`
	Name             string    `gorm:"column:name"`
	PpaerType        string    `gorm:"column:paper_type"`
	Province         int       `gorm:"column:province"`
	Grade            int       `gorm:"column:grade"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}

type OMaterial struct {
	ID               int       `gorm:"column:id"`
	Name             string    `gorm:"column:name"`
	Parent           int       `gorm:"column:parent"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}

type OKeypoint struct {
	ID               int       `gorm:"column:id"`
	Name             string    `gorm:"column:name"`
	Parent           int       `gorm:"column:parent"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}

func TranslatePaperType(str string) int {
	switch str {
	case "历年真题":
		return 1
	case "模拟题":
		return 2
	case "入学测验":
		return 3
	case "期末考试":
		return 4
	case "期中考试":
		return 5
	case "月考试卷":
		return 6
	case "其他类型":
		return 7
	case "单元测试":
		return 8
	default:
		return 0
	}
}

func TranslateProvince(str string) int {
	switch str {
	case "全国":
		return 1
	case "北京":
		return 2
	case "天津":
		return 3
	case "河北":
		return 4
	case "山西":
		return 5
	case "辽宁":
		return 6
	case "吉林":
		return 7
	case "上海":
		return 8
	case "江苏":
		return 9
	case "浙江":
		return 10
	case "安徽":
		return 11
	case "福建":
		return 12
	case "江西":
		return 13
	case "山东":
		return 14
	case "河南":
		return 15
	case "湖北":
		return 16
	case "湖南":
		return 17
	case "广东":
		return 18
	case "广西":
		return 19
	case "海南":
		return 20
	case "重庆":
		return 21
	case "四川":
		return 22
	case "贵州":
		return 23
	case "云南":
		return 24
	case "西藏":
		return 25
	case "陕西":
		return 26
	case "甘肃":
		return 27
	case "青海":
		return 28
	case "宁夏":
		return 29
	case "新疆":
		return 30
	case "黑龙江":
		return 31
	case "内蒙古":
		return 32
	default:
		return 0
	}
}

func TranslateGrade(str string) int {
	switch str {
	case "一年级":
		return 1
	case "二年级":
		return 2
	case "三年级":
		return 3
	case "四年级":
		return 4
	case "五年级":
		return 5
	case "六年级":
		return 6
	case "初一":
		return 7
	case "初二":
		return 8
	case "初三":
		return 9
	case "高一":
		return 10
	case "高二":
		return 11
	case "高三":
		return 12
	default:
		return 0
	}
}

func TranslateDiffculty(str string) int {
	switch str {
	case "容易":
		return 1
	case "较易":
		return 2
	case "一般":
		return 3
	case "较难":
		return 4
	case "困难":
		return 5
	default:
		return 0
	}
}
