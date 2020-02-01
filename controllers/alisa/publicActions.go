package alisa

import (
	"alisa-millionaire-lite/server/controllers"
	h "alisa-millionaire-lite/server/helpers"
	"encoding/json"
)

func WaitAuthCommand(reqData ReqData, resData ResData) ResData {
	rc := h.RedisClient()

	var sessionJSON SessionStore
	session, _ := rc.Get(reqData.Session.UserID).Result()
	_ = json.Unmarshal([]byte(session), &sessionJSON)

	if reqData.Request.Command == "Регистрация" {
		sessionJSON.Status = "register"
		sessionJSON.Step = "login"
		resData.Response.Text = "Желаемый логин"
	} else if reqData.Request.Command == "Авторизация" {
		sessionJSON.Status = "login"
		sessionJSON.Step = "login"
		resData.Response.Text = "Ваш логин"
	}

	SetStore(reqData.Session.UserID, sessionJSON)
	return resData
}

func Register(reqData ReqData, resData ResData) ResData {
	rc := h.RedisClient()

	var sessionJSON SessionStore
	session, _ := rc.Get(reqData.Session.UserID).Result()
	_ = json.Unmarshal([]byte(session), &sessionJSON)

	if sessionJSON.Step == "login" {
		if controllers.CheckLogin(reqData.Request.Command) {
			sessionJSON = ReStore(sessionJSON)
			resData.Response.Text = "Такой логин уже существует. Выбирете команду: Регистрация или Авторизация"
		} else {
			sessionJSON.Data = make(map[string]string)
			sessionJSON.Data["login"] = reqData.Request.Command
			sessionJSON.Step = "password"
			resData.Response.Text = "Желаемый пароль"
		}
	} else if sessionJSON.Step == "password" {
		sessionJSON.Data["password"] = reqData.Request.Command
		controllers.SaveNewUser(sessionJSON.Data)
		resData.Response.Text = "Вы авторизованы. Хотите начать игру?"
		sessionJSON = authStore(sessionJSON)
	}

	SetStore(reqData.Session.UserID, sessionJSON)
	return resData
}

func Login(reqData ReqData, resData ResData) ResData {
	rc := h.RedisClient()

	var sessionJSON SessionStore
	session, _ := rc.Get(reqData.Session.UserID).Result()
	_ = json.Unmarshal([]byte(session), &sessionJSON)

	if sessionJSON.Step == "login" {
		if controllers.CheckLogin(reqData.Request.Command) {
			sessionJSON.Data = make(map[string]string)
			sessionJSON.Data["login"] = reqData.Request.Command
			sessionJSON.Step = "password"
			resData.Response.Text = "Ваш пароль"
		} else {
			sessionJSON = ReStore(sessionJSON)
			resData.Response.Text = "Неверный логин. Выбирете команду: Регистрация или Авторизация"
		}
	} else if sessionJSON.Step == "password" {
		sessionJSON.Data["password"] = reqData.Request.Command
		user := controllers.Login(sessionJSON.Data)

		if user.Login == "" {
			sessionJSON = ReStore(sessionJSON)
			resData.Response.Text = "Неверный пароль. Выбирете команду: Регистрация или Авторизация"
		} else {
			resData.Response.Text = "Вы авторизованы. Хотите начать игру?"
			sessionJSON = authStore(sessionJSON)
		}
	}

	SetStore(reqData.Session.UserID, sessionJSON)
	return resData
}

func authStore(sessionJSON SessionStore) SessionStore {
	sessionJSON.Step = ""
	sessionJSON.Status = "waitRunGame"
	sessionJSON.IsAuth = true
	sessionJSON.Data = map[string]string{}
	return sessionJSON
}
