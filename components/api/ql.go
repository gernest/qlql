package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cznic/ql"
	"github.com/gernest/alien"
)

type Result struct {
	Error error       `json:"error"`
	Data  interface{} `json:"data"`
}

type Server struct {
	dbs []*ql.DB
}

func (s *Server) Open(w http.ResponseWriter, r *http.Request) {
}
func (s *Server) Info(w http.ResponseWriter, r *http.Request) {
	p := alien.GetParams(r)
	dbName := p.Get("dbname")
	if dbName == "" {
		rst := &Result{Error: errors.New("missing database name")}
		d, _ := json.Marshal(rst)
		w.Write(d)
		return
	}
	for _, v := range s.dbs {
		if v.Name() == dbName {
			info, err := v.Info()
			if err != nil {
				renderErr(w, err)
				return
			}
			renderJSON(w, info)
		}
	}
}

func renderErr(w http.ResponseWriter, err error) {
	rst := &Result{Error: err}
	d, _ := json.Marshal(rst)
	w.Write(d)
}

func renderJSON(w http.ResponseWriter, data interface{}) {
	d, err := json.Marshal(data)
	if err != nil {
		renderErr(w, err)
		return
	}
	w.Write(d)
}

func NewServer() *alien.Mux {
	s := &Server{}
	r := alien.New()
	r.Get("/info/:dbname", s.Info)
	r.Get("/open/:dbname", s.Open)
	return r
}
