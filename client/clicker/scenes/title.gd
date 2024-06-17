extends Control

@onready var login_btn = $TextureButton
@onready var join_btn = $TextureButton2
@onready var user_id_editor = $userIdEditor
@onready var user_pw_editor = $userPwEditor

@onready var popup = $Panel

# Called when the node enters the scene tree for the first time.
func _ready():
	pass # Replace with function body.


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	pass
	

func _on_texture_login_pressed():
	login_btn.disabled = true
	Network.request(":9001/login", {
		"user_id" : user_id_editor.text,
		"user_pw" : user_pw_editor.text
	}, _on_login_response)
	
	
func _on_login_response(result, res_code, headers, body, http_request):
	http_request.queue_free()
	login_btn.disabled = false
	
	if res_code != 200:
		return
	
	var json = JSON.parse_string(body.get_string_from_utf8())
	if json["error_code"] != Global.ERROR_NONE:
		var code = json["error_code"]
		popup.set_content(Global.error_msg[int(code)])
		popup.visible = true
		return
	
	Global.user_id = json["user_id"]
	Global.token = json["token"]
	get_tree().change_scene_to_file("res://scenes/main.tscn")


func _on_join_response(result, res_code, headers, body, http_request):
	http_request.queue_free()
	join_btn.disabled = false
	
	if res_code != 200:
		return
	
	var json = JSON.parse_string(body.get_string_from_utf8())
	if json["error_code"] != Global.ERROR_NONE:
		var code = json["error_code"]
		popup.set_content(Global.error_msg[int(code)])
		popup.visible = true
		return

	popup.set_content("회원가입이 완료되었습니다.")
	popup.visible = true 


func _on_texture_join_pressed():
	join_btn.disabled = true
	Network.request(":9001/join", {
		"user_id" : user_id_editor.text,
		"user_pw" : user_pw_editor.text
	}, _on_join_response)


func _on_panel_btn_click():
	popup.visible = false
