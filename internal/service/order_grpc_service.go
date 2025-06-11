package service

import (
	"context"
	"fmt"

	"github.com/alvesph/challenge-clean-architecture/internal/pb"
)

type OrderGRPCService struct {
	pb.UnimplementedOrderServiceServer
	Service *OrderService
}

func NewOrderGRPCService(service *OrderService) *OrderGRPCService {
	return &OrderGRPCService{
		Service: service,
	}
}

func (s *OrderGRPCService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.Service.FindAll()
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, o := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:         fmt.Sprintf("%v", o.ID),
			CustomerId: o.CustomerID,
			ProductId:  o.ProductID,
			Quantity:   fmt.Sprintf("%v", o.Quantity), // helper para converter int para string
			TotalPrice: o.TotalPrice,
			Status:     o.Status,
			CreatedAt:  o.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  o.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.ListOrdersResponse{Orders: pbOrders}, nil
}
