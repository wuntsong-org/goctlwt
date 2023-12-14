package eventlogic

import (
	"context"

	"github.com/wuntsong-org/goctlwt/example/rpc/hi/internal/svc"
	"github.com/wuntsong-org/goctlwt/example/rpc/hi/pb/hi"

	"github.com/wuntsong-org/go-zero-plus/core/logx"
)

type AskQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAskQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AskQuestionLogic {
	return &AskQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AskQuestionLogic) AskQuestion(in *hi.EventReq) (*hi.EventResp, error) {
	// todo: add your logic here and delete this line

	return &hi.EventResp{}, nil
}
