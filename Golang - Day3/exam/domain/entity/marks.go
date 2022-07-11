package entity

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

var (
	ErrMissingSubjectName = errors.New("Empty subject name")
	ErrMissingUUID        = errors.New("Invalid uuid")
	ErrMarksInvalid       = errors.New("Invalid Marks")
)

// foreign key comstraint???
type Marks struct {
	StudentId   uuid.UUID `json:"student_id"`
	SubjectName string    `json:"subject_name"`
	Marks       int       `json:"marks"`
}

// Validate marks
func (marks *Marks) Valdiate() error {
	switch {
	case marks.StudentId == uuid.Nil:
		return ErrMissingUUID
	case marks.SubjectName == "":
		return ErrMissingSubjectName
	case marks.Marks > 100 || marks.Marks < 0:
		return ErrMarksInvalid
	default:
		return nil
	}
}
