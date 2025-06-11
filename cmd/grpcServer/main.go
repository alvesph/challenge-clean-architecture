package main

import (
	"log"
	"net"

	"github.com/alvesph/challenge-clean-architecture/configs"
	"github.com/alvesph/challenge-clean-architecture/internal/entity"
	"github.com/alvesph/challenge-clean-architecture/internal/pb"
	"github.com/alvesph/challenge-clean-architecture/internal/repository"
	"github.com/alvesph/challenge-clean-architecture/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	db.AutoMigrate(&entity.Order{})

	orderRepo := repository.NewOrder(db)
	orderService := service.NewOrderService(orderRepo)
	ordergRPCService := service.NewOrderGRPCService(orderService)

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, ordergRPCService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		panic("Failed to listen on port " + cfg.GrpcPort + ": " + err.Error())
	}
	log.Printf("gRPC server listening on port %s", cfg.GrpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
