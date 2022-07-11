package interfaces

import (
	"exam/application"
	"exam/domain/entity"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	student1 = entity.Student{
		Id:        uuid.FromStringOrNil("e3547962-01aa-4418-8b03-21103a9cedc1"),
		FirstName: "Test",
		LastName:  "User",
		DOB:       "5/10/20",
		Address:   "Home",
	}
	student2 = entity.Student{
		Id:        uuid.FromStringOrNil("e9ea26fd-12ab-44e7-baf8-77890d7d27ac"),
		FirstName: "Test",
		LastName:  "User",
		DOB:       "5/10/20",
		Address:   "Home",
	}
)

type MockApp struct {
	mock.Mock
}

var _ application.StudentAppInterface = &MockApp{}

type TestsForStudentById struct {
	arg      uuid.UUID
	expected *entity.Student
}

func TestGetStudentByID(t *testing.T) {
	mockApp := new(MockApp)
	tests := []TestsForStudentById{
		{
			arg:      student1.Id,
			expected: &student1,
		},
		{
			arg:      student2.Id,
			expected: &student2,
		},
	}
	// what??
	mockApp.AssertExpectations(t)

	mockApp.On("GetStudentByID").Return(&student1, nil)
	testApp := NewStudentDependency(mockApp)

	for _, test := range tests {
		result, _ := testApp.stu.GetStudentByID(test.arg)
		// print(test)
		assert.Equal(t, result, test.expected)
	}

}

func TestGetAllStudents(t *testing.T) {
	mockApp := new(MockApp)
	studentsMock := []entity.Student{student1, student2}
	studentsExpected := []entity.Student{student1, student2}
	// Expectations
	mockApp.On("GetAllStudents").Return(studentsMock, nil)
	testApp := NewStudentDependency(mockApp)
	result, _ := testApp.stu.GetAllStudents()

	// Assertion on expectation created above
	mockApp.AssertExpectations(t)

	// Data assertion for all values
	for r, val := range result {
		assert.Equal(t, val, studentsExpected[r])
	}
}

// Stubs
func (mApp *MockApp) GetAllStudents() ([]entity.Student, error) {
	args := mApp.Called()
	result := args.Get(0)
	return result.([]entity.Student), args.Error(1)
}

func (mApp *MockApp) AddStudent(*entity.Student) (*entity.Student, error) {
	args := mApp.Called()
	result := args.Get(0)
	// need type assertion?
	return result.(*entity.Student), args.Error(1)

}
func (mApp *MockApp) GetStudentByID(uuid uuid.UUID) (*entity.Student, error) {
	args := mApp.Called()
	result := args.Get(0)
	// need type assertion?
	return result.(*entity.Student), args.Error(1)
}
