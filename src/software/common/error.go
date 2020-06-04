package common

type BgErr struct {
	ErrNo  int32
	ErrMsg string
}

func (e BgErr) Error() string {
	return e.ErrMsg
}

func (e BgErr) Is(err BgErr) bool {
	return e.ErrNo == err.ErrNo
}

func (e BgErr) ErrNoMsg() (int32, string) {
	return e.ErrNo, e.ErrMsg
}

var (
	Success       = BgErr{0, ""}
	BindErr       = BgErr{10001, "Bind Code Wrong!"}
	UserErr       = BgErr{10002, "User Wrong!"}
	PasswordErr   = BgErr{10003, "Password Wrong!"}
	OccupyErr     = BgErr{20001, "This Port Has Been Used!"}
	AbandonErr    = BgErr{20002, "This Port Is Abandoned!"}
	BillErr       = BgErr{20003, "User Have Bill To Pay!"}
	WrongBillErr  = BgErr{20004, "Not Your Bill!"}
	ParamErr      = BgErr{40001, "Param Error!"}
	InternalErr   = BgErr{40002, "Internal Error!"}
	TokenErr      = BgErr{40003, "Token Error!"}
	PermissionErr = BgErr{40004, "No Permission Error!"}
	DbErr         = BgErr{50001, "Db Error!"}
	SmsErr        = BgErr{50002, "Sms Error!"}
)

func CustomErr(err BgErr, custom error) BgErr {
	return BgErr{ErrNo: err.ErrNo, ErrMsg: err.ErrMsg + custom.Error()}
}
