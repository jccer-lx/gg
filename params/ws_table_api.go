package params

type WsTableAddParams struct {
	Title string `json:"title" validate:"required,min=1,max=20"`
}

func (p *WsTableAddParams) NewParams() GGParams {
	return new(WsTableAddParams)
}
