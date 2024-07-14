package services

import (
	"aapanavyapar-service-shopviewprovider/data-base/structs"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"strings"
)

func (viewServer *ViewProviderService) LoadBasicCategoriesInCash(ctx context.Context) error {

	err := viewServer.Data.GetAllBasicCategories(ctx, func(data structs.BasicCategoriesData) error {
		err := viewServer.Cash.Cash.HSet(ctx, "categories", data.Category, data.Marshal()).Err()
		if err != nil {
			return status.Errorf(codes.Internal, "unable to add data to hash of Cash  : %w", err)
		}
		return nil

	})
	if err != nil {
		return err

	}
	return nil

}

func (viewServer *ViewProviderService) LoadShopsInCash(ctx context.Context) error {

	err := viewServer.Data.GetAllShopsFromShopData(ctx, func(data structs.ShopData) error {

		err := viewServer.Cash.AddShopDataToCash(ctx, data.ShopId, data.Marshal())
		if err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		return err

	}
	return nil

}

func (viewServer *ViewProviderService) LoadProductsInCash(ctx context.Context) error {

	err := viewServer.Data.GetAllProductsFromProductData(ctx, func(data structs.ProductData) error {

		err := viewServer.Cash.AddProductDataToCash(ctx, data.ProductId.Hex(), data.Marshal())
		if err != nil {
			return err
		}

		err = viewServer.Cash.AddShopProductMapDataToCash(ctx, data.ShopId, data.ProductId.Hex())
		if err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		return err

	}
	return nil

}

func (viewServer *ViewProviderService) InitShopStream(ctx context.Context) error {
	err := viewServer.CreateShopStream(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "already exist") {
			fmt.Println("Cart Already Exist")
			return nil
		}
		return err
	}
	return nil
}

func (viewServer *ViewProviderService) InitProductStream(ctx context.Context) error {
	err := viewServer.CreateProductStream(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "already exist") {
			fmt.Println("Fav Already Exist")
			return nil
		}
		return err
	}
	return nil
}

func (viewServer *ViewProviderService) InitUpdateProductStream(ctx context.Context) error {
	err := viewServer.CreateUpdateProductStream(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "already exist") {
			fmt.Println("Update Stream Already Exist")
			return nil
		}
		return err
	}
	return nil
}

func (viewServer *ViewProviderService) CreateShopStream(ctx context.Context) error {
	return viewServer.Cash.Cash.XGroupCreateMkStream(ctx, os.Getenv("REDIS_STREAM_SHOP_NAME"), os.Getenv("REDIS_STREAM_SHOP_GROUP"), "$").Err()

}

func (viewServer *ViewProviderService) CreateProductStream(ctx context.Context) error {
	return viewServer.Cash.Cash.XGroupCreateMkStream(ctx, os.Getenv("REDIS_STREAM_PRODUCT_NAME"), os.Getenv("REDIS_STREAM_PRODUCT_GROUP"), "$").Err()

}

func (viewServer *ViewProviderService) CreateUpdateProductStream(ctx context.Context) error {
	return viewServer.Cash.Cash.XGroupCreateMkStream(ctx, os.Getenv("REDIS_STREAM_UPDATE_PRODUCT_NAME"), os.Getenv("REDIS_STREAM_UPDATE_PRODUCT_GROUP"), "$").Err()

}
