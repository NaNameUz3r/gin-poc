package controllers

import (
	"gin-poc/entity"
	"gin-poc/services"
	"gin-poc/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	GetAll() []entity.Video
}

type videoController struct {
	service services.VideoService
}

var validate *validator.Validate

func NewVideoController(service services.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("no-bad-word", validators.ValidateTitleIsOk)
	vc := videoController{
		service: service,
	}

	return &vc
}

func (controller *videoController) GetAll() []entity.Video {
	return controller.service.GetAll()
}

func (controller *videoController) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}

	controller.service.Save(video)
	return nil
}

// func must(err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
