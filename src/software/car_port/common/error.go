package common

type BgErr struct {
	errNo  int32
	errMsg string
}

func (e BgErr) Error() string {
	return e.errMsg
}

var (
	SuccessErr = BgErr{0, ""}
	ParamErr   = BgErr{40001, "Param Error!"}
	DbErr      = BgErr{50001, "Db Error!"}
)
