package service

import (
	"project-api-golang/dto"
	"project-api-golang/entity"
	"project-api-golang/errorhandler"
	"project-api-golang/repository"
)

type PostService interface {
	Create(req *dto.PostRequest) error
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(req *dto.PostRequest) error {
	post := entity.Post{
		UserID: req.UserID,
	}

	if req.Picture != nil {
		post.Picture = &req.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
