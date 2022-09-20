package db

// Pop a task from an implant
func PopTask() {
	Mu.Lock()
	defer Mu.Unlock()
	// TODO: Query implant
	// TODO: Get the task
	// TODO: Remove the task from the implant
}

// TODO: Get an Implant from Db
func Implant() {
	Mu.Lock()
	defer Mu.Unlock()
	// ...
}

// TODO: Add an Implant to the Db
func AddImplant() {
	Mu.Lock()
	defer Mu.Unlock()
}

// TODO: Get a Task from Db
func Task() {
	Mu.Lock()
	defer Mu.Unlock()
	// ...
}

// TODO: Add a Task to the Db
func AddTask() {

}
