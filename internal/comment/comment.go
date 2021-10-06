package comment

import "github.com/eatigo/gorm"

//Service - the struct for the comment service
type Service struct {
	DB *gorm.DB
}

//Comment - defines our comment struct
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

//CommentService - the interface for comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) (Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComent Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

//NewService - returns a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
