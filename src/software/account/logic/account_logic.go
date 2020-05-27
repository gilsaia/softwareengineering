package logic

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"software/account/message"
	"software/account/model"
	"software/common"
	"time"
)

type AccountLogic struct {
	ctx context.Context
}

func NewAccountLogic(ctx context.Context) AccountLogic {
	return AccountLogic{ctx: ctx}
}

func (logic AccountLogic) BackendAuth(keyId string, keySecret string) common.BgErr {
	err := message.BackendAuth(keyId, keySecret)
	if err != nil {
		return common.CustomErr(common.SmsErr, err)
	}
	return common.Success
}

func (logic AccountLogic) Verify(cellphone string) common.BgErr {
	ran := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	bindCode := fmt.Sprintf("%06v", ran.Int31n(1000000))
	user := model.User{
		Id:        common.GenId(),
		Cellphone: cellphone,
		BindCode:  bindCode,
	}
	err := message.SendSms(cellphone, bindCode)
	if err != nil {
		return common.CustomErr(common.SmsErr, err)
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return common.CustomErr(common.DbErr, err)
	}
	err = model.CreateUser(db, user)
	if err != nil {
		err = model.UpdateUser(db, cellphone, user)
		return common.CustomErr(common.DbErr, err)
	}
	return common.Success
}

func (logic AccountLogic) Register(cellphone string, password string, nickname string, bindcode string) common.BgErr {
	db, err := model.NewDbConnection()
	if err != nil {
		return common.CustomErr(common.DbErr, err)
	}
	user, err := model.GetUser(db, cellphone)
	if err != nil {
		return common.CustomErr(common.DbErr, err)
	}
	if user.BindCode != bindcode {
		return common.BindErr
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return common.CustomErr(common.InternalErr, err)
	}
	encodePwd := string(hash)
	user.Password = encodePwd
	user.Nickname = nickname
	err = model.UpdateUser(db, cellphone, user)
	if err != nil {
		return common.CustomErr(common.DbErr, err)
	}
	return common.Success
}

func (logic AccountLogic) Login(cellphone string, password string) (string, common.BgErr) {
	db, err := model.NewDbConnection()
	if err != nil {
		return "", common.CustomErr(common.DbErr, err)
	}
	user, err := model.GetUser(db, cellphone)
	if err != nil {
		return "", common.UserErr
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", common.UserErr
	}
	token, err := common.GetToken(user.Id)
	if err != nil {
		return "", common.CustomErr(common.TokenErr, err)
	}
	return token, common.Success
}
