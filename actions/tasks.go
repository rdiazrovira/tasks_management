package actions

import (
	"applying-tdd-with-buffalo/tasks_management/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

type TasksResource struct {
	buffalo.Resource
}

func (t TasksResource) Create(c buffalo.Context) error {
	task := models.Task{}
	if err := c.Bind(&task); err != nil {
		return errors.WithStack(err)
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	if err := tx.Create(&task); err != nil {
		return errors.Wrap(err, "error by trying to create a task")
	}

	return c.Render(http.StatusCreated, r.JSON(task))
}

func (t TasksResource) List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	tasks := models.Tasks{}
	if err := tx.All(&tasks); err != nil {
		return errors.Wrap(err, "error by trying to get all tasks")
	}

	return c.Render(http.StatusOK, r.JSON(tasks))
}
