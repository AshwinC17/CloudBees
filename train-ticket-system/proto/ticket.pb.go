// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: proto/ticket.proto

package proto

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

type TicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *TicketRequest) Reset() {
	*x = TicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketRequest) ProtoMessage() {}

func (x *TicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketRequest.ProtoReflect.Descriptor instead.
func (*TicketRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *TicketRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *TicketRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TicketRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type TicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receipt string `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
}

func (x *TicketResponse) Reset() {
	*x = TicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketResponse) ProtoMessage() {}

func (x *TicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketResponse.ProtoReflect.Descriptor instead.
func (*TicketResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *TicketResponse) GetReceipt() string {
	if x != nil {
		return x.Receipt
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserSeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Section string `protobuf:"bytes,2,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *UserSeat) Reset() {
	*x = UserSeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSeat) ProtoMessage() {}

func (x *UserSeat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSeat.ProtoReflect.Descriptor instead.
func (*UserSeat) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *UserSeat) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserSeat) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type UserSeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserSeat []*UserSeat `protobuf:"bytes,1,rep,name=user_seat,json=userSeat,proto3" json:"user_seat,omitempty"`
}

func (x *UserSeatResponse) Reset() {
	*x = UserSeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSeatResponse) ProtoMessage() {}

func (x *UserSeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSeatResponse.ProtoReflect.Descriptor instead.
func (*UserSeatResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *UserSeatResponse) GetUserSeat() []*UserSeat {
	if x != nil {
		return x.UserSeat
	}
	return nil
}

type SectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *SectionRequest) Reset() {
	*x = SectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionRequest) ProtoMessage() {}

func (x *SectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionRequest.ProtoReflect.Descriptor instead.
func (*SectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *SectionRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type ModifySeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	NewSection string `protobuf:"bytes,2,opt,name=new_section,json=newSection,proto3" json:"new_section,omitempty"`
}

func (x *ModifySeatRequest) Reset() {
	*x = ModifySeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatRequest) ProtoMessage() {}

func (x *ModifySeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatRequest.ProtoReflect.Descriptor instead.
func (*ModifySeatRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *ModifySeatRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ModifySeatRequest) GetNewSection() string {
	if x != nil {
		return x.NewSection
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{7}
}

var File_proto_ticket_proto protoreflect.FileDescriptor

var file_proto_ticket_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0d, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x6f, 0x22, 0x2a, 0x0a, 0x0e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x58, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x45, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x61, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x61, 0x74, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74,
	0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x11,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xae, 0x02, 0x0a,
	0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f,
	0x0a, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42,
	0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0a, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53,
	0x65, 0x61, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x68, 0x77,
	0x69, 0x6e, 0x43, 0x31, 0x37, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2d, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ticket_proto_rawDescOnce sync.Once
	file_proto_ticket_proto_rawDescData = file_proto_ticket_proto_rawDesc
)

func file_proto_ticket_proto_rawDescGZIP() []byte {
	file_proto_ticket_proto_rawDescOnce.Do(func() {
		file_proto_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ticket_proto_rawDescData)
	})
	return file_proto_ticket_proto_rawDescData
}

var file_proto_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_ticket_proto_goTypes = []interface{}{
	(*TicketRequest)(nil),     // 0: proto.TicketRequest
	(*TicketResponse)(nil),    // 1: proto.TicketResponse
	(*User)(nil),              // 2: proto.User
	(*UserSeat)(nil),          // 3: proto.UserSeat
	(*UserSeatResponse)(nil),  // 4: proto.UserSeatResponse
	(*SectionRequest)(nil),    // 5: proto.SectionRequest
	(*ModifySeatRequest)(nil), // 6: proto.ModifySeatRequest
	(*Empty)(nil),             // 7: proto.Empty
}
var file_proto_ticket_proto_depIdxs = []int32{
	2, // 0: proto.TicketRequest.user:type_name -> proto.User
	2, // 1: proto.UserSeat.user:type_name -> proto.User
	3, // 2: proto.UserSeatResponse.user_seat:type_name -> proto.UserSeat
	2, // 3: proto.ModifySeatRequest.user:type_name -> proto.User
	0, // 4: proto.TicketService.PurchaseTicket:input_type -> proto.TicketRequest
	2, // 5: proto.TicketService.GetReceipt:input_type -> proto.User
	5, // 6: proto.TicketService.GetUsersBySection:input_type -> proto.SectionRequest
	2, // 7: proto.TicketService.RemoveUser:input_type -> proto.User
	6, // 8: proto.TicketService.ModifySeat:input_type -> proto.ModifySeatRequest
	1, // 9: proto.TicketService.PurchaseTicket:output_type -> proto.TicketResponse
	1, // 10: proto.TicketService.GetReceipt:output_type -> proto.TicketResponse
	4, // 11: proto.TicketService.GetUsersBySection:output_type -> proto.UserSeatResponse
	7, // 12: proto.TicketService.RemoveUser:output_type -> proto.Empty
	7, // 13: proto.TicketService.ModifySeat:output_type -> proto.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_ticket_proto_init() }
func file_proto_ticket_proto_init() {
	if File_proto_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketRequest); i {
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
		file_proto_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketResponse); i {
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
		file_proto_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSeat); i {
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
		file_proto_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSeatResponse); i {
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
		file_proto_ticket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionRequest); i {
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
		file_proto_ticket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifySeatRequest); i {
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
		file_proto_ticket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ticket_proto_goTypes,
		DependencyIndexes: file_proto_ticket_proto_depIdxs,
		MessageInfos:      file_proto_ticket_proto_msgTypes,
	}.Build()
	File_proto_ticket_proto = out.File
	file_proto_ticket_proto_rawDesc = nil
	file_proto_ticket_proto_goTypes = nil
	file_proto_ticket_proto_depIdxs = nil
}