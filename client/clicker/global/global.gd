extends Node

const ERROR_NONE                  = 1000
const ERROR_INVALID_REQUEST       = 1001
const ERROR_NOT_EXIST_USER        = 1002
const ERROR_FAILED_GENERATE_TOKEN = 1003
const ERROR_ALREADY_EXIST_USER    = 1004
const ERROR_INVALID_USER          = 1005
const ERROR_INTERNAL_SERVER       = 1099

var error_msg = {
	ERROR_NONE: "문제 없음",
	ERROR_INVALID_REQUEST: "잘못된 요청입니다.",
	ERROR_NOT_EXIST_USER: "존재하지 않는 유저입니다.",
	ERROR_FAILED_GENERATE_TOKEN: "토큰 발급에 실패했습니다", 
	ERROR_ALREADY_EXIST_USER: "이미 등록된 유저입니다.",
	ERROR_INVALID_USER: "유효하지 않는 유저입니다.", 
	ERROR_INTERNAL_SERVER: "서버 내부적 오류입니다."
}

var token 
var user_id

# Called when the node enters the scene tree for the first time.
func _ready():
	token = ""
	user_id = ""
