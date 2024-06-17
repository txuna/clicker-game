extends Node2D

@onready var popup = $Panel
@onready var mining_btn = $MiningBtn

@onready var bar = $LevelBar

@onready var left_pos = $left
@onready var right_pos = $right

@onready var coin_value = $CoinValue

var _max_coin = 0 
var _coin = 0

# Called when the node enters the scene tree for the first time.
func _ready():
	popup.visible = false
	Network.request(":9003/user", {
		"user_id": Global.user_id,
		"token": Global.token
	}, _on_userinfo_response)


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	pass


func display_star():
	pass


func display_coin(coin):
	_coin += coin
	coin_value.text = "{coin}/{max_coin}".format({
		"coin" : str(_coin),
		"max_coin" : str(_max_coin)
	})
	
	

# 초기 유저 정보 요청
func _on_userinfo_response(result, res_code, headers, body, http_request):
	http_request.queue_free()
	if res_code != 200:
		return	
		
	var json = JSON.parse_string(body.get_string_from_utf8())
	if json["error_code"] != Global.ERROR_NONE:
		var code = json["error_code"]
		#print(code)
		popup.set_content(Global.error_msg[int(code)])
		popup.visible = true
		return

	var max_coin = int(json["max_coin"])
	var coin = int(json["coin"])
	
	bar.min_value = 0
	bar.max_value = max_coin
	bar.value = coin
	
	_max_coin = max_coin
	_coin = coin
	display_coin(0)


func _on_mining_response(result, res_code, headers, body, http_request):
	http_request.queue_free()
	if res_code != 200:
		return
	
	var json = JSON.parse_string(body.get_string_from_utf8())
	if json["error_code"] != Global.ERROR_NONE:
		var code = json["error_code"]
		popup.set_content(Global.error_msg[int(code)])
		popup.visible = true
		return
		
	var coin = int(json["coin"])
	bar.value += coin
	display_coin(coin)
	

func _on_mining_pressed():
	Network.request(":9003/mining", {
		"user_id" : Global.user_id, 
		"token" : Global.token
	}, _on_mining_response)


func _on_panel_btn_click():
	visible = false
