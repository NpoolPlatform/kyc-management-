// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/kyc-management.proto

package npool

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// request body and response
type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,100,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{0}
}

func (x *VersionResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type KycInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                  string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID              string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	AppID               string `protobuf:"bytes,14,opt,name=AppID,proto3" json:"AppID,omitempty"`
	FirstName           string `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName            string `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Region              string `protobuf:"bytes,5,opt,name=Region,proto3" json:"Region,omitempty"`
	CardType            string `protobuf:"bytes,6,opt,name=CardType,proto3" json:"CardType,omitempty"`
	CardID              string `protobuf:"bytes,7,opt,name=CardID,proto3" json:"CardID,omitempty"`
	FrontCardImg        string `protobuf:"bytes,8,opt,name=FrontCardImg,proto3" json:"FrontCardImg,omitempty"`
	BackCardImg         string `protobuf:"bytes,9,opt,name=BackCardImg,proto3" json:"BackCardImg,omitempty"`
	UserHandlingCardImg string `protobuf:"bytes,10,opt,name=UserHandlingCardImg,proto3" json:"UserHandlingCardImg,omitempty"`
	ReviewStatus        uint32 `protobuf:"varint,11,opt,name=ReviewStatus,proto3" json:"ReviewStatus,omitempty"`
	CreateAT            uint32 `protobuf:"varint,12,opt,name=CreateAT,proto3" json:"CreateAT,omitempty"`
	UpdateAT            uint32 `protobuf:"varint,13,opt,name=UpdateAT,proto3" json:"UpdateAT,omitempty"`
}

func (x *KycInfo) Reset() {
	*x = KycInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KycInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KycInfo) ProtoMessage() {}

func (x *KycInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KycInfo.ProtoReflect.Descriptor instead.
func (*KycInfo) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{1}
}

func (x *KycInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *KycInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *KycInfo) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *KycInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *KycInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *KycInfo) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *KycInfo) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *KycInfo) GetCardID() string {
	if x != nil {
		return x.CardID
	}
	return ""
}

func (x *KycInfo) GetFrontCardImg() string {
	if x != nil {
		return x.FrontCardImg
	}
	return ""
}

func (x *KycInfo) GetBackCardImg() string {
	if x != nil {
		return x.BackCardImg
	}
	return ""
}

func (x *KycInfo) GetUserHandlingCardImg() string {
	if x != nil {
		return x.UserHandlingCardImg
	}
	return ""
}

func (x *KycInfo) GetReviewStatus() uint32 {
	if x != nil {
		return x.ReviewStatus
	}
	return 0
}

func (x *KycInfo) GetCreateAT() uint32 {
	if x != nil {
		return x.CreateAT
	}
	return 0
}

func (x *KycInfo) GetUpdateAT() uint32 {
	if x != nil {
		return x.UpdateAT
	}
	return 0
}

type CreateKycRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *KycInfo `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateKycRecordRequest) Reset() {
	*x = CreateKycRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKycRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKycRecordRequest) ProtoMessage() {}

func (x *CreateKycRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKycRecordRequest.ProtoReflect.Descriptor instead.
func (*CreateKycRecordRequest) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{2}
}

func (x *CreateKycRecordRequest) GetInfo() *KycInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateKycRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *KycInfo `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateKycRecordResponse) Reset() {
	*x = CreateKycRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKycRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKycRecordResponse) ProtoMessage() {}

func (x *CreateKycRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKycRecordResponse.ProtoReflect.Descriptor instead.
func (*CreateKycRecordResponse) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{3}
}

func (x *CreateKycRecordResponse) GetInfo() *KycInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAllKycInfosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string   `protobuf:"bytes,1,opt,name=AppID,proto3" json:"AppID,omitempty"`
	KycIDs []string `protobuf:"bytes,2,rep,name=KycIDs,proto3" json:"KycIDs,omitempty"`
}

func (x *GetAllKycInfosRequest) Reset() {
	*x = GetAllKycInfosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllKycInfosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllKycInfosRequest) ProtoMessage() {}

