package models

// import (
// 	"dreamEbagPapers/helper"
// 	"regexp"
// 	"strconv"
// 	"strings"

// 	"github.com/astaxie/beego"
// 	"github.com/jinzhu/gorm"
// )

// type MaterialSimple struct {
// 	ID             int    `gorm:"primary_key;column:F_id;type:BIGINT(20)" json:"id"` //材料的id
// 	Author         string `gorm:"column:F_author;type:TEXT" json:"author"`           //材料作者
// 	Prompt         string `gorm:"column:F_prompt;type:TEXT" json:"prompt"`           //材料教师修改提示
// 	Content        string `gorm:"column:F_content;type:TEXT" json:"content"`         //材料内容
// 	Title          string `gorm:"column:F_title" json:"title"`                       //材料标题
// 	Type           int    `gorm:"column:F_type" json:"type"`                         //材料类型
// 	WordCount      int    `gorm:"column:F_word_count" json:"wordCount"`              //材料文章字数
// 	GradeID        int    `gorm:"column:F_grade_id" json:"gradeId"`                  //对应年级id（额外添加标签）
// 	LiteraryID     int    `gorm:"column:F_literary_id" json:"literaryId"`            //对应文体id（额外添加标签）
// 	WordNumberID   int    `gorm:"column:F_word_number_id" json:"wordNumberId"`       //对应字数id（额外添加标签）
// 	TeacherID      int    `gorm:"column:F_teacher_id" json:"teacherId"`              //对应老师的id
// 	ReferenceCount int    `gorm:"column:F_reference_count" json:"-"`                 //引用次数
// }

// func GenerateMaterialsForManager(managedb *gorm.DB) {
// 	abc()
// 	var material []Material
// 	managedb.Table("t_materials").Find(&material)

// 	for i := range material {
// 		var titleTxt = material[i].Title
// 		var contentTxt = material[i].Content
// 		var contentCount = 0

// 		if titleTxt != "" && !IsHaveEqualTitle(material[i].Title, material[i].GradeID) {
// 			contentTxt = strings.Replace(contentTxt, titleTxt, "", 1)
// 			contentTxt = trimUnNeedHTML(contentTxt)
// 			contentCount = strings.Count(trimHtml(contentTxt), "") - 1

// 			var authorTxt string
// 			authorStartCount := strings.Index(contentTxt, "<p style=\"text-indent: 2em; text-align: center;\">")
// 			authorEndCount := strings.Index(contentTxt, "</p>")

// 			if authorStartCount != -1 && authorStartCount < authorEndCount {
// 				authorTxt = trimHtml(string(contentTxt[authorStartCount:authorEndCount]))
// 			} else {
// 				authorTxt = ""
// 			}

// 			if authorTxt == titleTxt {
// 				authorTxt = ""
// 			}

// 			if strings.Contains(authorTxt, "（节选）") {
// 				authorTxt = ""
// 			}

// 			if len(authorTxt) > 15 {
// 				authorTxt = ""
// 			}

// 			if strings.Count(authorTxt, "")-1 <= 5 {
// 				contentTxt = strings.Replace(contentTxt, authorTxt, "", 1)
// 			}

// 			if strings.Contains(contentTxt, "乔 叶") {
// 				contentTxt = strings.Replace(contentTxt, "乔 叶", "", 1)
// 			}

// 			if strings.Contains(contentTxt, "（节选）") {
// 				reg := regexp.MustCompile("<p style=\"text-align: center;\"><strong>.*?</strong>.*?</p>")
// 				styleStr := reg.FindAllString(contentTxt, 1)
// 				for a := range styleStr {
// 					contentTxt = strings.Replace(contentTxt, styleStr[a], "", 1)
// 				}
// 			}

// 			contentTxt = trimUnNeedHTML(contentTxt)

// 			material[i].Author = authorTxt
// 			material[i].Prompt = ""
// 			material[i].Content = contentTxt

// 			material[i].ID = (int)(flakCurl.GetIntId())

// 			questionIds := strings.Split(material[i].QuestionIds, ",")
// 			var newQuestionIds string
// 			for q := range questionIds {
// 				questionID, _ := strconv.Atoi(questionIds[q])
// 				question := GetQuestionForID(questionID)

// 				if IsEfficientReadingQuestion(question.Content) || IsNoNeedReadingQuestion(question.Type) {
// 					questionId := (int)(flakCurl.GetIntId())
// 					if newQuestionIds == "" {
// 						newQuestionIds = newQuestionIds + strconv.Itoa(questionId)
// 					} else {
// 						newQuestionIds = newQuestionIds + "," + strconv.Itoa(questionId)
// 					}

