package actions

import (
	"applying-tdd-with-buffalo/tasks_management/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (as *ActionSuite) Test_TasksResource_Create() {
	tasks := models.Tasks{
		{},
		{Description: "Make a table with 4 chairs"},
		{Status: "Done"},
		{CompletionDate: time.Now()},
		{Requester: "James Bond"},
		{Executor: "John Smith"},
	}

	for tIndex, task := range tasks {
		res := as.JSON("/task").Post(task)
		as.Equal(http.StatusCreated, res.Code)

		t := models.Task{}
		as.NoError(as.DB.Last(&t))

		as.Equal(task.Description, t.Description, fmt.Sprintf("index: %v", tIndex))
		as.Equal(task.Status, t.Status, fmt.Sprintf("index: %v", tIndex))
		as.Equal(task.CompletionDate.Format("2006-01-02"), t.CompletionDate.Format("2006-01-02"), fmt.Sprintf("index: %v", tIndex))
		as.Equal(task.Requester, t.Requester, fmt.Sprintf("index: %v", tIndex))
		as.Equal(task.Executor, t.Executor, fmt.Sprintf("index: %v", tIndex))

		response := models.Task{}
		as.NoError(json.Unmarshal(res.Body.Bytes(), &response))

		as.Equal(task.Description, response.Description, fmt.Sprintf("index: %v", tIndex))
		as.Equal(task.Status, response.Status, fmt.Sprintf("index: %v", tIndex))
		as.Equal(task.CompletionDate.Format("2006-01-02"), t.CompletionDate.Format("2006-01-02"), fmt.Sprintf("index: %v", tIndex))
		as.Equal(task.Requester, response.Requester, fmt.Sprintf("index: %v", tIndex))
		as.Equal(task.Executor, response.Executor, fmt.Sprintf("index: %v", tIndex))
	}

	tasks = models.Tasks{}
	as.NoError(as.DB.All(&tasks))
	as.Len(tasks, 6)
}

func (as *ActionSuite) Test_TasksResource_List() {
	tasks := models.Tasks{
		{
			Description:    "Make a table with 4 chairs",
			Status:         "Done",
			CompletionDate: time.Now(),
			Requester:      "John Smith",
			Executor:       "James Bond",
		},
		{
			Description:    "Close the door",
			Status:         "Done",
			CompletionDate: time.Now(),
			Requester:      "John Smith",
			Executor:       "James Bond",
		},
	}
	as.NoError(as.DB.Create(&tasks))

	res := as.JSON("/tasks").Get()
	as.Equal(http.StatusOK, res.Code)

	response := models.Tasks{}
	as.NoError(json.Unmarshal(res.Body.Bytes(), &response))
	as.Len(response, 2)

	for rIndex := range response {
		as.Equal(tasks[rIndex].Description, response[rIndex].Description, fmt.Sprintf("index: %v", rIndex))
		as.Equal(tasks[rIndex].Status, response[rIndex].Status, fmt.Sprintf("index: %v", rIndex))
		as.Equal(tasks[rIndex].CompletionDate.Format("2006-01-02"), response[rIndex].CompletionDate.Format("2006-01-02"), fmt.Sprintf("index: %v", rIndex))
		as.Equal(tasks[rIndex].Requester, response[rIndex].Requester, fmt.Sprintf("index: %v", rIndex))
		as.Equal(tasks[rIndex].Executor, response[rIndex].Executor, fmt.Sprintf("index: %v", rIndex))
	}
}
