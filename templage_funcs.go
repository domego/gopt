package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

var funcMap = map[string]interface{}{
	"capitalizeFirst": CapitalizeFirst,
	"lowerFirst":      LowerFirst,
	"camel":           Camel,
	"upperCamel":      UpperCamel,
	"underline":       Underline,
	"contains":        contains,
	"match":           match,
	"incr":            incr,
	"upper":           strings.ToUpper,
	"lower":           strings.ToLower,
	"toSortedMap":     ToSortedMap,

	"isArray":        isArray,
	"isBuiltIn":      isBuiltIn,
	"isBuiltInArray": isBuiltInArray,
	"hasPrefix":      strings.HasPrefix,
	"nullableType":   nullableType,
}

func ToSortedMap(m interface{}) []MapItem {
	slice := []MapItem{}
	if m != nil {
		for k, v := range m.(map[interface{}]interface{}) {
			slice = append(slice, MapItem{
				Key:   k.(string),
				Value: v,
			})
			sort.Sort(byKey(slice))
		}
	}
	return slice
}

type MapItem struct {
	Key   string
	Value interface{}
}

type byKey []MapItem

func (a byKey) Len() int           { return len(a) }
func (a byKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func nullableType(s string) string {
	if strings.HasPrefix(s, "[]") {
		return s
	}
	if isBuiltIn(s) {
		return s
	}
	return "*" + s
}

func isBuiltIn(t string) bool {
	if t == "int" || t == "int32" || t == "int64" || t == "string" || t == "bool" || t == "float32" || t == "float64" || t == "uint" || t == "uint32" || t == "uint64" {
		return true
	}
	return false
}

func isBuiltInArray(t string) bool {
	if !strings.HasPrefix(t, "[]") {
		return false
	}
	return isBuiltIn(t[:2])
}

func isArray(t string) bool {
	return strings.HasPrefix(t, "[]")
}

func match(s1, s2 string) bool {
	if s1 == "*" {
		return true
	}
	if strings.Contains(s1, s2) {
		return true
	}
	return false
}

func incr(i int) int {
	return i + 1
}

func contains(o interface{}, s string) bool {
	s1 := fmt.Sprintf("%s", o)
	return strings.Contains(s1, s)
}

func CapitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	first := s[:1]
	return strings.ToUpper(first) + s[1:]
}

func LowerFirst(s string) string {
	if s == "" {
		return s
	}
	first := s[:1]
	return strings.ToLower(first) + s[1:]
}

func Camel(s string) string {
	bs := []byte{}
	captalize := false
	for i, c := range s {
		if i == 0 {
			bs = append(bs, byte(c))
			continue
		}
		if c == '_' {
			captalize = true
		} else {
			if captalize {
				c = unicode.ToUpper(c)
			}
			bs = append(bs, byte(c))
			captalize = false
		}
	}
	return LowerFirst(string(bs))
}

func UpperCamel(s string) string {
	return CapitalizeFirst(Camel(s))
}

func Underline(s string) string {
	bs := []byte{}
	for i, c := range s {
		if i == 0 {
			bs = append(bs, byte(c))
			continue
		}
		if unicode.IsLower(c) || c == '_' {
			bs = append(bs, byte(c))
		} else {
			bs = append(bs, '_', byte(unicode.ToLower(c)))
		}
	}
	return LowerFirst(string(bs))
}
