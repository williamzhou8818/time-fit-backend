package test

import (
	"fmt"
	"testing"

	"github.com/williamzhou8818/time-fit-backend/database"
)

func TestUser(t *testing.T) {
	user := database.GetAllCategory()
	if user == nil {
		t.Fail()
		return
	}

	fmt.Println(user)
	// if user.PassWd != "e10adc3949ba59abbe56e057f20f883e" {
	// 	fmt.Println(user.PassWd)
	// 	t.Fail()
	// 	return
	// }
	// user = database.GetUserByName("none")
	// if user != nil {
	// 	t.Fail()
	// 	return
	// }
}

// func TestCreateUser(t *testing.T) {
// 	name, pass := "大乔乔", "345678"
// 	database.CreateUser(name, pass)
// }

// func TestDeleteUser(t *testing.T) {
// 	name := "大乔乔"
// 	database.DeleteUser(name)
// }

// go test -v .\database\test\ -run=^TestUser$ -count=1
// go test -v .\database\test\ -run=^TestCreateUser$ -count=1
// go test -v .\database\test\ -run=^TestDeleteUser$ -count=1
