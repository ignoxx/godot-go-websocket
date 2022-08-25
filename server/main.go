package main

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	fmt.Println("starting to listen on port 8080")
	http.ListenAndServe("0.0.0.0:2312", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			fmt.Println("failed to upgrade http")
			return
		}
		fmt.Println("client connected " + fmt.Sprint(conn))
		go func() {
			defer conn.Close()
			fmt.Println("listening for messages")
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					fmt.Println("failed to read client data")
					return
				}
				fmt.Println("data received " + fmt.Sprint(msg))
				fmt.Println(fmt.Sprint(msg[0:1]))
				fmt.Println(fmt.Sprint(msg[1:2]))
				fmt.Println(fmt.Sprint(msg[2:3]))
				fmt.Println(fmt.Sprint(msg[3:4]))
				fmt.Println(fmt.Sprint(msg[4:8]))
				fmt.Println(fmt.Sprint(string(msg[8:])))



				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					fmt.Println("failed to write server message")
					return
				}
			}
		}()
	}))
}
