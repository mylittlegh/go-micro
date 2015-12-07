// Code generated by protoc-gen-go.
// source: go-micro/examples/server/proto/example/example.proto
// DO NOT EDIT!

/*
Package go_micro_srv_example is a generated protocol buffer package.

It is generated from these files:
	go-micro/examples/server/proto/example/example.proto

It has these top-level messages:
	Message
	Request
	Response
	StreamingRequest
	StreamingResponse
*/
package go_micro_srv_example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message struct {
	Say string `protobuf:"bytes,1,opt,name=say" json:"say,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StreamingRequest struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *StreamingRequest) Reset()                    { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()               {}
func (*StreamingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type StreamingResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *StreamingResponse) Reset()                    { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()               {}
func (*StreamingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.example.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.example.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.example.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.example.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.example.StreamingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Example service

type ExampleClient interface {
	Call(ctx context.Context, in *Request) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest) (Example_StreamClient, error)
}

type exampleClient struct {
	c           client.Client
	serviceName string
}

func NewExampleClient(serviceName string, c client.Client) ExampleClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.example"
	}
	return &exampleClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *exampleClient) Call(ctx context.Context, in *Request) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) Stream(ctx context.Context, in *StreamingRequest) (Example_StreamClient, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Stream", in)
	outCh := make(chan *StreamingResponse)
	stream, err := c.c.Stream(ctx, req, outCh)
	if err != nil {
		return nil, err
	}
	return &exampleStreamClient{stream, outCh}, nil
}

type Example_StreamClient interface {
	Next() (*StreamingResponse, error)
	client.Streamer
}

type exampleStreamClient struct {
	client.Streamer
	next chan *StreamingResponse
}

func (x *exampleStreamClient) Next() (*StreamingResponse, error) {
	out, ok := <-x.next
	if !ok {
		return nil, fmt.Errorf(`chan closed`)
	}
	return out, nil
}

// Server API for Example service

type ExampleHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, func(*StreamingResponse) error) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler) {
	s.Handle(s.NewHandler(hdlr))
}

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0x0d, 0xf7, 0x7a, 0xab, 0xa3, 0x82, 0x06, 0x51, 0x29, 0x28, 0x9a, 0x85, 0xba, 0x31,
	0x15, 0xf5, 0x0d, 0x44, 0x5c, 0xb9, 0xa9, 0x6b, 0x17, 0xb1, 0x0c, 0xa1, 0xd0, 0x24, 0x35, 0x93,
	0x16, 0x7d, 0x2c, 0xdf, 0x50, 0x48, 0xd3, 0xa2, 0x52, 0x71, 0x15, 0x98, 0xf3, 0x9d, 0x1f, 0x02,
	0x77, 0xda, 0x5d, 0x99, 0xba, 0xf2, 0xae, 0xc0, 0x77, 0x65, 0xda, 0x06, 0xa9, 0x20, 0xf4, 0x3d,
	0xfa, 0xa2, 0xf5, 0x2e, 0x4c, 0xd7, 0xf1, 0x95, 0xf1, 0xca, 0xf7, 0xb5, 0x93, 0xd1, 0x25, 0xc9,
	0xf7, 0x32, 0x69, 0xe2, 0x00, 0xb2, 0x27, 0x24, 0x52, 0x1a, 0xf9, 0x16, 0x2c, 0x48, 0x7d, 0x1c,
	0xb1, 0x53, 0x76, 0xb9, 0x29, 0x0e, 0x21, 0x2b, 0xf1, 0xad, 0x43, 0x0a, 0x7c, 0x1b, 0x96, 0x56,
	0x19, 0x9c, 0x84, 0x8d, 0x12, 0xa9, 0x75, 0x96, 0xa2, 0xc3, 0x90, 0x4e, 0xc2, 0x19, 0xec, 0x3e,
	0x07, 0x8f, 0xca, 0xd4, 0x56, 0x8f, 0xd6, 0x1d, 0x58, 0xaf, 0x5c, 0x67, 0x43, 0x44, 0x16, 0x42,
	0xc0, 0xde, 0x37, 0x24, 0x85, 0xfc, 0x64, 0x6e, 0x3e, 0x19, 0x64, 0x0f, 0xc3, 0x38, 0xfe, 0x08,
	0xcb, 0x7b, 0xd5, 0x34, 0xfc, 0x58, 0xce, 0x6d, 0x97, 0xa9, 0x25, 0x3f, 0xf9, 0x4b, 0x1e, 0x1a,
	0xc4, 0x1a, 0x7f, 0x81, 0xd5, 0x50, 0xcc, 0xcf, 0xe7, 0xd9, 0xdf, 0xcb, 0xf3, 0x8b, 0x7f, 0xb9,
	0x31, 0xfc, 0x9a, 0xbd, 0xae, 0xe2, 0x0f, 0xdf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x63, 0x02,
	0xbf, 0x5f, 0x99, 0x01, 0x00, 0x00,
}
