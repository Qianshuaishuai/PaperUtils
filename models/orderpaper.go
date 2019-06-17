package models

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

const (
	ORDER_BASE_URL = "/Users/DashuaiQian/Desktop/orderPaper/试题数据/智学网高中化学/化学/高一/"
)

func writeOrderSql() {
	abc()
	// rd, _ := ioutil.ReadDir("./河南试卷/英语/八年级/")

	// for r := range rd {
	// 	writePaperSql(rd[r].Name())
	// }
	// writePaperSql("2017-2018学年河南省周口市八年级（上）期末英语试卷.txt")

	// for g := range Grades {
	// 	sem = Sems[g]
	// 	grade = Grades[g]
	// 	rd, _ := ioutil.ReadDir("./内蒙古/英语/" + Grades[g])
	// 	for r := range rd {
	// 		writePaperSql(rd[r].Name())
	// 	}
	// }

	rd, _ := ioutil.ReadDir(ORDER_BASE_URL)

	// 已录到50
	for i := 1241; i < 2000; i++ {
		beego.Debug("记录点：" + strconv.Itoa(i))
		if i >= 50 {
			writeOrderPaperSql(rd[i].Name())
		}
	}

	// writeOrderPaperSql(rd[1222].Name())

	// for i := 0; i < 1241; i++ {
	// 	paperName := rd[i].Name()
	// 	paperNameData := []rune(paperName)
	// 	paperNameLength := len(paperNameData)

	// 	var count int
	// 	GetDb().Table("papers").Where("name = ?", string(paperNameData[0:paperNameLength-4])).Count(&count)

	// 	if count <= 0 {
	// 		beego.Debug("缺失试卷:" + paperName)
	// 		beego.Debug("缺失试卷:" + strconv.Itoa(i))
	// 	}
	// }

}

func writeOrderPaperSql(name string) {
	tx := GetDb().Begin()

	//新数据库表相关声明
	var paperSimple OrderPaper

	beego.Debug(ORDER_BASE_URL + name)
	paperBytes, _ := ioutil.ReadFile(ORDER_BASE_URL + name)

	var test PaperJsonSimple
	err := json.Unmarshal(paperBytes, &test)
	// if err != nil {
	// 	beego.Debug(err)
	// 	tx.Rollback()
	// 	return
	// }

	paperSimple.ID = int(flakCurl.GetIntId())
	paperSimple.OldID = test.Paper.ID
	paperSimple.CreationDate = time.Now()
	paperSimple.ModificationDate = time.Now()
	paperSimple.GradeID = translateGrade(test.Paper.Grade[0].Name)
	paperSimple.PaperType = translatePaperType(test.Paper.PaperType.Name)
	paperSimple.RealDate = getDateTimeFormatForInt(test.Paper.DateTime / 1000)
	paperSimple.Name = test.Paper.Title
	paperSimple.Score = test.AnalyseData.TotalScore
	paperSimple.SubjectID = 18 //(高中化学)

	//创建试卷表对象
	err = tx.Table("papers").Create(&paperSimple).Error
	if err != nil {
		beego.Debug(err)
		tx.Rollback()
		return
	}

	//创建试卷和省份关系
	var paperProvince OrderPaperProvince
	paperProvince.PaperID = paperSimple.ID
	paperProvince.ProvinceID = translateProvinceID(test.Paper.DefaultArea.Name)
	paperProvince.CreationDate = time.Now()
	paperProvince.ModificationDate = time.Now()
	err = tx.Table("paper_provinces").Create(&paperProvince).Error

	if err != nil {
		beego.Debug(err)
		tx.Rollback()
		return
	}

	//录入题目相关
	questionList := test.Pager.List

	for q := range questionList {
		exerciseID := int(flakCurl.GetIntId())

		//创建题和表的关系
		var exercisePaper OrderPaperQuestion
		exercisePaper.PaperID = paperSimple.ID
		exercisePaper.ExerciseID = exerciseID
		exercisePaper.CreationDate = time.Now()
		exercisePaper.ModificationDate = time.Now()

		err = tx.Table("exercise_papers").Create(&exercisePaper).Error

		if err != nil {
			beego.Debug(err)
			tx.Rollback()
			return
		}

		//创建题和知识点的关系
		keypointData := questionList[q].Knowledges
		for k := range keypointData {
			var keypointQuestion OrderKeyPointQuestion
			keypointQuestion.ExerciseID = exerciseID
			keypointQuestion.KeypointID = translateKeypoint(keypointData[k].Name)
			keypointQuestion.CreationDate = time.Now()
			keypointQuestion.ModificationDate = time.Now()

			err = tx.Table("exercise_keypoints").Create(&keypointQuestion).Error

			if err != nil {
				beego.Debug(err)
				tx.Rollback()
				return
			}
		}

		if len(questionList[q].SubQuestion) > 1 {
			beego.Debug("有大题")
		} else {
			var exercise OrderExercise
			exercise.ID = exerciseID
			exercise.Content = questionList[q].SubQuestion[0].Stem
			exercise.CreationDate = time.Now()
			exercise.ModificationDate = time.Now()
			exercise.Difficulty = questionList[q].Difficulty.Value
			exercise.ExerciseType = translateExerciseType(questionList[q].Section.Name)
			exercise.SubjectID = 18 //(高中化学)
			exercise.Score = questionList[q].Score
			exercise.Analysis = questionList[q].SubQuestion[0].OriginalStruct.AnalysisHTML

			err = tx.Table("exercise_info").Create(&exercise).Error

			if err != nil {
				beego.Debug(err)
				tx.Rollback()
				return
			}

			if translateIsOption(questionList[q].Section.Name) == 1 {
				optionDatas := questionList[q].SubQuestion[0].Options
				for o := range optionDatas {
					var question OrderQuestion
					question.ID = int(flakCurl.GetIntId())
					question.ExerciseID = exerciseID
					question.CreationDate = time.Now()
					question.ModificationDate = time.Now()
					question.QuestionScore = questionList[q].Score
					question.QuestionIndex = o
					question.QuestionType = 101
					question.Question = optionDatas[o].Desc

					err = tx.Table("exercise_question").Create(&question).Error

					if err != nil {
						beego.Debug(err)
						tx.Rollback()
						return
					}
				}
			}

			if translateIsOption(questionList[q].Section.Name) == 1 {
				answerData := questionList[q].SubQuestion[0].Answers
				for a := range answerData {

					answerStrDatas := []rune(answerData[a].Desc)
					answerLength := len(answerStrDatas)
					for i := 0; i < answerLength; i++ {
						min := i
						max := i + 1
						var answer OrderAnswer
						answer.ID = int(flakCurl.GetIntId())
						answer.ExerciseID = exerciseID
						answer.CreationDate = time.Now()
						answer.ModificationDate = time.Now()
						answer.Answer = string(answerStrDatas[min:max])
						// answer.
						answer.IsCorrect = 1

						err = tx.Table("exercise_answer").Create(&answer).Error

						if err != nil {
							beego.Debug(err)
							tx.Rollback()
							return
						}
					}

				}
			} else {
				answerData := questionList[q].SubQuestion[0].Answers
				for a := range answerData {
					var answer OrderAnswer
					answer.ID = int(flakCurl.GetIntId())
					answer.ExerciseID = exerciseID
					answer.CreationDate = time.Now()
					answer.ModificationDate = time.Now()
					answer.Answer = answerData[a].Desc

					err = tx.Table("exercise_answer").Create(&answer).Error

					if err != nil {
						beego.Debug(err)
						tx.Rollback()
						return
					}
				}
			}
		}

	}

	tx.Commit()
}

func changeFileCode(name string) {

}
