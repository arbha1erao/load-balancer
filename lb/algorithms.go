package lb

import "golang.org/x/exp/rand"

// TODO - nil pointer check

// Round Robin Algorithm
func (lb *LoadBalancer) RoundRobin() *Server {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	for {
		server := lb.Servers[lb.rrIndex]
		lb.rrIndex = (lb.rrIndex + 1) % len(lb.Servers)
		if *server.Active {
			return server
		}
	}
}

// Weighted Round Robin Algorithm
func (lb *LoadBalancer) WeightedRoundRobin() *Server {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	totalWeight := 0
	for _, server := range lb.Servers {
		if *server.Active {
			totalWeight += server.Weight
		}
	}

	randWeight := rand.Intn(totalWeight)
	for _, server := range lb.Servers {
		if *server.Active {
			randWeight -= server.Weight
			if randWeight <= 0 {
				return server
			}
		}
	}
	return nil
}

// Least Connections Algorithm
func (lb *LoadBalancer) LeastConnections() *Server {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	var leastConnServer *Server
	minConnections := int(^uint(0) >> 1)

	for _, server := range lb.Servers {
		if *server.Active && *server.Connections < minConnections {
			minConnections = *server.Connections
			leastConnServer = server
		}
	}
	return leastConnServer
}
