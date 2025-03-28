package service

type IUsageService interface {
}

type UsageService struct {
}

func NewUsageService() IUsageService {
	return &UsageService{}
}
