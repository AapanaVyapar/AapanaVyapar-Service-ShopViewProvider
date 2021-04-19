package services

import (
	redisDataBase "aapanavyapar-service-shopviewprovider/data-base/cash-services"
	mongoDataBase "aapanavyapar-service-shopviewprovider/data-base/data-services"
	"aapanavyapar-service-shopviewprovider/data-base/helpers"
	"aapanavyapar-service-shopviewprovider/data-base/structs"
	"aapanavyapar-service-shopviewprovider/pb"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"time"
)

type ViewProviderService struct {
	Data *mongoDataBase.MongoDataBase
	Cash *redisDataBase.CashDataBase
}

func NewViewProviderService(ctx context.Context) *ViewProviderService {
	mongoData := mongoDataBase.NewDataBase()
	redisData := redisDataBase.NewDataBase()

	view := ViewProviderService{
		Data: mongoData,
		Cash: redisData,
	}

	err := view.LoadBasicCategoriesInCash(ctx)
	if err != nil {
		panic(err)
	}

	err = view.LoadShopsInCash(ctx)
	if err != nil {
		panic(err)
	}

	err = view.LoadProductsInCash(ctx)
	if err != nil {
		panic(err)
	}

	err = view.InitShopStream(ctx)
	if err != nil {
		panic(err)
	}

	err = view.InitProductStream(ctx)
	if err != nil {
		panic(err)
	}

	return &view
}

func (viewServer *ViewProviderService) CreateShop(ctx context.Context, request *pb.CreateShopRequest) (*pb.CreateShopResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(ctx, request.GetToken(), os.Getenv("CREATE_SHOP_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	data := structs.ShopData{
		ShopId:         receivedToken.Audience,
		ShopName:       request.GetShopName(),
		ShopKeeperName: request.GetShopKeeperName(),
		Images:         request.GetImages(),
		PrimaryImage:   request.GetPrimaryImage(),
		Address: &structs.Address{
			FullName:      request.GetAddress().GetFullName(),
			HouseDetails:  request.GetAddress().GetHouseDetails(),
			StreetDetails: request.GetAddress().GetStreetDetails(),
			LandMark:      request.GetAddress().GetLandMark(),
			PinCode:       request.GetAddress().GetPinCode(),
			City:          request.GetAddress().GetCity(),
			State:         request.GetAddress().GetState(),
			Country:       request.GetAddress().GetCountry(),
			PhoneNo:       request.GetAddress().GetPhoneNo(),
		},
		Location: &structs.Location{
			Longitude: request.GetLocation().GetLongitude(),
			Latitude:  request.GetLocation().GetLatitude(),
		},
		Category:            request.GetCategory(),
		BusinessInformation: request.GetBusinessInformation(),
		OperationalHours: &structs.OperationalHours{
			Sunday:    [2]string{request.GetOperationalHours().GetSunday()[0], request.GetOperationalHours().GetSunday()[1]},
			Monday:    [2]string{request.GetOperationalHours().GetMonday()[0], request.GetOperationalHours().GetMonday()[1]},
			Tuesday:   [2]string{request.GetOperationalHours().GetTuesday()[0], request.GetOperationalHours().GetTuesday()[1]},
			Wednesday: [2]string{request.GetOperationalHours().GetWednesday()[0], request.GetOperationalHours().GetWednesday()[1]},
			Thursday:  [2]string{request.GetOperationalHours().GetThursday()[0], request.GetOperationalHours().GetThursday()[1]},
			Friday:    [2]string{request.GetOperationalHours().GetFriday()[0], request.GetOperationalHours().GetFriday()[1]},
			Saturday:  [2]string{request.GetOperationalHours().GetSaturday()[0], request.GetOperationalHours().GetSaturday()[1]},
		},
		Timestamp: time.Now().UTC(),
	}
	_, err = viewServer.Data.CreateShop(ctx, &data)
	if err != nil {
		return nil, err
	}

	marshalData := data.Marshal()

	err = viewServer.Cash.AddShopDataToCash(ctx, data.ShopId, marshalData)
	if err != nil {
		return nil, err
	}

	//array := structs.CashStructureProductArray{
	//	Products: []string{},
	//}
	//
	//err = viewServer.Cash.AddShopProductMapDataToCash(ctx, data.ShopId, array.Marshal())
	//if err != nil {
	//	return nil, err
	//}

	err = viewServer.Cash.AddShopInShopStream(ctx, marshalData)
	if err != nil {
		return nil, err
	}

	return &pb.CreateShopResponse{Ok: true}, nil
}

func (viewServer *ViewProviderService) DeleteShop(ctx context.Context, request *pb.DeleteShopRequest) (*pb.DeleteShopResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_SHOP_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	fmt.Println("Add Like : ", receivedToken)

	_, err = viewServer.Cash.GetShopDataFromCash(ctx, receivedToken.Audience)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Shop Not Found")
	}

	_, err = viewServer.Data.DelShopFromShopData(ctx, receivedToken.Audience)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail To Delete Shop")
	}

	err = viewServer.Cash.DelShopDataFromCash(ctx, receivedToken.Audience)
	if err != nil {
		return nil, err
	}

	// Call This After Deleting Respective Products Of Shop Form Cash Function In Stream.
	//err = viewServer.Cash.DelShopProductMapDataFromCash(ctx, receivedToken.Audience)
	//if err != nil {
	//	return nil, err
	//}

	//Note Remember To Delete Shop And Refresh Token Form Auth System Also

	err = viewServer.Cash.DelShopFromShopStream(ctx, receivedToken.Audience)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteShopResponse{Ok: true}, nil
}

