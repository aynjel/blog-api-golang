package posts

import (
	"anggi.blog/utils/errors"
)

type Post struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (post *Post) Validate() *errors.RestErr {
	if post.Title == "" {
		return errors.NewBadRequestError("invalid post title", "Domain")
	}
	if post.Content == "" {
		return errors.NewBadRequestError("invalid post content", "Domain")
	}
	if post.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id", "Domain")
	}
	return nil
}
