package models

import (
	"encoding/json"
	"fmt"
	"time"
)

func (ms *ModelSuite) Test_Task() {
	task := Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	}
	ms.NotEmpty(task.Description)
	ms.NotEmpty(task.Status)
	ms.NotEmpty(task.CompletionDate)
	ms.NotEmpty(task.Requester)
	ms.NotEmpty(task.Executor)
}

func (ms *ModelSuite) Test_Task_Unmarshal() {
	timeStr := "2020-08-10T08:27:00.000Z"
	testTime, err := time.Parse("2006-01-02T15:04:05.000Z", timeStr)
	ms.NoError(err)

	tcases := []struct {
		tJSON     string
		task      Task
		hasErrors bool
	}{
		{
			`testing...`,
			Task{},
			true,
		},
		{
			`{}`,
			Task{},
			false,
		},
		{
			`{"description": "Make a table with 4 chairs"}`,
			Task{
				Description: "Make a table with 4 chairs",
			},
			false,
		},
		{
			`{"status": "Done"}`,
			Task{
				Status: "Done",
			},
			false,
		},
		{
			`{"completion_date": "2020-08-10T08:27:00.000Z"}`,
			Task{
				CompletionDate: testTime,
			},
			false,
		},
		{
			`{"requester": "James Bond"}`,
			Task{
				Requester: "James Bond",
			},
			false,
		},
		{
			`{"executor": "John Smith"}`,
			Task{
				Executor: "John Smith",
			},
			false,
		},
	}

	for tIndex, tcase := range tcases {
		task := Task{}
		err = json.Unmarshal([]byte(tcase.tJSON), &task)
		if tcase.hasErrors {
			ms.Error(err, fmt.Sprintf("index: %v", tIndex))
			continue
		}
		ms.NoError(err, fmt.Sprintf("index: %v", tIndex))
		ms.Equal(task.Description, tcase.task.Description, fmt.Sprintf("index: %v", tIndex))
		ms.Equal(task.Status, tcase.task.Status, fmt.Sprintf("index: %v", tIndex))
		ms.Equal(task.CompletionDate, tcase.task.CompletionDate, fmt.Sprintf("index: %v", tIndex))
		ms.Equal(task.Requester, tcase.task.Requester, fmt.Sprintf("index: %v", tIndex))
		ms.Equal(task.Executor, tcase.task.Executor, fmt.Sprintf("index: %v", tIndex))
	}
}

func (ms *ModelSuite) Test_Task_Create() {

	task := Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	}
	ms.NoError(ms.DB.Create(&task))

	savedTask := Task{}
	ms.NoError(ms.DB.First(&savedTask))

	ms.Equal(task.Description, savedTask.Description)
	ms.Equal(task.Status, savedTask.Status)
	ms.Equal(task.CompletionDate.Format("2006-01-02"), savedTask.CompletionDate.Format("2006-01-02"))
	ms.Equal(task.Requester, savedTask.Requester)
	ms.Equal(task.Executor, savedTask.Executor)
}

func (ms *ModelSuite) Test_Task_List_Unmarshal() {
	timeStr := "2020-08-10T08:27:00.000Z"

	tcases := []struct {
		tJSON      string
		tasksCount int
		hasErrors  bool
	}{
		{
			`testing...`,
			0,
			true,
		},
		{
			`[
				{
				"description": "Make a table with 4 chairs",
				"status": "Done",
				"completion_date": "` + timeStr + `",
				"requester": "James Bond",
				"executor": "John Smith"
				}
			]`,
			1,
			false,
		},
		{
			`[
				{
					"description": "Make a table with 4 chairs",
					"status": "Done",
					"completion_date": "` + timeStr + `",
					"requester": "James Bond",
					"executor": "John Smith"
				},
				{
					"description": "Close the door",
					"status": "Done",
					"completion_date": "` + timeStr + `",
					"requester": "James Bond",
					"executor": "John Smith"
				}
			]`,
			2,
			false,
		},
	}

	for tIndex, tcase := range tcases {
		tasks := Tasks{}
		err := json.Unmarshal([]byte(tcase.tJSON), &tasks)

		if tcase.hasErrors {
			ms.Error(err, fmt.Sprintf("index: %v", tIndex))
			continue
		}

		ms.NoError(err, fmt.Sprintf("index: %v", tIndex))
		ms.Len(tasks, tcase.tasksCount)
	}
}

func (ms *ModelSuite) Test_Task_List() {
	task := Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "James Bond",
		Executor:       "John Smith",
	}
	ms.NoError(ms.DB.Create(&task))

	count, err := ms.DB.Count(Tasks{})
	ms.NoError(err)
	ms.Equal(1, count)

	task = Task{
		Description:    "Close the door",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "James Bond",
		Executor:       "John Smith",
	}
	ms.NoError(ms.DB.Create(&task))

	count, err = ms.DB.Count(Tasks{})
	ms.NoError(err)
	ms.Equal(2, count)

	tasks := Tasks{}
	ms.NoError(ms.DB.All(&tasks))
	ms.Equal("Make a table with 4 chairs", tasks[0].Description)
	ms.Equal("Close the door", tasks[1].Description)
}
