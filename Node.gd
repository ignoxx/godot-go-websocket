extends Node

var ws: WebSocketClient = WebSocketClient.new()

func _ready():
	print("connecting to go-ws-server")
	ws.connect("connection_established", self, "_connection_established")
	ws.connect("connection_closed", self, "_connection_closed")
	ws.connect("connection_error", self, "_connection_error")
	ws.connect("data_received", self, "_connection_data")

	var err = ws.connect_to_url("localhost:2312")

	if err != OK:
		print("failed to connect to server " + str(err))

func _connection_established():
	print("connection established")

func _connection_closed():
	print("connection closed")

func _connection_error():
	print("connection error")

func _connection_data():
	print("connection data")
	var packet = ws.get_peer(1).get_packet()
	var buffer = StreamPeerBuffer.new()
	buffer.set_data_array(packet)

	var packet_id = buffer.get_u8()
	print("received u8: " + str(packet_id))

func _process(delta):
	ws.poll()


func _on_Button_pressed():
	var buffer = StreamPeerBuffer.new()
	buffer.put_u8(42)
	buffer.put_u8(69)
	buffer.put_u8(255)
	buffer.put_u8(0)
	buffer.put_string("testtesttesttesttesttesttesttesttesttestte")
	var data =  buffer.get_data_array()

	ws.get_peer(1).put_packet(data)
