//所有筛选条件的数据插入配置文件
package models

// import (
// 	"github.com/jinzhu/gorm"
// )

// //年级结构体

// type Grade struct {
// 	GradeID       int    `gorm:"primary_key;column:F_grade_id" json:"gradeId"` //年级id
// 	Name          string `gorm:"column:F_name" json:"name"`                    //年级的名称
// 	MaterialCount int    `gorm:"column:F_material_count" json:"materialCount"` //这个年级下的文章总数
// }

// //文体结构体
// type Literary struct {
// 	LiteraryID    int    `gorm:"primary_key;column:F_literary_id" json:"literaryId"` //文体id
// 	Name          string `gorm:"column:F_name" json:"name"`                          //文体的名称
// 	MaterialCount int    `gorm:"column:F_material_count" json:"materialCount"`       //这个文体下的文章总数
// }

// //字数范围结构体
// type WordNumber struct {
// 	WordNumberID  int    `gorm:"primary_key;column:F_word_number_id" json:"wordNumberId"` //字数范围的id
// 	Name          string `gorm:"column:F_name" json:"name"`                               //字数范围的名称
// 	MaterialCount int    `gorm:"column:F_material_count" json:"materialCount"`            //这个字体范围下的文章总数
// }

// var H = []uint{10, 11, 13}                          //高中年级
// var M = []uint{7, 8, 9, 12}                         //初中年级
// var S = []int{-1, 1, 2, 3, 4, 5, 6}                 //小学年级
// var Sn = []string{"全部", "三年级", "四年级", "五年级", "六年级"} //小学年级

// var L = []int{-1, 1, 2, 3, 4, 5, 6, 7, 12, 13, 14}
// var Ln = []string{"全部", "记叙文", "散文", "小说", "诗歌", "说明文", "议论文", "应用文", "古文体", "寓言", "其它"}

// var W = []int{-1, 1, 2, 3, 4, 5, 6}
// var Wn = []string{"全部", "200字以下", "200-400字", "400-800字", "800-1000字", "1000-1200字", "1200字以上"}

// func AddGradeList(db *gorm.DB) {
// 	for index := 0; index < 5; index++ {
// 		var MGrade = Grade{}
// 		MGrade.GradeID = S[index]
// 		MGrade.Name = Sn[index]
// 		MGrade.MaterialCount = 33
// 		db.Create(MGrade)
// 	}
// }

// func AddLiteraryList(db *gorm.DB) {
// 	for i := range Ln {
// 		var MLiterary = Literary{}
// 		MLiterary.LiteraryID = L[i]
// 		MLiterary.Name = Ln[i]
// 		MLiterary.MaterialCount = 33
// 		db.Create(MLiterary)
// 	}
// }

// func AddWordNumberList(db *gorm.DB) {
// 	for i := range Wn {
// 		var MWordNumber = WordNumber{}
// 		MWordNumber.WordNumberID = W[i]
// 		MWordNumber.Name = Wn[i]
// 		MWordNumber.MaterialCount = 33
// 		db.Create(MWordNumber)
// 	}
// }
