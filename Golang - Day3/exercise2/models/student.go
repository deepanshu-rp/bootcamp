package models

import (
	"exercise2/config"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllStudents(students *[]Student) (err error) {
	if err := config.DB.Find(&students).Error; err != nil {
		return err
	}
	return nil
}

func AddNewStudent(student *Student) error {
	if err := config.DB.Create(&student).Error; err != nil {
		return err
	}
	return nil
}

func GetStudent(student *Student, id string) error {
	if err := config.DB.First(&student, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateMarks(id string, marks string) error {
	i, _ := strconv.Atoi(marks)
	config.DB.Model(&Student{}).Where("id = ?", id).Update("marks", i)
	return nil
}

func UpdateStudent(student *Student) error {
	if err := config.DB.Save(student).Error; err != nil {
		return err
	}
	return nil
}

func DeleteStudent(id string) error {
	if err := config.DB.Delete(&Student{}, id).Error; err != nil {
		return err
	}
	return nil
}
