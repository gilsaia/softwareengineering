package logic

import (
	"context"
	"flag"
	"os"
	"software/common"
	"testing"
)

var logic AccountLogic

func TestAccountLogic_Verify(t *testing.T) {

}

func TestAccountLogic_Login(t *testing.T) {
	var loginTests = []struct {
		cellphone string
		password  string
		result    common.BgErr
	}{
		{"15822797219", "12345678", common.Success},
		{"15822797218", "12345678", common.UserErr},
		{"15822797219", "1234567", common.PasswordErr},
	}
	for _, value := range loginTests {
		_, err := logic.Login(value.cellphone, value.password)
		if !err.Is(value.result) {
			t.Errorf("Login Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
	}
}

func TestAccountLogic_Register(t *testing.T) {
	var registerTests = []struct {
		cellphone string
		password  string
		nickname  string
		bindCode  string
		result    common.BgErr
	}{
		{"15822797219", "12345678", "gilsaia", "786989", common.Success},
		{"15822797219", "12345678", "gilsaia", "123456", common.BindErr},
	}
	for _, value := range registerTests {
		err := logic.Register(value.cellphone, value.password, value.nickname, value.bindCode)
		if !err.Is(value.result) {
			t.Errorf("Login Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
	}
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	logic = NewAccountLogic(ctx)
	flag.Parse()
	os.Exit(m.Run())
}
