package models

import "time"

type MaterialData struct {
	ErrorCode int         `json:"errorCode"`
	Result    []Knowledge `json:"result"`
}

type MaterialSimple struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	NodeType string           `json:"nodeType"`
	Children []MaterialSimple `json:"children"`
	Parent   bool             `json:"parent"`
}

type OMaterialSimple struct {
	ID               int       `gorm:"column:id"`
	Name             string    `gorm:"column:name"`
	Parent           int       `gorm:"column:parent"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}
