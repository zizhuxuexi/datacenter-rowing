// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: rowDataService.proto

package rowdata

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
	RowDataService_TrainingSummaryAdd_FullMethodName                  = "/idl.RowDataService/TrainingSummaryAdd"
	RowDataService_TrainingSummaryGetByTrainingName_FullMethodName    = "/idl.RowDataService/TrainingSummaryGetByTrainingName"
	RowDataService_TrainingSummaryGetByTrainDate_FullMethodName       = "/idl.RowDataService/TrainingSummaryGetByTrainDate"
	RowDataService_TrainingSummaryGetByEvent_FullMethodName           = "/idl.RowDataService/TrainingSummaryGetByEvent"
	RowDataService_TrainingSummaryGet_FullMethodName                  = "/idl.RowDataService/TrainingSummaryGet"
	RowDataService_AthleteTrainingDataAdd_FullMethodName              = "/idl.RowDataService/AthleteTrainingDataAdd"
	RowDataService_AthleteTrainingDataGetByName_FullMethodName        = "/idl.RowDataService/AthleteTrainingDataGetByName"
	RowDataService_AthleteTrainingDataGet_FullMethodName              = "/idl.RowDataService/AthleteTrainingDataGet"
	RowDataService_SampleMetricsAdd_FullMethodName                    = "/idl.RowDataService/SampleMetricsAdd"
	RowDataService_SampleMetricsGetByAthleteTrainingId_FullMethodName = "/idl.RowDataService/SampleMetricsGetByAthleteTrainingId"
)

// RowDataServiceClient is the client API for RowDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RowDataServiceClient interface {
	TrainingSummaryAdd(ctx context.Context, in *TrainingSummaryAddRequest, opts ...grpc.CallOption) (*TrainingSummaryAddResponse, error)
	TrainingSummaryGetByTrainingName(ctx context.Context, in *TrainingSummaryGetByTrainingNameRequest, opts ...grpc.CallOption) (*TrainingSummaryResponse, error)
	TrainingSummaryGetByTrainDate(ctx context.Context, in *TrainingSummaryGetByTrainDateRequest, opts ...grpc.CallOption) (*TrainingSummaryResponse, error)
	TrainingSummaryGetByEvent(ctx context.Context, in *TrainingSummaryGetByEventRequest, opts ...grpc.CallOption) (*TrainingSummaryResponse, error)
	TrainingSummaryGet(ctx context.Context, in *TrainingSummaryGetRequset, opts ...grpc.CallOption) (*TrainingSummaryGetResponse, error)
	AthleteTrainingDataAdd(ctx context.Context, in *AthleteTrainingDataAddRequest, opts ...grpc.CallOption) (*AthleteTrainingDataAddResponse, error)
	AthleteTrainingDataGetByName(ctx context.Context, in *AthleteTrainingDataGetByName, opts ...grpc.CallOption) (*AthleteTrainingDataResponse, error)
	AthleteTrainingDataGet(ctx context.Context, in *AthleteTrainingDataGetRequest, opts ...grpc.CallOption) (*AthleteTrainingDataGetResponse, error)
	SampleMetricsAdd(ctx context.Context, in *SampleMetricsModel, opts ...grpc.CallOption) (*SampleMetricsAddResponse, error)
	SampleMetricsGetByAthleteTrainingId(ctx context.Context, in *SampleMetricsGetByAthleteTrainingIdRequest, opts ...grpc.CallOption) (*SampleMetricsGetByAthleteTrainingIdResponse, error)
}

type rowDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRowDataServiceClient(cc grpc.ClientConnInterface) RowDataServiceClient {
	return &rowDataServiceClient{cc}
}

