package logic

import (
	"context"
	"flag"
	"os"
	"software/car_port/pb_gen"
	"software/common"
	"testing"
)

const (
	userId     = int64(304917993446440966)
	permission = 1
)

var (
	billLogic    *BillLogic
	carPortLogic *CarPortLogic
	clientLogic  *ClientLogic
	parkLogic    *ParkLogic
	userLogic    *UserLogic
)

func TestBillLogic_GetBill(t *testing.T) {
	var loginTests = []struct {
		billId int64
		bill   *pb_gen.Bill
		result common.BgErr
	}{
		{1, &pb_gen.Bill{
			Id:       1,
			UserId:   304917993446440966,
			ParkId:   1,
			Duration: "2m39.762921564s",
			Charge:   2,
			Status:   2,
		}, common.Success},
		{-1, nil, common.ParamErr},
		{10, nil, common.DbErr},
	}
	for _, value := range loginTests {
		bill, err := billLogic.GetBill(value.billId)
		if !err.Is(value.result) {
			t.Errorf("GetBill Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if value.bill != nil {
			if bill == nil {
				t.Errorf("Get Bill Wrong! Result: %v", value.bill)
			} else if value.bill.UserId != bill.UserId || value.bill.Status != bill.Status || value.bill.Id != bill.Id {
				t.Errorf("Bill Info Wrong! Result: %v Expect: %v", bill, value.bill)
			}
		}
	}
}

func TestBillLogic_MGetBill(t *testing.T) {
	var loginTests = []struct {
		count  int32
		num    int32
		result common.BgErr
	}{
		{1, 0, common.Success},
		{10, 0, common.Success},
		{10, 20, common.Success},
	}
	for _, value := range loginTests {
		bills, _, err := billLogic.MGetBill(value.count, value.num)
		if !err.Is(value.result) {
			t.Errorf("MGetBill Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if len(bills) > int(value.count) {
			t.Errorf("MGetBill Result Count Wrong! %d %d", len(bills), value.count)
		}
	}
}

func TestCarPortLogic_CreateCarPort(t *testing.T) {
	var loginTests = []struct {
		port   *pb_gen.CarPort
		result common.BgErr
	}{
		{&pb_gen.CarPort{
			State:     3,
			Longitude: 0,
			Latitude:  0,
			Park:      0,
		}, common.Success},
		{&pb_gen.CarPort{
			Id:        -1,
			State:     0,
			Longitude: 0,
			Latitude:  0,
			Park:      0,
		}, common.ParamErr},
	}
	for _, value := range loginTests {
		err := carPortLogic.CreateCarPort(value.port)
		if !err.Is(value.result) {
			t.Errorf("Create CarPort Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
	}
}

func TestCarPortLogic_UpdateCarPort(t *testing.T) {
	var loginTests = []struct {
		port   *pb_gen.CarPort
		result common.BgErr
	}{
		{&pb_gen.CarPort{
			Id:        1,
			State:     3,
			Longitude: 0,
			Latitude:  0,
			Park:      0,
		}, common.Success},
		{&pb_gen.CarPort{
			Id:        -1,
			State:     0,
			Longitude: 0,
			Latitude:  0,
			Park:      0,
		}, common.ParamErr},
	}
	for _, value := range loginTests {
		err := carPortLogic.UpdateCarPort(value.port)
		if !err.Is(value.result) {
			t.Errorf("Update CarPort Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
	}
}

func TestCarPortLogic_MGetCarPort(t *testing.T) {
	var loginTests = []struct {
		count  int32
		num    int32
		result common.BgErr
	}{
		{1, 0, common.Success},
		{10, 0, common.Success},
		{10, 20, common.Success},
	}
	for _, value := range loginTests {
		bills, _, err := carPortLogic.MGetCarPort(value.count, value.num)
		if !err.Is(value.result) {
			t.Errorf("MGetCarPort Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if len(bills) > int(value.count) {
			t.Errorf("MGetCarPort Result Count Wrong! %d %d", len(bills), value.count)
		}
	}
}

func TestCarPortLogic_GetCarPort(t *testing.T) {
	var loginTests = []struct {
		portId int64
		port   *pb_gen.CarPort
		result common.BgErr
	}{
		{1, &pb_gen.CarPort{
			Id:        1,
			State:     1,
			Longitude: 0,
			Latitude:  0,
			Park:      0,
		}, common.Success},
		{-1, nil, common.ParamErr},
		{2, nil, common.DbErr},
	}
	for _, value := range loginTests {
		port, err := carPortLogic.GetCarPort(value.portId)
		if !err.Is(value.result) {
			t.Errorf("GetCarPort Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if value.port != nil {
			if port == nil {
				t.Errorf("Get CarPort Wrong! Result: %v", value.port)
			} else if value.port.Id != port.Id || value.port.Park != port.Park {
				t.Errorf("Port Info Wrong! Result: %v Expect: %v", port, value.port)
			}
		}
	}
}

func TestParkLogic_CreatePark(t *testing.T) {
	var loginTests = []struct {
		park   *pb_gen.Park
		result common.BgErr
	}{
		{&pb_gen.Park{
			Name:      "??",
			Latitude:  0,
			Longitude: 0,
			Charge:    0,
		}, common.Success},
		{&pb_gen.Park{
			Id:        -1,
			Name:      "??",
			Latitude:  0,
			Longitude: 0,
			Charge:    0,
		}, common.ParamErr},
	}
	for _, value := range loginTests {
		err := parkLogic.CreatePark(value.park)
		if !err.Is(value.result) {
			t.Errorf("Create Park Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
	}
}

func TestParkLogic_UpdatePark(t *testing.T) {
	var loginTests = []struct {
		park   *pb_gen.Park
		result common.BgErr
	}{
		{&pb_gen.Park{
			Id:        1,
			Name:      "??",
			Latitude:  0,
			Longitude: 0,
			Charge:    0,
		}, common.Success},
		{&pb_gen.Park{
			Id:        -1,
			Name:      "??",
			Latitude:  0,
			Longitude: 0,
			Charge:    0,
		}, common.ParamErr},
	}
	for _, value := range loginTests {
		err := parkLogic.UpdatePark(value.park)
		if !err.Is(value.result) {
			t.Errorf("Update Park Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
	}
}

func TestParkLogic_MGetPark(t *testing.T) {
	var loginTests = []struct {
		count  int32
		num    int32
		result common.BgErr
	}{
		{1, 0, common.Success},
		{10, 0, common.Success},
		{10, 20, common.Success},
	}
	for _, value := range loginTests {
		bills, _, err := parkLogic.MGetPark(value.count, value.num)
		if !err.Is(value.result) {
			t.Errorf("MGetPark Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if len(bills) > int(value.count) {
			t.Errorf("MGetPark Result Count Wrong! %d %d", len(bills), value.count)
		}
	}
}

func TestParkLogic_GetPark(t *testing.T) {
	var loginTests = []struct {
		parkId int64
		park   *pb_gen.Park
		result common.BgErr
	}{
		{1, &pb_gen.Park{
			Id:        1,
			Name:      "??",
			Latitude:  0,
			Longitude: 0,
			Charge:    0,
		}, common.Success},
		{-1, nil, common.ParamErr},
		{2, nil, common.DbErr},
	}
	for _, value := range loginTests {
		port, err := parkLogic.GetPark(value.parkId)
		if !err.Is(value.result) {
			t.Errorf("GetPark Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if value.park != nil {
			if port == nil {
				t.Errorf("Get Park Wrong! Result: %v", value.park)
			}
		}
	}
}

func TestUserLogic_UpdateUser(t *testing.T) {
	var loginTests = []struct {
		user   *pb_gen.User
		result common.BgErr
	}{
		{&pb_gen.User{
			Id:        304917993446440966,
			Cellphone: "15822797219",
			NickName:  "gilsaia",
			BindCode:  "123456",
		}, common.Success},
		{&pb_gen.User{
			Id:        -1,
			Cellphone: "15822797219",
			NickName:  "gilsaia",
			BindCode:  "123456",
		}, common.ParamErr},
	}
	for _, value := range loginTests {
		err := userLogic.UpdateUser(value.user, "")
		if !err.Is(value.result) {
			t.Errorf("Update User Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
	}
}

func TestUserLogic_MGetUser(t *testing.T) {
	var loginTests = []struct {
		count  int32
		num    int32
		result common.BgErr
	}{
		{1, 0, common.Success},
		{10, 0, common.Success},
		{10, 20, common.Success},
	}
	for _, value := range loginTests {
		bills, _, err := userLogic.MGetUser(value.count, value.num)
		if !err.Is(value.result) {
			t.Errorf("MGetUser Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if len(bills) > int(value.count) {
			t.Errorf("MGetUser Result Count Wrong! %d %d", len(bills), value.count)
		}
	}
}

func TestUserLogic_GetUser(t *testing.T) {
	var loginTests = []struct {
		userId int64
		user   *pb_gen.User
		result common.BgErr
	}{
		{304917993446440966, &pb_gen.User{
			Id:        304917993446440966,
			Cellphone: "15822797219",
			NickName:  "gilsaia",
			BindCode:  "123456",
		}, common.Success},
		{-1, nil, common.ParamErr},
		{2, nil, common.DbErr},
	}
	for _, value := range loginTests {
		user, err := userLogic.GetUser(value.userId, "")
		if !err.Is(value.result) {
			t.Errorf("GetUser Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if value.user != nil {
			if user == nil {
				t.Errorf("Get User Wrong! Result: %v", value.user)
			}
		}
	}
}

func TestClientLogic_UserInfo(t *testing.T) {
	var loginTests = []struct {
		cellphone string
		result    common.BgErr
	}{
		{"15822797219", common.Success},
		{"1", common.DbErr},
	}
	for _, value := range loginTests {
		_, _, _, err := clientLogic.UserInfo(value.cellphone)
		if !err.Is(value.result) {
			t.Errorf("UserInfo Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
	}

}

func TestClientLogic_BillInfo(t *testing.T) {
	var loginTests = []struct {
		billId int64
		bill   *pb_gen.Bill
		result common.BgErr
	}{
		{1, &pb_gen.Bill{
			Id:       1,
			UserId:   304917993446440966,
			ParkId:   1,
			Duration: "2m39.762921564s",
			Charge:   2,
			Status:   2,
		}, common.Success},
		{-1, nil, common.DbErr},
		{10, nil, common.DbErr},
	}
	for _, value := range loginTests {
		bill, err := clientLogic.BillInfo(value.billId)
		if !err.Is(value.result) {
			t.Errorf("BillInfo Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if value.bill != nil {
			if bill == nil {
				t.Errorf("BillInfo Wrong! Result: %v", value.bill)
			} else if value.bill.UserId != bill.UserId || value.bill.Status != bill.Status || value.bill.Id != bill.Id {
				t.Errorf("BillInfo Wrong! Result: %v Expect: %v", bill, value.bill)
			}
		}
	}
}

func TestClientLogic_MetaInfo(t *testing.T) {
	var loginTests = []struct {
		count  int32
		result common.BgErr
	}{
		{0, common.Success},
	}
	for _, value := range loginTests {
		_, parks, err := clientLogic.MetaInfo()
		if !err.Is(value.result) {
			t.Errorf("MetaInfo Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if len(parks) < int(value.count) {
			t.Errorf("MetaInfo Result Count Wrong! %d %d", len(parks), value.count)
		}
	}
}

func TestClientLogic_ParkInfo(t *testing.T) {
	var loginTests = []struct {
		parkId   int64
		parkName string
		result   common.BgErr
	}{
		{1, "??", common.Success},
		{-1, "", common.DbErr},
		{2, "", common.DbErr},
	}
	for _, value := range loginTests {
		parkName, _, err := clientLogic.ParkInfo(value.parkId)
		if !err.Is(value.result) {
			t.Errorf("ParkInfo Result Test Wrong! Error: %d %s", err.ErrNo, err.ErrMsg)
		}
		if value.parkName != parkName {
			t.Errorf("ParkInfo Wrong!")
		}
	}
}

func TestClientLogic_ClientProcess(t *testing.T) {
	err := clientLogic.Park(1)
	if !err.Is(common.Success) {
		t.Fatalf("ClientProcess Wrong! Error:%d %s", err.ErrNo, err.ErrMsg)
	}
	billId, err := clientLogic.PickUp(1)
	if !err.Is(common.Success) {
		t.Fatalf("ClientProcess Wrong! Error:%d %s", err.ErrNo, err.ErrMsg)
	}
	err = clientLogic.Park(1)
	if !err.Is(common.BillErr) {
		t.Fatalf("ClientProcess Wrong! Error:%d %s", err.ErrNo, err.ErrMsg)
	}
	err = clientLogic.Pay(2)
	if !err.Is(common.WrongBillErr) {
		t.Fatalf("ClientProcess Wrong! Error:%d %s", err.ErrNo, err.ErrMsg)
	}
	err = clientLogic.Pay(billId)
	if !err.Is(common.Success) {
		t.Fatalf("ClientProcess Wrong! Error:%d %s", err.ErrNo, err.ErrMsg)
	}
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userId", userId)
	ctx = context.WithValue(ctx, "permission", permission)
	billLogic, _ = NewBillLogic(ctx)
	carPortLogic, _ = NewCarPortLogic(ctx)
	clientLogic, _ = NewClientLogic(ctx)
	parkLogic, _ = NewParkLogic(ctx)
	userLogic, _ = NewUserLogic(ctx)
	flag.Parse()
	os.Exit(m.Run())
}
