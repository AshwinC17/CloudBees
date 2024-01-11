// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceClient interface {
	PurchaseTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketResponse, error)
	GetReceipt(ctx context.Context, in *User, opts ...grpc.CallOption) (*TicketResponse, error)
	GetUsersBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*UserSeatResponse, error)
	RemoveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*Empty, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) PurchaseTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketResponse, error) {
	out := new(TicketResponse)
	err := c.cc.Invoke(ctx, "/proto.TicketService/PurchaseTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetReceipt(ctx context.Context, in *User, opts ...grpc.CallOption) (*TicketResponse, error) {
	out := new(TicketResponse)
	err := c.cc.Invoke(ctx, "/proto.TicketService/GetReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetUsersBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*UserSeatResponse, error) {
	out := new(UserSeatResponse)
	err := c.cc.Invoke(ctx, "/proto.TicketService/GetUsersBySection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) RemoveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.TicketService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.TicketService/ModifySeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations must embed UnimplementedTicketServiceServer
// for forward compatibility
type TicketServiceServer interface {
	PurchaseTicket(context.Context, *TicketRequest) (*TicketResponse, error)
	GetReceipt(context.Context, *User) (*TicketResponse, error)
	GetUsersBySection(context.Context, *SectionRequest) (*UserSeatResponse, error)
	RemoveUser(context.Context, *User) (*Empty, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*Empty, error)
	mustEmbedUnimplementedTicketServiceServer()
}

// UnimplementedTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTicketServiceServer struct {
}

func (UnimplementedTicketServiceServer) PurchaseTicket(context.Context, *TicketRequest) (*TicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTicketServiceServer) GetReceipt(context.Context, *User) (*TicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedTicketServiceServer) GetUsersBySection(context.Context, *SectionRequest) (*UserSeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersBySection not implemented")
}
func (UnimplementedTicketServiceServer) RemoveUser(context.Context, *User) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTicketServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTicketServiceServer) mustEmbedUnimplementedTicketServiceServer() {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TicketService/PurchaseTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).PurchaseTicket(ctx, req.(*TicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TicketService/GetReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetReceipt(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetUsersBySection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetUsersBySection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TicketService/GetUsersBySection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetUsersBySection(ctx, req.(*SectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TicketService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).RemoveUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TicketService/ModifySeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TicketService_PurchaseTicket_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _TicketService_GetReceipt_Handler,
		},
		{
			MethodName: "GetUsersBySection",
			Handler:    _TicketService_GetUsersBySection_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TicketService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TicketService_ModifySeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ticket.proto",
}
