package entity

import (
	"errors"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var (
	ErrMissingFirstName = errors.New("Empty first name")
	ErrMissingLastName  = errors.New("Empty last name")
	ErrMissingAddress   = errors.New("Empty Address")
	ErrMissingDOB       = errors.New("Empty DOB")
)

type Student struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	DOB       string    `json:"dob"`
	Address   string    `json:"address"`
}

func (student *Student) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4().String()
	return scope.SetColumn("id", id)
}

// Validate
func (student *Student) Valdiate() error {
	switch {
	case student.FirstName == "":
		return ErrMissingFirstName
	case student.LastName == "":
		return ErrMissingLastName
	case student.Address == "":
		return ErrMissingAddress
	case student.DOB == "":
		return ErrMissingDOB
	default:
		return nil
	}
}
