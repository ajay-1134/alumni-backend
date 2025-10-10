package service

import (
	"log"

	"github.com/ajay-1134/alumni-backend/internal/dto"
	"github.com/ajay-1134/alumni-backend/internal/ports/repository"
	"github.com/ajay-1134/alumni-backend/internal/ports/service"
)

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) service.PostService {
	return &postService{repo: repo}
}

func (ps *postService) CreatePost(postReq *dto.PostDto) error {
	return ps.repo.Create(dtoToPost(postReq))
}

func (ps *postService) GetAllPosts() ([]*dto.PostDto, error) {
	posts, err := ps.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var postsDto []*dto.PostDto
	for _, p := range posts {
		postsDto = append(postsDto, postToDto(p))
	}

	return postsDto, nil
}

func (ps *postService) GetAllPostsWithUserId(userId uint) ([]*dto.PostDto, error) {
	posts, err := ps.repo.GetAllWithUserId(userId)
	if err != nil {
		return nil, err
	}
	var postsDto []*dto.PostDto
	for _, p := range posts {
		postsDto = append(postsDto, postToDto(p))
	}

	return postsDto, nil
}

func (ps *postService) UpdatePost(postID uint, postDto *dto.PostDto) error {
	post,err := ps.repo.GetPostById(postID)
	if err != nil {
		log.Printf("post does not exist with id %v",postID)
		return err
	}

	return ps.repo.Update(post,buildUpdatedPost(postDto))
}


func (ps *postService) DeletePost(postID uint) error {
	_,err := ps.repo.GetPostById(postID)
	if err != nil {
		log.Printf("post does not exist with id %v",postID)
		return err
	}

	return ps.repo.Delete(postID)
}