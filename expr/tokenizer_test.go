package expr

import (
	"reflect"
	"testing"
)

func TestNewTokenizer(t *testing.T) {
	for _, test := range []struct {
		expression  string
		text        []rune
		shouldError bool
	}{
		{"", []rune{}, true},
		{"1 > 0", []rune{49, 32, 62, 32, 48}, false},
	} {
		tokenizer, err := NewTokenizer(test.expression, nil)
		if test.shouldError && err != nil {
			continue
		} else if test.shouldError && err == nil {
			t.Fatal("expected test to error but it didn't")
		} else if !test.shouldError && err != nil {
			t.Fatal("expected test to succeed but it didn't")
		}
		if !reflect.DeepEqual(test.text, tokenizer.text) {
			t.Fatalf("expected %v but got %v for text", test.text, tokenizer.text)
		}
		if len(test.text) != tokenizer.length {
			t.Fatalf("expected %v but got %v for length", test.text, tokenizer.length)
		}
	}
}

func TestNextChar(t *testing.T) {
	for _, test := range []struct {
		expression string
		expected   []rune
	}{
		{"abc", []rune{'a', 'b', 'c', '\000', '\000'}},
		{"1 > 0", []rune{'1', ' ', '>', ' ', '0', '\000'}},
	} {
		tokenizer, err := NewTokenizer(test.expression, nil)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		for _, expectedCh := range test.expected {
			if expectedCh != tokenizer.ch {
				t.Fatalf("expected %v but got %v", expectedCh, tokenizer.ch)
			}
			tokenizer.NextChar()
		}
	}
}

func TestSetPosition(t *testing.T) {
	for _, test := range []struct {
		position int
	}{
		{50},
		{0},
	} {
		tokenizer, err := NewTokenizer("1 > 0", nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		tokenizer.SetPosition(test.position)
		if tokenizer.position != test.position {
			t.Fatalf("expected %v as the position but got %v", test.position, tokenizer.position)
		}
	}
}

func TestHasNext(t *testing.T) {
	for _, test := range []struct {
		expression string
		position   int
		expected   bool
	}{
		{"1 > 0", 50, false},
		{"1", 1, false},
		{"1 > 0", 0, true},
		{"1", 0, true},
	} {
		tokenizer, err := NewTokenizer(test.expression, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		tokenizer.position = test.position
		if actual := tokenizer.HasNext(); actual != test.expected {
			t.Fatalf("expected %v but got %v", test.expected, actual)
		}
	}
}
