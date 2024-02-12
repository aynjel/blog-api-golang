package services

import (
	"anggi.blog/domain/posts"
	"anggi.blog/utils/errors"
)

func CreatePost(post posts.Post) (*posts.Post, *errors.RestErr) {
	if err := post.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Message, "Service")
	}

	if err := post.Save(); err != nil {
		return nil, err
	}

	return &post, nil
}

func GetPosts() ([]posts.Post, *errors.RestErr) {
	post := posts.Post{}
	return post.GetPosts()
}

func GetPost(postID int64) (*posts.Post, *errors.RestErr) {
	post, err := posts.GetPost(postID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func UpdatePost(post posts.Post) (*posts.Post, *errors.RestErr) {
	if err := post.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Message, "Service")
	}

	if err := post.Update(); err != nil {
		return nil, err
	}

	return &post, nil
}

func DeletePost(postID int64) *errors.RestErr {
	post := posts.Post{ID: postID}
	return post.Delete()
}
