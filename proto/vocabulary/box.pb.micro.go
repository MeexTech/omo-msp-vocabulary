// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/vocabulary/box.proto

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

// Client API for BoxService service

type BoxService interface {
	AddOne(ctx context.Context, in *ReqBoxAdd, opts ...client.CallOption) (*ReplyBoxInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyBoxInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetAll(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyBoxList, error)
	GetListByUser(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyBoxList, error)
	GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyBoxList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateBase(ctx context.Context, in *ReqBoxUpdate, opts ...client.CallOption) (*ReplyBoxInfo, error)
	AppendKeywords(ctx context.Context, in *ReqBoxKeywords, opts ...client.CallOption) (*ReplyBoxInfo, error)
	SubtractKeywords(ctx context.Context, in *ReqBoxKeywords, opts ...client.CallOption) (*ReplyBoxInfo, error)
	AppendUsers(ctx context.Context, in *ReqBoxKeywords, opts ...client.CallOption) (*ReplyBoxInfo, error)
	SubtractUsers(ctx context.Context, in *ReqBoxKeywords, opts ...client.CallOption) (*ReplyBoxInfo, error)
	UpdateUsers(ctx context.Context, in *ReqBoxKeywords, opts ...client.CallOption) (*ReplyBoxInfo, error)
	UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, opts ...client.CallOption) (*ReplyInfo, error)
}

type boxService struct {
	c    client.Client
	name string
}

func NewBoxService(name string, c client.Client) BoxService {
	return &boxService{
		c:    c,
		name: name,
	}
}

