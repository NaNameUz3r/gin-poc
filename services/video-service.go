package services

import "gin-poc/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	GetAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func NewVideoService() VideoService {
	return &videoService{}
}

func (svc *videoService) Save(v entity.Video) entity.Video {
	svc.videos = append(svc.videos, v)
	return v
}

func (svc *videoService) GetAll() []entity.Video {
	return svc.videos
}
