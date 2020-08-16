package models

import "time"

func (ms *ModelSuite) Test_TaskStorage_List() {
	storage := TaskStorage{}
	ms.Len(storage.List(), 0)

	storage.Add(Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	ms.Len(storage.List(), 1)

	storage.Add(Task{
		Description:    "Close the door",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	ms.Len(storage.List(), 2)
}

func (ms *ModelSuite) Test_TaskStorage_Add() {
	storage := TaskStorage{}
	ms.Len(storage.List(), 0)

	storage.Add(Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	ms.Len(storage.List(), 1)
}

func (ms *ModelSuite) Test_TaskStorage_Clear() {
	storage := TaskStorage{}
	task := Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	}
	storage.Add(task)
	ms.Len(storage.List(), 1)

	storage.Clear()
	ms.Len(storage.List(), 0)
}

func (ms *ModelSuite) Test_TaskStorage_PendingList() {
	storage := TaskStorage{}
	ms.Len(storage.PendingList(), 0)

	storage.Add(Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	ms.Len(storage.PendingList(), 0)

	storage.Add(Task{
		Description: "Close the door",
		Status:      "Pending",
		Requester:   "John Smith",
		Executor:    "James Bond",
	})
	ms.Len(storage.PendingList(), 1)

	storage.Add(Task{
		Description: "Open the window",
		Status:      "Pending",
		Requester:   "John Smith",
		Executor:    "James Bond",
	})
	ms.Len(storage.PendingList(), 2)
}

func (ms *ModelSuite) Test_TaskStorage_DoneTaskList() {
	storage := TaskStorage{}
	ms.Len(storage.DoneTaskList(), 0)

	storage.Add(Task{
		Description: "Open the window",
		Status:      "Pending",
		Requester:   "John Smith",
		Executor:    "James Bond",
	})
	ms.Len(storage.DoneTaskList(), 0)

	storage.Add(Task{
		Description:    "Close the door",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	ms.Len(storage.DoneTaskList(), 1)

	storage.Add(Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	ms.Len(storage.DoneTaskList(), 2)
}

func (ms *ModelSuite) Test_TaskStorage_DoneTasks_InDateRange() {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	yesterday := time.Now().AddDate(0, 0, -1)
	today := time.Now()

	storage := TaskStorage{}
	ms.Len(storage.DoneTasksInDateRange(sevenDaysAgo, today), 0)

	storage.Add(Task{
		Description:    "Close the door",
		Status:         "Done",
		CompletionDate: today,
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	ms.Len(storage.DoneTasksInDateRange(sevenDaysAgo, today), 0)

	storage.Add(Task{
		Description:    "Open the window",
		Status:         "Done",
		CompletionDate: yesterday,
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	ms.Len(storage.DoneTasksInDateRange(sevenDaysAgo, today), 1)
}

func (ms *ModelSuite) Test_TaskStorage_DoneTasks_ByExecutor() {
	storage := TaskStorage{}
	ms.Len(storage.DoneTasksByExecutor("James Bond"), 0)

	storage.Add(Task{
		Description:    "Open the window",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	ms.Len(storage.DoneTasksByExecutor("Peter Howard"), 0)
	ms.Len(storage.DoneTasksByExecutor("James Bond"), 1)
}

func (ms *ModelSuite) Test_TaskStorage_Tasks_ByRequester() {
	storage := TaskStorage{}
	ms.Len(storage.TasksByRequester("John Smith"), 0)

	storage.Add(Task{
		Description:    "Open the window",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "John Smith",
		Executor:       "James Bond",
	})
	ms.Len(storage.TasksByRequester("Peter Howard"), 0)
	ms.Len(storage.TasksByRequester("John Smith"), 1)
}
