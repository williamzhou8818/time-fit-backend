package database

import (
	"fmt"

	"github.com/williamzhou8818/time-fit-backend/util"
)

type Category struct {
	Id    int    `json:"id" gorm:"column:id;primaryKey"`
	Title string `json:"title" gorm:"column:title"`
	Color string `json:"color" gorm:"column:color"`
}

func (Category) TableName() string {
	return "category"
}

var (
	_all_cate_field = util.GetGormFields(Category{})
)

// 读取所有category表格的数据
func GetAllCategory() []Category {
	db := GetTimeToFitDBConnection()
	var category []Category
	if err := db.Find(&category).Error; err != nil {
		fmt.Printf("Error retrieving categories: %v\n", err)
		return nil
	}
	return category
}
