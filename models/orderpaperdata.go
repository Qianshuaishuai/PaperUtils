package models

import "time"

type OrderPaper struct {
	ID               int       `gorm:"column:id"`
	Name             string    `gorm:"column:name"`
	Score            float64   `gorm:"column:score"`
	PaperType        int       `gorm:"column:paper_type_id"`
	GradeID          int       `gorm:"column:grade_id"`
	SubjectID        int       `gorm:"column:subject_id"`
	OldID            string    `gorm:"column:old_id"`
	RealDate         time.Time `gorm:"column:real_date"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}

type OrderPaperProvince struct {
	ID               int       `gorm:"column:id"`
	PaperID          int       `gorm:"column:paper_id"`
	ProvinceID       int       `gorm:"column:province_id"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}

type OrderKeyPointQuestion struct {
	ID               int       `gorm:"column:id"`
	ExerciseID       int       `gorm:"column:exercise_id"`
	KeypointID       int       `gorm:"column:keypoint_id"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}

type OrderPaperQuestion struct {
	ID               int       `gorm:"column:id"`
	ExerciseID       int       `gorm:"column:exercise_id"`
	PaperID          int       `gorm:"column:paper_id"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}

type OrderExercise struct {
	ID               int       `gorm:"column:id"`
	Content          string    `gorm:"column:content"`
	Analysis         string    `gorm:"column:analysis"`
	SubjectID        int       `gorm:"column:subject_id"`
	ExerciseType     int       `gorm:"column:exercise_type"`
	Score            float64   `gorm:"column:score"`
	Difficulty       int       `gorm:"column:difficulty_level"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}

type OrderQuestion struct {
	ID               int       `gorm:"column:id"`
	ExerciseID       int       `gorm:"column:exercise_id"`
	Question         string    `gorm:"column:question"`
	QuestionIndex    int       `gorm:"column:question_index"`
	QuestionType     int       `gorm:"column:question_type"`
	QuestionScore    float64   `gorm:"column:question_score"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}

type OrderAnswer struct {
	ID               int       `gorm:"column:id"`
	ExerciseID       int       `gorm:"column:exercise_id"`
	QuestionID       int       `gorm:"column:question_id"`
	IsCorrect        int       `gorm:"column:is_correct"`
	Answer           string    `gorm:"column:answer"`
	Analysis         string    `gorm:"column:analysis"`
	CreationDate     time.Time `gorm:"column:creation_date"`
	ModificationDate time.Time `gorm:"column:modification_date"`
}
