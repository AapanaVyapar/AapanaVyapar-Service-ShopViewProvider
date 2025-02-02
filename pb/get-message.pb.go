// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: get-message.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetShopDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,70,opt,name=ApiKey,proto3" json:"ApiKey,omitempty"`
	Token  string `protobuf:"bytes,71,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *GetShopDetailsRequest) Reset() {
	*x = GetShopDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShopDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShopDetailsRequest) ProtoMessage() {}

func (x *GetShopDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShopDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetShopDetailsRequest) Descriptor() ([]byte, []int) {
	return file_get_message_proto_rawDescGZIP(), []int{0}
}

func (x *GetShopDetailsRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *GetShopDetailsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetShopDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopName            string            `protobuf:"bytes,72,opt,name=ShopName,proto3" json:"ShopName,omitempty"`
	ShopKeeperName      string            `protobuf:"bytes,73,opt,name=ShopKeeperName,proto3" json:"ShopKeeperName,omitempty"`
	Images              []string          `protobuf:"bytes,74,rep,name=Images,proto3" json:"Images,omitempty"`
	PrimaryImage        string            `protobuf:"bytes,75,opt,name=PrimaryImage,proto3" json:"PrimaryImage,omitempty"`
	Address             *Address          `protobuf:"bytes,76,opt,name=Address,proto3" json:"Address,omitempty"`
	Location            *Location         `protobuf:"bytes,77,opt,name=Location,proto3" json:"Location,omitempty"`
	Category            []Category        `protobuf:"varint,78,rep,packed,name=Category,proto3,enum=Category" json:"Category,omitempty"`
	BusinessInformation string            `protobuf:"bytes,79,opt,name=BusinessInformation,proto3" json:"BusinessInformation,omitempty"`
	OperationalHours    *OperationalHours `protobuf:"bytes,80,opt,name=OperationalHours,proto3" json:"OperationalHours,omitempty"`
	Ratings             []*RatingOfShop   `protobuf:"bytes,81,rep,name=ratings,proto3" json:"ratings,omitempty"`
	Timestamp           string            `protobuf:"bytes,82,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetShopDetailsResponse) Reset() {
	*x = GetShopDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShopDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShopDetailsResponse) ProtoMessage() {}

func (x *GetShopDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShopDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetShopDetailsResponse) Descriptor() ([]byte, []int) {
	return file_get_message_proto_rawDescGZIP(), []int{1}
}

func (x *GetShopDetailsResponse) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

func (x *GetShopDetailsResponse) GetShopKeeperName() string {
	if x != nil {
		return x.ShopKeeperName
	}
	return ""
}

func (x *GetShopDetailsResponse) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *GetShopDetailsResponse) GetPrimaryImage() string {
	if x != nil {
		return x.PrimaryImage
	}
	return ""
}

func (x *GetShopDetailsResponse) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *GetShopDetailsResponse) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *GetShopDetailsResponse) GetCategory() []Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *GetShopDetailsResponse) GetBusinessInformation() string {
	if x != nil {
		return x.BusinessInformation
	}
	return ""
}

func (x *GetShopDetailsResponse) GetOperationalHours() *OperationalHours {
	if x != nil {
		return x.OperationalHours
	}
	return nil
}

func (x *GetShopDetailsResponse) GetRatings() []*RatingOfShop {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *GetShopDetailsResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type GetOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,83,opt,name=ApiKey,proto3" json:"ApiKey,omitempty"`
	Token  string `protobuf:"bytes,84,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *GetOrdersRequest) Reset() {
	*x = GetOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersRequest) ProtoMessage() {}

func (x *GetOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersRequest) Descriptor() ([]byte, []int) {
	return file_get_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrdersRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *GetOrdersRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId           string   `protobuf:"bytes,85,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Status            Status   `protobuf:"varint,86,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	ProductId         string   `protobuf:"bytes,87,opt,name=productId,proto3" json:"productId,omitempty"`
	DeliveryTimeStamp string   `protobuf:"bytes,88,opt,name=deliveryTimeStamp,proto3" json:"deliveryTimeStamp,omitempty"`
	OrderTimeStamp    string   `protobuf:"bytes,89,opt,name=orderTimeStamp,proto3" json:"orderTimeStamp,omitempty"`
	Price             float32  `protobuf:"fixed32,90,opt,name=price,proto3" json:"price,omitempty"`
	Quantity          uint32   `protobuf:"varint,91,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ProductName       string   `protobuf:"bytes,92,opt,name=productName,proto3" json:"productName,omitempty"`
	ProductImage      string   `protobuf:"bytes,93,opt,name=productImage,proto3" json:"productImage,omitempty"`
	Address           *Address `protobuf:"bytes,94,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetOrdersResponse) Reset() {
	*x = GetOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersResponse) ProtoMessage() {}

func (x *GetOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetOrdersResponse) Descriptor() ([]byte, []int) {
	return file_get_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrdersResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *GetOrdersResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_PENDING
}

func (x *GetOrdersResponse) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *GetOrdersResponse) GetDeliveryTimeStamp() string {
	if x != nil {
		return x.DeliveryTimeStamp
	}
	return ""
}

func (x *GetOrdersResponse) GetOrderTimeStamp() string {
	if x != nil {
		return x.OrderTimeStamp
	}
	return ""
}

func (x *GetOrdersResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetOrdersResponse) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *GetOrdersResponse) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *GetOrdersResponse) GetProductImage() string {
	if x != nil {
		return x.ProductImage
	}
	return ""
}

