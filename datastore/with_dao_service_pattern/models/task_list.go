package models

type TaskList struct {
	ID int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
}

type CreateOrUpdateTaskListRequest struct {
	Name string `json:"name,omitempty"`
}

type TaskListDao interface {
	GetAll() ([]TaskList, error)
	Get(id int64) (*TaskList, error)
	Create(t CreateOrUpdateTaskListRequest) (*TaskList, error)
	Update(t CreateOrUpdateTaskListRequest) (*TaskList, error)
	Delete(id int64) error
}

type DatastoreTaskListDao struct {
	// TODO
}

func (d *DatastoreTaskListDao) GetAll() ([]TaskList, error) {
	panic("implement me")
}

func (d *DatastoreTaskListDao) Get(id int64) (*TaskList, error) {
	panic("implement me")
}

func (d *DatastoreTaskListDao) Create(t CreateOrUpdateTaskListRequest) (*TaskList, error) {
	panic("implement me")
}

func (d *DatastoreTaskListDao) Update(t CreateOrUpdateTaskListRequest) (*TaskList, error) {
	panic("implement me")
}

func (d *DatastoreTaskListDao) Delete(id int64) error {
	panic("implement me")
}
