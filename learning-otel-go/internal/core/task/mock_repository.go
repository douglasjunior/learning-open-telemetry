package task

type MockTaskRepository struct {
	tasks   []Task
	created []Task
}

func NewMockTaskRepository() *MockTaskRepository {
	return &MockTaskRepository{
		tasks:   []Task{},
		created: []Task{},
	}
}

func (m *MockTaskRepository) Create(t Task) error {
	m.tasks = append(m.tasks, t)
	m.created = append(m.created, t)
	return nil
}

func (m *MockTaskRepository) GetAll() ([]Task, error) {
	return m.tasks, nil
}

func (m *MockTaskRepository) Update(id string, newTask Task) error {
	for i, t := range m.tasks {
		if t.Id == id {
			m.tasks[i] = newTask
			return nil
		}
	}
	return nil
}

func (m *MockTaskRepository) Delete(id string) error {
	for i, t := range m.tasks {
		if t.Id == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return nil
}
