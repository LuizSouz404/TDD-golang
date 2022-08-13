package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Nome   string
	Perfil Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}

func TestReflection(t *testing.T) {
	cases := []struct {
		Nome         string
		Input        interface{}
		CallsWaiting []string
	}{
		{
			"Struct com um campo tipo string",
			struct {
				Nome string
			}{"Luiz"},
			[]string{"Luiz"},
		},
		{
			"Struct com dois campos tipo string",
			struct {
				Nome   string
				Cidade string
			}{"Luiz", "Santos"},
			[]string{"Luiz", "Santos"},
		},
		{
			"Campos aninhados",
			Person{
				"Luiz",
				Perfil{20, "Santos"},
			},
			[]string{"Luiz", "Santos"},
		},
		{
			"Ponteiros para coisas",
			&Person{
				"Luiz",
				Perfil{20, "Santos"},
			},
			[]string{"Luiz", "Santos"},
		},
		{
			"Slices",
			[]Perfil{
				{20, "Santos"},
				{34, "São Vicente"},
			},
			[]string{"Santos", "São Vicente"},
		},
		{
			"Arrays",
			[2]Perfil{
				{20, "Santos"},
				{34, "São Vicente"},
			},
			[]string{"Santos", "São Vicente"},
		},
	}

	for _, test := range cases {
		t.Run(test.Nome, func(t *testing.T) {
			var result []string

			through(test.Input, func(input string) {
				result = append(result, input)
			})

			if !reflect.DeepEqual(result, test.CallsWaiting) {
				t.Errorf("Reflection\ngot: %v\nexpect: %v", result, test.CallsWaiting)
			}
		})
	}

	t.Run("With Maps", func(t *testing.T) {
		mapA := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var result []string
		through(mapA, func(input string) {
			result = append(result, input)
		})

		checkIfHave(t, result, "Bar")
		checkIfHave(t, result, "Boz")
	})

}

func checkIfHave(t *testing.T, array []string, value string) {
	have := false
	for _, x := range array {
		if x == value {
			have = true
		}
	}

	if !have {
		t.Errorf("Expect have %s in %+v, but dont have", value, array)
	}
}
