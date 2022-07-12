package persistence

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CustomerGormRepo struct {
	db *gorm.DB
}

func NewCustomerGormRepo(db *gorm.DB) *CustomerGormRepo {
	return &CustomerGormRepo{db: db}
}

var _ repository.CustomerRepository = &CustomerGormRepo{}

func (cgr *CustomerGormRepo) AddCustomer(customer *entity.Customer) (*entity.Customer, error) {
	if err := cgr.db.Create(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (cgr *CustomerGormRepo) GetCustomerByID(uuid uuid.UUID) (*entity.Customer, error) {
	var customer entity.Customer

	if err := cgr.db.Where("customer_id = ?", uuid).Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
