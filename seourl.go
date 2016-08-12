//copy from https://www.socketloop.com/tutorials/golang-format-strings-to-seo-friendly-url-example
package gografana

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strings"
	"unicode"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func SEOURL(s string) string {
	seoStr := strings.ToLower(s)

	// convert all spaces to dash
	regE := regexp.MustCompile("[[:space:]]")
	seoStrByte := regE.ReplaceAll([]byte(seoStr), []byte("-"))
	seoStr = string(seoStrByte) // convert []byte to string

	// remove all blanks such as tab
	regE = regexp.MustCompile("[[:blank:]]")
	seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte(""))
	seoStr = string(seoStrByte) // convert []byte to string

	regE = regexp.MustCompile("[!:-@[-`{-~]")
	seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte(""))
	seoStr = string(seoStrByte) // convert []byte to string

	regE = regexp.MustCompile("/[^\x20-\x7F]/")
	seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte(""))
	seoStr = string(seoStrByte) // convert []byte to string

	regE = regexp.MustCompile("`&(amp;)?#?[a-z0-9]+;`i")
	seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte("-"))
	seoStr = string(seoStrByte) // convert []byte to string

	regE = regexp.MustCompile("`&([a-z])(acute|uml|circ|grave|ring|cedil|slash|tilde|caron|lig|quot|rsquo);`i")
	seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte("\\1"))
	seoStr = string(seoStrByte) // convert []byte to string

	regE = regexp.MustCompile("`[^a-z0-9]`i")
	seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte("-"))
	seoStr = string(seoStrByte) // convert []byte to string

	regE = regexp.MustCompile("`[-]+`")
	seoStrByte = regE.ReplaceAll([]byte(seoStr), []byte("-"))
	seoStr = string(seoStrByte) // convert []byte to string

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	seoStr, _, _ = transform.String(t, seoStr)

	return strings.TrimSpace(seoStr)
}
