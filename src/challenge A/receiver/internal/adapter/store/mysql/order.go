package mysql

import (
	"context"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/entities"
)

func (m MySQLStore) CreateOrder(ctx context.Context, order entities.Order) (entities.Order, error) {
	o := mapToOrderModel(order)
	// Execute a prepared statement to insert the order into the database
	if err := m.db.WithContext(ctx).Create(&o).Error; err != nil {
		return entities.Order{}, err
	}
	return mapFromOrderModel(o), nil
}
