package server

import (
	"net/http"
	"sync"
)

// contexte de requete
type Param struct { Key, Value string }

// structure context
type Context struct {
	W http.ResponseWriter
	R *http.Request
	params [8]Param
	nparams int
}

// fonctions du contexte SolidE
// attribution de valeur d'initialisation au contexte
func (c *Context) update(w http.ResponseWriter, r *http.Request) {
	c.W = w
	c.R = r
	c.nparams = 0
}

// liberation de l'espace memoire
func (c *Context) release() {
	c.W = nil
	c.R = nil
	c.nparams = 0
}

// creation du pool de memoire contexte
var contextPool = sync.Pool{
	New: func() any { return new(Context) },
}
