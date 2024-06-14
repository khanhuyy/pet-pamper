package handler

type IFulfillmentService interface {
	IAuthHandler
}
type Handler struct {
	authService IAuthHandler
}

func NewHandler() *Handler {
	authService := NewAuthHandler()
	return &Handler{
		authService: authService,
	}
}
