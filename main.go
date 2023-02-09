package main

import "fmt"

type RoundRobinBalancer struct {
	servers  int
	clients  int
	requests int
	s        []int
	stat     int
}

func newRoundRobinBalancer(servers, clients, requests int) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		servers:  servers,
		clients:  clients,
		requests: requests,
		s:        make([]int, servers),
	}
}

func (r *RoundRobinBalancer) Init(numbersOfServer int) {
	r.servers = numbersOfServer
}

func (r *RoundRobinBalancer) GiveStat() []int {
	if r.stat == 0 {
		for i := range r.s {
			r.s[i] = 0
		}

	}
	return r.s
}

func (r *RoundRobinBalancer) GiveNode() []int {
	r.stat = r.clients * r.requests
	if r.stat < r.servers && r.stat >= 1 {
		r.s[0] = 1
	}
	if r.stat == r.servers {
		r.s[0] = 1
		r.s[1] = 1
		r.s[2] = 1
	}
	if r.stat > 1 && r.stat < r.servers {
		r.s[0] = 1
		r.s[1] = 1
		r.s[2] = 0
	}

	return r.s
}

func main() {
	r := newRoundRobinBalancer(3, 1, 2)
	r.Init(3)
	r.GiveStat()
	r.GiveNode()
	fmt.Println(r.stat, r.s)
}
