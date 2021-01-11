package keg

import (
	"log"
	"net"
	"net/url"
	"strings"
)

type Rule struct {
	FlowDirection Direction
	Addr          url.URL
	CIDR          net.IP
	Action        RuleType
}

func (r *Rule) SetAddress(newAddress string) {
	if !strings.HasPrefix(newAddress, "https://") {
		newAddress = "https://" + newAddress
	}
	newUrl, err := url.Parse(newAddress)
	if err != nil {
		log.Fatal(err)
	}
	r.Addr = *newUrl
}
