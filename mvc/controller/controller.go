package controller

type Controller struct {
	DB string
}

func NewController() *Controller {
	var c Controller
	return &c
}