func (viewServer *ViewProviderService) AddProduct(ctx context.Context, request *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_SHOP_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	fmt.Println("Add Like : ", receivedToken)

	_, err = viewServer.Cash.GetShopDataFromCash(ctx, receivedToken.Audience)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Shop Not Found")
	}

	productData := structs.ProductData{
		ShopId:           receivedToken.Audience,
		ShopName:         receivedToken.Subject,
		Title:            request.GetTitle(),
		ShortDescription: request.GetShortDescription(),
		Description:      request.GetDescription(),
		ShippingInfo:     request.GetShippingInfo(),
		Stock:            request.GetStock(),
		Price:            request.GetPrice(),
		Offer:            request.GetOffer(),
		Images:           request.GetImages(),
		Category:         request.GetCategory(),
		Timestamp:        time.Now().UTC(),
	}

	id, err := viewServer.Data.CreateProduct(ctx, productData)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Unable To Add Product")
	}

	productData.ProductId = id
	marshalData := productData.Marshal()

	err = viewServer.Cash.AddProductDataToCash(ctx, id.Hex(), marshalData)
	if err != nil {
		return nil, err
	}

	// Note Remember To Add Product To ShopProduct Map In Stream
	err = viewServer.Cash.AddShopProductMapDataToCash(ctx, productData.ShopId, productData.ProductId.Hex())
	if err != nil {
		return nil, err
	}

	err = viewServer.Cash.AddProductInProductStream(ctx, marshalData)
	if err != nil {
		return nil, err
	}

	return &pb.AddProductResponse{Ok: true}, nil

}

func (viewServer *ViewProviderService) DelProduct(ctx context.Context, request *pb.DelProductRequest) (*pb.DelProductResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_SHOP_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	product, err := viewServer.Cash.GetProductDataFromCash(ctx, request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Shop Not Found")
	}

	productData := structs.ProductData{}
	structs.UnmarshalProductData([]byte(product), &productData)

	if productData.ShopId != receivedToken.Audience {
		// Report For Inquiry Of ShopKeeper
		return nil, status.Errorf(codes.Unauthenticated, "You Are In Danger Now You Tried To Delete Product Which Not Belongs To You.. !!")
	}

	err = viewServer.Data.DelProductFromProductData(ctx, receivedToken.Audience, productData.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Unable To Delete Product")
	}

	err = viewServer.Cash.DelProductDataFromCash(ctx, productData.ProductId.Hex())
	if err != nil {
		return nil, err
	}

	// Note Remember To Del Product From ShopProduct Map In Stream

	err = viewServer.Cash.DelProductFromProductStream(ctx, productData.ProductId.Hex())
	if err != nil {
		return nil, err
	}

	return &pb.DelProductResponse{Ok: true}, nil

}

func (viewServer *ViewProviderService) GetShopDetails(ctx context.Context, request *pb.GetShopDetailsRequest) (*pb.GetShopDetailsResponse, error) {
	return nil, nil

}

func (viewServer *ViewProviderService) UpdateShopPrimaryImage(ctx context.Context, request *pb.UpdateShopPrimaryImageRequest) (*pb.UpdateShopPrimaryImageResponse, error) {
	return nil, nil

}
func (viewServer *ViewProviderService) UpdateShopKeeperName(ctx context.Context, request *pb.UpdateShopKeeperNameRequest) (*pb.UpdateShopKeeperNameResponse, error) {
	return nil, nil

}

func (viewServer *ViewProviderService) UpdateShopAddressAndLocation(ctx context.Context, request *pb.UpdateShopAddressAndLocationRequest) (*pb.UpdateShopAddressAndLocationResponse, error) {
	return nil, nil

}

func (viewServer *ViewProviderService) UpdateCategory(ctx context.Context, request *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	return nil, nil

}

func (viewServer *ViewProviderService) UpdateBusinessInfo(ctx context.Context, request *pb.UpdateBusinessInfoRequest) (*pb.UpdateBusinessInfoResponse, error) {
	return nil, nil

}

func (viewServer *ViewProviderService) UpdateOperationalHours(ctx context.Context, request *pb.UpdateOperationalHoursRequest) (*pb.UpdateOperationalHoursResponse, error) {
	return nil, nil

}
