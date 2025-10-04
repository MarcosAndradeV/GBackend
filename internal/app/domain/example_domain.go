package domain

import "context"

type ExampleUser struct {
	ID       int
	Name     string
}

type UserRepository interface {
	Create(c context.Context, user *ExampleUser) error
	Fetch(c context.Context) ([]ExampleUser, error)
}
