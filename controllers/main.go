package controllers

import "paytabs-task/services"

type Controllers struct {
	services *services.Services
}

func NewController(services *services.Services) *Controllers {
	return &Controllers{
		services: services,
	}
}
