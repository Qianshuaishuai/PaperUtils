package models

type ZsPaperSimple struct {
	Status     int `json:"status"`
	UpdateTime int `json:"updateTime"`
	From       int `json:"from"`
	Name       int `json:"name"`
	Source     int `json:"source"`
	CourseId   int `json:"courseId"`
	CreateTime int `json:"createTime"`
	Grade      int `json:"grade"`
	Score      int `json:"score"`
	Version    int `json:"version"`
	Time       int `json:"time"`
	Addition   int `json:"addition"`
	Type       int `json:"type"`
	ID         int `json:"id"`
	Subject    int `json:"subject"`
}
