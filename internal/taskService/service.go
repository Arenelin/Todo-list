package taskService

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo}
}

func (s *TaskService) GetTasks() ([]Task, error) {
	return s.repo.GetTasks()
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) UpdateTaskById(id uint, task TaskUpdate) (Task, error) {
	return s.repo.UpdateTaskById(id, task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskById(id)
}
