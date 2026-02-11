package service

import (
	"testing"
)

func TestRegisterUser(t *testing.T) {
	svc := NewService()

	tests := []struct {
		name     string
		username string
		isNil    bool
	}{
		{
			name:     "Успешная регистрация",
			username: "vanya",
			isNil:    false,
		},
		{
			name:     "Повторное имя (ошибка)",
			username: "vanya",
			isNil:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := svc.RegisterUser(tt.username)
			if (got == nil) != tt.isNil {
				t.Errorf("RegisterUser() для %s: получили nil=%v, хотели nil=%v", tt.username, got == nil, tt.isNil)
			}
		})
	}
}

func TestCreatePost(t *testing.T) {
	svc := NewService()
	user := svc.RegisterUser("author1")

	t.Run("Создание поста существующим юзером", func(t *testing.T) {
		post := svc.CreatePost(user.ID, "Привет, это мой пост")

		if post == nil {
			t.Fatal("Ожидали создание поста, получили nil")
		}
		if len(svc.GetPosts()) != 1 {
			t.Errorf("Ожидали 1 пост в списке, получили %d", len(svc.GetPosts()))
		}
	})

	t.Run("Ошибка: несуществующий автор", func(t *testing.T) {
		post := svc.CreatePost(999, "Текст")
		if post != nil {
			t.Error("Пост не должен создаваться для автора с ID 999")
		}
	})
}

func TestLike(t *testing.T) {
	svc := NewService()
	author := svc.RegisterUser("boss")
	post := svc.CreatePost(author.ID, "Лайкни меня")
	liker := svc.RegisterUser("fan")

	tests := []struct {
		name     string
		postID   int
		userID   int
		expected bool
	}{
		{
			name:     "Первый лайк",
			postID:   post.ID,
			userID:   liker.ID,
			expected: true,
		},
		{
			name:     "Повторный лайк",
			postID:   post.ID,
			userID:   liker.ID,
			expected: false,
		},
		{
			name:     "Лайк несуществующего поста",
			postID:   404,
			userID:   liker.ID,
			expected: false,
		},
		{
			name:     "Лайк от несуществующего юзера",
			postID:   post.ID,
			userID:   777,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := svc.Like(tt.postID, tt.userID)
			if got != tt.expected {
				t.Errorf("%s: Like() = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
} // бля не хуйя я по тестом не понял
