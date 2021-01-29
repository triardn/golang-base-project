package repository

import "math/rand"

// IPersonRepository interface for hello repository
type IPersonRepository interface {
	RandomPerson() (personName string)
}

// PersonRepository struct for hello repository
type PersonRepository struct {
	opt Option
}

// NewPersonRepository initiate new hello repository
func NewPersonRepository(opt Option) IPersonRepository {
	return &PersonRepository{
		opt: opt,
	}
}

// RandomPerson function that return
func (hr *PersonRepository) RandomPerson() (personName string) {
	allPersonNames := []string{"Albert", "Happy", "Steve", "Anthony", "Sebastian", "Edward", "John", "Howard", "Michael", "Phil"}

	selectedIndex := rand.Intn(len(allPersonNames))

	personName = allPersonNames[selectedIndex]

	return
}
