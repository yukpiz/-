package main

import (
	firebase "firebase.google.com/go"
	firebasedb "firebase.google.com/go/db"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"time"
)

func main() {
	opt := option.WithCredentialsFile("./redish-dev-firebase-adminsdk-hjd6d-91e202b209.json")
	config := &firebase.Config{
		DatabaseURL: "https://redish-dev.firebaseio.com/",
	}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(client)

	//ref(client)
	//set1(client)
	//set2(client) //rooms
	//push1(client) //messages
	update1(client)
	//get1(client)
	//get2(client)
	//get3(client)
	//get4(client)

	//transaction1(client)
	//transaction2(client)
}

func ref(client *firebasedb.Client) {
	ref := client.NewRef("messages/room_key1/message_id/text")
	log.Println(ref)
	//&{text /messages/room_key1/message_id/text [messages room_key1 message_id text] 0xc420192b10}

	ref = client.NewRef("messages").Child("room_key1").Child("message_id").Child("text")
	log.Println(ref)
	//&{text /messages/room_key1/message_id/text [messages room_key1 message_id text] 0xc420192b10}

	ref = client.NewRef("messages/room_key1/message_id/text").Parent().Child("text")
	log.Println(ref)
	//&{text /messages/room_key1/message_id/text [messages room_key1 message_id text] 0xc420192b10}

	ref, err := ref.Push(context.Background(), "(」・ω・)」うー!")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ref)
	// &{-LIIk8kp2HLKobUA-kxc /messages/room_key1/message_id/text/-LIIk8kp2HLKobUA-kxc [messages room_key1 message_id text -LIIk8kp2HLKobUA-kxc] 0xc4201acb10}
}

func push1(client *firebasedb.Client) {
	type Message struct {
		UserId int    `json:"user_id"`
		Text   string `json:"text"`
	}

	ref := client.NewRef("messages/room_key")
	for i := 1; i <= 3; i += 1 {
		message := Message{
			UserId: i,
			Text:   fmt.Sprintf("text%d", i),
		}
		childRef, err := ref.Push(context.Background(), message)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(childRef)
	}
}

func set2(client *firebasedb.Client) {
	type Room struct {
		PartnerId       int    `json:"partner_id"`
		LastMessageText string `json:"last_message_text"`
	}

	ref := client.NewRef("rooms")
	rooms := map[string]Room{}
	for i := 1; i <= 3; i += 1 {
		room := Room{
			PartnerId:       i + 10,
			LastMessageText: fmt.Sprintf("text%d", i),
		}
		rooms[fmt.Sprintf("user%d", i)] = room
	}
	err := ref.Set(context.Background(), rooms)
	if err != nil {
		log.Fatal(err)
	}
}

func set1(client *firebasedb.Client) {
	type Message struct {
		MessageId int    `json:"message_id"`
		Text      string `json:"text"`
	}

	ref := client.NewRef("messages/room_key1")
	message1 := Message{
		MessageId: 1,
		Text:      "text1",
	}
	err := ref.Set(context.Background(), message1)
	if err != nil {
		log.Fatal(err)
	}

	message2 := Message{
		MessageId: 2,
		Text:      "text2",
	}
	err = ref.Set(context.Background(), message2)
	if err != nil {
		log.Fatal(err)
	}
}

func update1(client *firebasedb.Client) {
	ref := client.NewRef("rooms/user1")
	err := ref.Update(context.Background(), map[string]interface{}{
		"last_message_text": "updated!",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func transaction1(client *firebasedb.Client) {
	type Room struct {
		PartnerId       int    `json:"partner_id"`
		LastMessageText string `json:"last_message_text"`
	}

	ref := client.NewRef("rooms")

	transaction := func(node firebasedb.TransactionNode) (interface{}, error) {
		var rooms map[string]Room
		node.Unmarshal(&rooms)

		newRoom := Room{
			PartnerId:       14,
			LastMessageText: "これはトランザクションで追加されたデータ",
		}
		rooms["user4"] = newRoom
		return rooms, nil
	}

	err := ref.Transaction(context.Background(), transaction)
	if err != nil {
		log.Fatal(err)
	}
}

func transaction2(client *firebasedb.Client) {
	type Room struct {
		PartnerId       int    `json:"partner_id"`
		LastMessageText string `json:"last_message_text"`
	}

	ref := client.NewRef("rooms")

	transaction := func(node firebasedb.TransactionNode) (interface{}, error) {
		time.Sleep(3 * time.Second) //トランザクションに時間をかける為のSleep()
		var rooms map[string]Room
		node.Unmarshal(&rooms)

		newRoom := Room{
			PartnerId:       14,
			LastMessageText: "これはトランザクションで更新されたデータ",
		}
		rooms["user4"] = newRoom
		return rooms, nil
	}

	go ref.Transaction(context.Background(), transaction)
	// push
	time.Sleep(3 * time.Second)
}

func get1(client *firebasedb.Client) {
	type Message struct {
		UserId int    `json:"user_id"`
		Text   string `json:"text"`
	}

	var messages map[string]Message
	ref := client.NewRef("messages/room_key")
	err := ref.Get(context.Background(), &messages)
	if err != nil {
		log.Fatal(err)
	}
	for key, message := range messages {
		fmt.Printf("%s: %+v\n", key, message)
	}
}

func get2(client *firebasedb.Client) {
	type Message struct {
		UserId int    `json:"user_id"`
		Text   string `json:"text"`
	}

	ref := client.NewRef("messages/room_key")
	results, err := ref.OrderByChild("user_id").GetOrdered(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		var message Message
		err := result.Unmarshal(&message)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s: %+v\n", result.Key(), message)
	}
}

func get3(client *firebasedb.Client) {
	type Message struct {
		UserId int    `json:"user_id"`
		Text   string `json:"text"`
	}

	var messages map[string]Message
	ref := client.NewRef("messages/room_key")
	err := ref.OrderByChild("text").Get(context.Background(), &messages)
	if err != nil {
		log.Fatal(err)
	}

	for key, message := range messages {
		log.Printf("%s: %+v\n", key, message)
	}
}

func get4(client *firebasedb.Client) {
	type Message struct {
		UserId int    `json:"user_id"`
		Text   string `json:"text"`
	}

	var messages map[string]Message
	ref := client.NewRef("messages/room_key")
	err := ref.OrderByChild("user_id").EqualTo(2).Get(context.Background(), &messages)
	if err != nil {
		log.Fatal(err)
	}

	for key, message := range messages {
		log.Printf("%s: %+v\n", key, message)
	}
}
