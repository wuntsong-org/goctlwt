package svc

type ServiceContext struct {
	{{.middleware}}
}

func NewServiceContext() *ServiceContext {
	return &ServiceContext{
		{{.middlewareAssignment}}
	}
}
