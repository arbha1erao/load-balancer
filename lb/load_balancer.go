package lb

import (
	"sync"
)

type Server struct {
	URL         string
	Weight      int
	Active      *bool
	Connections *int
}

type LoadBalancer struct {
	Servers []*Server
	mutex   sync.Mutex
	rrIndex int
}
