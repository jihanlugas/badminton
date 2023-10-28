package utils

import "regexp"

var regFormatHp *regexp.Regexp

func init() {
	regFormatHp = regexp.MustCompile(`(^\+?628)|(^0?8){1}`)
}

func FormatPhoneTo62(phone string) string {
	formatPhone := regFormatHp.ReplaceAllString(phone, "628")
	return formatPhone
}
