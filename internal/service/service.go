package service

import (
	"fmt"
	"microblog/internal/models"
)

type Service struct {
	users      map[int]*models.User
	posts      []*models.Post
	nextUserID int
	nextPostID int
}

func NewService() *Service {
	return &Service{
		users:      make(map[int]*models.User),
		posts:      make([]*models.Post, 0),
		nextUserID: 1,
		nextPostID: 1,
	}
}

func (s *Service) RegisterUser(userName string) *models.User {
	for _, user := range s.users {
		if user.Name == userName {
			fmt.Println("такое имя занято")
			return nil
		}
	}
	NewUser := &models.User{
		Name: userName,
		ID:   s.nextUserID,
	}
	s.users[s.nextUserID] = NewUser
	s.nextUserID++

	return NewUser

}

func (s *Service) CreatePost(IDAuthor int, text string) *models.Post {
	if _, ok := s.users[IDAuthor]; !ok {
		fmt.Println("такой пользователь не найден")
		return nil
	}
	NewPost := &models.Post{
		ID:       s.nextPostID,
		IdAuthor: IDAuthor,
		Text:     text,
		Like:     make([]int, 0),
	}
	s.posts = append(s.posts, NewPost)
	s.nextPostID++
	return NewPost
}

func (s *Service) Like(postID int, userID int) bool {
	if _, ok := s.users[userID]; !ok {
		fmt.Println("нет пользователя")
		return false
	}
	for i := range s.posts {
		if s.posts[i].ID == postID {
			for _, ID := range s.posts[i].Like {
				if ID == userID {
					fmt.Println("пользователь ставил лайк")
					return false

				}
			}
			s.posts[i].Like = append(s.posts[i].Like, userID)
			fmt.Print("лайк поставлен")

			fmt.Println("кол-во лайков", len(s.posts[i].Like))
			return true
		}
	}
	fmt.Println("пост не найден")
	return false

}

func (s *Service) GetPosts() []*models.Post {
	return s.posts
}
