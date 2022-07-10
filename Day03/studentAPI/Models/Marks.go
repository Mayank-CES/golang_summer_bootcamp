package Models

import (
	"fmt"
	"studentAPI/Config"

	_ "github.com/go-sql-driver/mysql"
)


//GetMarksByID ... Fetch only one user by Id
func GetMarksByID(marks *Marks, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(marks).Error; err != nil {
		return err
	}
	return nil
}

//UpdateMarks ... Update marks
func UpdateMarks(marks *Marks, id string) (err error) {
	fmt.Println(marks)
	Config.DB.Save(marks)
	return nil
}


