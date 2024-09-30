package router

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"strings"

	"intercom/internal/events/contact"
	"intercom/internal/handler"
	"intercom/pkg/intercom"
)

func HandlerWrapper[T any](h handler.HandlerFunc[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h(w, r)
		if resp.Status == 0 {
			switch {
			case resp.Error != nil, resp.InternalError != nil:
				resp.Status = http.StatusInternalServerError
				break
			case resp.Data != nil && r.Method == http.MethodPost:
				resp.Status = http.StatusCreated
				break
			case resp.Data != nil:
				resp.Status = http.StatusOK
			}
		}
		w.WriteHeader(resp.Status)

		if resp.Error != nil {
			fmt.Println(resp.Error)
			json.NewEncoder(w).Encode(resp.Error)
			return
		}

		if resp.InternalError != nil {
			fmt.Println(resp.Error)
			return
		}

		json.NewEncoder(w).Encode(resp.Data)
	}
}

func RegisterHandlers() {
	contactIndexHandler := contact.NewIndexHandler(intercom.NewClient())
	http.HandleFunc(
		fmt.Sprintf("%s %s", contactIndexHandler.Method(), contactIndexHandler.Route()),
		HandlerWrapper(contactIndexHandler.Handle),
	)
}

func RegisterStatic(files embed.FS) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := "app/build/index.html"
		if r.URL.Path != "/" {
			filePath = "app/build" + r.URL.Path
		}

		content, err := files.ReadFile(filePath)
		if err != nil {
			log.Println(err)
			http.NotFound(w, r)
			return
		}

		ext := fileExtension(filePath)
		if mimeType := mime.TypeByExtension(ext); mimeType != "" {
			w.Header().Set("Content-Type", mimeType)
		}

		w.Write(content)
	})
}

func fileExtension(filePath string) string {
	parts := strings.Split(filePath, ".")
	if len(parts) == 0 {
		return ""
	}
	return fmt.Sprintf(".%s", parts[len(parts)-1])
}
