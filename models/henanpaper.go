package models

import (
	"dreamEbagPapers/helper"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	BASE_URL    = "./内蒙古/英语/"
	OFFSET_TIME = 1531395763200
)

type PaperSimple struct {
	PaperID        int    `gorm:"column:F_paper_id"`
	Name           string `gorm:"column:F_name"`
	ShortName      string `gorm:"column:F_short_name"`
	PaperType      int    `gorm:"column:F_paper_type"`
	Difficulty     int    `gorm:"column:F_difficulty"`
	FullScore      int    `gorm:"column:F_full_score"`
	Date           string `gorm:"column:F_date"`
	Semester       int    `gorm:"column:F_semester_id"`
	CourseID       int    `gorm:"column:F_course_id"`
	ReferenceCount int    `gorm:"column:F_reference_count"`
	BrowseCount    int    `gorm:"column:F_browse_count"`
	ResourceType   int    `gorm:"column:F_resource_type"`
}

type PaperSet struct {
	SetID       int    `gorm:"column:F_set_id"`
	PaperID     int    `gorm:"column:F_paper_id"`
	Name        string `gorm:"column:F_paper_name"`
	QuestionIDs string `gorm:"column:F_question_ids"`
}

type PaperSetChapter struct {
	Name          string  `gorm:"column:F_name"`
	Detail        string  `gorm:"column:F_detail"`
	QuestionCount int     `gorm:"column:F_question_count"`
	Time          int     `gorm:"column:F_time"`
	PresetScore   float64 `gorm:"column:F_preset_score"`
	SetID         int     `gorm:"column:F_set_id"`
	ChapterID     int     `gorm:"column:F_chapter_id"`
}

type QuestionSimple struct {
	QuestionID          int    `gorm:"column:F_question_id"`
	CourseID            int    `gorm:"column:F_course_id"`
	Content             string `gorm:"column:F_content"`
	Accessories         string `gorm:"column:F_accessories"`
	Solution            string `gorm:"column:F_solution"`
	SolutionAccessories string `gorm:"column:F_solution_accessories"`

	Difficulty int     `gorm:"column:F_difficulty;" json:"difficulty"`
	Source     string  `gorm:"column:F_source;size:80" json:"source"`   //问题的来源 试卷名称
	Score      float64 `gorm:"column:F_score;type:FLOAT(4,1)" json:"-"` //问题分数 最大值999.9
	ScoreInt   int     `gorm:"-" json:"score"`                          //这个问题的总分

	ShortSource   string `gorm:"column:F_short_source"`
	CorrectAnswer string `gorm:"column:F_correct_answer"`

	Type         int `gorm:"column:F_type"`
	SetID        int `gorm:"column:F_set_id"`
	ResourceType int `gorm:"column:F_resource_type"`
}

type LargeQuestionSimple struct {
	BigQuestionID int     `gorm:"column:F_big_question_id"`
	Content       string  `gorm:"column:F_content"`
	Accessories   string  `gorm:"column:F_accessories"`
	QuestionIDs   string  `gorm:"column:F_question_ids"`
	Score         float64 `gorm:"column:F_score"`
}

type PaperProvince struct {
	PaperID    int `gorm:"column:paper_F_paper_id"`
	ProvinceID int `gorm:"column:province_F_province_id"`
}

type NewOption struct {
	Option OptionSimple `json:"option"`
}

type OptionSimple struct {
	Options    []string `json:"options"`
	OptionType int      `json:"optionType"`
}

type KeyPointQuestion struct {
	KeyPointID int `gorm:"column:keypoint_F_keypoint_id"`
	QuestionID int `gorm:"column:question_F_question_id"`
}

var Grades = []string{"七年级/", "八年级/", "九年级/"}
var Sems = []int{12, 10, 8}
var sem = 12
var grade = "七年级/"

func writeSql() {
	abc()
	// rd, _ := ioutil.ReadDir("./河南试卷/英语/八年级/")

	// for r := range rd {
	// 	writePaperSql(rd[r].Name())
	// }
	// writePaperSql("2017-2018学年河南省周口市八年级（上）期末英语试卷.txt")

	for g := range Grades {
		sem = Sems[g]
		grade = Grades[g]
		rd, _ := ioutil.ReadDir("./内蒙古/英语/" + Grades[g])
		for r := range rd {
			writePaperSql(rd[r].Name())
		}
	}
}

