package models

import (
	"regexp"
	"strings"
)

func translateProvinceID(name string) int {
	var IDs []int
	GetDb().Table("provinces").Where("name = ?", name).Pluck("id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translatePaperType(name string) int {
	var IDs []int
	GetDb().Table("paper_types").Where("name = ?", name).Pluck("id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translateGrade(name string) int {
	var IDs []int
	GetDb().Table("grades").Where("name = ?", name).Pluck("id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translateKeypoint(name string) int {
	var IDs []int
	GetDb().Table("keypoints").Where("name = ?", name).Pluck("id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translateExerciseType(name string) int {
	var IDs []int
	GetDb().Table("exercise_types").Where("name = ?", name).Pluck("id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translateIsOption(name string) int {
	switch name {
	case "单选题", "双选题", "多选题":
		return 1
	case "填空题":
		return 2
	case "计算题", "推断题", "简答题", "实验题":
		return 3
	default:
		return 0
	}
}

func judgeIsCorrectForOption(index int, answerStr string) bool {
	answerIndexStr := [4]string{"A", "B", "C", "D"}
	answerStrDatas := []rune(answerStr)
	answerLength := len(answerStrDatas)
	for i := 0; i < answerLength; i++ {
		min := i
		max := i + 1
		correctAnswer := string(answerStrDatas[min:max])

		if answerIndexStr[index] == correctAnswer {
			return true
		}

	}

	return false
}

func translateContent(typeName, content string) (newContent string) {
	switch typeName {
	case "单选题", "双选题", "多选题":
		r, _ := regexp.Compile("<div(.+?)><table(.+?)></table></div>")
		filtTxt := r.FindAllString(content, 2)
		newContent = content
		for f := range filtTxt {
			newContent = strings.Replace(newContent, filtTxt[f], "", 1)
		}
		return newContent
	default:
		return content
	}

	return ""
}
