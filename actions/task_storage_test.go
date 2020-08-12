package actions

import (
	"applying-tdd-with-buffalo/tasks_management/models"
	"time"
)

func (as *ActionSuite) Test_TaskStorage_AllTasks() {
	task := models.Task{
		Description:    "Make a table with 4 chairs",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "James Bond",
		Executor:       "John Smith",
	}
	as.NoError(as.DB.Create(&task))

	count, err := as.DB.Count(models.Tasks{})
	as.NoError(err)
	as.Equal(1, count)

	task = models.Task{
		Description:    "Close the door",
		Status:         "Done",
		CompletionDate: time.Now(),
		Requester:      "James Bond",
		Executor:       "John Smith",
	}
	as.NoError(as.DB.Create(&task))

	count, err = as.DB.Count(models.Tasks{})
	as.NoError(err)
	as.Equal(2, count)

	/*tasks := models.Tasks{}
	as.NoError(AllTasks(as.DB, &tasks))
	as.Equal("Make a table with 4 chairs", tasks[0].Description)
	as.Equal("Close the door", tasks[1].Description)*/
}
