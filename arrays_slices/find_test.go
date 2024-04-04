package arraysslices

import (
	"strings"
	"testing"
)

type Person struct {
	Name string
}

func TestFind(t *testing.T) {
	t.Run("find number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firsEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		AssertTrue(t, found)
		AssertEqual(t, firsEvenNumber, 2)
	})

	t.Run("Find person", func(t *testing.T) {
		people := []Person{
			Person{Name: "John Doe"},
			Person{Name: "João Silva"},
			Person{Name: "Bruno"},
		}

		person, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Bruno")
		})
		AssertTrue(t, found)
		AssertEqual(t, person, Person{Name: "Bruno"})
	})
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %+v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %+v, want false", got)
	}
}
