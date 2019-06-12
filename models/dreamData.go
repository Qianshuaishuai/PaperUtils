package models

// import (
// 	"dreamEbagPapers/helper"
// 	"encoding/json"
// 	"io/ioutil"
// 	"regexp"
// 	"strconv"
// 	"strings"

// 	"github.com/astaxie/beego"
// 	"github.com/jinzhu/gorm"
// )

// var teacherID = "0" //默认插入数据的老师id为0.以0或非0来识别是否为老师新增资源
// var checkTitleStr = "<p style=\"text-indent: 2em; text-align: center;\"><strong>"
// var flakCurl MSnowflakCurl

// // var keyPointID = 1
// var optionID = 1
// var gradeArray = [6]int{1, 2, 3, 4, 5, 6}

// //多次填空答案转换
// var answerArray = [20]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T"}

// //默认
// var materialType = 10019
// var subject = 1
// var version = 1

// type Material struct {
// 	ID           int    `gorm:"primary_key;column:F_id;type:BIGINT(20)" json:"id"`              //材料的id
// 	Author       string `gorm:"column:F_author;type:TEXT" json:"author"`                        //材料作者
// 	Prompt       string `gorm:"column:F_prompt;type:TEXT" json:"prompt"`                        //材料教师修改提示
// 	Content      string `gorm:"column:F_content;type:TEXT" json:"content"`                      //材料内容
// 	Type         int    `gorm:"column:F_type" json:"type"`                                      //材料类型
// 	AccessorieID int    `gorm:"column:F_accessories_id;type:INT UNSIGNED" json:"accessoriesId"` //材料附属资源id
// 	Title        string `gorm:"column:F_title" json:"title"`                                    //材料标题
// 	WordCount    int    `gorm:"column:F_word_count" json:"wordCount"`                           //材料文章字数
// 	QuestionIds  string `gorm:"column:F_question_ids" json:"questionIds"`                       //材料题目id
// 	//以下为增加的筛选标签
// 	GradeID      int    `gorm:"column:F_grade_id" json:"gradeId"`               //对应年级id（额外添加标签）
// 	LiteraryID   int    `gorm:"column:F_literary_id" json:"literaryId"`         //对应文体id（额外添加标签）
// 	WordNumberID int    `gorm:"column:F_word_number_id" json:"wordNumberId"`    //对应字数id（额外添加标签）
// 	TeacherID    string `gorm:"column:F_teacher_id;type:TEXT" json:"teacherId"` //对应字数id（额外添加标签）
// }

// type Accessorie struct {
// 	AccessorieID int    `gorm:"column:F_accessorie_id;type:INT UNSIGNED" json:"accessorieId"` //材料附属资源id
// 	Type         int    `gorm:"column:F_type" json:"type"`                                    //材料附属资源类型
// 	Label        string `gorm:"column:F_label" json:"label"`                                  //材料附属资源标签
// 	Content      string `gorm:"column:F_content" json:"content"`                              //材料附属资源内容
// }

// type Question struct {
// 	QuestionID           int    `gorm:"column:F_question_id;type:BIGINT(20)" json:"questionId"`                        //材料题目id
// 	Type                 int    `gorm:"column:F_type" json:"type"`                                                     //题目类型
// 	Subject              int    `gorm:"column:F_subject" json:"subject"`                                               //所属科目
// 	Content              string `gorm:"column:F_content;type:TEXT" json:"content"`                                     //题目正文
// 	Solution             string `gorm:"column:F_solution" json:"solution"`                                             //题目解析
// 	SolutionAccessorieID int    `gorm:"column:F_solution_accessorie_id;type:INT UNSIGNED" json:"solutionAccessorieId"` //题目解析附加资源id
// 	V                    int    `gorm:"column:F_v" json:"v"`                                                           //数据库版本
// 	Difficulty           int    `gorm:"column:F_difficulty" json:"difficulty"`                                         //难度系数
// 	KeyPointID           int    `gorm:"column:F_key_point_id;type:INT UNSIGNED" json:"keyPointId"`                     //考点id
// 	OptionIDs            string `gorm:"column:F_option_id" json:"optionIds"`                                           //
// 	AnswerIDs            string `gorm:"column:F_answer_id" json:"anwserIds"`                                           //答案Id
// }

