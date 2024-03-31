package service

type CallbackService interface {
}

type CallbackServiceImpl struct {
}

func NewCallbackService() CallbackService {
	return &CallbackServiceImpl{}
}
