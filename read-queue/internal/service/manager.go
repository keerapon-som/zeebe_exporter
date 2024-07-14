package service

type Manager struct {
	VariableManager
	JobManager
}

func NewManager() *Manager {
	return &Manager{
		VariableManager: &variableManager{},
		JobManager:      &jobManager{},
	}
}
