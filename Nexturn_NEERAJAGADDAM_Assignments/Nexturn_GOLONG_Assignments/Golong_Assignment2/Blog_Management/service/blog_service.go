package service

import (
	"blog-management/model"
	"blog-management/repository"
)

type BlogService struct {
	BlogRepo *repository.BlogRepository
}

// NewBlogService initializes a new BlogService
func NewBlogService(blogRepo *repository.BlogRepository) *BlogService {
	return &BlogService{BlogRepo: blogRepo}
}

// CreateBlog handles business logic for creating a blog
func (service *BlogService) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	// Add any business logic here, if needed
	return service.BlogRepo.CreateBlog(blog)
}

// GetBlog handles business logic for fetching a blog by ID
func (service *BlogService) GetBlog(id int) (*model.Blog, error) {
	return service.BlogRepo.GetBlog(id)
}

// GetAllBlogs handles business logic for fetching all blogs
func (service *BlogService) GetAllBlogs() ([]model.Blog, error) {
	return service.BlogRepo.GetAllBlogs()
}

// UpdateBlog handles business logic for updating a blog
func (service *BlogService) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	return service.BlogRepo.UpdateBlog(blog)
}

// DeleteBlog handles business logic for deleting a blog
func (service *BlogService) DeleteBlog(id int) error {
	return service.BlogRepo.DeleteBlog(id)
}
