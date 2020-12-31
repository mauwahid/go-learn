package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/mauwahid/order/server/pb"
)

const (
	port           = ":50051"
	orderBatchSize = 3
)

var orderMap = make(map[string]pb.Order)

type server struct {
	orderMap map[string]*pb.Order
}

func (s *server) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrappers.StringValue, error) {
	log.Printf("Order added. ID %v", orderReq.Id)
	orderMap[orderReq.Id] = *orderReq
	return &wrappers.StringValue{Value: "Order Added : " + orderReq.Id}, nil
}

func (s *server) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*pb.Order, error) {
	ord, exist := orderMap[orderId.Value]
	if exist {
		return &ord, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "order doesn't exist : ", orderId)
}

func
