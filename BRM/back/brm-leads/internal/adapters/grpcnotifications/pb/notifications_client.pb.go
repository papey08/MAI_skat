// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: notifications_client.proto

package pb

import (
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

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId  uint64          `protobuf:"varint,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Type       string          `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Date       int64           `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
	Viewed     bool            `protobuf:"varint,5,opt,name=viewed,proto3" json:"viewed,omitempty"`
	NewLead    *NewLeadInfo    `protobuf:"bytes,6,opt,name=new_lead,json=newLead,proto3" json:"new_lead,omitempty"`
	ClosedLead *ClosedLeadInfo `protobuf:"bytes,7,opt,name=closed_lead,json=closedLead,proto3" json:"closed_lead,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notifications_client_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Notification) GetCompanyId() uint64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *Notification) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Notification) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Notification) GetViewed() bool {
	if x != nil {
		return x.Viewed
	}
	return false
}

func (x *Notification) GetNewLead() *NewLeadInfo {
	if x != nil {
		return x.NewLead
	}
	return nil
}

func (x *Notification) GetClosedLead() *ClosedLeadInfo {
	if x != nil {
		return x.ClosedLead
	}
	return nil
}

type NewLeadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeadId        uint64 `protobuf:"varint,1,opt,name=lead_id,json=leadId,proto3" json:"lead_id,omitempty"`
	ClientCompany uint64 `protobuf:"varint,2,opt,name=client_company,json=clientCompany,proto3" json:"client_company,omitempty"`
}

func (x *NewLeadInfo) Reset() {
	*x = NewLeadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewLeadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewLeadInfo) ProtoMessage() {}

func (x *NewLeadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewLeadInfo.ProtoReflect.Descriptor instead.
func (*NewLeadInfo) Descriptor() ([]byte, []int) {
	return file_notifications_client_proto_rawDescGZIP(), []int{1}
}

func (x *NewLeadInfo) GetLeadId() uint64 {
	if x != nil {
		return x.LeadId
	}
	return 0
}

func (x *NewLeadInfo) GetClientCompany() uint64 {
	if x != nil {
		return x.ClientCompany
	}
	return 0
}

type ClosedLeadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId            uint64 `protobuf:"varint,1,opt,name=ad_id,json=adId,proto3" json:"ad_id,omitempty"`
	ProducerCompany uint64 `protobuf:"varint,2,opt,name=producer_company,json=producerCompany,proto3" json:"producer_company,omitempty"`
	Answered        bool   `protobuf:"varint,3,opt,name=answered,proto3" json:"answered,omitempty"`
}

func (x *ClosedLeadInfo) Reset() {
	*x = ClosedLeadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClosedLeadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClosedLeadInfo) ProtoMessage() {}

func (x *ClosedLeadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClosedLeadInfo.ProtoReflect.Descriptor instead.
func (*ClosedLeadInfo) Descriptor() ([]byte, []int) {
	return file_notifications_client_proto_rawDescGZIP(), []int{2}
}

func (x *ClosedLeadInfo) GetAdId() uint64 {
	if x != nil {
		return x.AdId
	}
	return 0
}

func (x *ClosedLeadInfo) GetProducerCompany() uint64 {
	if x != nil {
		return x.ProducerCompany
	}
	return 0
}

func (x *ClosedLeadInfo) GetAnswered() bool {
	if x != nil {
		return x.Answered
	}
	return false
}

type CreateNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *CreateNotificationRequest) Reset() {
	*x = CreateNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotificationRequest) ProtoMessage() {}

func (x *CreateNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotificationRequest.ProtoReflect.Descriptor instead.
func (*CreateNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notifications_client_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNotificationRequest) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

var File_notifications_client_proto protoreflect.FileDescriptor

var file_notifications_client_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x64,
	0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0,
	0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x12, 0x2b,
	0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x64, 0x12, 0x34, 0x0a, 0x0b, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x4c, 0x65, 0x61,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x4c, 0x65, 0x61,
	0x64, 0x22, 0x4d, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x22, 0x6c, 0x0a, 0x0e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x61, 0x64, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x64, 0x22, 0x52,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0c, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x66, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifications_client_proto_rawDescOnce sync.Once
	file_notifications_client_proto_rawDescData = file_notifications_client_proto_rawDesc
)

func file_notifications_client_proto_rawDescGZIP() []byte {
	file_notifications_client_proto_rawDescOnce.Do(func() {
		file_notifications_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifications_client_proto_rawDescData)
	})
	return file_notifications_client_proto_rawDescData
}

var file_notifications_client_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_notifications_client_proto_goTypes = []interface{}{
	(*Notification)(nil),              // 0: ads.Notification
	(*NewLeadInfo)(nil),               // 1: ads.NewLeadInfo
	(*ClosedLeadInfo)(nil),            // 2: ads.ClosedLeadInfo
	(*CreateNotificationRequest)(nil), // 3: ads.CreateNotificationRequest
	(*emptypb.Empty)(nil),             // 4: google.protobuf.Empty
}
var file_notifications_client_proto_depIdxs = []int32{
	1, // 0: ads.Notification.new_lead:type_name -> ads.NewLeadInfo
	2, // 1: ads.Notification.closed_lead:type_name -> ads.ClosedLeadInfo
	0, // 2: ads.CreateNotificationRequest.notification:type_name -> ads.Notification
	3, // 3: ads.NotificationsService.CreateNotification:input_type -> ads.CreateNotificationRequest
	4, // 4: ads.NotificationsService.CreateNotification:output_type -> google.protobuf.Empty
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_notifications_client_proto_init() }
func file_notifications_client_proto_init() {
	if File_notifications_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifications_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_notifications_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewLeadInfo); i {
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
		file_notifications_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClosedLeadInfo); i {
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
		file_notifications_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNotificationRequest); i {
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
			RawDescriptor: file_notifications_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notifications_client_proto_goTypes,
		DependencyIndexes: file_notifications_client_proto_depIdxs,
		MessageInfos:      file_notifications_client_proto_msgTypes,
	}.Build()
	File_notifications_client_proto = out.File
	file_notifications_client_proto_rawDesc = nil
	file_notifications_client_proto_goTypes = nil
	file_notifications_client_proto_depIdxs = nil
}
