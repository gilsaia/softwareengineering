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
	Success  = BgErr{0, ""}
	ParamErr = BgErr{40001, "Param Error!"}
	DbErr    = BgErr{50001, "Db Error!"}
)

func CustomErr(err BgErr, custom error) BgErr {
	return BgErr{ErrNo: err.ErrNo, ErrMsg: err.ErrMsg + custom.Error()}
}
