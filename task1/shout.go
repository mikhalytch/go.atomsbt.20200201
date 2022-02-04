package main

import "strings"

type Animal interface {
	Sound() string
}

func Shout(as ...Animal) string {
	b := strings.Builder{}
	for i := 0; i < len(as); i++ {
		a := as[i]
		if i != 0 {
			b.WriteByte('\n')
		}
		b.WriteString(a.Sound())
	}
	return b.String()
}
