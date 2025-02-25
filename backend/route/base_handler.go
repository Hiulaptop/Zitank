package route

import (
	"Zitank/models"
	"Zitank/repositories"

	"github.com/go-chi/jwtauth/v5"
	"github.com/jmoiron/sqlx"
)

type BaseHandler struct {
	albumRepository models.AlbumRepository
	musicRepository models.MusicRepository
	orderRepository models.OrderRepository
	postRepository  models.PostRepository
	roomRepository  models.RoomRepository
	userRepositor   models.UserRepository
	TokenAuth       *jwtauth.JWTAuth
}

func NewBaseHandler(DB *sqlx.DB, TokenAuth *jwtauth.JWTAuth) *BaseHandler {
	return &BaseHandler{
		albumRepository: repositories.NewAlbumRepo(DB),
		musicRepository: repositories.NewMusicRepo(DB),
		orderRepository: repositories.NewOrderRepo(DB),
		postRepository:  repositories.NewPostRepo(DB),
		roomRepository:  repositories.NewRoomRepo(DB),
		userRepositor:   repositories.NewUserRepo(DB),
		TokenAuth:       TokenAuth,
	}
}
