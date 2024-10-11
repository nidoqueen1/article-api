package controller

import (
	"github.com/nidoqueen1/article-api/service"
	"github.com/sirupsen/logrus"
)

// Implements handler functions
type handler struct {
	service service.IService
	logger  *logrus.Logger
}

// Initializes handler layer with its dependencies
func Init(s service.IService, l *logrus.Logger) *handler {
	return &handler{
		service: s,
		logger:  l,
	}
}
