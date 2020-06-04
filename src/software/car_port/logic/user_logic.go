package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"software/car_port/model"
	"software/car_port/pb_gen"
	"software/common"
)

type UserLogic struct {
	ctx context.Context
}

func NewUserLogic(ctx context.Context) (*UserLogic, common.BgErr) {
	if err := common.AuthPermission(ctx, common.PermissionAdmin); !err.Is(common.Success) {
		return nil, err
	}
	return &UserLogic{ctx: ctx}, common.Success
}

func (logic UserLogic) UpdateUser(user *pb_gen.User, password string) common.BgErr {
	if user.Id < 0 {
		return common.ParamErr
	}
	mUser := model.User{
		Id:        user.Id,
		Cellphone: user.Cellphone,
		Nickname:  user.NickName,
		Status:    int(user.Status),
	}
	if password != "" {
		encodePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return common.CustomErr(common.InternalErr, err)
		}
		mUser.Password = string(encodePassword)
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return common.CustomErr(common.DbErr, err)
	}
	err = model.UpdateUser(db, mUser)
	if err != nil {
		return common.CustomErr(common.DbErr, err)
	}
	return common.Success
}

func (logic UserLogic) GetUser(userId int64, cellphone string) (*pb_gen.User, common.BgErr) {
	if userId < 0 {
		return nil, common.ParamErr
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return nil, common.CustomErr(common.DbErr, err)
	}
	user := model.User{}
	if userId != 0 {
		user, err = model.GetUser(db, userId)
	} else {
		user, err = model.GetUserByCellphone(db, cellphone)
	}
	if err != nil {
		return nil, common.CustomErr(common.DbErr, err)
	}
	return packUser(user), common.Success
}

func (logic UserLogic) MGetUser(count int32, num int32) ([]*pb_gen.User, int32, common.BgErr) {
	db, err := model.NewDbConnection()
	if err != nil {
		return nil, 0, common.CustomErr(common.DbErr, err)
	}
	offset := num * count
	users, tableCount, err := model.MGetUser(db, offset, count)
	if err != nil {
		return nil, 0, common.CustomErr(common.DbErr, err)
	}
	var userList []*pb_gen.User
	for _, value := range users {
		userList = append(userList, packUser(value))
	}
	return userList, tableCount, common.Success
}

func packUser(user model.User) *pb_gen.User {
	return &pb_gen.User{
		Id:        user.Id,
		Cellphone: user.Cellphone,
		NickName:  user.Nickname,
		BindCode:  user.BindCode,
		Status:    pb_gen.UserStatus(user.Status),
	}
}
