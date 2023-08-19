package middleware

type middleware struct {
}

func (m *middleware) PanicRecovery() {

}
func NewMiddleware() *middleware {
	return &middleware{}
}
