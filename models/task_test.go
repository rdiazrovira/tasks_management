package models

import "time"

func (ms *ModelSuite) Test_Task_Create() {
	tasks := Tasks{
		{
			Description:    "First task",
			Status:         "Done",
			CompletionDate: time.Now(),
			Requester:      "Rodolfo",
			Executor:       "Rodolfo",
		},
	}

	for _, task := range tasks {
		ms.NoError(ms.DB.Create(&task))

		t := Task{}
		ms.NoError(ms.DB.Last(&t))
		ms.Equal(t.Description, task.Description)
		ms.Equal(t.Status, task.Status)
		ms.Equal(t.CompletionDate.Format("2006-01-02"), task.CompletionDate.Format("2006-01-02"))
		ms.Equal(t.Requester, task.Requester)
	}
}
