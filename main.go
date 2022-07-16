package main

import (
	"fmt"
	"github.com/mattermost/mattermost-server/v5/model"
	"log"
	"net/http"
)

const URL = "http://your-mattermost-url.com"
const ApiUrl = "http://your-mattermost-url.com/api/v4"
const PersonalAccessToken = ""

func main() {
	c := model.Client4{
		Url:        URL,
		ApiUrl:     ApiUrl,
		HttpClient: &http.Client{},
		AuthToken:  PersonalAccessToken,
		AuthType:   model.HEADER_BEARER,
		HttpHeader: nil,
	}

	// すでに登録されている絵文字を 10件取得する
	emojiList, resp := c.GetEmojiList(0, 10)
	if resp.Error != nil {
		log.Fatalln(resp.Error)
	}

	fmt.Println("登録されている絵文字の一例")
	for _, emoji := range emojiList {
		fmt.Println(emoji.Name)
	}

	// me で自分自身のデータを取得することができる
	user, resp := c.GetUser("me", "")
	if resp.Error != nil {
		log.Fatalln(resp.Error)
	}

	fmt.Printf("%v\n", user.ToJson())

}
