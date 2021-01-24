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

func (r *room) run() {
	// goroutineであれば他に影響がないため無限ループでも問題ない
	for {
		select {
		case client := <-r.join:
			// 参加
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// すべてのクライアントにメッセージを転送
			for client := range r.clients {
				select {
				case client.send <- msg:
					// メッセージを送信
				default:
					// 送信に失敗
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
