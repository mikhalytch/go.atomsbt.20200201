package main

import (
	"fmt"
	"testing"
)

func TestSound(t *testing.T) {
	tests := []struct {
		inAnimals []Animal
		want      string
	}{
		{[]Animal{Dog{}, Cat{}}, `"Гав"
"Мяу"`},
		{nil, ``},
		{[]Animal{}, ``},
		{[]Animal{Cat{}, Cat{}}, `"Мяу"
"Мяу"`},
		{[]Animal{*Create(DogType)}, `"Гав"`},
		{[]Animal{*Create(CatType), *Create(CatType)}, `"Мяу"
"Мяу"`},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got := Shout(test.inAnimals...)
			want := test.want
			if got != want {
				t.Fatalf("Got %v animal sounds, want %v", got, want)
			}
		})
	}
}
