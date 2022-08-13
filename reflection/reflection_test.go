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
			"Struct com um campo string",
			struct {
				Nome string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct com dois campos string",
			struct {
				Nome   string
				Cidade string
			}{"Chris", "Londres"},
			[]string{"Chris", "Londres"},
		},
		{
			"Struct com um campo string e int",
			struct {
				Nome  string
				Idade int
			}{"Chris", 22},
			[]string{"Chris"},
		},
		{
			"Struct com um campo string e uma struct",
			Person{
				"Luiz",
				Perfil{
					Idade:  20,
					Cidade: "Santos",
				},
			},
			[]string{"Luiz", "Santos"},
		},
		{
			"Struct com um pointer",
			&Person{
				"Luiz",
				Perfil{
					Idade:  20,
					Cidade: "Santos",
				},
			},
			[]string{"Luiz", "Santos"},
		},
		{
			"Slices",
			[]Perfil{
				{20, "Santos"},
				{19, "São Vicente"},
			},
			[]string{"Santos", "São Vicente"},
		},
		{
			"Arrays",
			[2]Perfil{
				{20, "Santos"},
				{19, "São Vicente"},
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

}

func through(x interface{}, fn func(input string)) {
	value := getValue(x)

	quantityOfValue := 0
	var getField func(int) reflect.Value

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		quantityOfValue = value.NumField()
		getField = value.Field
	case reflect.Slice, reflect.Array:
		quantityOfValue = value.Len()
		getField = value.Index
	}

	for i := 0; i < quantityOfValue; i++ {
		through(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
