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
