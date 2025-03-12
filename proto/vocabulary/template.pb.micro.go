// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/vocabulary/template.proto

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

// Client API for TemplateService service

type TemplateService interface {
	AddOne(ctx context.Context, in *ReqTemplateAdd, opts ...client.CallOption) (*ReplyTemplateInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTemplateInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, opts ...client.CallOption) (*ReplyTemplateInfo, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyTemplateList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
}

type templateService struct {
	c    client.Client
	name string
}

func NewTemplateService(name string, c client.Client) TemplateService {
	return &templateService{
		c:    c,
		name: name,
	}
}

func (c *templateService) AddOne(ctx context.Context, in *ReqTemplateAdd, opts ...client.CallOption) (*ReplyTemplateInfo, error) {
	req := c.c.NewRequest(c.name, "TemplateService.AddOne", in)
	out := new(ReplyTemplateInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTemplateInfo, error) {
	req := c.c.NewRequest(c.name, "TemplateService.GetOne", in)
	out := new(ReplyTemplateInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TemplateService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateService) UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, opts ...client.CallOption) (*ReplyTemplateInfo, error) {
	req := c.c.NewRequest(c.name, "TemplateService.UpdateByFilter", in)
	out := new(ReplyTemplateInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyTemplateList, error) {
	req := c.c.NewRequest(c.name, "TemplateService.GetListByFilter", in)
	out := new(ReplyTemplateList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "TemplateService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TemplateService service

type TemplateServiceHandler interface {
	AddOne(context.Context, *ReqTemplateAdd, *ReplyTemplateInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyTemplateInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	UpdateByFilter(context.Context, *ReqUpdateFilter, *ReplyTemplateInfo) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyTemplateList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
}

func RegisterTemplateServiceHandler(s server.Server, hdlr TemplateServiceHandler, opts ...server.HandlerOption) error {
	type templateService interface {
		AddOne(ctx context.Context, in *ReqTemplateAdd, out *ReplyTemplateInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyTemplateInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, out *ReplyTemplateInfo) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyTemplateList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
	}
	type TemplateService struct {
		templateService
	}
	h := &templateServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TemplateService{h}, opts...))
}

type templateServiceHandler struct {
	TemplateServiceHandler
}

func (h *templateServiceHandler) AddOne(ctx context.Context, in *ReqTemplateAdd, out *ReplyTemplateInfo) error {
	return h.TemplateServiceHandler.AddOne(ctx, in, out)
}

func (h *templateServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyTemplateInfo) error {
	return h.TemplateServiceHandler.GetOne(ctx, in, out)
}

func (h *templateServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.TemplateServiceHandler.RemoveOne(ctx, in, out)
}

func (h *templateServiceHandler) UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, out *ReplyTemplateInfo) error {
	return h.TemplateServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *templateServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyTemplateList) error {
	return h.TemplateServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *templateServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.TemplateServiceHandler.GetStatistic(ctx, in, out)
}
