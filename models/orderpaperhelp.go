package models

import (
	"regexp"
	"strings"

	"github.com/astaxie/beego"
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
	answerIndexStr := [11]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
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

func isContainSmall(content string) bool {
	r, _ := regexp.Compile("（[0-9]）(.+?)（")
	filtTxt := r.FindAllString(content, 15)

	if len(filtTxt) > 0 {
		return true
	} else {
		ar, _ := regexp.Compile(`\([0-9]\)`)
		afiltTxt := ar.FindAllString(content, 15)

		if len(afiltTxt) > 0 {
			return true
		}
	}

	return false
}

func getSmallQuestionForNotOption(content string) (newContent string, questionData []string) {
	r, _ := regexp.Compile("（[0-9]）")
	filtTxt := r.FindAllString(content, 15)

	if len(filtTxt) > 0 {
		newQuestionData := make([]string, 0)
		newContent = content
		for f := range filtTxt {
			if f < len(filtTxt)-1 {
				startStr := filtTxt[f]
				endStr := filtTxt[f+1]

				qr, _ := regexp.Compile(startStr + "(.+?)" + endStr)
				fr := qr.FindAllString(content, 2)

				if len(fr) <= 0 {
					beego.Debug(filtTxt)
					beego.Debug(startStr + "(.+?)" + endStr)
					beego.Debug(fr)
					beego.Debug(content)
					beego.Debug("小题提干切割出错001")
					continue
				}

				cacheStr := []rune(fr[0])
				newQuestion := string(cacheStr[0 : len(cacheStr)-3])
				newQuestionData = append(newQuestionData, newQuestion)
			} else {
				startStr := strings.Index(content, filtTxt[f])

				if startStr < 0 {
					beego.Debug(filtTxt)
					beego.Debug(content)
					beego.Debug("小题提干切割出错002")
					continue
				}
				prefix := []byte(content)[0:startStr]
				// 将子串之前的字符串转换成[]rune
				rs := []rune(string(prefix))
				// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
				startStr = len(rs)

				cacheStr := []rune(content)

				newQuestion := string(cacheStr[startStr:len(cacheStr)])

				newQuestionData = append(newQuestionData, newQuestion)
			}
		}

		for n := range newQuestionData {
			newContent = strings.Replace(newContent, newQuestionData[n], "", 1)
		}

		return newContent, newQuestionData
	} else {
		ar, _ := regexp.Compile(`\([0-9]\)`)
		afiltTxt := ar.FindAllString(content, 15)

		filtTxt = afiltTxt

		newQuestionData := make([]string, 0)
		newContent = content
		for f := range filtTxt {
			if f < len(filtTxt)-1 {
				newReg := ""

				switch f {
				case 0:
					newReg = `\(1\)(.+?)\(2\)`
					break
				case 1:
					newReg = `\(2\)(.+?)\(3\)`
					break
				case 2:
					newReg = `\(3\)(.+?)\(4\)`
					break
				case 3:
					newReg = `\(4\)(.+?)\(5\)`
					break
				case 4:
					newReg = `\(5\)(.+?)\(6\)`
					break
				case 5:
					newReg = `\(6\)(.+?)\(7\)`
					break
				case 6:
					newReg = `\(7\)(.+?)\(8\)`
					break
				case 7:
					newReg = `\(8\)(.+?)\(9\)`
					break
				case 8:
					newReg = `\(9\)(.+?)\(10\)`
					break
				}

				qr, _ := regexp.Compile(newReg)
				fr := qr.FindAllString(content, 2)

				if len(fr) <= 0 {
					beego.Debug("小题提干切割出错001")
					continue
				}

				cacheStr := []rune(fr[0])
				newQuestion := string(cacheStr[0 : len(cacheStr)-3])
				newQuestionData = append(newQuestionData, newQuestion)
			} else {
				startStr := strings.Index(content, filtTxt[f])

				if startStr < 0 {
					beego.Debug("小题提干切割出错002")
					continue
				}
				prefix := []byte(content)[0:startStr]
				// 将子串之前的字符串转换成[]rune
				rs := []rune(string(prefix))
				// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
				startStr = len(rs)

				cacheStr := []rune(content)

				newQuestion := string(cacheStr[startStr:len(cacheStr)])

				newQuestionData = append(newQuestionData, newQuestion)
			}
		}

		for n := range newQuestionData {
			newContent = strings.Replace(newContent, newQuestionData[n], "", 1)
		}

		return newContent, newQuestionData
	}

	return "", nil
}

func getSmallAnswerForNotOption(data string) (answerData []string) {
	if strings.Contains(data, "(1)") {
		data = strings.Replace(data, "(1)", "（1）", 5)
	}

	if strings.Contains(data, "(2)") {
		data = strings.Replace(data, "(2)", "（2）", 5)
	}

	if strings.Contains(data, "(3)") {
		data = strings.Replace(data, "(3)", "（3）", 5)
	}

	if strings.Contains(data, "(4)") {
		data = strings.Replace(data, "(4)", "（4）", 5)
	}

	if strings.Contains(data, "(5)") {
		data = strings.Replace(data, "(5)", "（5）", 5)
	}

	if strings.Contains(data, "(6)") {
		data = strings.Replace(data, "(6)", "（6）", 5)
	}

	if strings.Contains(data, "(7)") {
		data = strings.Replace(data, "(7)", "（7）", 5)
	}

	if strings.Contains(data, "(8)") {
		data = strings.Replace(data, "(8)", "（8）", 5)
	}

	if strings.Contains(data, "(9)") {
		data = strings.Replace(data, "(9)", "（9）", 5)
	}

	r, _ := regexp.Compile("<p>（[0-9]）(.+?)</p>")
	filtTxt := r.FindAllString(data, 15)

	if len(filtTxt) > 0 {
		return filtTxt
	}

	r, _ = regexp.Compile("<p(.+?)>（[0-9]）(.+?)</p>")
	filtTxt = r.FindAllString(data, 15)

	if len(filtTxt) > 0 {
		return filtTxt
	}

	ar, _ := regexp.Compile("<li(.+?)>（[0-9]）(.+?)</li>")
	afiltTxt := ar.FindAllString(data, 15)

	if len(afiltTxt) > 0 {
		return afiltTxt
	}

	beego.Debug(data)
	beego.Debug(filtTxt)
	beego.Debug("获取答案有异样")

	return nil
}
