package actions

import (
	"applying-tdd-with-buffalo/tasks_management/models"

	"github.com/gobuffalo/pop/v5"
)

func AllTasks(tx *pop.Connection, tasks *models.Tasks) error {
	if err := tx.All(tasks); err != nil {
		return err
	}
	return nil
}
