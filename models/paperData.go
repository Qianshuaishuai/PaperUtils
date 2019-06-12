package models

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var (
	accessKey = os.Getenv("uz6OX00UqM7sCbt5OTZVgDz_mmjUn2XOPDYvjHzp")
	secretKey = os.Getenv("Iap1keCxKnnMj6cxZaTt4O2Ql7EC_qvy5LXF8RhB")
	bucket    = os.Getenv("lbtest")
)

type randItem struct {
	start float64
	end   float64
}

type info struct {
	Start float64
	End   float64
}

func GetRandItem(items map[interface{}]float64) interface{} {
	nums := make(map[interface{}]randItem)

	var s float64
	for k, f := range items {
		if f <= 0 {
			continue
		}

		nums[k] = randItem{
			start: s,
			end:   s + f,
		}

		s = nums[k].end
	}

	rnd := rand.Float64() * s

	for k, f := range nums {
		if f.start <= rnd && f.end > rnd {
			return k
		}
	}

	return nil
}

type PaperQuestion struct {
	Content  string `gorm:"column:F_content;type:TEXT" json:"content"` //题目内容
	Type     int    `gorm:"column:F_type" json:"type"`                 //题目类型
	CourseID int    `gorm:"column:F_course_id" json:"courseId"`        //题目科目id
}

type KeyPointSimple struct {
	KeyPointID float64 `gorm:"primary_key;column:F_keypoint_id" json:"keyPointId"` //考点id
	Name       string  `gorm:"column:F_name" json:"name"`                          //考点名字
	Type       int     `gorm:"column:F_type" json:"type"`                          //考点类型
}

type DataABC struct {
	ID                 int    `gorm:"column:F_poerty_id"`
	GradeID            int    `gorm:"column:F_grade_id"`
	DynastyID          int    `gorm:"column:F_dynasty_id"`
	Name               string `gorm:"column:F_title"`
	Author             string `gorm:"column:F_author"`
	Content            string `gorm:"column:F_content"`
	AuthorIntroduction string `gorm:"column:F_author_introduction"`
	Background         string `gorm:"column:F_background"`
	Translation        string `gorm:"column:F_translation"`
	Comment            string `gorm:"column:F_conmment"`
	Appreciation       string `gorm:"column:F_appreciation"`
	Audio              string `gorm:"column:F_Audio"`
}

var dataFileNames = [3]string{"七年级上", "七年级下", "八年级上"}

func getPaperQuestion(db *gorm.DB) {
	// 设置上传凭证有效期
	putPolicy := storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      "http://api.example.com/qiniu/upload/callback",
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		CallbackBodyType: "application/json",
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)
}

func exampleUploadToken() {
	// 设置上传凭证有效期
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)

	// 覆盖上传凭证
	// 需要覆盖的文件名
	keyToOverwrite := "qiniu.mp4"
	putPolicy = storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
	}
	upToken = putPolicy.UploadToken(mac)
	fmt.Println(upToken)

	// 自定义上传回复凭证
	putPolicy = storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	upToken = putPolicy.UploadToken(mac)
	fmt.Println(upToken)

	// 带回调业务服务器的凭证(JSON方式)
	putPolicy = storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      "http://api.example.com/qiniu/upload/callback",
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		CallbackBodyType: "application/json",
	}
	upToken = putPolicy.UploadToken(mac)
	fmt.Println(upToken)

	// 带回调业务服务器的凭证（URL方式）
	putPolicy = storage.PutPolicy{
		Scope:        bucket,
		CallbackURL:  "http://api.example.com/qiniu/upload/callback",
		CallbackBody: "key=$(key)&hash=$(etag)&bucket=$(bucket)&fsize=$(fsize)&name=$(x:name)",
	}
	upToken = putPolicy.UploadToken(mac)
	fmt.Println(upToken)

	// 带数据处理的凭证
	saveMp4Entry := base64.URLEncoding.EncodeToString([]byte(bucket + ":avthumb_test_target.mp4"))
	saveJpgEntry := base64.URLEncoding.EncodeToString([]byte(bucket + ":vframe_test_target.jpg"))
	//数据处理指令，支持多个指令
	avthumbMp4Fop := "avthumb/mp4|saveas/" + saveMp4Entry
	vframeJpgFop := "vframe/jpg/offset/1|saveas/" + saveJpgEntry
	//连接多个操作指令
	persistentOps := strings.Join([]string{avthumbMp4Fop, vframeJpgFop}, ";")
	pipeline := "test"
	putPolicy = storage.PutPolicy{
		Scope:               bucket,
		PersistentOps:       persistentOps,
		PersistentPipeline:  pipeline,
		PersistentNotifyURL: "http://api.example.com/qiniu/pfop/notify",
	}
	upToken = putPolicy.UploadToken(mac)
	fmt.Println(upToken)
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

