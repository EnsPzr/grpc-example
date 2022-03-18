package main

import (
	"context"
	"github.com/enspzr/grpc-example/protos/product"
	"google.golang.org/grpc"
	"log"
	"net"
)

var items = []*product.Product{
	{
		Id:    1,
		Name:  "Lahmacun",
		Price: 12,
	}, {
		Id:    2,
		Name:  "Adana Dürüm",
		Price: 18,
	},
}

type server struct {
	product.ProductServiceServer
}

func (s *server) List(ctx context.Context, in *product.ProductListRequest) (*product.ProductListResponse, error) {
	return &product.ProductListResponse{
		Items: items,
	}, nil
}

func (s *server) Post(ctx context.Context, vm *product.Product) (*product.Product, error) {
	items = append(items, vm)
	return vm, nil
}

func main() {
	log.Println("Server running ...")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	product.RegisterProductServiceServer(srv, &server{})

	log.Fatalln(srv.Serve(lis))
}
