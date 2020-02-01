package alisa

import (
	"alisa-millionaire-lite/server/controllers"
	h "alisa-millionaire-lite/server/helpers"
	"alisa-millionaire-lite/server/models"
	"encoding/json"
	"strconv"
	"strings"
)

func WaitRunGame(reqData ReqData, resData ResData) ResData {
	rc := h.RedisClient()

	var sessionJSON SessionStore
	session, _ := rc.Get(reqData.Session.UserID).Result()
	_ = json.Unmarshal([]byte(session), &sessionJSON)

	if reqData.Request.Command == "Да" {
		sessionJSON.Data["GameID"] = h.GenerateRandomString(16)
		sessionJSON.Data["countAnswer"] = "0"

		task := controllers.GetRandomTask(sessionJSON.Data["GameID"])
		sessionJSON.Data["TaskID"] = strconv.Itoa(task.ID)
		sessionJSON.Status = "waitAnswer"
		resData.Response.Text = prepareTask(task, sessionJSON.Data["countAnswer"])
	}

	SetStore(reqData.Session.UserID, sessionJSON)
	return resData
}

func WaitAnswer(reqData ReqData, resData ResData) ResData {
	rc := h.RedisClient()

	var sessionJSON SessionStore
	session, _ := rc.Get(reqData.Session.UserID).Result()
	_ = json.Unmarshal([]byte(session), &sessionJSON)

	taskId, _ := strconv.Atoi(sessionJSON.Data["TaskID"])
	task := models.GetTaskFromID(taskId)

	varinatIndex := getVarinatIndex(reqData.Request.Command, task)
	isValidAnswer := models.SetAnswer(sessionJSON.Data["GameID"], taskId, varinatIndex)

	if isValidAnswer {
		countAnswer, _ := strconv.Atoi(sessionJSON.Data["countAnswer"])
		sessionJSON.Data["countAnswer"] = strconv.Itoa(countAnswer + 1)

		newTask := controllers.GetRandomTask(sessionJSON.Data["GameID"])
		sessionJSON.Data["TaskID"] = strconv.Itoa(newTask.ID)

		if newTask.Title != "" {
			sessionJSON.Status = "waitAnswer"
			resData.Response.Text = prepareTask(newTask, sessionJSON.Data["countAnswer"])
		} else {
			sessionJSON = ReStoreAuth(sessionJSON)
			resData.Response.Text = `
				Вопросы закончились.\n
				Набранное кол-во баллов: ` + sessionJSON.Data["countAnswer"] + `\n
				Хотите сыграть еще раз?
			`
		}
	} else {
		sessionJSON = ReStoreAuth(sessionJSON)
		resData.Response.Text = `
			К сожелению вы ошиблись.\n
			Набранное кол-во баллов: ` + sessionJSON.Data["countAnswer"] + `\n
			Хотите сыграть еще раз?
		`
	}

	SetStore(reqData.Session.UserID, sessionJSON)
	return resData
}

func prepareTask(task models.Task, countAnswer string) string {
	var taskString string

	taskString += task.Title + `\n
	------------------\n
	Варианты ответов: \n
	а) ` + task.Variants[0] + `
	б) ` + task.Variants[1] + `
	в) ` + task.Variants[2] + `
	г) ` + task.Variants[3] + `
	------------------
	Кол-во ответов: ` + countAnswer

	return taskString
}

func getVarinatIndex(text string, task models.Task) int {
	if task.Variants[0] == text {
		return 1
	} else if task.Variants[1] == text {
		return 2
	} else if task.Variants[2] == text {
		return 3
	} else if task.Variants[3] == text {
		return 4
	}

	if strings.Contains(text, "а") {
		return 1
	} else if strings.Contains(text, "б") {
		return 2
	} else if strings.Contains(text, "в") {
		return 3
	} else if strings.Contains(text, "г") {
		return 4
	}

	return -1
}