func updateDateTime(name string) {
	var test PaperJsonSimple

	paperBytes, _ := ioutil.ReadFile(BASE_URL + grade + name)

	err := json.Unmarshal(paperBytes, &test)
	if err != nil {
		beego.Debug(err)
		return
	}

	rune_str := []rune(name)
	nameLength := len(rune_str)

	fName := string(rune_str[0 : nameLength-4])

	dateTimeStr := getDateTimeForInt(test.Paper.DateTime / 1000)

	dbErr := GetDb().Table("t_papers").Where("F_name = ?", fName).Update("F_date", dateTimeStr)
	beego.Debug(dbErr)
}

func writePaperSql(name string) {
	var chapterList []PaperAmountDistribution
	var chapterCodeList []string
	var paper PaperSimple

	allCount := 0
	allDifficulty := 0
	beego.Debug(BASE_URL + grade + name)
	paperBytes, _ := ioutil.ReadFile(BASE_URL + grade + name)

	// var paperObj interface{}
	var test PaperJsonSimple
	var totalScore float64
	// json.Unmarshal(paperBytes, &paperObj)
	err := json.Unmarshal(paperBytes, &test)
	if err != nil {
		beego.Debug(err)
		return
	}

	paper.CourseID = 55

	totalScore = test.AnalyseData.TotalScore

	Title := test.Paper.Title

	paperTypeName := test.Paper.PaperType.Name

	// year := test.Paper.Year
	dateTime := test.Paper.DateTime

	id := (int)(flakCurl.GetIntId())

	//省份录入
	var paperProvince PaperProvince
	paperProvince.PaperID = id
	paperProvince.ProvinceID = 1319
	GetDb().Table("t_paper_province").Create(&paperProvince)

	//录入题目相关
	questionList := test.Pager.List
	var questionIDs []int
	if paper.CourseID == 55 {
		questionList = rebuildEnglishQuestionList(questionList)
	}

	for m := range questionList {
		if len(questionList[m].SubQuestion) > 1 {
			bigQusetionID, bigQuestionScore, questionCount, difficulty := newBigQuestion(questionList[m])
			questionIDs = append(questionIDs, bigQusetionID)
			allCount = allCount + questionCount
			allDifficulty = allDifficulty + difficultyTranslate(difficulty)
			chapterCode := questionList[m].Section.Code
			if checkInList(chapterCode, chapterCodeList) {
				count := len(chapterCodeList)
				if count > 0 {
					chapterList[count-1].QuestionCount = chapterList[count-1].QuestionCount + questionCount
					chapterList[count-1].Score = chapterList[count-1].Score + bigQuestionScore
				}
			} else {
				var chapter PaperAmountDistribution
				chapter.Name = questionList[m].Section.Name
				chapter.QuestionCount = questionCount
				chapter.Score = bigQuestionScore
				chapterList = append(chapterList, chapter)
				chapterCodeList = append(chapterCodeList, chapterCode)
			}

		} else {
			allCount = allCount + 1
			var question QuestionSimple
			chapterCode := questionList[m].Section.Code
			if checkInList(chapterCode, chapterCodeList) {
				count := len(chapterCodeList)
				if count > 0 {
					chapterList[count-1].QuestionCount = chapterList[count-1].QuestionCount + 1
					chapterList[count-1].Score = chapterList[count-1].Score + questionList[m].SubQuestion[0].Score
				}
			} else {
				var chapter PaperAmountDistribution
				chapter.Name = questionList[m].Section.Name
				chapter.QuestionCount = 1
				chapter.Score = questionList[m].SubQuestion[0].Score
				chapterList = append(chapterList, chapter)
				chapterCodeList = append(chapterCodeList, chapterCode)
			}
			questionID := int(flakCurl.GetIntId())

			question.QuestionID = questionID
			question.CourseID = 55 //记得改
			question.Content = dealTexContent(questionList[m].SubQuestion[0].Stem)
			question.Accessories = getAccessories(questionList[m].SubQuestion[0].Options)
			question.Solution = ""
			question.SolutionAccessories = "[]"
			question.Difficulty = difficultyTranslate(questionList[m].Difficulty.Value)
			question.Source = ""
			question.Score = questionList[m].Score
			question.ShortSource = ""
			question.CorrectAnswer = dealTexContent(questionList[m].SubQuestion[0].Answers[0].Desc)
			question.Type = getType(questionList[m].Section.CategoryName)
			question.SetID = 0
			question.ResourceType = 101

			if question.Type == 1 || question.Type == 2 || question.Type == 3 {
				question.Content = dealWithOptionContent(question.Content)
			} else if question.Type == 61 || question.Type == 51 {
				question.Content = dealWithInputContent(question.Content)
			}

			if question.Type == 61 && question.CourseID == 54 {
				if strings.Contains(question.CorrectAnswer, "tex") {
					question.Type = 50
				}
			}

			allDifficulty = allDifficulty + difficultyTranslate(question.Difficulty)

			questionIDs = append(questionIDs, questionID)
			GetDb().Table("t_questions").Create(&question)

			//knowLedge
			for k := range questionList[m].Knowledges {

				var keyPointIDs []int
				keyPointID := 0
				GetDb().Table("t_keypoints").Where("F_name = ?", questionList[m].Knowledges[k].Name).Pluck("F_keypoint_id", &keyPointIDs)

				for k := range keyPointIDs {
					if keyPointIDs[k] > 2000000 {
						keyPointID = keyPointIDs[k]
					}
				}

				if keyPointID != 0 {
					var keyPointQuestion KeyPointQuestion
					keyPointQuestion.QuestionID = questionID
					keyPointQuestion.KeyPointID = keyPointID
					GetDb().Table("t_keypoint_question").Create(&keyPointQuestion)
				}
			}
		}

	}

	setDatas := test.AnalyseData.PaperAmountDistribution

	var paperSet PaperSet
	setID := int(flakCurl.GetIntId())
	paperSet.Name = Title
	paperSet.PaperID = id
	paperSet.SetID = setID

	paperSet.QuestionIDs = helper.TransformIntArrToString(questionIDs)

	GetDb().Table("t_paper_question_sets").Create(&paperSet)
	if len(setDatas) <= 0 || paper.CourseID == 55 {
		for c := range chapterList {
			chapterID := int(flakCurl.GetIntId())
			var paperSetChapters PaperSetChapter
			paperSetChapters.ChapterID = chapterID
			paperSetChapters.Detail = ""
			paperSetChapters.Name = chapterList[c].Name
			paperSetChapters.PresetScore = chapterList[c].Score
			paperSetChapters.QuestionCount = chapterList[c].QuestionCount
			paperSetChapters.SetID = setID

			GetDb().Table("t_paper_question_set_chapters").Create(&paperSetChapters)
		}
	} else {
		for s := range setDatas {
			questionCount := setDatas[s].QuestionCount
			if setDatas[s].Name == "阅读理解" {
				questionCount = questionCount / 5
			} else if setDatas[s].Name == "完形填空" {
				questionCount = questionCount / 10
			}
			chapterID := int(flakCurl.GetIntId())
			var paperSetChapters PaperSetChapter
			paperSetChapters.ChapterID = chapterID
			paperSetChapters.Detail = ""
			paperSetChapters.Name = setDatas[s].Name
			paperSetChapters.PresetScore = setDatas[s].Score
			paperSetChapters.QuestionCount = questionCount
			paperSetChapters.SetID = setID

			GetDb().Table("t_paper_question_set_chapters").Create(&paperSetChapters)
		}
	}

	paper.Name = Title
	paper.PaperID = id
	paper.ShortName = ""
	paper.PaperType = getPaperType(paperTypeName)
	paper.Difficulty = (allDifficulty / allCount)
	paper.FullScore = int(totalScore)
	paper.Date = getDateTimeForInt(dateTime / 1000)
	paper.Semester = sem
	paper.ResourceType = 101

	if strings.Contains(paper.Name, "中考") && !strings.Contains(paper.Name, "期中考试") {
		paper.Semester = 6
		paper.CourseID = 72
		beego.Debug(paper.Name)
	}

	beego.Debug(id)

	GetDb().Table("t_papers").Create(&paper)

}

