package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Ifeoluwa"},
			[]string{"Ifeoluwa"},
		},
		{
			"struct with two string fields",
			struct {
				Name  string
				State string
			}{"Ifeoluwa", "Lagos"},
			[]string{"Ifeoluwa", "Lagos"},
		},
		{
			"struct with a non string field",
			struct {
				Name string
				Age  int
			}{"Ifeoluwa", 20},
			[]string{"Ifeoluwa"},
		},
		{
			"Nested fields",
			Person{
				"Ifeoluwa",
				Profile{20, "Lagos"},
			},
			[]string{"Ifeoluwa", "Lagos"},
		},
		{
			"Pointers to things",
			&Person{
				"Ifeoluwa",
				Profile{20, "Lagos"},
			},
			[]string{"Ifeoluwa", "Lagos"},
		},
		{
			"slices",
			[]Profile{
				{20, "Ifeoluwa"},
				{19, "Ifeanyi"},
			},
			[]string{"Ifeoluwa", "Ifeanyi"},
		},
		{
			"Arrays",
			[2]Profile{
				{20, "Lagos"},
				{19, "Abia"},
			},
			[]string{"Lagos", "Abia"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("expected %v but got %v", test.ExpectedCalls, got)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
		// assertContains(t, got, "Box")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{20, "Lagos"}
			aChannel <- Profile{19, "Abia"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Lagos", "Abia"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{20, "Lagos"}, Profile{19, "Abia"}
		}

		var got []string
		want := []string{"Lagos", "Abia"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