// type SolutionAccessorie struct {
// 	SolutionAccessorieID int    `gorm:"column:F_solution_accessorie_id;type:BIGINT(20)" json:"solutionAccessorieId"` //题目附属资源id
// 	Type                 int    `gorm:"column:F_type" json:"type"`                                                   //题目附属资源类型
// 	Label                string `gorm:"column:F_label" json:"label"`                                                 //题目附属资源标签
// 	Content              string `gorm:"column:F_content;type:TEXT" json:"content"`                                   //题目附属资源内容
// }

// type Option struct {
// 	OptionID      int    `gorm:"column:F_option_id;type:BIGINT(20)" json:"optionId"`     //
// 	OptionContent string `gorm:"column:F_option_content;type:TEXT" json:"optionContent"` //
// }

// type KeyPoint struct {
// 	KeyPointID int    `gorm:"column:F_key_point_id;type:BIGINT(20)" json:"keyPointId"` //考点id
// 	Name       string `gorm:"column:F_name;type:TEXT" json:"name"`                     //考点名称
// }

// type Answer struct {
// 	AnswerID      int    `gorm:"column:F_answer_id;type:BIGINT(20)" json:"anwserId"`     //答案Id
// 	AnswerContent string `gorm:"column:F_answer_content;type:TEXT" json:"answerContent"` //答案名称
// }

// func GenerateMaterialsForzuoye(db *gorm.DB, suffix string, index int) {
// 	materialBytes, _ := ioutil.ReadFile("./res/一起作业网高效阅读数据/" + strconv.Itoa(index) + "/" + suffix)
// 	beego.Debug("./res/一起作业网高效阅读数据/" + strconv.Itoa(index) + "/" + suffix)
// 	idLength := len(suffix)
// 	idName := string([]byte(suffix)[:idLength-5])

// 	var materialObj interface{}
// 	json.Unmarshal(materialBytes, &materialObj)

// 	if materialObj != nil {
// 		materialMap := materialObj.(map[string]interface{})
// 		materialResultMap := materialMap["result"].(map[string]interface{})
// 		materialIDMap := materialResultMap[idName].(map[string]interface{})
// 		materialQueMap := materialIDMap["questions"].([]interface{})
// 		materialQueeMap := materialQueMap[0].(map[string]interface{})
// 		materialContentMap := materialQueeMap["content"].(map[string]interface{})

// 		//subject
// 		materialSubContenObj := materialContentMap["subContents"].([]interface{})

// 		var materialBean = Material{}
// 		var difficultyInt int
// 		var contentTxt string
// 		var questionContentTxt string
// 		var questionType int
// 		var solutionTxt = ""
// 		var solutionAccessorieID = 0
// 		var keyPointID = 0
// 		var wordNumberId = 0
// 		var titleTxt = ""
// 		var contentCount = 0
// 		var oldType = 0

// 		//数据解析
// 		if materialQueeMap["difficultyInt"] != nil {
// 			difficultyInt = (int)(materialQueeMap["difficultyInt"].(float64))
// 		}

// 		if materialContentMap["content"] != nil {
// 			contentTxt = materialContentMap["content"].(string)
// 			contentTxtNoHtml := trimHtml(contentTxt)
// 			contentCount = strings.Count(contentTxtNoHtml, "") - 1
// 			wordNumberId = getWordNumberType(contentCount)
// 		}

// 		if strings.Contains(contentTxt, checkTitleStr) {
// 			titleStartCount := strings.Index(contentTxt, "<strong>") + 8
// 			titleEndCount := strings.Index(contentTxt, "</strong>")
// 			titleTxt = trimHtml(string(contentTxt[titleStartCount:titleEndCount]))

