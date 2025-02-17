package storage

import (
	"Todo/models"
	"sync"
)

var (
	tasks  = make(map[int]models.Task)
	mu     sync.RWMutex
	nextID = 1
)

func PostTask(task models.Task) models.Task {
	mu.Lock()
	defer mu.Unlock()

	task.ID = nextID
	tasks[nextID] = task
	nextID++

	return task
}

func GetList() map[int]models.Task {
	mu.Lock()
	defer mu.Unlock()
	return tasks
}

func UpdateTasks(id int, updTask models.Task) (models.Task, bool) {
	mu.Lock()
	defer mu.Unlock()

	task, exists := tasks[id]
	if !exists {
		return models.Task{}, false
	}

	if updTask.Header != "" {
		task.Header = updTask.Header
	}
	if updTask.Content != "" {
		task.Content = updTask.Content
	}
	task.IsDone = updTask.IsDone

	tasks[id] = task
	return task, true
}

func DeleteTask(id int) bool {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := tasks[id]; !exists {
		return false
	}
	delete(tasks, id)
	return true
}
