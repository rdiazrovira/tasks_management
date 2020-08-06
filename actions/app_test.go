package actions

func (as *ActionSuite) Test_Routes() {

	routes := []struct {
		method  string
		path    string
		handler string
	}{
		{"GET", "/tasks/", "applying-tdd-with-buffalo/tasks_management/actions.TasksResource.Create"},
	}

	for _, route := range as.App.Routes() {

		found := false
		for _, r := range routes {
			if r.method == route.Method && r.path == route.Path && r.handler == route.HandlerName {
				found = true
			}
		}

		as.True(found)
	}
}
