package models

type Student struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	DOB      string `json:"dob"`
	Address  string `json:"address"`
	Subject  string `json:"subject"`
	Marks    uint   `json:"marks"`
}
