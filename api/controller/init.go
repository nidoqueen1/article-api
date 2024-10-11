package controller

import (
	"github.com/nidoqueen1/article-api/service"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service service.IService
	logger  *logrus.Logger
}

func Init(s service.IService, l *logrus.Logger) *handler {
	return &handler{
		service: s,
		logger:  l,
	}
}
