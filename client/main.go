package main

import (
	"aapanavyapar-service-shopviewprovider/data-base/helpers"
	"aapanavyapar-service-shopviewprovider/pb"
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Fali to load enviromental variables")
	}

	serverAddress := "0.0.0.0:" + os.Getenv("Port")
	flag.Parse()
	log.Printf("dialing to server  : %s", serverAddress)

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot dial server ", err)
	}

	server := pb.NewViewProviderServiceClient(conn)

	token, _ := helpers.GenerateAuthToken("9ef8d870-f900-430a-9737-b540cdbd1352", "Gurukrupa Art", "9ef8d870-f900-430a-9737-b540cdbd1352", true, []int{helpers.External})

	resp, err := server.UpdateProductData(context.Background(), &pb.UpdateProductDataRequest{
		Token:        token,
		ApiKey:       os.Getenv("API_KEY_FOR_WEB"),
		Title:        "Android Smart Phone",
		Description:  "test",
		ShippingInfo: "hi",
		Stock:        3,
		Price:        1000000,
		Offer:        2,
		Images:       []string{"https://Image1.com", "https://Image2.com"},
		Category:     []pb.Category{pb.Category_MENS_CLOTHING},
		ProductId:    "60a76bb7f5ab012effedcc5e",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.String())

}
