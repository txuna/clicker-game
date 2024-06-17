extends Node

const url = "http://127.0.0.1"
const headers = ["Content-Type: application/json"]
const GET = 1
const POST = 2

# Called when the node enters the scene tree for the first time.
func _ready():
	pass # Replace with function body.

func request(path, data, callback):
	var json = JSON.stringify(data)
	var http = HTTPRequest.new()
	http.timeout = 1.0
	add_child(http)
	http.request_completed.connect(callback.bind(http))
	http.request(url+path, headers, HTTPClient.METHOD_POST, json)
