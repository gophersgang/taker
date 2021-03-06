package taker

// Task represents a job that can be run.
// You are encouraged to amend this interface accordingly in order to reference
// results of the task.
type Task interface {
	// Run synchronously executes the task.
	// In order to simulate multiple return values while guaranteeing
	// type-safety, amend the interface with additional properties.
	Run() error
}

type task struct {
	run func() error
}

func (r *task) Run() error {
	return r.run()
}

// Wrap creates a new task from a run function.
// This is useful in case you want to wrap an arbitrary function into a task.
func Wrap(run func() error) Task {
	return &task{run}
}
