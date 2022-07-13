package entity

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var (
	ErrInvalidJSONBody = errors.New("Invalid JSON body")
)

type Customer struct {
	CustomerId      uuid.UUID `json:"customer_id" gorm:"customer_id"`
	CustomerName    string    `json:"customer_name" gorm:"customer_name"`
	CustomerAddress string    `json:"customer_address" gorm:"customer_address"`
}

func (customer *Customer) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()
	return scope.SetColumn("customer_id", id)
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func (c Customer) TableName() string {
	return "customer"
}

func (c Customer) ValidateInput() error {
	err := validation.ValidateStruct(&c,
		validation.Field(&c.CustomerName, validation.Required, validation.Length(1, 20)),
		validation.Field(&c.CustomerAddress, validation.Required, validation.Length(1, 45)),
	)
	if err != nil {
		return ErrInvalidJSONBody
	}
	return nil
}
