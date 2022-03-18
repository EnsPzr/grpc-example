package main

import (
	"context"
	"github.com/enspzr/grpc-example/protos/product"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	log.Println("Client running ...")
	listProducts()
	postProducts()
	listProducts()
}

func listProducts() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := product.NewProductServiceClient(conn)

	request := &product.ProductListRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.List(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Response:", response.String())
}

func postProducts() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := product.NewProductServiceClient(conn)

	request := &product.Product{
		Id:    3,
		Name:  "Kuşbaşı",
		Price: 20,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Post(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Response:", response.String())
}
