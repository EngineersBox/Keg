package keg

import "net"

type Gateway struct {
	RuleSet  []Rule
	ToCIDR   net.IP
	FromCIDR net.IP
}
