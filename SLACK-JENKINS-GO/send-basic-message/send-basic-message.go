package main

import (
	"fmt"
	
	"github.com/slack-go/slack"
)

func main() {
	api := slack.New("xoxb-4871465838097-4882677902880-xa2deKSX6EtH17BTG2fpQ8Xv")

	channelID, timestamp, err := api.PostMessage(
		"C04R5SA22KF",
		slack.MsgOptionText("Hello word",false),
	)

	if err != nil{
		fmt.Println("%s\n",err)
		return
	}

	 fmt.Printf("Message sent successfully to channel %s at %s",channelID ,timestamp)
}