// 			materialBean.ID = (int)(flakCurl.GetIntId())

// 			if index == 2 || index == 1 {
// 				index = 3
// 			}

// 			materialBean.WordCount = contentCount
// 			materialBean.Type = materialType
// 			materialBean.GradeID = index
// 			materialBean.LiteraryID = -1
// 			materialBean.WordNumberID = wordNumberId
// 			materialBean.AccessorieID = 0
// 			materialBean.TeacherID = teacherID
// 			// materialBean.QuestionIDs = questionID

// 			materialBean.Title = titleTxt

// 			contentTxt = strings.Replace(contentTxt, titleTxt, "", 1)

// 			contentTxt = trimUnNeedHTML(contentTxt)

// 			// r, _ := regexp.Compile("<p style=\"text-indent: 2em; text-align: center;\"><strong>(.*)</strong></p>")
// 			// filtTxt := r.FindAllString(contentTxt, 2)

// 			// for i := range filtTxt {
// 			// 	contentTxt = strings.Replace(contentTxt, filtTxt[i], "", 1)
// 			// }

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

// 			contentTxt = trimUnNeedHTML(contentTxt)

// 			materialBean.Author = authorTxt
// 			materialBean.Prompt = ""
// 			materialBean.Content = contentTxt

// 			var questionIds string
// 			for i := range materialSubContenObj {
// 				questionId := (int)(flakCurl.GetIntId())
// 				if i == 0 {
// 					questionIds = questionIds + strconv.Itoa(questionId)
// 				} else {
// 					questionIds = questionIds + "," + strconv.Itoa(questionId)
// 				}

// 				materialSubContenMap := materialSubContenObj[i].(map[string]interface{})

// 				if materialSubContenMap["subContentTypeId"] != nil {
// 					oldType = (int)(materialSubContenMap["subContentTypeId"].(float64))
// 					questionType = (int)(materialSubContenMap["subContentTypeId"].(float64))
// 					questionType = translateDreamDataQuestionType(questionType)
// 				}

// 				if materialSubContenMap["content"] != nil {
// 					questionContentTxt = materialSubContenMap["content"].(string)
// 					if oldType == 10 {
// 						questionContentTxt = questionContentTxt + "<br>"
// 						materialOptionsObj := materialSubContenMap["options"].([]interface{})
// 						for a := range materialOptionsObj {
// 							materialOptionMap := materialOptionsObj[a].(map[string]interface{})
// 							optionContent := materialOptionMap["option"].(string)
// 							beego.Debug(a)
// 							questionContentTxt = questionContentTxt + "&nbsp;&nbsp;" + answerArray[a] + "." + optionContent
// 						}
// 					}
// 				}

// 				materialAnswerObj := materialSubContenMap["answers"].([]interface{})
// 				var answerIds string
// 				for h := range materialAnswerObj {
// 					answerId := (int)(flakCurl.GetIntId())
// 					if h == 0 {
// 						answerIds = answerIds + strconv.Itoa(answerId)
// 					} else {
// 						answerIds = answerIds + "," + strconv.Itoa(answerId)
// 					}

// 					materialAnswerMap := materialAnswerObj[h].(map[string]interface{})
// 					var answer = Answer{}
// 					answer.AnswerID = answerId
// 					answer.AnswerContent = materialAnswerMap["answer"].(string)

// 					if questionType == 10009 {
// 						optionCount := string(answer.AnswerContent[0:1])
// 						answerCount := string(answer.AnswerContent[2:3])
// 						answerInt, _ := strconv.Atoi(answerCount)
// 						answerInt = answerInt - len(materialAnswerObj)

// 						answer.AnswerContent = optionCount + "," + strconv.Itoa(answerInt)
// 					}

