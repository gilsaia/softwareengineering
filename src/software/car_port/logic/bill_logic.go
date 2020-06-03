package logic

import (
	"context"
	"software/car_port/model"
	"software/car_port/pb_gen"
	"software/common"
)

type BillLogic struct {
	ctx context.Context
}

func NewBillLogic(ctx context.Context) (*BillLogic, common.BgErr) {
	if err := common.AuthPermission(ctx, common.PermissionAdmin); !err.Is(common.Success) {
		return nil, err
	}
	return &BillLogic{ctx: ctx}, common.Success
}

func (logic BillLogic) GetBill(billId int64) (*pb_gen.Bill, common.BgErr) {
	if billId < 0 {
		return nil, common.ParamErr
	}
	db, err := model.NewDbConnection()
	if err != nil {
		return nil, common.DbErr
	}
	mBill, err := model.GetBill(db, billId)
	if err != nil {
		return nil, common.DbErr
	}
	bill := packBill(mBill)
	return bill, common.Success
}

func (logic BillLogic) MGetBill(count int32, num int32) ([]*pb_gen.Bill, int32, common.BgErr) {
	db, err := model.NewDbConnection()
	if err != nil {
		return nil, 0, common.DbErr
	}
	offset := num * count
	bills, tableCount, err := model.MGetBill(db, offset, count)
	if err != nil {
		return nil, 0, common.CustomErr(common.DbErr, err)
	}
	var billList []*pb_gen.Bill
	for _, value := range bills {
		billList = append(billList, packBill(value))
	}
	return billList, tableCount, common.Success
}

func packBill(bill model.Bill) *pb_gen.Bill {
	return &pb_gen.Bill{
		Id:       bill.Id,
		UserId:   bill.UserId,
		ParkId:   bill.ParkId,
		Duration: bill.Duration,
		Charge:   int32(bill.Charge),
		Status:   pb_gen.BillStatus(bill.Status),
	}
}
