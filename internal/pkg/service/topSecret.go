package service

import "starshipCommsResolver/internal/pkg/ports"

type service struct {
	ports ports.CommunicationServices
}

func NewService(port ports.CommunicationServices) *service {
	return &service{
		ports: port,
	}
}

func GetLocation(...float32) (x, y float32) {

	return 1, 2
}

func GetMessage(...[]string) string {
	return ""
}
