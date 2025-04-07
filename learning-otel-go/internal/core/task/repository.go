package task

type TaskRepository interface {
	Create(task Task) error
	GetAll() ([]Task, error)
	Update(id string, task Task) error
	Delete(id string) error
}