func rebuildEnglishQuestionList(questionList []Question) (newQuestionList []Question) {
	for q := range questionList {
		for i := 1; i < len(questionList)-q; i++ {
			iNumber, _ := strconv.Atoi(questionList[i].Number)
			bNumber, _ := strconv.Atoi(questionList[i-1].Number)
			if iNumber < bNumber {
				questionList[i], questionList[i-1] = questionList[i-1], questionList[i]
			}
		}
	}

	return questionList
}

func checkInList(code string, codeList []string) bool {
	for c := range codeList {
		if code == codeList[c] {
			return true
		}
	}
	return false
}

func newBigQuestion(data Question) (int, float64, int, int) {
	var largeQuestionScore float64
	var largeQuestionSimple LargeQuestionSimple
	allDifficulty := 0
	bigQuestionID := int(flakCurl.GetIntId())

	smallQuestionList := data.SubQuestion
	var smallQuestionIDs []int
	for s := range smallQuestionList {
		var question QuestionSimple
		questionID := int(flakCurl.GetIntId())

		question.QuestionID = questionID
		question.CourseID = 55
		if data.Section.Name == "阅读理解" || data.Section.Name == "信息匹配" || data.Section.Name == "任务型阅读" {
			// index := s + 1
			question.Content = "<p style='padding:0 30px'>" + dealTexContent(smallQuestionList[s].Stem) + "</p>"
		} else if data.Section.Name == "选词填空" || data.Section.Name == "完成句子" || data.Section.Name == "选词填空-短文" || data.Section.Name == "补全对话" {
			// index := s + 1
			question.Content = "<p style='padding:0 30px'>" + dealTexContent(smallQuestionList[s].Stem) + "</p>"
		} else if data.Section.Name == "完形填空" {
			// index := s + 1
			question.Content = "<p style='padding:0 30px'>" + dealTexContent(smallQuestionList[s].Stem) + "</p>"
		} else {
			question.Content = dealTexContent(smallQuestionList[s].Stem)
		}
		question.Accessories = getAccessories(smallQuestionList[s].Options)
		question.Solution = ""
		question.SolutionAccessories = "[]"
		question.Difficulty = difficultyTranslate(data.Difficulty.Value)
		question.Source = ""
		question.Score = smallQuestionList[s].Score
		question.ShortSource = ""
		question.CorrectAnswer = smallQuestionList[s].Answers[0].Desc
		question.Type = getType(data.Section.CategoryName)
		question.SetID = 0
		question.ResourceType = 101

		largeQuestionScore = largeQuestionScore + question.Score

		if question.Type == 1 || question.Type == 2 || question.Type == 3 {
			question.Content = dealWithOptionContent(question.Content)
		} else if question.Type == 61 || question.Type == 51 {
			question.Content = dealWithInputContent(question.Content)
		}

		if question.Type == 61 && question.CourseID == 54 {
			if strings.Contains(question.CorrectAnswer, "tex") {
				question.Type = 50
			}
		}

		allDifficulty = allDifficulty + difficultyTranslate(question.Difficulty)
		smallQuestionIDs = append(smallQuestionIDs, questionID)

		//knowLedge
		for k := range data.Knowledges {
			var keyPointIDs []int
			keyPointID := 0
			GetDb().Table("t_keypoints").Where("F_name = ?", data.Knowledges[k].Name).Pluck("F_keypoint_id", &keyPointIDs)

			for k := range keyPointIDs {
				if keyPointIDs[k] > 2000000 {
					keyPointID = keyPointIDs[k]
				}
			}

			if keyPointID != 0 {
				var keyPointQuestion KeyPointQuestion
				keyPointQuestion.QuestionID = questionID
				keyPointQuestion.KeyPointID = keyPointID
				GetDb().Table("t_keypoint_question").Create(&keyPointQuestion)
			}
		}

		GetDb().Table("t_questions").Create(&question)
	}

	largeQuestionSimple.BigQuestionID = bigQuestionID
	if len(data.Materials) == 0 {
		largeQuestionSimple.Content = ""
	} else {
		largeQuestionSimple.Content = data.Materials[0].Html
	}
	largeQuestionSimple.Accessories = "[]"
	largeQuestionSimple.Score = data.Score
	largeQuestionSimple.QuestionIDs = helper.TransformIntArrToString(smallQuestionIDs)

	GetDb().Table("t_large_questions").Create(&largeQuestionSimple)

	return bigQuestionID, largeQuestionScore, len(smallQuestionList), allDifficulty
}

