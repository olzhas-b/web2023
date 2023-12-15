package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/olzhas-b/social-media/internal/helper"
	"github.com/olzhas-b/social-media/internal/models"
	"net/http"
	"strconv"
)

func (srv *Server) GetPosts(c *gin.Context) {
	searchText := c.Query("search")

	results, err := srv.core.Posts().List(
		searchText,
		helper.GetUserID(c),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (srv *Server) GetPost(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, err := srv.core.Posts().ByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (srv *Server) AddPosts(c *gin.Context) {
	var post models.Posts
	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.UserID = helper.GetUserID(c)
	post.Author = helper.GetUsername(c)
	err = srv.core.Posts().Add(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (srv *Server) RemovePosts(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = srv.core.Posts().Remove(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