func getRandomForDifficulty() float64 {
	rand.Seed(time.Now().UnixNano())

	mItems := map[interface{}]float64{
		info{
			Start: 0,
			End:   2.5,
		}: 0.05,
		info{
			Start: 2.6,
			End:   4.0,
		}: 0.05,
		info{
			Start: 4.1,
			End:   5.1,
		}: 0.1,
		info{
			Start: 5.1,
			End:   7.0,
		}: 0.25,
		info{
			Start: 7.1,
			End:   9.0,
		}: 0.55,
	}

	item := GetRandItem(mItems).(info)

	switch item.Start {
	case 0:
		return 2
	case 2.6:
		return 3
	case 4.1:
		return 4.5
	case 5.1:
		return 6
	case 7.1:
		return 8
	default:
		return 2
	}
}

func isEnglish(courseId int) bool {
	if courseId == 15 || courseId == 32 || courseId == 44 || courseId == 55 || courseId == 72 {
		return true
	}
	return false
}

func readFile(path string) string {

	var buf string

	input, _ := os.Open(path)

	reader := bufio.NewReader(input)

	buff := make([]byte, 1)
	for {
		n, _ := reader.Read(buff)
		if n == 0 {
			break
		}
		buf = buf + string(buff[0:n])
	}
	return buf
}

func trimTex(content string) string {
	reg1 := regexp.MustCompile("\\[tex.*?\\](.|\n|\f|\r)*?\\[/tex\\]")
	styleStr := reg1.FindAllString(content, -1)
	for b := range styleStr {
		content = strings.Replace(content, styleStr[b], "", 1)
	}
	return content
}

func trimImg(content string) string {
	reg1 := regexp.MustCompile("\\[img.*?\\](.|\n|\f|\r)*?\\[/img\\]")
	styleStr := reg1.FindAllString(content, -1)
	for b := range styleStr {
		content = strings.Replace(content, styleStr[b], "", 1)
	}
	return content
}

func trimBrackets(content string) string {
	reg1 := regexp.MustCompile("（.*?）")
	styleStr := reg1.FindAllString(content, -1)
	for b := range styleStr {
		content = strings.Replace(content, styleStr[b], "", 1)
	}
	return content
}

func trimBracket(content string) string {
	reg1 := regexp.MustCompile("\\(.*?\\)")
	styleStr := reg1.FindAllString(content, -1)
	for b := range styleStr {
		content = strings.Replace(content, styleStr[b], "", 1)
	}

	return content
}

func translate(content string) string {
	reg1 := regexp.MustCompile("\\[")
	styleStr := reg1.FindAllString(content, -1)
	for b := range styleStr {
		content = strings.Replace(content, styleStr[b], "<", 1)
	}
	reg2 := regexp.MustCompile("\\]")
	styleStr2 := reg2.FindAllString(content, -1)
	for c := range styleStr2 {
		content = strings.Replace(content, styleStr2[c], ">", 1)
	}
	return content
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func base64toFile(base64String string) {
	ddd, _ := base64.StdEncoding.DecodeString(base64String) //成图片文件并把文件写入到buffer
	err2 := ioutil.WriteFile("./test.pdf", ddd, 0666)       //buffer输出到jpg文件中（不做处理，直接写到文件）
	beego.Debug(err2)
}

func writeQuestion(str string) {
	var filename = "./question.txt"
	var f *os.File
	var err1 error

	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}

	check(err1)
	n, err1 := io.WriteString(f, "\n"+str) //写入文件(字符串)
	check(err1)
	fmt.Printf("写入 %d 个字节n", n)

}

/*
 * 读取目录内的文件
 */
func readDirectory(dir string) (b bool, fl []string) {
	//检查目录是否存在
	if !IsDirExist(dir) {
		return false, nil
	}

	files, _ := ioutil.ReadDir(dir)

	var fileList []string
	fileList = make([]string, len(files))

	i := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fileList[i] = file.Name()
			i++
		}
	}

	ret := false
	if len(fileList) > 0 {
		ret = true
	}

	return ret, fileList
}

func IsDirExist(dir string) bool {
	fi, err := os.Stat(dir)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}
