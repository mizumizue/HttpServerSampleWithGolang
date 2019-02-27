package models

import "time"

type Task struct {
	ID int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Done bool `json:"done,omitempty"`
	Deadline time.Time `json:"deadline,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
}

type CreateTaskRequest struct {
	Name string `json:"name,omitempty"`
}

type UpdateTaskRequest struct {
	Name string `json:"name,omitempty"`
	Done bool `json:"done,omitempty"`
	Deadline time.Time `json:"deadline,omitempty"`
}

type TaskDao interface {
	GetAll() ([]Task, error)
	Get(id int64) (*Task, error)
	Create(t CreateTaskRequest) (*Task, error)
	Update(t UpdateTaskRequest) (*Task, error)
	Delete(id int64) error
}

type DatastoreTaskDao struct {
	
}

func (d *DatastoreTaskDao) GetAll() ([]Task, error) {
	panic("implement me")
}

func (d *DatastoreTaskDao) Get(id int64) (*Task, error) {
	panic("implement me")
}

func (d *DatastoreTaskDao) Create(t CreateTaskRequest) (*Task, error) {
	panic("implement me")
}

func (d *DatastoreTaskDao) Update(t UpdateTaskRequest) (*Task, error) {
	panic("implement me")
}

func (d *DatastoreTaskDao) Delete(id int64) error {
	panic("implement me")
}


