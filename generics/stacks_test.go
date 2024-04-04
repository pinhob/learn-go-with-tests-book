package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("test stack of numbers", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(1)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 1)

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(2)
		myStackOfInts.Push(3)
		firstValue, _ := myStackOfInts.Pop()
		secondValue, _ := myStackOfInts.Pop()
		AssertEqual(t, firstValue+secondValue, 5)
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
