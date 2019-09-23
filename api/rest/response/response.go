package response

const (
	OK = 0
	BadRequest = 40002
	ErrorAuthCheckTokenInvalid = 40003
	ErrorAuthCheckTokenTimetout = 40004
	InternalServerError = 50000
)

var statusText = map[int]string {
	OK: 		"Success",
	BadRequest: "Bad request params",
	InternalServerError: "Server internal error",
}

type RespJsonObj struct {
	Code 	int  	`json:"code"`
	Msg  	string 	`json:"msg"`
	Data  	interface{} `json:"data"`
}

func StatusText(code int) string {
	return statusText[code]
}

func HttpStatusCode(code int) int {
	return code / 100
}