package main

// import (
// 	"fmt"
// 	"log"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// const (
// 	token                     = "7165645088:AAEOqDNdEMuCHMyHlCbY1lki6wgdLpkjXJ0" // Replace with your bot API token
// 	groupID            int64  = 7112653641                                       // Replace with your group ID
// 	targetUserID              = 7112653641
// 	SuperGroupUsername string = "MartinLin121" // Replace with the user ID you want to check
// )

// func main() {
// 	// Create a new updater

// 	bot, err := tgbotapi.NewBotAPI(token)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ChatConfigWithUser := tgbotapi.ChatConfigWithUser{ChatID: groupID,
// 		SuperGroupUsername: SuperGroupUsername,
// 		UserID:             targetUserID}

// 	catconfig := tgbotapi.GetChatMemberConfig{
// 		ChatConfigWithUser: ChatConfigWithUser,
// 	}
// 	// Get group chat members
// 	members, err := bot.GetChatMember(catconfig)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check if target user is a member of the group
// 	isMember := false
// 	fmt.Println(members)
// 	// for _, member := range members.User.UserName {
// 	// 	fmt.Println(member)
// 	// }

// 	if isMember {
// 		fmt.Printf("User %d is a member of the group.\n", targetUserID)
// 	} else {
// 		fmt.Printf("User %d is not a member of the group.\n", targetUserID)
// 	}
// }

// package main

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
)

func main() {

	b, err := bot.New("7165645088:AAEOqDNdEMuCHMyHlCbY1lki6wgdLpkjXJ0")
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}
	param := bot.GetChatMemberParams{
		ChatID: "-1002138153950",
		UserID: 7112653641,
	}
	param1 := bot.GetChatParams{
		ChatID: "-1002138153950",
	}

	user, _ := b.GetChatMember(context.Background(), &param)
	chat, _ := b.GetChat(context.Background(), &param1)
	//bot, _ := b.GetMe(context.Background())
	//GetMe(context.Background())
	fmt.Printf("chat: %#v\n", chat)
	fmt.Printf("User: %#v\n", user)
	// fmt.Printf("bot: %#v\n", bot)
}
