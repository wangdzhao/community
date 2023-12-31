// Code generated by goctl. DO NOT EDIT.
// Source: discovery.proto

package discoveryclient

import (
	"context"
	"github.com/wangdzhao/discovery/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = pb.Request
	Response = pb.Response

	Discovery interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultDiscovery struct {
		cli zrpc.Client
	}
)

func NewDiscovery(cli zrpc.Client) Discovery {
	return &defaultDiscovery{
		cli: cli,
	}
}

func (m *defaultDiscovery) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewDiscoveryClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