func (c *rowDataServiceClient) TrainingSummaryAdd(ctx context.Context, in *TrainingSummaryAddRequest, opts ...grpc.CallOption) (*TrainingSummaryAddResponse, error) {
	out := new(TrainingSummaryAddResponse)
	err := c.cc.Invoke(ctx, RowDataService_TrainingSummaryAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rowDataServiceClient) TrainingSummaryGetByTrainingName(ctx context.Context, in *TrainingSummaryGetByTrainingNameRequest, opts ...grpc.CallOption) (*TrainingSummaryResponse, error) {
	out := new(TrainingSummaryResponse)
	err := c.cc.Invoke(ctx, RowDataService_TrainingSummaryGetByTrainingName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rowDataServiceClient) TrainingSummaryGetByTrainDate(ctx context.Context, in *TrainingSummaryGetByTrainDateRequest, opts ...grpc.CallOption) (*TrainingSummaryResponse, error) {
	out := new(TrainingSummaryResponse)
	err := c.cc.Invoke(ctx, RowDataService_TrainingSummaryGetByTrainDate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rowDataServiceClient) TrainingSummaryGetByEvent(ctx context.Context, in *TrainingSummaryGetByEventRequest, opts ...grpc.CallOption) (*TrainingSummaryResponse, error) {
	out := new(TrainingSummaryResponse)
	err := c.cc.Invoke(ctx, RowDataService_TrainingSummaryGetByEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rowDataServiceClient) TrainingSummaryGet(ctx context.Context, in *TrainingSummaryGetRequset, opts ...grpc.CallOption) (*TrainingSummaryGetResponse, error) {
	out := new(TrainingSummaryGetResponse)
	err := c.cc.Invoke(ctx, RowDataService_TrainingSummaryGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rowDataServiceClient) AthleteTrainingDataAdd(ctx context.Context, in *AthleteTrainingDataAddRequest, opts ...grpc.CallOption) (*AthleteTrainingDataAddResponse, error) {
	out := new(AthleteTrainingDataAddResponse)
	err := c.cc.Invoke(ctx, RowDataService_AthleteTrainingDataAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rowDataServiceClient) AthleteTrainingDataGetByName(ctx context.Context, in *AthleteTrainingDataGetByName, opts ...grpc.CallOption) (*AthleteTrainingDataResponse, error) {
	out := new(AthleteTrainingDataResponse)
	err := c.cc.Invoke(ctx, RowDataService_AthleteTrainingDataGetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rowDataServiceClient) AthleteTrainingDataGet(ctx context.Context, in *AthleteTrainingDataGetRequest, opts ...grpc.CallOption) (*AthleteTrainingDataGetResponse, error) {
	out := new(AthleteTrainingDataGetResponse)
	err := c.cc.Invoke(ctx, RowDataService_AthleteTrainingDataGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rowDataServiceClient) SampleMetricsAdd(ctx context.Context, in *SampleMetricsModel, opts ...grpc.CallOption) (*SampleMetricsAddResponse, error) {
	out := new(SampleMetricsAddResponse)
	err := c.cc.Invoke(ctx, RowDataService_SampleMetricsAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rowDataServiceClient) SampleMetricsGetByAthleteTrainingId(ctx context.Context, in *SampleMetricsGetByAthleteTrainingIdRequest, opts ...grpc.CallOption) (*SampleMetricsGetByAthleteTrainingIdResponse, error) {
	out := new(SampleMetricsGetByAthleteTrainingIdResponse)
	err := c.cc.Invoke(ctx, RowDataService_SampleMetricsGetByAthleteTrainingId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RowDataServiceServer is the server API for RowDataService service.
// All implementations must embed UnimplementedRowDataServiceServer
// for forward compatibility
type RowDataServiceServer interface {
	TrainingSummaryAdd(context.Context, *TrainingSummaryAddRequest) (*TrainingSummaryAddResponse, error)
	TrainingSummaryGetByTrainingName(context.Context, *TrainingSummaryGetByTrainingNameRequest) (*TrainingSummaryResponse, error)
	TrainingSummaryGetByTrainDate(context.Context, *TrainingSummaryGetByTrainDateRequest) (*TrainingSummaryResponse, error)
	TrainingSummaryGetByEvent(context.Context, *TrainingSummaryGetByEventRequest) (*TrainingSummaryResponse, error)
	TrainingSummaryGet(context.Context, *TrainingSummaryGetRequset) (*TrainingSummaryGetResponse, error)
	AthleteTrainingDataAdd(context.Context, *AthleteTrainingDataAddRequest) (*AthleteTrainingDataAddResponse, error)
	AthleteTrainingDataGetByName(context.Context, *AthleteTrainingDataGetByName) (*AthleteTrainingDataResponse, error)
	AthleteTrainingDataGet(context.Context, *AthleteTrainingDataGetRequest) (*AthleteTrainingDataGetResponse, error)
	SampleMetricsAdd(context.Context, *SampleMetricsModel) (*SampleMetricsAddResponse, error)
	SampleMetricsGetByAthleteTrainingId(context.Context, *SampleMetricsGetByAthleteTrainingIdRequest) (*SampleMetricsGetByAthleteTrainingIdResponse, error)
	mustEmbedUnimplementedRowDataServiceServer()
}

// UnimplementedRowDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRowDataServiceServer struct {
}

func (UnimplementedRowDataServiceServer) TrainingSummaryAdd(context.Context, *TrainingSummaryAddRequest) (*TrainingSummaryAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainingSummaryAdd not implemented")
}
func (UnimplementedRowDataServiceServer) TrainingSummaryGetByTrainingName(context.Context, *TrainingSummaryGetByTrainingNameRequest) (*TrainingSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainingSummaryGetByTrainingName not implemented")
}
func (UnimplementedRowDataServiceServer) TrainingSummaryGetByTrainDate(context.Context, *TrainingSummaryGetByTrainDateRequest) (*TrainingSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainingSummaryGetByTrainDate not implemented")
}
func (UnimplementedRowDataServiceServer) TrainingSummaryGetByEvent(context.Context, *TrainingSummaryGetByEventRequest) (*TrainingSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainingSummaryGetByEvent not implemented")
}
func (UnimplementedRowDataServiceServer) TrainingSummaryGet(context.Context, *TrainingSummaryGetRequset) (*TrainingSummaryGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainingSummaryGet not implemented")
}
func (UnimplementedRowDataServiceServer) AthleteTrainingDataAdd(context.Context, *AthleteTrainingDataAddRequest) (*AthleteTrainingDataAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AthleteTrainingDataAdd not implemented")
}
func (UnimplementedRowDataServiceServer) AthleteTrainingDataGetByName(context.Context, *AthleteTrainingDataGetByName) (*AthleteTrainingDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AthleteTrainingDataGetByName not implemented")
}
func (UnimplementedRowDataServiceServer) AthleteTrainingDataGet(context.Context, *AthleteTrainingDataGetRequest) (*AthleteTrainingDataGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AthleteTrainingDataGet not implemented")
}
func (UnimplementedRowDataServiceServer) SampleMetricsAdd(context.Context, *SampleMetricsModel) (*SampleMetricsAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SampleMetricsAdd not implemented")
}
func (UnimplementedRowDataServiceServer) SampleMetricsGetByAthleteTrainingId(context.Context, *SampleMetricsGetByAthleteTrainingIdRequest) (*SampleMetricsGetByAthleteTrainingIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SampleMetricsGetByAthleteTrainingId not implemented")
}
func (UnimplementedRowDataServiceServer) mustEmbedUnimplementedRowDataServiceServer() {}

// UnsafeRowDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RowDataServiceServer will
// result in compilation errors.
type UnsafeRowDataServiceServer interface {
	mustEmbedUnimplementedRowDataServiceServer()
}

func RegisterRowDataServiceServer(s grpc.ServiceRegistrar, srv RowDataServiceServer) {
	s.RegisterService(&RowDataService_ServiceDesc, srv)
}

func _RowDataService_TrainingSummaryAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainingSummaryAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RowDataServiceServer).TrainingSummaryAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RowDataService_TrainingSummaryAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RowDataServiceServer).TrainingSummaryAdd(ctx, req.(*TrainingSummaryAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RowDataService_TrainingSummaryGetByTrainingName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainingSummaryGetByTrainingNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RowDataServiceServer).TrainingSummaryGetByTrainingName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RowDataService_TrainingSummaryGetByTrainingName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RowDataServiceServer).TrainingSummaryGetByTrainingName(ctx, req.(*TrainingSummaryGetByTrainingNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RowDataService_TrainingSummaryGetByTrainDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainingSummaryGetByTrainDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RowDataServiceServer).TrainingSummaryGetByTrainDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RowDataService_TrainingSummaryGetByTrainDate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RowDataServiceServer).TrainingSummaryGetByTrainDate(ctx, req.(*TrainingSummaryGetByTrainDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RowDataService_TrainingSummaryGetByEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainingSummaryGetByEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RowDataServiceServer).TrainingSummaryGetByEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RowDataService_TrainingSummaryGetByEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RowDataServiceServer).TrainingSummaryGetByEvent(ctx, req.(*TrainingSummaryGetByEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RowDataService_TrainingSummaryGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainingSummaryGetRequset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RowDataServiceServer).TrainingSummaryGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RowDataService_TrainingSummaryGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RowDataServiceServer).TrainingSummaryGet(ctx, req.(*TrainingSummaryGetRequset))
	}
	return interceptor(ctx, in, info, handler)
}

func _RowDataService_AthleteTrainingDataAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AthleteTrainingDataAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RowDataServiceServer).AthleteTrainingDataAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RowDataService_AthleteTrainingDataAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RowDataServiceServer).AthleteTrainingDataAdd(ctx, req.(*AthleteTrainingDataAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RowDataService_AthleteTrainingDataGetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AthleteTrainingDataGetByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RowDataServiceServer).AthleteTrainingDataGetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RowDataService_AthleteTrainingDataGetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RowDataServiceServer).AthleteTrainingDataGetByName(ctx, req.(*AthleteTrainingDataGetByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _RowDataService_AthleteTrainingDataGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AthleteTrainingDataGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RowDataServiceServer).AthleteTrainingDataGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RowDataService_AthleteTrainingDataGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RowDataServiceServer).AthleteTrainingDataGet(ctx, req.(*AthleteTrainingDataGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RowDataService_SampleMetricsAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleMetricsModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RowDataServiceServer).SampleMetricsAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RowDataService_SampleMetricsAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RowDataServiceServer).SampleMetricsAdd(ctx, req.(*SampleMetricsModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _RowDataService_SampleMetricsGetByAthleteTrainingId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleMetricsGetByAthleteTrainingIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RowDataServiceServer).SampleMetricsGetByAthleteTrainingId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RowDataService_SampleMetricsGetByAthleteTrainingId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RowDataServiceServer).SampleMetricsGetByAthleteTrainingId(ctx, req.(*SampleMetricsGetByAthleteTrainingIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RowDataService_ServiceDesc is the grpc.ServiceDesc for RowDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RowDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "idl.RowDataService",
	HandlerType: (*RowDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrainingSummaryAdd",
			Handler:    _RowDataService_TrainingSummaryAdd_Handler,
		},
		{
			MethodName: "TrainingSummaryGetByTrainingName",
			Handler:    _RowDataService_TrainingSummaryGetByTrainingName_Handler,
		},
		{
			MethodName: "TrainingSummaryGetByTrainDate",
			Handler:    _RowDataService_TrainingSummaryGetByTrainDate_Handler,
		},
		{
			MethodName: "TrainingSummaryGetByEvent",
			Handler:    _RowDataService_TrainingSummaryGetByEvent_Handler,
		},
		{
			MethodName: "TrainingSummaryGet",
			Handler:    _RowDataService_TrainingSummaryGet_Handler,
		},
		{
			MethodName: "AthleteTrainingDataAdd",
			Handler:    _RowDataService_AthleteTrainingDataAdd_Handler,
		},
		{
			MethodName: "AthleteTrainingDataGetByName",
			Handler:    _RowDataService_AthleteTrainingDataGetByName_Handler,
		},
		{
			MethodName: "AthleteTrainingDataGet",
			Handler:    _RowDataService_AthleteTrainingDataGet_Handler,
		},
		{
			MethodName: "SampleMetricsAdd",
			Handler:    _RowDataService_SampleMetricsAdd_Handler,
		},
		{
			MethodName: "SampleMetricsGetByAthleteTrainingId",
			Handler:    _RowDataService_SampleMetricsGetByAthleteTrainingId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rowDataService.proto",
}
