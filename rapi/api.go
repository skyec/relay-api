package rapi

type TransformCb func() (err error)

type Api struct {
	endpoint string
}

func New(endpoint string) *Api {
	return &Api{
		endpoint: "/api" + endpoint,
	}
}

func (a *Api) NewEndpoint(method, endpoint string) *Api {
}

func (a *Api) InternalPath(endpoint string) *Api {
}

func (a *Api) TransformRequest(internal, external interface{}) *Api {
}

func (a *Api) TransformResponse(internal, external interface{}) *Api {
}

func (a *Api) TransformRequestCb(callback TransformCb) *Api {
	err := callback()
	return
}

func (a *Api) TransformResponseCb(callback TransformCb) *Api {
	err := callback()
	return
}