// 					//option
// 					optionIDs := strings.Split(question.OptionIDs, ",")
// 					var newOptionIds string
// 					if question.OptionIDs != "" {
// 						for o := range optionIDs {
// 							optionID, _ := strconv.Atoi(optionIDs[o])
// 							optionId := (int)(flakCurl.GetIntId())
// 							option := GetOptionForID(optionID)
// 							if !strings.Contains(option.OptionContent, "img") {
// 								if o == 0 {
// 									newOptionIds = newOptionIds + strconv.Itoa(optionId)
// 								} else {
// 									newOptionIds = newOptionIds + "," + strconv.Itoa(optionId)
// 								}
// 								if strings.Contains(option.OptionContent, "<!-- -->") {
// 									txtDatas := strings.Split(option.OptionContent, "<!-- -->")
// 									var optionSmallTxt = make([]string, len(txtDatas))
// 									for t := range txtDatas {
// 										optionSmallTxt[t] = "\"" + helper.InterfaceToString(txtDatas[t]) + "\""
// 									}
// 									option.OptionContent = "[" + strings.Join(optionSmallTxt, ",") + "]"
// 								} else if strings.Contains(option.OptionContent, "55965e902ed9b6276116bc6d") {
// 									option.OptionContent = "T"
// 								} else if strings.Contains(option.OptionContent, "55965e942ed9b62763dddc64") {
// 									option.OptionContent = "F"
// 								} else if strings.Contains(option.OptionContent, "55965e942ed9b62763dddc64") {
// 									option.OptionContent = "F"
// 								}
// 								// beego.Debug(option)
// 								beego.Debug(optionIDs)
// 								option.OptionID = optionId
// 								GetDb().Create(option)
// 							}
// 						}
// 					}

// 					//answer
// 					answerIDs := strings.Split(question.AnswerIDs, ",")
// 					var newAnswerIds string

// 					for a := range answerIDs {
// 						answerID, _ := strconv.Atoi(answerIDs[a])
// 						answerId := (int)(flakCurl.GetIntId())

// 						if a == 0 {
// 							newAnswerIds = newAnswerIds + strconv.Itoa(answerId)
// 						} else {
// 							newAnswerIds = newAnswerIds + "," + strconv.Itoa(answerId)
// 						}

// 						answer := GetAnswerForID(answerID)
// 						if question.Type == 10021 {
// 							optionCount := string(answer.AnswerContent[0:1])
// 							answerCount := string(answer.AnswerContent[2:3])
// 							answerInt, _ := strconv.Atoi(answerCount)
// 							answerInt = answerInt - len(answerIDs)

// 							answer.AnswerContent = optionCount + "," + strconv.Itoa(answerInt)
// 						}

// 						if question.Type == 10019 {
// 							answerNumber, err := strconv.Atoi(answer.AnswerContent)
// 							if err == nil {
// 								// answerContentTxt := materialSubContenMap["options"].([]interface{})[answerNumber].(map[string]interface{})["option"].(string)
// 								// answerInt, _ := strconv.Atoi(answerContentTxt)
// 								// beego.Debug("answerInt:" + answerContentTxt)
// 								answer.AnswerContent = answerArray[answerNumber]
// 							}
// 						}

// 						answer.AnswerID = answerId
// 						GetDb().Create(answer)
// 					}

// 					if question.Type == 10021 {
// 						question.Type = 10009
// 					}

// 					if question.Type == 10019 {
// 						question.Type = 10005
// 					}

// 					question.QuestionID = questionId
// 					question.AnswerIDs = newAnswerIds
// 					question.OptionIDs = newOptionIds
// 					if question.Type != 10020 {
// 						GetDb().Create(question)
// 					}
// 				}
// 			}
// 			material[i].WordCount = contentCount
// 			material[i].QuestionIds = newQuestionIds
// 			material[i].WordNumberID = getWordNumberType(contentCount)
// 			GetDb().Create(material[i])
// 		}
// 	}
// }

// func FixAnswers() {
// 	var questions []Question
// 	GetDb().Table("t_questions").Where("F_content = ?", "").Find(&questions)
// 	for q := range questions {
// 		var newContent = "<p>连一连</p>"
// 		beego.Debug(q)
// 		GetDb().Table("t_questions").Where("F_content = ?", "").Update("F_content", newContent)
// 	}
// 	beego.Debug(len(questions))
// }

// //通过questionID搜索对应的题目资源
// func GetQuestionForID(questionID int) (question Question) {
// 	GetManageDb().Where("F_question_id= ?", questionID).Find(&question)
// 	return question
// }

// //通过optionID搜索题目选项
// func GetOptionForID(optionID int) (option Option) {
// 	GetManageDb().Where("F_option_id= ?", optionID).Find(&option)
// 	return option
// }

// //通过answerID搜索题目选项
// func GetAnswerForID(answerID int) (answer Answer) {
// 	GetManageDb().Where("F_answer_id= ?", answerID).Find(&answer)
// 	return answer
// }

// //判断是否是“非高效阅读类题目”
// func IsEfficientReadingQuestion(str string) bool {
// 	if strings.Contains(str, "读音") || strings.Contains(str, "正确的字") || strings.Contains(str, "拼音") || strings.Contains(str, "对应词语") || strings.Contains(str, "近义词") || strings.Contains(str, "反义词") || strings.Contains(str, "正确读音") {
// 		return false
// 	}
// 	return true
// }

// //判断是否是不要的问题
// func IsNoNeedReadingQuestion(questionType int) bool {
// 	if questionType >= 8 && questionType <= 11 {
// 		return false
// 	}
// 	return true
// }

// //判断是否有同样的标题
// func IsHaveEqualTitle(title string, grade int) bool {
// 	var count = 0
// 	GetDb().Table("t_materials").Where("F_title = ? AND F_grade_id = ?", title, grade).Count(&count)
// 	beego.Debug(count)
// 	if count == 0 {
// 		return false
// 	} else {
// 		return true
// 	}
// }
