package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	// connect to server
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://localhost:8080/ws")
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()

	go func() {
		defer wg.Done()
		for {
			// receive data from server
			data, _, err := wsutil.ReadServerData(conn)
			if err != nil {
				log.Printf("error reading data from server: %s", err)
				return
			}
			fmt.Printf("GOT: %s\n", data)
		}
	}()

	go func() {
		defer wg.Done()
		counter := 0
		for {
			time.Sleep(time.Second)
			counter++
			// send data to server
			_ = wsutil.WriteClientText(conn, []byte(fmt.Sprintf("MESSAGE %d", counter)))
			fmt.Printf("SENT: MESSAGE %d\n", counter)
		}
	}()
}
