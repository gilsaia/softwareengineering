package common

type BgErr struct {
	errNo  int32
	errMsg string
}

func (e BgErr) Error() string {
	return e.errMsg
}

func (e BgErr) Is(err BgErr) bool {
	return e.errNo == err.errNo
}

func (e BgErr) ErrNoMsg() (int32, string) {
	return e.errNo, e.errMsg
}

var (
	Success  = BgErr{0, ""}
	ParamErr = BgErr{40001, "Param Error!"}
	DbErr    = BgErr{50001, "Db Error!"}
)
