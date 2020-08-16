package actions

import (
	"applying-tdd-with-buffalo/tasks_management/models"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

type TasksResource struct {
	buffalo.Resource
}

var storage models.TaskStorage

func (t TasksResource) Create(c buffalo.Context) error {
	task := models.Task{}
	if err := c.Bind(&task); err != nil {
		return errors.WithStack(err)
	}

	storage.Add(task)
	return c.Render(http.StatusCreated, r.JSON(task))
}

func (t TasksResource) List(c buffalo.Context) error {
	if c.Param("requester") != "" {
		return c.Render(http.StatusOK, r.JSON(storage.TasksByRequester(c.Param("requester"))))
	}

	return c.Render(http.StatusOK, r.JSON(storage.List()))
}

func (t TasksResource) PendingList(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(storage.PendingList()))
}

func (t TasksResource) DoneTaskList(c buffalo.Context) error {
	var from, to time.Time
	if c.Param("from") != "" {
		from, _ = time.Parse(time.RFC3339, c.Param("from"))
	}

	if c.Param("to") != "" {
		to, _ = time.Parse(time.RFC3339, c.Param("to"))
	}

	if !from.IsZero() && !to.IsZero() {
		return c.Render(http.StatusOK, r.JSON(storage.DoneTasksInDateRange(from, to)))
	}

	if c.Param("executor") != "" {
		return c.Render(http.StatusOK, r.JSON(storage.DoneTasksByExecutor(c.Param("executor"))))
	}

	return c.Render(http.StatusOK, r.JSON(storage.DoneTaskList()))
}
