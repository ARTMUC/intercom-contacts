package main

import (
	"fmt"
	"log"
	"net/http"

	webview "github.com/webview/webview_go"
	"intercom/frontend"
	"intercom/internal/router"
	"intercom/pkg/intercom"
)

func main() {
	router.RegisterStatic(frontend.Files)
	router.RegisterHandlers()

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.Bind("ListContacts", func() ([]intercom.ContactResponse, error) {
		fmt.Println("test")
		return intercom.NewClient().ListContacts()
	})
	w.SetTitle("AWS Intercom App")
	w.SetSize(1200, 900, webview.HintNone)

	w.Navigate("http://localhost:8080/index.html")
	w.Run()
}