func dealWithInputContent(content string) string {
	content = strings.Replace(content, "______", "[input=type:blank,size:6][/input]", -1)
	return content
}

func dealTexContent(content string) string {
	content = strings.Replace(content, "\\(", "[tex]$$", -1)
	content = strings.Replace(content, "\\)", "$$%[/tex]", -1)
	return content
}

func dealWithOptionContent(content string) string {
	var newContent string
	startACount := 0
	startBCount := strings.Index(content, "<table class=")
	// endACount := strings.Index(content, "</table>")
	// endBCount := len([]rune(content))
	newContent = content
	if startACount < startBCount {
		startTxt := string(content[startACount:startBCount])
		newContent = startTxt + "</div>"
	}

	// endTxt := string(content[endACount:endBCount])
	return newContent
}

func getQuestionIDs(count int, questionIDs []int) (newQuestionIDs []int, newAllQuestionIDs []int) {
	var c int
	for c = 0; c < count; c++ {
		newQuestionIDs = append(newQuestionIDs, questionIDs[c])
	}

	for c = 0; c < count; c++ {
		a := 0
		questionIDs = append(questionIDs[:a], questionIDs[a+1:]...)
	}

	return newQuestionIDs, questionIDs
}

func getType(name string) int {
	if name == "单选题" {
		return 1
	} else if name == "多选题" {
		return 2
	} else if name == "不定项选择题" {
		return 3
	} else if name == "完形填空" {
		return 4
	} else if name == "判断题" {
		return 5
	} else if name == "阅读理解七选五" {
		return 6
	} else if name == "问答题" {
		return 13
	} else if name == "英语材料回答" {
		return 14
	} else if name == "材料回答" {
		return 15
	} else if name == "改错题" {
		return 16
	} else if name == "填空题" {
		return 61
	} else if name == "主观题" {
		return 50
	}
	return 0
}

