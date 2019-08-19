package store

import "github.com/asdine/storm"

type Store struct {
	db *storm.DB
}

var awsService aws.AwsService

type Task struct {
	Key   int
	Value []byte
}

// setup session : aws account credentials
func (s *Store) Init(path string) error {

	return nil
}

// get s3 bucket contents
// return list of links
func (s *Store) AllTasks() ([]*Task, error) {
	res := []*Task{}

	return res, nil
}

// upload image file to aws
// save link
func (s *Store) CreateTask(t string) (int, error) {

	return -1, nil
}

func (s *Store) DeleteTask(key int) error {
	return nil
}
