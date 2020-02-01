package controllers

import (
	"alisa-millionaire-lite/server/models"
)

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func SaveNewUser(data map[string]string) models.User {
	user := models.CreateUser(data["login"], data["password"])
	return user
}

func CheckLogin(login string) bool {
	user := models.GetUserFromLogin(login)
	if user.Login == "" {
		return false
	} else {
		return true
	}
}

func Login(data map[string]string) models.User {
	user := models.GetUserAuth(data["login"], data["password"])
	return user
}
