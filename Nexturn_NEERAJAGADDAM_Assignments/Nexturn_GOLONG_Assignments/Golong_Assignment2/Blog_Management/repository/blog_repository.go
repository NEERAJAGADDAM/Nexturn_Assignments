package repository

import (
	"blog-management/model"
	"database/sql"
	"fmt"
)

type BlogRepository struct {
	DB *sql.DB
}

// NewBlogRepository initializes a new repository
func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{DB: db}
}

// CreateBlog inserts a new blog into the database
func (repo *BlogRepository) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare("INSERT INTO blogs (title, content, author, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(blog.Title, blog.Content, blog.Author, blog.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	blog.ID = int(id)
	return blog, nil
}

// GetBlog retrieves a blog by ID
func (repo *BlogRepository) GetBlog(id int) (*model.Blog, error) {
	row := repo.DB.QueryRow("SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?", id)

	blog := &model.Blog{}
	err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no blog found with id %d: %w", id, err)
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}
	return blog, nil
}

// GetAllBlogs retrieves all blogs
func (repo *BlogRepository) GetAllBlogs() ([]model.Blog, error) {
	rows, err := repo.DB.Query("SELECT id, title, content, author, timestamp FROM blogs")
	if err != nil {
		return nil, fmt.Errorf("failed to query rows: %w", err)
	}
	defer rows.Close()

	var blogs []model.Blog
	for rows.Next() {
		var blog model.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return blogs, nil
}

// UpdateBlog updates an existing blog
func (repo *BlogRepository) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare("UPDATE blogs SET title = ?, content = ?, author = ?, timestamp = ? WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(blog.Title, blog.Content, blog.Author, blog.Timestamp, blog.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	return blog, nil
}

// DeleteBlog deletes a blog by ID
func (repo *BlogRepository) DeleteBlog(id int) error {
	stmt, err := repo.DB.Prepare("DELETE FROM blogs WHERE id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}
