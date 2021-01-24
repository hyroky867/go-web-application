package main

type room struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	forward chan []byte

	// joinはチャットルームに参加しようとしているクライアントのためのチャネルです。
	join chan *client

	// leaveはチャットルームから退室しようとしているクライアントのチャネルです。
	leave chan *client

	// clients には在籍している術とのクライアントが保持されます。
	clients map[*client]bool
}