func (c *boxService) AddOne(ctx context.Context, in *ReqBoxAdd, opts ...client.CallOption) (*ReplyBoxInfo, error) {
	req := c.c.NewRequest(c.name, "BoxService.AddOne", in)
	out := new(ReplyBoxInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyBoxInfo, error) {
	req := c.c.NewRequest(c.name, "BoxService.GetOne", in)
	out := new(ReplyBoxInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "BoxService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) GetAll(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyBoxList, error) {
	req := c.c.NewRequest(c.name, "BoxService.GetAll", in)
	out := new(ReplyBoxList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) GetListByUser(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyBoxList, error) {
	req := c.c.NewRequest(c.name, "BoxService.GetListByUser", in)
	out := new(ReplyBoxList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyBoxList, error) {
	req := c.c.NewRequest(c.name, "BoxService.GetByFilter", in)
	out := new(ReplyBoxList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "BoxService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) UpdateBase(ctx context.Context, in *ReqBoxUpdate, opts ...client.CallOption) (*ReplyBoxInfo, error) {
	req := c.c.NewRequest(c.name, "BoxService.UpdateBase", in)
	out := new(ReplyBoxInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) AppendKeywords(ctx context.Context, in *ReqBoxKeywords, opts ...client.CallOption) (*ReplyBoxInfo, error) {
	req := c.c.NewRequest(c.name, "BoxService.AppendKeywords", in)
	out := new(ReplyBoxInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) SubtractKeywords(ctx context.Context, in *ReqBoxKeywords, opts ...client.CallOption) (*ReplyBoxInfo, error) {
	req := c.c.NewRequest(c.name, "BoxService.SubtractKeywords", in)
	out := new(ReplyBoxInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) AppendUsers(ctx context.Context, in *ReqBoxKeywords, opts ...client.CallOption) (*ReplyBoxInfo, error) {
	req := c.c.NewRequest(c.name, "BoxService.AppendUsers", in)
	out := new(ReplyBoxInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) SubtractUsers(ctx context.Context, in *ReqBoxKeywords, opts ...client.CallOption) (*ReplyBoxInfo, error) {
	req := c.c.NewRequest(c.name, "BoxService.SubtractUsers", in)
	out := new(ReplyBoxInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) UpdateUsers(ctx context.Context, in *ReqBoxKeywords, opts ...client.CallOption) (*ReplyBoxInfo, error) {
	req := c.c.NewRequest(c.name, "BoxService.UpdateUsers", in)
	out := new(ReplyBoxInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxService) UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "BoxService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BoxService service

type BoxServiceHandler interface {
	AddOne(context.Context, *ReqBoxAdd, *ReplyBoxInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyBoxInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetAll(context.Context, *RequestInfo, *ReplyBoxList) error
	GetListByUser(context.Context, *RequestInfo, *ReplyBoxList) error
	GetByFilter(context.Context, *RequestFilter, *ReplyBoxList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateBase(context.Context, *ReqBoxUpdate, *ReplyBoxInfo) error
	AppendKeywords(context.Context, *ReqBoxKeywords, *ReplyBoxInfo) error
	SubtractKeywords(context.Context, *ReqBoxKeywords, *ReplyBoxInfo) error
	AppendUsers(context.Context, *ReqBoxKeywords, *ReplyBoxInfo) error
	SubtractUsers(context.Context, *ReqBoxKeywords, *ReplyBoxInfo) error
	UpdateUsers(context.Context, *ReqBoxKeywords, *ReplyBoxInfo) error
	UpdateByFilter(context.Context, *ReqUpdateFilter, *ReplyInfo) error
}

func RegisterBoxServiceHandler(s server.Server, hdlr BoxServiceHandler, opts ...server.HandlerOption) error {
	type boxService interface {
		AddOne(ctx context.Context, in *ReqBoxAdd, out *ReplyBoxInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyBoxInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetAll(ctx context.Context, in *RequestInfo, out *ReplyBoxList) error
		GetListByUser(ctx context.Context, in *RequestInfo, out *ReplyBoxList) error
		GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyBoxList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateBase(ctx context.Context, in *ReqBoxUpdate, out *ReplyBoxInfo) error
		AppendKeywords(ctx context.Context, in *ReqBoxKeywords, out *ReplyBoxInfo) error
		SubtractKeywords(ctx context.Context, in *ReqBoxKeywords, out *ReplyBoxInfo) error
		AppendUsers(ctx context.Context, in *ReqBoxKeywords, out *ReplyBoxInfo) error
		SubtractUsers(ctx context.Context, in *ReqBoxKeywords, out *ReplyBoxInfo) error
		UpdateUsers(ctx context.Context, in *ReqBoxKeywords, out *ReplyBoxInfo) error
		UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, out *ReplyInfo) error
	}
	type BoxService struct {
		boxService
	}
	h := &boxServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BoxService{h}, opts...))
}

type boxServiceHandler struct {
	BoxServiceHandler
}

func (h *boxServiceHandler) AddOne(ctx context.Context, in *ReqBoxAdd, out *ReplyBoxInfo) error {
	return h.BoxServiceHandler.AddOne(ctx, in, out)
}

func (h *boxServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyBoxInfo) error {
	return h.BoxServiceHandler.GetOne(ctx, in, out)
}

func (h *boxServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.BoxServiceHandler.RemoveOne(ctx, in, out)
}

func (h *boxServiceHandler) GetAll(ctx context.Context, in *RequestInfo, out *ReplyBoxList) error {
	return h.BoxServiceHandler.GetAll(ctx, in, out)
}

func (h *boxServiceHandler) GetListByUser(ctx context.Context, in *RequestInfo, out *ReplyBoxList) error {
	return h.BoxServiceHandler.GetListByUser(ctx, in, out)
}

func (h *boxServiceHandler) GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyBoxList) error {
	return h.BoxServiceHandler.GetByFilter(ctx, in, out)
}

func (h *boxServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.BoxServiceHandler.GetStatistic(ctx, in, out)
}

func (h *boxServiceHandler) UpdateBase(ctx context.Context, in *ReqBoxUpdate, out *ReplyBoxInfo) error {
	return h.BoxServiceHandler.UpdateBase(ctx, in, out)
}

func (h *boxServiceHandler) AppendKeywords(ctx context.Context, in *ReqBoxKeywords, out *ReplyBoxInfo) error {
	return h.BoxServiceHandler.AppendKeywords(ctx, in, out)
}

func (h *boxServiceHandler) SubtractKeywords(ctx context.Context, in *ReqBoxKeywords, out *ReplyBoxInfo) error {
	return h.BoxServiceHandler.SubtractKeywords(ctx, in, out)
}

func (h *boxServiceHandler) AppendUsers(ctx context.Context, in *ReqBoxKeywords, out *ReplyBoxInfo) error {
	return h.BoxServiceHandler.AppendUsers(ctx, in, out)
}

func (h *boxServiceHandler) SubtractUsers(ctx context.Context, in *ReqBoxKeywords, out *ReplyBoxInfo) error {
	return h.BoxServiceHandler.SubtractUsers(ctx, in, out)
}

func (h *boxServiceHandler) UpdateUsers(ctx context.Context, in *ReqBoxKeywords, out *ReplyBoxInfo) error {
	return h.BoxServiceHandler.UpdateUsers(ctx, in, out)
}

func (h *boxServiceHandler) UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, out *ReplyInfo) error {
	return h.BoxServiceHandler.UpdateByFilter(ctx, in, out)
}
