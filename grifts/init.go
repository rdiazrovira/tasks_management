package grifts

import (
	"applying-tdd-with-buffalo/tasks_management/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
