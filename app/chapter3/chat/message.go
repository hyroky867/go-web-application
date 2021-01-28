package main

import "time"

// messageはひとつのメッセージを表します。
type message struct {
	Name      string
	Message   string
	When      time.Time
	AvatarURL string
}
