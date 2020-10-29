package book

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"time"
)
 //var rnd =rand.New(rand.NewSource(time.Now().UnixNano()))

type message struct{
	Data string `json:"data"`
	Type string `json:"type"`
}
func socketHandler(ws *websocket.Conn){
	ch :=make(chan struct{})
	go func(ws *websocket.Conn) {
		defer close(ch)
		for{
			var msg message
			err :=websocket.JSON.Receive(ws,&msg)
			if err != nil {
				log.Fatal(err)
				break
			}
			fmt.Printf("Received message: %s\n", msg.Data)
		}

	}(ws)
	loop:
	for{
		select{
		case <-ch:
			fmt.Println("Connection was terminated!")
			break loop
		default:
			//id :=rnd.Intn(7)+1
			book := Book{ID:1,BookName:"XYZ",Writers:[]string{"Xyz","XYZ2"}, CopiesAvailable:20}
			err :=websocket.JSON.Send(ws,&book)
			if err != nil{
				log.Fatal(err)
				break
			}
			time.Sleep(10 *time.Second)
		}

	}

}