func (x *GetAllKycInfosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllKycInfosRequest.ProtoReflect.Descriptor instead.
func (*GetAllKycInfosRequest) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllKycInfosRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetAllKycInfosRequest) GetKycIDs() []string {
	if x != nil {
		return x.KycIDs
	}
	return nil
}

type GetAllKycInfosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*KycInfo `protobuf:"bytes,1,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetAllKycInfosResponse) Reset() {
	*x = GetAllKycInfosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllKycInfosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllKycInfosResponse) ProtoMessage() {}

func (x *GetAllKycInfosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllKycInfosResponse.ProtoReflect.Descriptor instead.
func (*GetAllKycInfosResponse) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllKycInfosResponse) GetInfos() []*KycInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type UpdateKycRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *KycInfo `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateKycRequest) Reset() {
	*x = UpdateKycRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateKycRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateKycRequest) ProtoMessage() {}

func (x *UpdateKycRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateKycRequest.ProtoReflect.Descriptor instead.
func (*UpdateKycRequest) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateKycRequest) GetInfo() *KycInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateKycResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *KycInfo `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateKycResponse) Reset() {
	*x = UpdateKycResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateKycResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateKycResponse) ProtoMessage() {}

func (x *UpdateKycResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateKycResponse.ProtoReflect.Descriptor instead.
func (*UpdateKycResponse) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateKycResponse) GetInfo() *KycInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UploadKycImgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ImgType   string `protobuf:"bytes,2,opt,name=ImgType,proto3" json:"ImgType,omitempty"`
	ImgBase64 string `protobuf:"bytes,3,opt,name=ImgBase64,proto3" json:"ImgBase64,omitempty"`
	AppID     string `protobuf:"bytes,4,opt,name=AppID,proto3" json:"AppID,omitempty"`
}

func (x *UploadKycImgRequest) Reset() {
	*x = UploadKycImgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadKycImgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadKycImgRequest) ProtoMessage() {}

func (x *UploadKycImgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadKycImgRequest.ProtoReflect.Descriptor instead.
func (*UploadKycImgRequest) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{8}
}

func (x *UploadKycImgRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UploadKycImgRequest) GetImgType() string {
	if x != nil {
		return x.ImgType
	}
	return ""
}

func (x *UploadKycImgRequest) GetImgBase64() string {
	if x != nil {
		return x.ImgBase64
	}
	return ""
}

func (x *UploadKycImgRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

type UploadKycImgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// return ImgID
	Info string `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UploadKycImgResponse) Reset() {
	*x = UploadKycImgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadKycImgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadKycImgResponse) ProtoMessage() {}

func (x *UploadKycImgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadKycImgResponse.ProtoReflect.Descriptor instead.
func (*UploadKycImgResponse) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{9}
}

func (x *UploadKycImgResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type GetKycImgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID string `protobuf:"bytes,1,opt,name=AppID,proto3" json:"AppID,omitempty"`
	ImgID string `protobuf:"bytes,2,opt,name=ImgID,proto3" json:"ImgID,omitempty"`
}

func (x *GetKycImgRequest) Reset() {
	*x = GetKycImgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKycImgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKycImgRequest) ProtoMessage() {}

func (x *GetKycImgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKycImgRequest.ProtoReflect.Descriptor instead.
func (*GetKycImgRequest) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{10}
}

func (x *GetKycImgRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetKycImgRequest) GetImgID() string {
	if x != nil {
		return x.ImgID
	}
	return ""
}

type GetKycImgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// return ImgBase64
	Info string `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetKycImgResponse) Reset() {
	*x = GetKycImgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_kyc_management_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKycImgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKycImgResponse) ProtoMessage() {}

func (x *GetKycImgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_kyc_management_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKycImgResponse.ProtoReflect.Descriptor instead.
func (*GetKycImgResponse) Descriptor() ([]byte, []int) {
	return file_npool_kyc_management_proto_rawDescGZIP(), []int{11}
}

func (x *GetKycImgResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_npool_kyc_management_proto protoreflect.FileDescriptor

var file_npool_kyc_management_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6b, 0x79, 0x63, 0x2d, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6b, 0x79,
	0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0xa1, 0x03, 0x0a, 0x07, 0x4b, 0x79, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x72,
	0x64, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x61, 0x72, 0x64, 0x49,
	0x44, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6d,
	0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x49, 0x6d, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x61, 0x63, 0x6b, 0x43, 0x61, 0x72,
	0x64, 0x49, 0x6d, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x61, 0x63, 0x6b,
	0x43, 0x61, 0x72, 0x64, 0x49, 0x6d, 0x67, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6d, 0x67, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69,
	0x6e, 0x67, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6d, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x54, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x54, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x54, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x54, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b,
	0x79, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x79, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x49, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x79, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x45, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x79, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4b, 0x79, 0x63,
	0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x4b, 0x79, 0x63, 0x49, 0x44,
	0x73, 0x22, 0x4a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x79, 0x63, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x79, 0x63,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4b,
	0x79, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x42, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x79, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x43, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x79, 0x63, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x7b, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x4b, 0x79, 0x63, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6d, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x6d, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x67, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x49, 0x6d, 0x67, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x12, 0x14, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x22, 0x2a, 0x0a, 0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x79, 0x63,
	0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4b, 0x79, 0x63, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x67,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x67, 0x49, 0x44, 0x22,
	0x27, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4b, 0x79, 0x63, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xe9, 0x05, 0x0a, 0x0d, 0x4b, 0x79, 0x63,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5a, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e,
	0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x8a, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x2e, 0x6b, 0x79, 0x63,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4b, 0x79, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x6b, 0x79, 0x63, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x3a, 0x01, 0x2a, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x79,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x28, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4b, 0x79, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x79, 0x63, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x6c, 0x6c,
	0x2f, 0x6b, 0x79, 0x63, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x12, 0x23, 0x2e, 0x6b, 0x79, 0x63,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x6b, 0x79, 0x63, 0x3a, 0x01, 0x2a,
	0x12, 0x7e, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x79, 0x63, 0x49, 0x6d, 0x67,
	0x12, 0x26, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4b, 0x79, 0x63, 0x49, 0x6d,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x4b, 0x79, 0x63, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x6b, 0x79, 0x63, 0x2f, 0x69, 0x6d, 0x67, 0x3a, 0x01, 0x2a,
	0x12, 0x72, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4b, 0x79, 0x63, 0x49, 0x6d, 0x67, 0x12, 0x23, 0x2e,
	0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x79, 0x63, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x79, 0x63, 0x49, 0x6d, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6b, 0x79, 0x63, 0x2f, 0x69, 0x6d,
	0x67, 0x3a, 0x01, 0x2a, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6b, 0x79, 0x63, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_kyc_management_proto_rawDescOnce sync.Once
	file_npool_kyc_management_proto_rawDescData = file_npool_kyc_management_proto_rawDesc
)

func file_npool_kyc_management_proto_rawDescGZIP() []byte {
	file_npool_kyc_management_proto_rawDescOnce.Do(func() {
		file_npool_kyc_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_kyc_management_proto_rawDescData)
	})
	return file_npool_kyc_management_proto_rawDescData
}

var file_npool_kyc_management_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_npool_kyc_management_proto_goTypes = []interface{}{
	(*VersionResponse)(nil),         // 0: kyc.management.v1.VersionResponse
	(*KycInfo)(nil),                 // 1: kyc.management.v1.KycInfo
	(*CreateKycRecordRequest)(nil),  // 2: kyc.management.v1.CreateKycRecordRequest
	(*CreateKycRecordResponse)(nil), // 3: kyc.management.v1.CreateKycRecordResponse
	(*GetAllKycInfosRequest)(nil),   // 4: kyc.management.v1.GetAllKycInfosRequest
	(*GetAllKycInfosResponse)(nil),  // 5: kyc.management.v1.GetAllKycInfosResponse
	(*UpdateKycRequest)(nil),        // 6: kyc.management.v1.UpdateKycRequest
	(*UpdateKycResponse)(nil),       // 7: kyc.management.v1.UpdateKycResponse
	(*UploadKycImgRequest)(nil),     // 8: kyc.management.v1.UploadKycImgRequest
	(*UploadKycImgResponse)(nil),    // 9: kyc.management.v1.UploadKycImgResponse
	(*GetKycImgRequest)(nil),        // 10: kyc.management.v1.GetKycImgRequest
	(*GetKycImgResponse)(nil),       // 11: kyc.management.v1.GetKycImgResponse
	(*emptypb.Empty)(nil),           // 12: google.protobuf.Empty
}
var file_npool_kyc_management_proto_depIdxs = []int32{
	1,  // 0: kyc.management.v1.CreateKycRecordRequest.Info:type_name -> kyc.management.v1.KycInfo
	1,  // 1: kyc.management.v1.CreateKycRecordResponse.Info:type_name -> kyc.management.v1.KycInfo
	1,  // 2: kyc.management.v1.GetAllKycInfosResponse.Infos:type_name -> kyc.management.v1.KycInfo
	1,  // 3: kyc.management.v1.UpdateKycRequest.Info:type_name -> kyc.management.v1.KycInfo
	1,  // 4: kyc.management.v1.UpdateKycResponse.Info:type_name -> kyc.management.v1.KycInfo
	12, // 5: kyc.management.v1.KycManagement.Version:input_type -> google.protobuf.Empty
	2,  // 6: kyc.management.v1.KycManagement.CreateKycRecord:input_type -> kyc.management.v1.CreateKycRecordRequest
	4,  // 7: kyc.management.v1.KycManagement.GetAllKycInfos:input_type -> kyc.management.v1.GetAllKycInfosRequest
	6,  // 8: kyc.management.v1.KycManagement.UpdateKyc:input_type -> kyc.management.v1.UpdateKycRequest
	8,  // 9: kyc.management.v1.KycManagement.UploadKycImg:input_type -> kyc.management.v1.UploadKycImgRequest
	10, // 10: kyc.management.v1.KycManagement.GetKycImg:input_type -> kyc.management.v1.GetKycImgRequest
	0,  // 11: kyc.management.v1.KycManagement.Version:output_type -> kyc.management.v1.VersionResponse
	3,  // 12: kyc.management.v1.KycManagement.CreateKycRecord:output_type -> kyc.management.v1.CreateKycRecordResponse
	5,  // 13: kyc.management.v1.KycManagement.GetAllKycInfos:output_type -> kyc.management.v1.GetAllKycInfosResponse
	7,  // 14: kyc.management.v1.KycManagement.UpdateKyc:output_type -> kyc.management.v1.UpdateKycResponse
	9,  // 15: kyc.management.v1.KycManagement.UploadKycImg:output_type -> kyc.management.v1.UploadKycImgResponse
	11, // 16: kyc.management.v1.KycManagement.GetKycImg:output_type -> kyc.management.v1.GetKycImgResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_npool_kyc_management_proto_init() }
func file_npool_kyc_management_proto_init() {
	if File_npool_kyc_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_kyc_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_npool_kyc_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KycInfo); i {
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
		file_npool_kyc_management_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKycRecordRequest); i {
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
		file_npool_kyc_management_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKycRecordResponse); i {
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
		file_npool_kyc_management_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllKycInfosRequest); i {
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
		file_npool_kyc_management_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllKycInfosResponse); i {
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
		file_npool_kyc_management_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateKycRequest); i {
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
		file_npool_kyc_management_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateKycResponse); i {
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
		file_npool_kyc_management_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadKycImgRequest); i {
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
		file_npool_kyc_management_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadKycImgResponse); i {
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
		file_npool_kyc_management_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKycImgRequest); i {
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
		file_npool_kyc_management_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKycImgResponse); i {
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
			RawDescriptor: file_npool_kyc_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_kyc_management_proto_goTypes,
		DependencyIndexes: file_npool_kyc_management_proto_depIdxs,
		MessageInfos:      file_npool_kyc_management_proto_msgTypes,
	}.Build()
	File_npool_kyc_management_proto = out.File
	file_npool_kyc_management_proto_rawDesc = nil
	file_npool_kyc_management_proto_goTypes = nil
	file_npool_kyc_management_proto_depIdxs = nil
}
