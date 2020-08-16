package actions

import (
	"applying-tdd-with-buffalo/tasks_management/models"
	"net/http"
	"time"
)

func (as *ActionSuite) Test_TasksResource_Create() {
	storage.Clear()

	task := models.Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	}
	res := as.JSON("/tasks").Post(task)
	as.Equal(http.StatusCreated, res.Code)

	createdTask := models.Task{}
	res.Bind(&createdTask)
	as.Equal(task.Description, createdTask.Description)
	as.Equal(task.Status, createdTask.Status)
	as.Equal(task.CompletionDate.Format("2006-01-02"), createdTask.CompletionDate.Format("2006-01-02"))
	as.Equal(task.Requester, createdTask.Requester)
	as.Equal(task.Executor, createdTask.Executor)

	as.Len(storage.List(), 1)
}

func (as *ActionSuite) Test_TasksResource_List() {
	storage.Clear()

	res := as.JSON("/tasks").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks := []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	res = as.JSON("/tasks").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 1)
}

func (as *ActionSuite) Test_TasksResource_PendingList() {
	storage.Clear()

	res := as.JSON("/tasks/pending").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks := []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	res = as.JSON("/tasks/pending").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description: "Close the door",
		Status:      "Pending",
		Requester:   "John Smith",
		Executor:    "James Bond",
	})
	res = as.JSON("/tasks/pending").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 1)
}

func (as *ActionSuite) Test_TasksResource_DoneTaskList() {
	storage.Clear()

	res := as.JSON("/tasks/done").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks := []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description: "Open the window",
		Status:      "Pending",
		Requester:   "John Smith",
		Executor:    "James Bond",
	})
	res = as.JSON("/tasks/done").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description:    "Close the door",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	res = as.JSON("/tasks/done").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 1)
}

func (as *ActionSuite) Test_TasksResource_DoneTasks_InDateRange() {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	yesterday := time.Now().AddDate(0, 0, -1)
	today := time.Now()

	storage.Clear()

	res := as.JSON("/tasks/done?from=%v&to=%v", sevenDaysAgo.Format(time.RFC3339), today.Format(time.RFC3339)).Get()
	as.Equal(http.StatusOK, res.Code)

	tasks := []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description:    "Close the door",
		Status:         "Done",
		CompletionDate: today,
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	res = as.JSON("/tasks/done?from=%v&to=%v", sevenDaysAgo.Format(time.RFC3339), today.Format(time.RFC3339)).Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description:    "Open the window",
		Status:         "Done",
		CompletionDate: yesterday,
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	res = as.JSON("/tasks/done?from=%v&to=%v", sevenDaysAgo.Format(time.RFC3339), today.Format(time.RFC3339)).Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 1)
}

func (as *ActionSuite) Test_TasksResource_DoneTasks_ByExecutor() {
	storage.Clear()

	res := as.JSON("/tasks/done?executor=%v", "James Bond").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks := []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description:    "Open the window",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	res = as.JSON("/tasks/done?executor=%v", "Peter Howard").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description:    "Close the door",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "Peter Howard",
	})
	res = as.JSON("/tasks/done?executor=%v", "Peter Howard").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 1)
}

func (as *ActionSuite) Test_TasksResource_Tasks_ByRequester() {
	storage.Clear()

	res := as.JSON("/tasks/?requester=%v", "John Smith").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks := []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description:    "Open the window",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	res = as.JSON("/tasks/?requester=%v", "Peter Howard").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 0)

	storage.Add(models.Task{
		Description:    "Close the door",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "Peter Howard",
		Executor:       "James Bond",
	})
	res = as.JSON("/tasks/?requester=%v", "Peter Howard").Get()
	as.Equal(http.StatusOK, res.Code)

	tasks = []models.Task{}
	res.Bind(&tasks)
	as.Len(tasks, 1)
}