// 					if (materialSubContenMap["subContentTypeId"].(float64)) == 10 {
// 						answerNumber, err := strconv.Atoi(answer.AnswerContent)
// 						if err == nil {
// 							// answerContentTxt := materialSubContenMap["options"].([]interface{})[answerNumber].(map[string]interface{})["option"].(string)
// 							// answerInt, _ := strconv.Atoi(answerContentTxt)
// 							// beego.Debug("answerInt:" + answerContentTxt)
// 							answer.AnswerContent = answerArray[answerNumber]
// 						}
// 					}

// 					db.Create(answer)
// 				}

// 				var optionIds string
// 				if materialSubContenMap["options"] != nil {
// 					materialOptionsObj := materialSubContenMap["options"].([]interface{})
// 					for a := range materialOptionsObj {
// 						optionId := (int)(flakCurl.GetIntId())
// 						if a == 0 {
// 							optionIds = optionIds + strconv.Itoa(optionId)
// 						} else {
// 							optionIds = optionIds + "," + strconv.Itoa(optionId)
// 						}

// 						materialOptionMap := materialOptionsObj[a].(map[string]interface{})
// 						var option = Option{}
// 						option.OptionID = optionId
// 						option.OptionContent = materialOptionMap["option"].(string)
// 						if strings.Contains(option.OptionContent, "<!-- -->") {
// 							txtDatas := strings.Split(option.OptionContent, "<!-- -->")
// 							var optionSmallTxt = make([]string, len(txtDatas))
// 							for t := range txtDatas {
// 								optionSmallTxt[t] = "\"" + helper.InterfaceToString(txtDatas[t]) + "\""
// 							}
// 							option.OptionContent = "[" + strings.Join(optionSmallTxt, ",") + "]"
// 							beego.Debug(option.OptionContent)
// 						} else if strings.Contains(option.OptionContent, "55965e902ed9b6276116bc6d") {
// 							option.OptionContent = "T"
// 						} else if strings.Contains(option.OptionContent, "55965e942ed9b62763dddc64") {
// 							option.OptionContent = "F"
// 						}
// 						db.Create(option)
// 					}

// 				} else {
// 					optionIds = ""
// 				}

// 				var question = Question{}
// 				question.QuestionID = questionId
// 				question.Content = questionContentTxt
// 				question.V = version
// 				question.Difficulty = difficultyInt
// 				question.KeyPointID = keyPointID
// 				question.Solution = solutionTxt
// 				question.SolutionAccessorieID = solutionAccessorieID
// 				question.Subject = subject
// 				question.Type = questionType
// 				question.AnswerIDs = answerIds
// 				question.OptionIDs = optionIds
// 				if questionType != 10020 {
// 					db.Create(question)
// 				}
// 			}

// 			materialBean.QuestionIds = questionIds
// 			db.Create(materialBean)
// 		}
// 	}
// }

// func readAllFile(db *gorm.DB) {
// 	abc()
// 	for _, i := range gradeArray {
// 		fpFile, _ := ioutil.ReadDir("./res/一起作业网高效阅读数据/" + strconv.Itoa(i) + "/")
// 		for _, file := range fpFile {
// 			GenerateMaterialsForzuoye(db, file.Name(), i)
// 		}
// 	}
// }

// func translateDreamDataQuestionType(oldType int) int {
// 	if oldType == 1 {
// 		return 10001
// 	} else if oldType == 4 {
// 		return 10005
// 	} else if oldType == 5 {
// 		return 10004
// 	} else if oldType == 7 {
// 		return 10009
// 	} else if oldType == 8 {
// 		return 10020
// 	} else if oldType == 9 {
// 		return 10008
// 	} else if oldType == 10 {
// 		return 10005
// 	}
// 	return 10001
// }

// func trimHtml(src string) string {
// 	//将HTML标签全转换成小写
// 	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
// 	src = re.ReplaceAllStringFunc(src, strings.ToLower)
// 	//去除STYLE
// 	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
// 	src = re.ReplaceAllString(src, "")
// 	//去除SCRIPT
// 	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
// 	src = re.ReplaceAllString(src, "")
// 	//去除所有尖括号内的HTML代码，并换成换行符
// 	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
// 	src = re.ReplaceAllString(src, "\n")
// 	//去除连续的换行符
// 	re, _ = regexp.Compile("\\s{2,}")
// 	src = re.ReplaceAllString(src, "\n")

