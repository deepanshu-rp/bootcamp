package persistence

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type orderGormRepo struct {
	db *gorm.DB
}

func NewOrderGormRepo(db *gorm.DB) *orderGormRepo {
	return &orderGormRepo{db: db}
}

var _ repository.OrderRepository = &orderGormRepo{}

func (ogr *orderGormRepo) GetOrderByID(uuid uuid.UUID) ([]entity.OrderDetail, error) {
	var orders []entity.OrderDetail
	if err := ogr.db.Where(map[string]interface{}{"order_id": uuid, "order_status": "placed"}).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (ogr *orderGormRepo) AddOrder(order *entity.Order) (*entity.Order, error) {
	if err := ogr.db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
func (ogr *orderGormRepo) AddOrderDetail(orderDetail *entity.OrderDetail) (*entity.OrderDetail, error) {
	if err := ogr.db.Create(orderDetail).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}
