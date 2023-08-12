package logic

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/wangdzhao/community/discovery/pb"

	"github.com/wangdzhao/community/rebot/internal/svc"
	"github.com/wangdzhao/community/rebot/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

type RebotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRebotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RebotLogic {
	return &RebotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RebotLogic) Rebot(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

func ClientDiscovery() {
	// 连接到 gRPC 服务器
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDiscoveryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用 RPC 方法
	resp, err := client.Ping(ctx, &pb.Request{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println("Response:", resp.Pong)
}
