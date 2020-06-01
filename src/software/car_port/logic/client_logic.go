package logic

import (
	"context"
	"software/car_port/model"
	"software/common"
)

type ClientLogic struct {
	ctx context.Context
}

func NewClientLogic(ctx context.Context) (*ClientLogic, common.BgErr) {
	if err := common.AuthPermission(ctx, common.PermissionVisitor); !err.Is(common.Success) {
		return nil, err
	}
	return &ClientLogic{ctx: ctx}, common.Success
}

func (logic ClientLogic) UserInfo(cellphone string) (nickname string, bgErr common.BgErr) {
	db, err := model.NewDbConnection()
	if err != nil {
		return "", common.CustomErr(common.DbErr, err)
	}
	user, err := model.GetUserByCellphone(db, cellphone)
	if err != nil {
		return "", common.CustomErr(common.DbErr, err)
	}
	return user.Nickname, common.Success
}
