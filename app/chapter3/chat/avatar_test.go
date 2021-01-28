package main

func TestAuthAvatar(t *testing) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetBeginAuthURL(client)

	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、AuthAvatar.getAvatarURLはErrNoAvatarURLを返すべきです")
	}

	// 値をセット
	testURL := "http://url-to-avatar/"
	client.userData = map[string]interface{}{
		"avatar_url": testURL,
	}
	url, err = authAvatar.GetBeginAuthURL(client)
	if err != nil {
		t.Error("値が存在する場合、AuthAvatar.GetAvatarURLはエラーを返すべきではありません")
	} else {
		if url != testURL {
			t.Error("AuthAvatar.GetAvatarURLは正しいURLをかえすべきです")
		}
	}
}
