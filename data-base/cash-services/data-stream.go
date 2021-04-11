package data_base

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

func (dataBase *CashDataBase) AddShopInShopStream(ctx context.Context, data []byte) error {

	err := dataBase.Cash.XAdd(ctx, &redis.XAddArgs{
		Stream:       os.Getenv("REDIS_STREAM_SHOP_NAME"),
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       []string{"shop", string(data), "operation", "+"},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *CashDataBase) DelShopFromShopStream(ctx context.Context, shopId string) error {

	err := dataBase.Cash.XAdd(ctx, &redis.XAddArgs{
		Stream:       os.Getenv("REDIS_STREAM_SHOP_NAME"),
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       []string{"shop", shopId, "operation", "-"},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *CashDataBase) AddProductInProductStream(ctx context.Context, data []byte) error {

	err := dataBase.Cash.XAdd(ctx, &redis.XAddArgs{
		Stream:       os.Getenv("REDIS_STREAM_PRODUCT_NAME"),
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       []string{"shop", string(data), "operation", "+"},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *CashDataBase) DelProductFromProductStream(ctx context.Context, productId string) error {

	err := dataBase.Cash.XAdd(ctx, &redis.XAddArgs{
		Stream:       os.Getenv("REDIS_STREAM_PRODUCT_NAME"),
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       []string{"product", productId, "operation", "-"},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}
