package dns

import (
	"github.com/miekg/dns"
)

type Server struct {
	server *dns.Server
	mux    *dns.ServeMux
}
