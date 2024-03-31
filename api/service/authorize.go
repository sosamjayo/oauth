package service

type AuthorizeService interface {
}

type AuthorizeServiceImpl struct {
}

func NewAuthorizeService() AuthorizeService {
	return &AuthorizeServiceImpl{}
}
