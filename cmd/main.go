package main

import (
	"fmt"
	"microblog/internal/handlers"
	"microblog/internal/service"
	"net/http"
)

func main() {
	svc := service.NewService()
	h := handlers.NewHandler(svc)
	http.HandleFunc("/register", h.Register)
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreatePost(w, r)

		} else {
			h.GetPosts(w, r)
		}
	})
	http.HandleFunc("/posts/", h.LikePost)
	fmt.Println("серевер слушает порт 9091")
	http.ListenAndServe(":9091", nil)
}
