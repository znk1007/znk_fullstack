package pool

type LoginJob interface {
	login() error
}

type AuthJob interface {
	login
}
