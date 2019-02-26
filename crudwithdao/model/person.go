package model

import (
	"cloud.google.com/go/datastore"
	"context"
)

type Person struct {
	ID	int64	`datastore:"-"json:"id"`
	Name	string	`json:"name"`
	Age	int	`json:"age"`
	Sex	string	`json:"sex"`
	Address	string	`json:"address"`
}

type PersonDao interface {
	Get(id int64) (Person, error)
	GetAll() ([]Person, error)
	Create(p Person) (Person, error)
	Update(p Person) (Person, error)
	Delete(id int64) error
}

type DatastorePersonDao struct {
	ctx context.Context
	projectId string
	kind string
	client *datastore.Client
}

func (dao *DatastorePersonDao) initClient() error {
	client, err := datastore.NewClient(dao.ctx, dao.projectId)
	if err != nil {
		return err
	}
	dao.client = client
	return nil
}

func (dao *DatastorePersonDao) Get(id int64) (Person, error) {
	if err := dao.initClient(); err != nil {
		return Person{}, err
	}
	defer dao.client.Close()

	key := datastore.IDKey("Person", id, nil)
	var p Person
	if err := dao.client.Get(dao.ctx, key, &p); err != nil {
		return p, err
	}
	return p, nil
}

func (dao *DatastorePersonDao) GetAll() ([]Person, error) {
	if err := dao.initClient(); err != nil {
		return []Person{}, err
	}
	defer dao.client.Close()

	query := datastore.NewQuery(dao.kind)
	var ps []Person
	if _, err := dao.client.GetAll(dao.ctx, query, &ps); err != nil {
		return ps, err
	}
	return ps, nil
}

func (dao *DatastorePersonDao) Create(p Person) (Person, error) {
	if err := dao.initClient(); err != nil {
		return Person{}, err
	}
	defer dao.client.Close()

	key := datastore.IncompleteKey(dao.kind, nil)
	insertedKey, err := dao.client.Put(dao.ctx, key, p)
	if err != nil {
		return p, err
	}
	p.ID = insertedKey.ID
	return p, nil
}

func (dao *DatastorePersonDao) Update(p Person) (Person, error) {
	if err := dao.initClient(); err != nil {
		return Person{}, err
	}
	defer dao.client.Close()

	key := datastore.IDKey(dao.kind, p.ID, nil)
	if _, err := dao.client.Put(dao.ctx, key, p); err != nil {
		return p, err
	}
	return p, nil
}

func (dao *DatastorePersonDao) Delete(id int64) error {
	if err := dao.initClient(); err != nil {
		return err
	}
	defer dao.client.Close()

	key := datastore.IDKey("Person", id, nil)
	if err := dao.client.Delete(dao.ctx, key); err != nil {
		return err
	}
	return nil
}

func NewPersonDao(ctx context.Context, projectId string) PersonDao {
	return &DatastorePersonDao{
		ctx: ctx,
		projectId: projectId,
		kind: "Person",
	}
}
