package adminctrl

type umsMenuController struct {
}

type UmsMenuController interface {
}

func NewUmsMenuController() UmsMenuController {
	return &umsMenuController{}
}
