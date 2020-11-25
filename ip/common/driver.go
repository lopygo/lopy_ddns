package common

import "regexp"

type IDriver interface {
	//
	Resolve() (string, error)
}

func IsIpv4(ip string) bool {
	if len(ip) < 7 {
		return false
	}
	match, err := regexp.MatchString(`^(((\d{1,2})|(1\d{2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{2})|(2[0-4]\d)|(25[0-5]))$`, ip)
	if err != nil {
		return false
	}

	return match
}

func IsIpv6(ip string) bool {
	if len(ip) < 3 {
		return false
	}
	match, err := regexp.MatchString(`^([a-f0-9]{1,4}(:[a-f0-9]{1,4}){7}|[a-f0-9]{1,4}(:[a-f0-9]{1,4}){0,7}::[a-f0-9]{0,4}(:[a-f0-9]{1,4}){0,7})$`, ip)
	if err != nil {
		return false
	}

	return match
}
