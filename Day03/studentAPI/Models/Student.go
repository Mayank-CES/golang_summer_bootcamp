package Models

import (
	"fmt"
	"studentAPI/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllStudents Fetch all student data
func GetAllStudents(students *[]Student) (err error) {
	if err = Config.DB.Find(students).Error; err != nil {
		return err
	}
	return nil
}

//CreateStudent ... Insert New data
func CreateStudent(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

//GetStudentByID ... Fetch only one student by Id
func GetStudentByID(student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

// UpdateInfo ... Update student
func UpdateInfo(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

//DeleteStudent ... Delete student
func DeleteStudent(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}
