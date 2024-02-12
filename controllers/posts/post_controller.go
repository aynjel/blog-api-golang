package posts

import (
	"net/http"
	"strconv"

	"anggi.blog/domain/posts"
	"anggi.blog/services"
	"anggi.blog/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post posts.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		restErr := errors.NewBadRequestError(err.Error(), "Controller")
		c.IndentedJSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreatePost(post)
	if saveErr != nil {
		c.IndentedJSON(saveErr.Status, saveErr)
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"post": result})
}

func GetPosts(c *gin.Context) {
	posts, getErr := services.GetPosts()
	if getErr != nil {
		c.IndentedJSON(getErr.Status, getErr)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"posts": posts})
}

func GetPost(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid post id", "Controller")
		c.IndentedJSON(restErr.Status, restErr)
		return
	}
	post, getErr := services.GetPost(postID)
	if getErr != nil {
		c.IndentedJSON(getErr.Status, getErr)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"post": post})
}

func UpdatePost(c *gin.Context) {
	var post posts.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		restErr := errors.NewBadRequestError(err.Error(), "Controller")
		c.IndentedJSON(restErr.Status, restErr)
		return
	}
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid post id", "Controller")
		c.IndentedJSON(restErr.Status, restErr)
		return
	}
	post.ID = postID
	result, updateErr := services.UpdatePost(post)
	if updateErr != nil {
		c.IndentedJSON(updateErr.Status, updateErr)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"post": result})
}

func DeletePost(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid post id", "Controller")
		c.IndentedJSON(restErr.Status, restErr)
		return
	}
	deleteErr := services.DeletePost(postID)
	if deleteErr != nil {
		c.IndentedJSON(deleteErr.Status, deleteErr)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"status": "deleted"})
}
