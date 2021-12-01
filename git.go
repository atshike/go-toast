package main

import (
	"fmt"
	"github.com/go-playground/webhooks/v6/github"
	"net/http"
)

const (
	path = "F:\\go\\go-toast"
)

func main() {
	hook, _ := github.New()
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("a1")
		payload, err := hook.Parse(r, github.PushEvent)
		if err != nil {
			fmt.Println("a2")
			return
		}

		switch payload.(type) {
		case github.PushPayload:
			fmt.Println("a3")
			pushEvent := payload.(github.PushPayload)
			fmt.Println(pushEvent)
		}
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Print("error!")
		return
	}
}
