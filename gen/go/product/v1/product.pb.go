// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: product/v1/product.proto

package productv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetProductRequest is the request message for the GetProduct method.
type GetProductRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the product to get.
	ProductId     string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_product_v1_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_product_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

// GetProductResponse is the response message for the GetProduct method.
type GetProductResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The product.
	Product       *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductResponse) Reset() {
	*x = GetProductResponse{}
	mi := &file_product_v1_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResponse) ProtoMessage() {}

func (x *GetProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResponse.ProtoReflect.Descriptor instead.
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_product_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

// Product is a product.
type Product struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the product.
	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// The name of the product.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The price of the product.
	Price         float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_product_v1_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_v1_product_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_product_v1_product_proto protoreflect.FileDescriptor

var file_product_v1_product_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0x52, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x32, 0x5f, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xad, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x72, 0x61, 0x73, 0x65, 0x72, 0x2d, 0x69, 0x73, 0x62, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x71,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x50, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_product_v1_product_proto_rawDescOnce sync.Once
	file_product_v1_product_proto_rawDescData []byte
)

func file_product_v1_product_proto_rawDescGZIP() []byte {
	file_product_v1_product_proto_rawDescOnce.Do(func() {
		file_product_v1_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_v1_product_proto_rawDesc), len(file_product_v1_product_proto_rawDesc)))
	})
	return file_product_v1_product_proto_rawDescData
}

var file_product_v1_product_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_product_v1_product_proto_goTypes = []any{
	(*GetProductRequest)(nil),  // 0: product.v1.GetProductRequest
	(*GetProductResponse)(nil), // 1: product.v1.GetProductResponse
	(*Product)(nil),            // 2: product.v1.Product
}
var file_product_v1_product_proto_depIdxs = []int32{
	2, // 0: product.v1.GetProductResponse.product:type_name -> product.v1.Product
	0, // 1: product.v1.ProductService.GetProduct:input_type -> product.v1.GetProductRequest
	1, // 2: product.v1.ProductService.GetProduct:output_type -> product.v1.GetProductResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_v1_product_proto_init() }
func file_product_v1_product_proto_init() {
	if File_product_v1_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_v1_product_proto_rawDesc), len(file_product_v1_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_v1_product_proto_goTypes,
		DependencyIndexes: file_product_v1_product_proto_depIdxs,
		MessageInfos:      file_product_v1_product_proto_msgTypes,
	}.Build()
	File_product_v1_product_proto = out.File
	file_product_v1_product_proto_goTypes = nil
	file_product_v1_product_proto_depIdxs = nil
}
