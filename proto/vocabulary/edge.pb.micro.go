// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/vocabulary/edge.proto

package vocabulary

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VEdgeService service

type VEdgeService interface {
	AddOne(ctx context.Context, in *ReqVEdgeAdd, opts ...client.CallOption) (*ReplyVEdgeInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyVEdgeInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateInfo(ctx context.Context, in *ReqVEdgeUpdate, opts ...client.CallOption) (*ReplyVEdgeInfo, error)
	GetAll(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyVEdgeList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, opts ...client.CallOption) (*ReplyInfo, error)
}

type vEdgeService struct {
	c    client.Client
	name string
}

func NewVEdgeService(name string, c client.Client) VEdgeService {
	return &vEdgeService{
		c:    c,
		name: name,
	}
}

func (c *vEdgeService) AddOne(ctx context.Context, in *ReqVEdgeAdd, opts ...client.CallOption) (*ReplyVEdgeInfo, error) {
	req := c.c.NewRequest(c.name, "VEdgeService.AddOne", in)
	out := new(ReplyVEdgeInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vEdgeService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyVEdgeInfo, error) {
	req := c.c.NewRequest(c.name, "VEdgeService.GetOne", in)
	out := new(ReplyVEdgeInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vEdgeService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "VEdgeService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vEdgeService) UpdateInfo(ctx context.Context, in *ReqVEdgeUpdate, opts ...client.CallOption) (*ReplyVEdgeInfo, error) {
	req := c.c.NewRequest(c.name, "VEdgeService.UpdateInfo", in)
	out := new(ReplyVEdgeInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vEdgeService) GetAll(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyVEdgeList, error) {
	req := c.c.NewRequest(c.name, "VEdgeService.GetAll", in)
	out := new(ReplyVEdgeList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vEdgeService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "VEdgeService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vEdgeService) UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "VEdgeService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VEdgeService service

type VEdgeServiceHandler interface {
	AddOne(context.Context, *ReqVEdgeAdd, *ReplyVEdgeInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyVEdgeInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	UpdateInfo(context.Context, *ReqVEdgeUpdate, *ReplyVEdgeInfo) error
	GetAll(context.Context, *RequestInfo, *ReplyVEdgeList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *ReqUpdateFilter, *ReplyInfo) error
}

func RegisterVEdgeServiceHandler(s server.Server, hdlr VEdgeServiceHandler, opts ...server.HandlerOption) error {
	type vEdgeService interface {
		AddOne(ctx context.Context, in *ReqVEdgeAdd, out *ReplyVEdgeInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyVEdgeInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		UpdateInfo(ctx context.Context, in *ReqVEdgeUpdate, out *ReplyVEdgeInfo) error
		GetAll(ctx context.Context, in *RequestInfo, out *ReplyVEdgeList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, out *ReplyInfo) error
	}
	type VEdgeService struct {
		vEdgeService
	}
	h := &vEdgeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&VEdgeService{h}, opts...))
}

type vEdgeServiceHandler struct {
	VEdgeServiceHandler
}

func (h *vEdgeServiceHandler) AddOne(ctx context.Context, in *ReqVEdgeAdd, out *ReplyVEdgeInfo) error {
	return h.VEdgeServiceHandler.AddOne(ctx, in, out)
}

func (h *vEdgeServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyVEdgeInfo) error {
	return h.VEdgeServiceHandler.GetOne(ctx, in, out)
}

func (h *vEdgeServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.VEdgeServiceHandler.RemoveOne(ctx, in, out)
}

func (h *vEdgeServiceHandler) UpdateInfo(ctx context.Context, in *ReqVEdgeUpdate, out *ReplyVEdgeInfo) error {
	return h.VEdgeServiceHandler.UpdateInfo(ctx, in, out)
}

func (h *vEdgeServiceHandler) GetAll(ctx context.Context, in *RequestInfo, out *ReplyVEdgeList) error {
	return h.VEdgeServiceHandler.GetAll(ctx, in, out)
}

func (h *vEdgeServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.VEdgeServiceHandler.GetStatistic(ctx, in, out)
}

func (h *vEdgeServiceHandler) UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, out *ReplyInfo) error {
	return h.VEdgeServiceHandler.UpdateByFilter(ctx, in, out)
}
