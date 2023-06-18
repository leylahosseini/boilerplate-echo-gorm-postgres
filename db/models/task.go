package models

type Task struct {
	//gorm.Model
	ID          uint   `json:"primary_key"`
	Name        string `json:"name";"not null"`
	Description string `json:"description","not null"`
	Completed   bool   `json:"completed","not null"`
}
