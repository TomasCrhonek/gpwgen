package main

import (
	"strings"
	"testing"
)

func BenchmarkPwgenNoSpecial(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		pwgen(12, false)
	}
}

func BenchmarkPwgenSpecial(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		pwgen(12, true)
	}
}

func TestPwgenSmallPassNoSpecial(t *testing.T) {
	for i := 0; i < 20; i++ {
		p := pwgen(i, false)
		if len(p) < 12 {
			t.Fail()
		}
	}
}

func TestPwgenSmallPassSpecial(t *testing.T) {
	for i := 0; i < 20; i++ {
		p := pwgen(i, true)
		if len(p) < 12 {
			t.Fail()
		}
	}
}

func TestPwgenSpecial(t *testing.T) {
	p := pwgen(12, true)
	for i := 0; i < len(spchr); i++ {
		if strings.Contains(p, spchr[i]) {
			return
		}
	}
	t.Fail()
}
