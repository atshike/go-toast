package main

import (
	"fmt"
	"github.com/go-playground/webhooks/v6/github"
	"net/http"
)

const (
	path = "F:\\webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret("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"))
	fmt.Print(hook)

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("doing...")
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		fmt.Print(payload)

		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}

		switch payload.(type) {
		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Print("error!")
		return 
	}
}
