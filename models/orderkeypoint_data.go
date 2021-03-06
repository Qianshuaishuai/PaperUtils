package models

import "time"

type OKeyPoint struct {
	ID               int       `gorm:"column:id"`
	Name             string    `gorm:"column:name"`
	Parent           int       `gorm:"column:parent"`
	SubjectID        int       `gorm:"column:subject_id"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}
