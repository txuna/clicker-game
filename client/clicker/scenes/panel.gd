extends CanvasLayer

@onready var content_label = $ColorRect/Content

signal btn_click

# Called when the node enters the scene tree for the first time.
func _ready():
	visible = false


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	pass

func set_content(str):
	content_label.text = str


func _on_button_pressed():
	emit_signal("btn_click")