func (x *GetOrdersResponse) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type GetProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,95,opt,name=ApiKey,proto3" json:"ApiKey,omitempty"`
	Token  string `protobuf:"bytes,96,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *GetProductsRequest) Reset() {
	*x = GetProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsRequest) ProtoMessage() {}

func (x *GetProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsRequest.ProtoReflect.Descriptor instead.
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return file_get_message_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductsRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *GetProductsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId    string     `protobuf:"bytes,97,opt,name=productId,proto3" json:"productId,omitempty"`
	ProductName  string     `protobuf:"bytes,98,opt,name=productName,proto3" json:"productName,omitempty"`
	Description  string     `protobuf:"bytes,99,opt,name=Description,proto3" json:"Description,omitempty"`
	ShippingInfo string     `protobuf:"bytes,100,opt,name=ShippingInfo,proto3" json:"ShippingInfo,omitempty"`
	Stock        uint32     `protobuf:"varint,101,opt,name=Stock,proto3" json:"Stock,omitempty"`
	Price        float32    `protobuf:"fixed32,102,opt,name=Price,proto3" json:"Price,omitempty"`
	Offer        uint32     `protobuf:"varint,103,opt,name=Offer,proto3" json:"Offer,omitempty"`
	Images       []string   `protobuf:"bytes,104,rep,name=Images,proto3" json:"Images,omitempty"`
	Category     []Category `protobuf:"varint,105,rep,packed,name=Category,proto3,enum=Category" json:"Category,omitempty"`
}

func (x *GetProductsResponse) Reset() {
	*x = GetProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsResponse) ProtoMessage() {}

func (x *GetProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsResponse.ProtoReflect.Descriptor instead.
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return file_get_message_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductsResponse) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *GetProductsResponse) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *GetProductsResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetProductsResponse) GetShippingInfo() string {
	if x != nil {
		return x.ShippingInfo
	}
	return ""
}

func (x *GetProductsResponse) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *GetProductsResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetProductsResponse) GetOffer() uint32 {
	if x != nil {
		return x.Offer
	}
	return 0
}

func (x *GetProductsResponse) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *GetProductsResponse) GetCategory() []Category {
	if x != nil {
		return x.Category
	}
	return nil
}

var File_get_message_proto protoreflect.FileDescriptor

var file_get_message_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x65, 0x74, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x6f, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x46, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x47, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xc2, 0x03, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x68, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x48, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53,
	0x68, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x70, 0x4b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x49, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x53, 0x68, 0x6f, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x4a, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x4b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x4c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x25, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x4d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x4e, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x30, 0x0a,
	0x13, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x4f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3d, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x10, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x27,
	0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x51, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x66, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x07,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x52, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x18, 0x53, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x54, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xde, 0x02, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x55, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x56, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x57, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x58, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x59, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x5b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x5c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x5d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x5e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x5f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x60, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9c, 0x02, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x62, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x18,
	0x67, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x68, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x69, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x30, 0x0a, 0x26, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x61, 0x70, 0x61, 0x6e, 0x61, 0x76, 0x79, 0x61, 0x70, 0x61, 0x72, 0x2e,
	0x61, 0x61, 0x70, 0x61, 0x6e, 0x61, 0x76, 0x79, 0x61, 0x70, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_message_proto_rawDescOnce sync.Once
	file_get_message_proto_rawDescData = file_get_message_proto_rawDesc
)

func file_get_message_proto_rawDescGZIP() []byte {
	file_get_message_proto_rawDescOnce.Do(func() {
		file_get_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_message_proto_rawDescData)
	})
	return file_get_message_proto_rawDescData
}

var file_get_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_get_message_proto_goTypes = []interface{}{
	(*GetShopDetailsRequest)(nil),  // 0: GetShopDetailsRequest
	(*GetShopDetailsResponse)(nil), // 1: GetShopDetailsResponse
	(*GetOrdersRequest)(nil),       // 2: GetOrdersRequest
	(*GetOrdersResponse)(nil),      // 3: GetOrdersResponse
	(*GetProductsRequest)(nil),     // 4: GetProductsRequest
	(*GetProductsResponse)(nil),    // 5: GetProductsResponse
	(*Address)(nil),                // 6: Address
	(*Location)(nil),               // 7: Location
	(Category)(0),                  // 8: Category
	(*OperationalHours)(nil),       // 9: OperationalHours
	(*RatingOfShop)(nil),           // 10: RatingOfShop
	(Status)(0),                    // 11: Status
}
var file_get_message_proto_depIdxs = []int32{
	6,  // 0: GetShopDetailsResponse.Address:type_name -> Address
	7,  // 1: GetShopDetailsResponse.Location:type_name -> Location
	8,  // 2: GetShopDetailsResponse.Category:type_name -> Category
	9,  // 3: GetShopDetailsResponse.OperationalHours:type_name -> OperationalHours
	10, // 4: GetShopDetailsResponse.ratings:type_name -> RatingOfShop
	11, // 5: GetOrdersResponse.status:type_name -> Status
	6,  // 6: GetOrdersResponse.address:type_name -> Address
	8,  // 7: GetProductsResponse.Category:type_name -> Category
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_get_message_proto_init() }
func file_get_message_proto_init() {
	if File_get_message_proto != nil {
		return
	}
	file_common_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_get_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShopDetailsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_get_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShopDetailsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_get_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_get_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_get_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_get_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_get_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_message_proto_goTypes,
		DependencyIndexes: file_get_message_proto_depIdxs,
		MessageInfos:      file_get_message_proto_msgTypes,
	}.Build()
	File_get_message_proto = out.File
	file_get_message_proto_rawDesc = nil
	file_get_message_proto_goTypes = nil
	file_get_message_proto_depIdxs = nil
}
