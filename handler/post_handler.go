package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"project-api-golang/dto"
	"project-api-golang/errorhandler"
	"project-api-golang/helper"
	"project-api-golang/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *postHandler {
	return &postHandler{
		service: service,
	}
}

func (h *postHandler) Create(c *gin.Context) {
	var post dto.PostRequest
	if err := c.ShouldBind(&post); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if post.Picture != nil {
		if err := os.MkdirAll("/public/picture", 0755); err != nil {
			errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)
		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, newFileName)
	}

	userID, _ := c.Get("userID")
	post.UserID = userID.(int)

	if err := h.service.Create(&post); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Sukses mengupload",
	})

	c.JSON(http.StatusCreated, res)

}