func getAccessories(data []Option) string {
	var newOption NewOption
	for d := range data {
		newOption.Option.Options = append(newOption.Option.Options, dealTexContent(data[d].Desc))
	}

	newOption.Option.OptionType = 101

	paramJson, _ := json.Marshal(newOption)

	return string(paramJson)
}

func getPaperType(name string) int {
	if name == "真题" {
		return 1
	} else if name == "模拟题" {
		return 2
	} else if name == "同步练习" {
		return 3
	} else if name == "单元测验" {
		return 4
	} else if name == "月考试卷" {
		return 5
	} else if name == "期中考试" {
		return 6
	} else if name == "期末考试" {
		return 7
	} else if name == "竞赛" {
		return 8
	} else if name == "专项训练" {
		return 9
	}

	return 0
}

func getDateTime(timeStr string) string {
	// t := time.Now()
	// y, m, d := t.Date()
	timeStr = timeStr + "-01-01 00:00:00"
	t, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	datetime := t.Format("2006-01-02 15:04:05")

	return datetime
}

func getDateTimeForInt(dateTime int64) string {
	timeLayout := "2006-01-02 15:04:05"
	dataTimeStr := time.Unix(dateTime, 0).Format(timeLayout)
	return dataTimeStr
}

func getDateTimeFormatForInt(dateTime int64) time.Time {
	dataTime := time.Unix(dateTime, 0)
	return dataTime
}

func difficultyTranslate(difficulty int) int {
	if difficulty == 3 {
		return 5
	} else if difficulty == 4 || difficulty == 5 {
		return 2
	} else if difficulty == 1 || difficulty == 2 {
		return 6
	}
	return 5
}
