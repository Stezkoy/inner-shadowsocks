package main

import (
	"errors"
	"net"
	"strings"
)

type Whitelist struct {
	enable     bool
	domainlist []string
	iplist     []net.IPNet
	logger     func(string, ...interface{})
}

func (w *Whitelist) check_ip(ip net.IP) bool {
	if !w.enable {
		return true
	}
	for _, ipnet := range w.iplist {
		if ipnet.Contains(ip) {
			return true
		}
	}
	return false
}

func (w *Whitelist) check_domain(d string) bool {
	if !w.enable {
		return true
	}
	for _, domain := range w.domainlist {
		if strings.HasSuffix(d, domain) {
			return true
		}
	}
	return false
}

func (w *Whitelist) check(data []byte) error {
	switch data[0] {
	case 0x01:
		if w.check_ip(net.IP(data[1 : 1+net.IPv4len])) {
			w.logger("[whitelist] Белый список ipv4 принят.")
			return nil
		}
		w.logger("[whitelist] Белый список ipv4 отклонен.")
		return errors.New("IPv4 не в белом списке.")
	case 0x03:
		if w.check_domain(string(data[2 : 2+data[1]])) {
			w.logger("[whitelist] Белый список доменов принят.")
			return nil
		}
		w.logger("[whitelist] Отклонение домена из белого списка.")
		return errors.New("Домен не в белом листе.")
	case 0x04:
		if w.check_ip(net.IP(data[1 : 1+net.IPv6len])) {
			w.logger("[whitelist] Белый список ipv6 принят.")
			return nil
		}
		w.logger("[whitelist] Белый список ipv6 отклонен.")
		return errors.New("IPv6 не в белом списке.")
	}
	return errors.New("Неизвестная ошибка.")
}
