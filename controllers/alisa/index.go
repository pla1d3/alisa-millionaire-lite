package alisa

import (
	h "alisa-millionaire-lite/server/helpers"
	"fmt"
	"io/ioutil"

	"encoding/json"
	"net/http"
)

type ReqData struct {
	Request struct {
		Command string `json:"command"`
	} `json:"request"`
	Session Session `json:"session"`
	Version string  `json:"version"`
}

type ResData struct {
	Response struct {
		Text       string `json:"text"`
		Tts        string `json:"tts"`
		EndSession bool   `json:"end_session"`
	} `json:"response"`
	Session Session `json:"session"`
	Version string  `json:"version"`
}

type Session struct {
	New       bool   `json:"new"`
	MessageID int    `json:"message_id"`
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
}

type SessionStore struct {
	Status string            `json:"status"`
	Step   string            `json:"step"`
	IsAuth bool              `json:"isAuth"`
	Data   map[string]string `json:"data"`
}

func AlisaHook(res http.ResponseWriter, req *http.Request) {
	rc := h.RedisClient()

	var reqData ReqData
	body, _ := ioutil.ReadAll(req.Body)
	_ = json.Unmarshal(body, &reqData)

	var sessionJSON SessionStore
	session, _ := rc.Get(reqData.Session.UserID).Result()
	_ = json.Unmarshal([]byte(session), &sessionJSON)

	var resData ResData
	if sessionJSON.IsAuth {
		resData = privateActions(reqData, resData)
	} else {
		resData = publicActions(reqData, resData)
	}

	resData.Session = reqData.Session
	resData.Version = reqData.Version
	h.WriteJSON(res, resData)
}

func publicActions(reqData ReqData, resData ResData) ResData {
	rc := h.RedisClient()

	var sessionJSON SessionStore
	session, _ := rc.Get(reqData.Session.UserID).Result()
	_ = json.Unmarshal([]byte(session), &sessionJSON)

	if reqData.Session.New {
		resData.Response.Text = "Вы запустили игру Millionaire Lite. Выбирете команду: Регистрация или Авторизация"
		resData.Response.Tts = "Вы запустили игру Millionaire Lite. Выбирете команду: Регистрация или Авторизация"
		sessionJSON.Status = "waitAuthCommand"
		SetStore(reqData.Session.UserID, sessionJSON)
		return resData
	}

	switch sessionJSON.Status {
	case "waitAuthCommand":
		resData = WaitAuthCommand(reqData, resData)
		break
	case "register":
		resData = Register(reqData, resData)
		break
	case "login":
		resData = Login(reqData, resData)
		break
	}

	return resData
}

func privateActions(reqData ReqData, resData ResData) ResData {
	rc := h.RedisClient()

	var sessionJSON SessionStore
	session, _ := rc.Get(reqData.Session.UserID).Result()
	_ = json.Unmarshal([]byte(session), &sessionJSON)

	if reqData.Session.New {
		resData.Response.Text = "Вы авторизованы. Хотите начать игру?"
		resData.Response.Tts = "Вы авторизованы. Хотите начать игру?"
		sessionJSON.Status = "waitRunGame"
		SetStore(reqData.Session.UserID, sessionJSON)
		return resData
	}

	fmt.Println(sessionJSON.Status)

	switch sessionJSON.Status {
	case "waitRunGame":
		resData = WaitRunGame(reqData, resData)
		break
	case "waitAnswer":
		resData = WaitAnswer(reqData, resData)
		break
	}

	return resData
}

func SetStore(userID string, sessionJSON SessionStore) {
	rc := h.RedisClient()
	redisFormat, _ := json.Marshal(sessionJSON)
	rc.Set(userID, redisFormat, 0)
}

func ReStore(sessionJSON SessionStore) SessionStore {
	sessionJSON.Status = "waitAuthCommand"
	sessionJSON.Step = ""
	sessionJSON.Data = map[string]string{}
	return sessionJSON
}

func ReStoreAuth(sessionJSON SessionStore) SessionStore {
	sessionJSON.Status = "waitRunGame"
	sessionJSON.Step = ""
	sessionJSON.Data = map[string]string{}
	return sessionJSON
}