// 	src = strings.Replace(src, " ", "", -1)
// 	src = strings.Replace(src, "\n", "", -1)

// 	return strings.TrimSpace(src)
// }

// func getWordNumberType(count int) int {
// 	if count < 200 {
// 		return 1
// 	} else if count >= 200 && count < 400 {
// 		return 2
// 	} else if count >= 400 && count < 800 {
// 		return 3
// 	} else if count >= 800 && count < 1000 {
// 		return 4
// 	} else if count >= 1000 && count < 1200 {
// 		return 5
// 	} else if count >= 1200 {
// 		return 6
// 	} else {
// 		return -1
// 	}
// }

// func trimUnNeedHTML(src string) string {
// 	styleStr1 := "<p style=\"text-indent: 2em; text-align: center;\"><strong></strong></p>"
// 	styleStr2 := "<p style=\"text-indent: 2em; text-align: center;\"><strong></strong><br></p>"
// 	styleStr3 := "<p style=\"text-indent: 2em; text-align: center;\"></p>"
// 	styleStr4 := "<p style=\"text-indent: 2em; text-align: center;\"><br></p>"
// 	styleStr5 := "<p style=\"text-indent: 0em; text-align: center;\"><strong><br></strong></p>"
// 	styleStr12 := "<p style=\"text-align: center; text-indent: 2em;\"><strong></strong></p>"
// 	styleStr13 := "<p style=\"text-align: center;\"><strong></strong></p>"
// 	reg1 := regexp.MustCompile("<rp>.*?</rp><rt style=\"text-align: center;\"><strong>.*?</strong></rt><rp><strong>.*?</strong></rp>")
// 	reg2 := regexp.MustCompile("<rp>.*?</rp><rt>.*?</rt><rp>.*?</rp>")
// 	reg3 := regexp.MustCompile("<rt>.*?</rt>")
// 	reg4 := regexp.MustCompile("<rp>.*?</rp><rt >.*?</rt><rp>.*?</rp>")
// 	reg5 := regexp.MustCompile("<rt style=\"  text-align: center;\"><strong>.*?</strong></rt>")
// 	reg6 := regexp.MustCompile("<rt >.*?</rt>")
// 	styleStr6 := reg1.FindAllString(src, -1)
// 	styleStr7 := reg2.FindAllString(src, -1)
// 	styleStr8 := reg3.FindAllString(src, -1)
// 	styleStr9 := reg4.FindAllString(src, -1)
// 	styleStr10 := reg5.FindAllString(src, -1)
// 	styleStr11 := reg6.FindAllString(src, -1)
// 	src = strings.Replace(src, styleStr1, "", 1)
// 	src = strings.Replace(src, styleStr2, "", 1)
// 	src = strings.Replace(src, styleStr3, "", 1)
// 	src = strings.Replace(src, styleStr4, "", 1)
// 	src = strings.Replace(src, styleStr5, "", 1)
// 	src = strings.Replace(src, styleStr12, "", 1)
// 	src = strings.Replace(src, styleStr13, "", 1)

// 	for a := range styleStr6 {
// 		src = strings.Replace(src, styleStr6[a], "", 1)
// 	}

// 	for b := range styleStr7 {
// 		src = strings.Replace(src, styleStr7[b], "", 1)
// 	}

// 	for c := range styleStr8 {
// 		src = strings.Replace(src, styleStr8[c], "", 1)
// 	}

// 	for d := range styleStr9 {
// 		src = strings.Replace(src, styleStr9[d], "", 1)
// 	}

// 	for e := range styleStr10 {
// 		src = strings.Replace(src, styleStr10[e], "", 1)
// 	}

// 	for f := range styleStr11 {
// 		src = strings.Replace(src, styleStr11[f], "", 1)
// 	}

// 	return src
// }
