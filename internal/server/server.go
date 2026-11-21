package server

// importation de module
import (
	"net/http"
)

// RFC 1
type HandleFunction func(c *Context) error

func New(h HandleFunction) *Server {
	return &Server{handler: h}
}

type Server struct {
	handler HandleFunction
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := contextPool.Get().(*Context)
	c.update(w, r)
	s.dispatch(c)
	c.release()
	contextPool.Put(c)
}

func (s *Server) dispatch(c *Context) {
	if err := s.handler(c); err != nil {
		http.Error(c.W, "Erreur interne", http.StatusInternalServerError)
	}
}