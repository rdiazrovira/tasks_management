package models

import "time"

type TaskStorage []Task

func (t TaskStorage) List() []Task {
	return t
}

func (t *TaskStorage) Add(task Task) {
	*t = append(*t, task)
}

func (t *TaskStorage) Clear() {
	*t = TaskStorage{}
}

func (t TaskStorage) PendingList() []Task {
	tasks := []Task{}
	for _, ts := range t {
		if ts.Status == "Pending" && ts.CompletionDate.IsZero() {
			tasks = append(tasks, ts)
		}
	}
	return tasks
}

func (t TaskStorage) DoneTaskList() []Task {
	tasks := []Task{}
	for _, ts := range t {
		if ts.Status == "Done" {
			tasks = append(tasks, ts)
		}
	}
	return tasks
}

func (t TaskStorage) DoneTasksInDateRange(from, to time.Time) []Task {
	tasks := []Task{}
	for _, ts := range t.DoneTaskList() {
		if ts.CompletionDate.After(from) && ts.CompletionDate.Before(to) {
			tasks = append(tasks, ts)
		}
	}
	return tasks
}

func (t TaskStorage) DoneTasksByExecutor(executor string) []Task {
	tasks := []Task{}
	for _, ts := range t.DoneTaskList() {
		if ts.Executor == executor {
			tasks = append(tasks, ts)
		}
	}
	return tasks
}

func (t TaskStorage) TasksByRequester(requester string) []Task {
	tasks := []Task{}
	for _, ts := range t.List() {
		if ts.Requester == requester {
			tasks = append(tasks, ts)
		}
	}
	return tasks
}
