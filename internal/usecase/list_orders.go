package usecase

import (
	"github.com/gabrielmq/clean-arch-go/internal/entity"
)

type ListOrderOutput struct {
	Orders []OrderOutput `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() (ListOrderOutput, error) {
	orders, err := l.OrderRepository.ListOrders()
	if err != nil {
		print(err)
		return ListOrderOutput{}, err
	}

	var output ListOrderOutput
	for _, order := range orders {
		output.Orders = append(output.Orders, OrderOutput{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return output, nil
}
