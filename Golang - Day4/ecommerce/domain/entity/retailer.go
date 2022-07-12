package entity

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Retailer struct {
	RetailerId     uuid.UUID `json:"retailer_id" gorm:"retailer_id"`
	RetailerName   string    `json:"retailer_name" gorm:"retailer_name"`
	RetailerRating int       `json:"retailer_rating" gorm:"retailer_rating"`
}

func (retailer *Retailer) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4().String()
	return scope.SetColumn("retailer_id", id)
}

func (r Retailer) TableName() string {
	return "retailer"
}
