package database

import (
	"fmt"
)

type TimsLists struct {
	Id         int    `json:"id" gorm:"column:id;primaryKey"`
	CategoryId int    `json:"category_id" gorm:"column:category_id;"`
	Sex        string `json:"sex" gorm:"column:sex"`
	LockerId   string `json:"locker_id" gorm:"column:locker_id"`
	TimeRange  string `json:"time_range" gorm:"column:time_range"`
	IsOccupied string `json:"is_occupied" gorm:"column:is_occupied"`
}

func (TimsLists) TableName() string {
	return "timslists"
}

func GetTimsListByCateID() []TimsLists {
	db := GetTimeToFitDBConnection()
	var timsLists []TimsLists
	if err := db.Find(&timsLists).Error; err != nil {
		fmt.Printf("Error retrieving timsLists: %v\n", err)
		return nil
	}
	return timsLists
}
