package handlers

import (
	"github.com/yuki-inoue-eng/authenticator/internal/models/repositories"
)

type Handler struct {
	reps repositories.Repositories
}

func NewHandlers(repositories repositories.Repositories) Handler {
	return Handler{
		reps: repositories,
	}
}



