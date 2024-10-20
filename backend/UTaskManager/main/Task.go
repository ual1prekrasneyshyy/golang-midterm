package main

type Task struct {
	Id          int
	Title       string
	Description string
	Status      string // Possible values: IN_PROGRESS (default), COMPLETED, FAILED
	CompletedBy User
}

func completeTask(t Task) {
	t.Status = "COMPLETED"
}

func refuseFromTask(t Task) {
	t.Status = "FAILED"
}
