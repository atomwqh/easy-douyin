// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: relation.proto

package relation

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

const (
	Relation_Favorite_FullMethodName         = "/relation.Relation/Favorite"
	Relation_FavoriteList_FullMethodName     = "/relation.Relation/FavoriteList"
	Relation_FollowerList_FullMethodName     = "/relation.Relation/FollowerList"
	Relation_FriendList_FullMethodName       = "/relation.Relation/FriendList"
	Relation_GetFollowCount_FullMethodName   = "/relation.Relation/GetFollowCount"
	Relation_GetFollowerCount_FullMethodName = "/relation.Relation/GetFollowerCount"
	Relation_IsFollowing_FullMethodName      = "/relation.Relation/IsFollowing"
)

// RelationClient is the client API for Relation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationClient interface {
	Favorite(ctx context.Context, in *FavoriteRequest, opts ...grpc.CallOption) (*FavoriteResponse, error)
	FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error)
	FollowerList(ctx context.Context, in *FollowerListReq, opts ...grpc.CallOption) (*FollowerListResp, error)
	FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error)
	GetFollowCount(ctx context.Context, in *FollowCountReq, opts ...grpc.CallOption) (*FollowCountResp, error)
	GetFollowerCount(ctx context.Context, in *FollowerCountReq, opts ...grpc.CallOption) (*FollowerCountResp, error)
	IsFollowing(ctx context.Context, in *IsFollowingReq, opts ...grpc.CallOption) (*IsFollowingResp, error)
}

type relationClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationClient(cc grpc.ClientConnInterface) RelationClient {
	return &relationClient{cc}
}

func (c *relationClient) Favorite(ctx context.Context, in *FavoriteRequest, opts ...grpc.CallOption) (*FavoriteResponse, error) {
	out := new(FavoriteResponse)
	err := c.cc.Invoke(ctx, Relation_Favorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	out := new(FavoriteListResp)
	err := c.cc.Invoke(ctx, Relation_FavoriteList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) FollowerList(ctx context.Context, in *FollowerListReq, opts ...grpc.CallOption) (*FollowerListResp, error) {
	out := new(FollowerListResp)
	err := c.cc.Invoke(ctx, Relation_FollowerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error) {
	out := new(FriendListResp)
	err := c.cc.Invoke(ctx, Relation_FriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GetFollowCount(ctx context.Context, in *FollowCountReq, opts ...grpc.CallOption) (*FollowCountResp, error) {
	out := new(FollowCountResp)
	err := c.cc.Invoke(ctx, Relation_GetFollowCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GetFollowerCount(ctx context.Context, in *FollowerCountReq, opts ...grpc.CallOption) (*FollowerCountResp, error) {
	out := new(FollowerCountResp)
	err := c.cc.Invoke(ctx, Relation_GetFollowerCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) IsFollowing(ctx context.Context, in *IsFollowingReq, opts ...grpc.CallOption) (*IsFollowingResp, error) {
	out := new(IsFollowingResp)
	err := c.cc.Invoke(ctx, Relation_IsFollowing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServer is the server API for Relation service.
// All implementations must embed UnimplementedRelationServer
// for forward compatibility
type RelationServer interface {
	Favorite(context.Context, *FavoriteRequest) (*FavoriteResponse, error)
	FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error)
	FollowerList(context.Context, *FollowerListReq) (*FollowerListResp, error)
	FriendList(context.Context, *FriendListReq) (*FriendListResp, error)
	GetFollowCount(context.Context, *FollowCountReq) (*FollowCountResp, error)
	GetFollowerCount(context.Context, *FollowerCountReq) (*FollowerCountResp, error)
	IsFollowing(context.Context, *IsFollowingReq) (*IsFollowingResp, error)
	mustEmbedUnimplementedRelationServer()
}

// UnimplementedRelationServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServer struct {
}

func (UnimplementedRelationServer) Favorite(context.Context, *FavoriteRequest) (*FavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Favorite not implemented")
}
func (UnimplementedRelationServer) FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedRelationServer) FollowerList(context.Context, *FollowerListReq) (*FollowerListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowerList not implemented")
}
func (UnimplementedRelationServer) FriendList(context.Context, *FriendListReq) (*FriendListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendList not implemented")
}
func (UnimplementedRelationServer) GetFollowCount(context.Context, *FollowCountReq) (*FollowCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowCount not implemented")
}
func (UnimplementedRelationServer) GetFollowerCount(context.Context, *FollowerCountReq) (*FollowerCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerCount not implemented")
}
func (UnimplementedRelationServer) IsFollowing(context.Context, *IsFollowingReq) (*IsFollowingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFollowing not implemented")
}
func (UnimplementedRelationServer) mustEmbedUnimplementedRelationServer() {}

// UnsafeRelationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServer will
// result in compilation errors.
type UnsafeRelationServer interface {
	mustEmbedUnimplementedRelationServer()
}

func RegisterRelationServer(s grpc.ServiceRegistrar, srv RelationServer) {
	s.RegisterService(&Relation_ServiceDesc, srv)
}

func _Relation_Favorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).Favorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_Favorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).Favorite(ctx, req.(*FavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_FavoriteList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).FavoriteList(ctx, req.(*FavoriteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_FollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).FollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_FollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).FollowerList(ctx, req.(*FollowerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_FriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).FriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_FriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).FriendList(ctx, req.(*FriendListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GetFollowCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GetFollowCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_GetFollowCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GetFollowCount(ctx, req.(*FollowCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GetFollowerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GetFollowerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_GetFollowerCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GetFollowerCount(ctx, req.(*FollowerCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_IsFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFollowingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).IsFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relation_IsFollowing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).IsFollowing(ctx, req.(*IsFollowingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Relation_ServiceDesc is the grpc.ServiceDesc for Relation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Relation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relation.Relation",
	HandlerType: (*RelationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Favorite",
			Handler:    _Relation_Favorite_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _Relation_FavoriteList_Handler,
		},
		{
			MethodName: "FollowerList",
			Handler:    _Relation_FollowerList_Handler,
		},
		{
			MethodName: "FriendList",
			Handler:    _Relation_FriendList_Handler,
		},
		{
			MethodName: "GetFollowCount",
			Handler:    _Relation_GetFollowCount_Handler,
		},
		{
			MethodName: "GetFollowerCount",
			Handler:    _Relation_GetFollowerCount_Handler,
		},
		{
			MethodName: "IsFollowing",
			Handler:    _Relation_IsFollowing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation.proto",
}