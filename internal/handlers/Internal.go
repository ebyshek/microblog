package handlers

import (
	"encoding/json"
	"microblog/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "не то мне надо", http.StatusMethodNotAllowed)
		return
	}
	var newuser struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&newuser); err != nil {
		http.Error(w, "ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	user := h.service.RegisterUser(newuser.Name)
	if user == nil {
		http.Error(w, "есть такой пользователь", http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "дерьмо какое-то", http.StatusBadRequest)
		return
	}

	var newpost struct {
		Text     string `json:"text"`
		AuthorID int    `json:"author_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&newpost); err != nil {
		http.Error(w, "ошибка", http.StatusBadRequest)
		return
	}
	post := h.service.CreatePost(newpost.AuthorID, newpost.Text)
	if post == nil {
		http.Error(w, "оштбка", http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	getposts := h.service.GetPosts()
	if getposts == nil {
		http.Error(w, "постов нет", http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getposts)
}

func (h Handler) LikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "да бля не то", http.StatusBadRequest)
		return
	}
	epta := strings.Split(r.URL.Path, "/")
	epta2 := strings.Split(r.URL.Path, "/")
	postIDstr := epta[2]
	postID, err := strconv.Atoi(postIDstr)
	userIDstr := epta2[4]
	userID, err := strconv.Atoi(userIDstr)

	if err != nil {
		http.Error(w, "нужно число", http.StatusBadRequest)
		return
	}
	postLike := h.service.Like(postID, userID)
	if !postLike {
		http.Error(w, "поста или чела нет", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("лайк"))
}
