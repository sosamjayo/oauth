package service

type TokenService interface {
}

type TokenServiceImpl struct {
}

func NewTokenService() TokenService {
	return &TokenServiceImpl{}
}
