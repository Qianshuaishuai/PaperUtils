package models

import (
	"strconv"

	"github.com/tidwall/gjson"
)

func GetZsPrimarySchoolPaper() {
	papersUrl := "http://eschoolbag.readboy.com:8000/api/papers?from=2&year=2018&sn=ebagtest&page=1&count=100"

	_, data := queryApiForZs(papersUrl, "papers")

	respData := data.Get("paper")

	respData.ForEach(func(key, value gjson.Result) bool {
		ID := value.Get("id").Int()
		Name := value.Get("name").String()
		GetZsPaperDetail(ID, Name)
		return true
	})
}

func GetZsPaperDetail(id int64, name string) {
	idStr := strconv.FormatInt(id, 10)
	paperDetailUrl := "http://eschoolbag.readboy.com:8000/api/paper/" + idStr + "?sn=ebagtest"

	_, data := queryApiForZs(paperDetailUrl, "paper-detail")

	TranslatePaperData(data)
}

func TranslatePaperData(data *gjson.Result) {
	var paperSimple PaperSimple
	paperSimple.CourseID = translateZsCourseID(data.Get("courseId").Int())
	paperSimple.Name = data.Get("name").String()

	paperSimple.PaperID = (int)(flakCurl.GetIntId())
	paperSimple.ShortName = ""
	paperSimple.PaperType = translateZsPaperType(data.Get("type").Int())
	paperSimple.Difficulty = 0
	paperSimple.FullScore = int(data.Get("score").Int() / 10)
	paperSimple.Date = getDateTimeForInt(data.Get("updateTime").Int())
	paperSimple.Semester = translateZsPaperSem(data.Get("grade").Int())
	paperSimple.ResourceType = 555
}

func translateZsCourseID(courseID int64) int {
	switch courseID {
	case 66:
		return 31
	case 65:
		return 30
	case 67:
		return 32
	}
	return 0
}

func translateZsPaperType(typeID int64) int {
	switch typeID {
	case 101:
		return 4
	case 102:
		return 6
	case 103:
		return 7
	}
	return 0
}

func translateZsPaperSem(gradeID int64) int {
	switch gradeID {
	case 1:
		return 24
	case 2:
		return 22
	case 3:
		return 20
	case 4:
		return 18
	case 5:
		return 16
	case 6:
		return 14
	}
	return 0
}
