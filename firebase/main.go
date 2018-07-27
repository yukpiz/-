package main

import (
	firebase "firebase.google.com/go"
	firebasedb "firebase.google.com/go/db"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
)

func main() {
	opt := option.WithCredentialsFile("../../../../../Downloads/project-dev-firebase-adminsdk-hjd6d-11249c6fca.json")
	config := &firebase.Config{
		DatabaseURL: "https://project-dev.firebaseio.com/",
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
	//set(client)
	//push(client)
	//update(client)

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

func push(client *firebasedb.Client) {
	type Message struct {
		Text string `json:"text"`
	}

	ref := client.NewRef("messages/room_key1")
	message1 := Message{
		Text: "text1",
	}
	childRef1, err := ref.Push(context.Background(), message1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(childRef1)
	// &{-LIGxTD7kFRpywkYiBbf /messages/room_key1/-LIGxTD7kFRpywkYiBbf [messages room_key1 -LIGxTD7kFRpywkYiBbf] 0xc420192e10}

	message2 := Message{
		Text: "text2",
	}
	childRef2, err := ref.Push(context.Background(), message2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(childRef2)
	// &{-LIGxTFjzYnI58lJ9ymB /messages/room_key1/-LIGxTFjzYnI58lJ9ymB [messages room_key1 -LIGxTFjzYnI58lJ9ymB] 0xc420192e10}
}

func set(client *firebasedb.Client) {
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

func update(client *firebasedb.Client) {
	type Message struct {
		Text string `json:"text"`
	}

	//ref := client.NewRef("messages/room_key1")
	//messages := map[string]Message{
	//	"message1": Message{
	//		Text: "text1",
	//	},
	//	"message2": Message{
	//		Text: "text2",
	//	},
	//}
	//messages := map[string]string{
	//	"message1": "text1",
	//	"message2": "text2",
	//}
	//err := ref.Update(context.Background(), messages)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func transaction1(client *firebasedb.Client) {
	log.Println("Start!")

	type Message struct {
		MessageId      int    `json:"message_id"`
		Text           string `json:"text"`
		OptionalField1 string `json:"optional_field1"`
		OptionalField2 int    `json:"optional_field2"`
		OptionalField3 bool   `json:"optional_field3"`
	}
	ref := client.NewRef("test/transaction/messages1")

	// Push message 1
	ref.Push(context.Background(), Message{
		MessageId: 1,
		Text:      "Test Message1",
	})

	// Push message 2
	ref.Push(context.Background(), Message{
		MessageId: 2,
		Text:      "Test Message2",
	})

	transaction := func(node firebasedb.TransactionNode) (interface{}, error) {
		var messages map[string]Message
		node.Unmarshal(&messages)

		log.Printf("%+v\n", messages)
		return messages, nil
	}

	ref.Transaction(context.Background(), transaction)
	log.Println("Finished!")
}

func transaction2(client *firebasedb.Client) {
	log.Println("Start!")

	type Message struct {
		MessageId      int     `json:"message_id"`
		Text           string  `json:"text"`
		OptionalField1 *string `json:"optional_field1"`
		OptionalField2 *int    `json:"optional_field2"`
		OptionalField3 *bool   `json:"optional_field3"`
	}
	ref := client.NewRef("test/transaction/messages2")

	// Push message 1
	ref.Push(context.Background(), Message{
		MessageId: 1,
		Text:      "Test Message1",
	})

	// Push message 2
	ref.Push(context.Background(), Message{
		MessageId: 2,
		Text:      "Test Message2",
	})

	transaction := func(node firebasedb.TransactionNode) (interface{}, error) {
		var messages map[string]Message
		node.Unmarshal(&messages)

		log.Printf("%+v\n", messages)
		return messages, nil
	}

	ref.Transaction(context.Background(), transaction)
	log.Println("Finished!")
}
