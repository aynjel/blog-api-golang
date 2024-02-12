package posts

import (
	"anggi.blog/datasources/sqlite/crud_db"
	"anggi.blog/utils/errors"
)

const (
	queryInsertPost = "INSERT INTO posts(title, content, user_id) VALUES(?, ?, ?)"
	queryGetPosts   = "SELECT id, title, content, user_id FROM posts"
	queryGetPost    = "SELECT id, title, content, user_id FROM posts WHERE id = ?"
	queryUpdatePost = "UPDATE posts SET title=?, content=? WHERE id=?"
	queryDeletePost = "DELETE FROM posts WHERE id=?"
)

func (post *Post) Save() *errors.RestErr {
	stmt, err := crud_db.Client.Begin()
	if err != nil {
		return errors.NewInternalServerError("error when trying to save post")
	}
	defer stmt.Rollback()

	insertResult, saveErr := stmt.Exec(queryInsertPost, post.Title, post.Content, post.UserID)
	if saveErr != nil {
		return errors.NewInternalServerError("error when trying to save post")
	}

	postID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error when trying to save post")
	}
	post.ID = postID

	err = stmt.Commit()
	if err != nil {
		return errors.NewInternalServerError("error when trying to save post")
	}

	return nil
}

func (post *Post) Get() *errors.RestErr {
	stmt, err := crud_db.Client.Prepare(queryGetPost)
	if err != nil {
		return errors.NewInternalServerError("error when trying to get post")
	}
	defer stmt.Close()

	result := stmt.QueryRow(post.ID)
	if getErr := result.Scan(&post.ID, &post.Title, &post.Content, &post.UserID); getErr != nil {
		return errors.NewInternalServerError("error when trying to get post")
	}

	return nil
}

func (post *Post) Update() *errors.RestErr {
	stmt, err := crud_db.Client.Prepare(queryUpdatePost)
	if err != nil {
		return errors.NewInternalServerError("error when trying to update post")
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Content, post.ID)
	if err != nil {
		return errors.NewInternalServerError("error when trying to update post")
	}

	return nil
}

func (post *Post) Delete() *errors.RestErr {
	stmt, err := crud_db.Client.Prepare(queryDeletePost)
	if err != nil {
		return errors.NewInternalServerError("error when trying to delete post")
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.ID)
	if err != nil {
		return errors.NewInternalServerError("error when trying to delete post")
	}

	return nil
}
