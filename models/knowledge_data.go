package models

type KnowledgeData struct {
	ErrorCode int         `json:"errorCode"`
	Result    []Knowledge `json:"result"`
}

type Knowledge struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	NodeType string      `json:"nodeType"`
	Children []Knowledge `json:"children"`
	Parent   bool        `json:"parent"`
}

type KeyPoint struct {
	KeyPointID int    `gorm:"column:F_keypoint_id"`
	Name       string `gorm:"column:F_name"`
	Type       int    `gorm:"column:F_type"`
}
