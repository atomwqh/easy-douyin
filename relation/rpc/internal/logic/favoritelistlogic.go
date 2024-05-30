package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *relation.FavoriteListReq) (*relation.FavoriteListResp, error) {
	favorite, err := l.svcCtx.RelationModel.FindFavorite(l.ctx, in.ActUser)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &relation.FavoriteListResp{
				StatusCode: constants.ServiceOKCode,
				StatusMsg:  constants.ServiceOK,
				List:       nil,
			}, nil
		} else {
			l.Logger.Error("数据库查询用户关注列表错")
			return nil, err
		}
	}

	List := make([]*relation.UserInfo, 0)
	for _, item := range favorite {
		one, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			UserId:  in.Uid,
			ActorId: int64(item.FollowingId),
		})
		if err != nil {
			l.Logger.Error("查询用户信息错")
			return nil, err
		}
		var coverUrl string
		var vid int64
		video, err := l.svcCtx.VideoModel.FindLastByUId(l.ctx, int64(one.User.Id))
		if err != nil {
			coverUrl = ""
			vid = 0
		} else {
			coverUrl = video.CoverUrl
			vid = int64(video.Id)
		}
		List = append(List, &relation.UserInfo{
			Id:              int64(one.User.Id),
			Name:            one.User.Name,
			Gender:          int64(one.User.Gender),
			Avatar:          *one.User.Avatar,
			Dec:             *one.User.Signature,
			BackgroundImage: *one.User.BackgroundImage,
			VideoId:         vid,
			CoverUrl:        coverUrl,
			FollowCount:     *one.User.FollowCount,
			FollowerCount:   *one.User.FollowerCount,
			IsFollow:        one.User.IsFollow,
			TotalFavorited:  *one.User.TotalFavorited,
			WorkCount:       *one.User.WorkCount,
			FavoriteCount:   *one.User.FavoriteCount,
			FriendCount:     one.User.FriendCount,
		})
	}
	return &relation.FavoriteListResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		List:       List,
	}, nil
}