// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package peersv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PeerDiscoveryServiceClient is the client API for PeerDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerDiscoveryServiceClient interface {
	GetPeers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Peers, error)
	GetPeer(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*PeerInfo, error)
	AddPeer(ctx context.Context, in *PeerInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemovePeer(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type peerDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerDiscoveryServiceClient(cc grpc.ClientConnInterface) PeerDiscoveryServiceClient {
	return &peerDiscoveryServiceClient{cc}
}

func (c *peerDiscoveryServiceClient) GetPeers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Peers, error) {
	out := new(Peers)
	err := c.cc.Invoke(ctx, "/aperture.common.peers.v1.PeerDiscoveryService/GetPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerDiscoveryServiceClient) GetPeer(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*PeerInfo, error) {
	out := new(PeerInfo)
	err := c.cc.Invoke(ctx, "/aperture.common.peers.v1.PeerDiscoveryService/GetPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerDiscoveryServiceClient) AddPeer(ctx context.Context, in *PeerInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/aperture.common.peers.v1.PeerDiscoveryService/AddPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerDiscoveryServiceClient) RemovePeer(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/aperture.common.peers.v1.PeerDiscoveryService/RemovePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerDiscoveryServiceServer is the server API for PeerDiscoveryService service.
// All implementations should embed UnimplementedPeerDiscoveryServiceServer
// for forward compatibility
type PeerDiscoveryServiceServer interface {
	GetPeers(context.Context, *emptypb.Empty) (*Peers, error)
	GetPeer(context.Context, *PeerRequest) (*PeerInfo, error)
	AddPeer(context.Context, *PeerInfo) (*emptypb.Empty, error)
	RemovePeer(context.Context, *PeerRequest) (*emptypb.Empty, error)
}

// UnimplementedPeerDiscoveryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPeerDiscoveryServiceServer struct {
}

func (UnimplementedPeerDiscoveryServiceServer) GetPeers(context.Context, *emptypb.Empty) (*Peers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeers not implemented")
}
func (UnimplementedPeerDiscoveryServiceServer) GetPeer(context.Context, *PeerRequest) (*PeerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeer not implemented")
}
func (UnimplementedPeerDiscoveryServiceServer) AddPeer(context.Context, *PeerInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeer not implemented")
}
func (UnimplementedPeerDiscoveryServiceServer) RemovePeer(context.Context, *PeerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePeer not implemented")
}

// UnsafePeerDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerDiscoveryServiceServer will
// result in compilation errors.
type UnsafePeerDiscoveryServiceServer interface {
	mustEmbedUnimplementedPeerDiscoveryServiceServer()
}

func RegisterPeerDiscoveryServiceServer(s grpc.ServiceRegistrar, srv PeerDiscoveryServiceServer) {
	s.RegisterService(&PeerDiscoveryService_ServiceDesc, srv)
}

func _PeerDiscoveryService_GetPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerDiscoveryServiceServer).GetPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aperture.common.peers.v1.PeerDiscoveryService/GetPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerDiscoveryServiceServer).GetPeers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerDiscoveryService_GetPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerDiscoveryServiceServer).GetPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aperture.common.peers.v1.PeerDiscoveryService/GetPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerDiscoveryServiceServer).GetPeer(ctx, req.(*PeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerDiscoveryService_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerDiscoveryServiceServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aperture.common.peers.v1.PeerDiscoveryService/AddPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerDiscoveryServiceServer).AddPeer(ctx, req.(*PeerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerDiscoveryService_RemovePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerDiscoveryServiceServer).RemovePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aperture.common.peers.v1.PeerDiscoveryService/RemovePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerDiscoveryServiceServer).RemovePeer(ctx, req.(*PeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PeerDiscoveryService_ServiceDesc is the grpc.ServiceDesc for PeerDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeerDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aperture.common.peers.v1.PeerDiscoveryService",
	HandlerType: (*PeerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPeers",
			Handler:    _PeerDiscoveryService_GetPeers_Handler,
		},
		{
			MethodName: "GetPeer",
			Handler:    _PeerDiscoveryService_GetPeer_Handler,
		},
		{
			MethodName: "AddPeer",
			Handler:    _PeerDiscoveryService_AddPeer_Handler,
		},
		{
			MethodName: "RemovePeer",
			Handler:    _PeerDiscoveryService_RemovePeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aperture/common/peers/v1/peers.proto",
}
