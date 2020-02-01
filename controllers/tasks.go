package controllers

import (
	"alisa-millionaire-lite/server/models"
)

type TaskData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func GetRandomTask(GameID string) models.Task {
	task := models.GetRandomTask(GameID)
	return task
}

func GetTaskFromID(TaskID int) models.Task {
	task := models.GetTaskFromID(TaskID)
	return task
}

/*
func CheckAnswer(res http.ResponseWriter, req *http.Request) {
	var data struct {
		GameID       string `json:"gameId"`
		TaskID       int    `json:"taskId"`
		VarinatIndex int    `json:"varinatIndex"`
	}
	body, _ := ioutil.ReadAll(req.Body)
	_ = json.Unmarshal(body, &data)

	isValid := models.SetAnswer(data.GameID, data.TaskID, data.VarinatIndex)
	res.Write([]byte(strconv.FormatBool(isValid)))
}
*/
