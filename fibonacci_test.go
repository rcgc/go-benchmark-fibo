package main

import "testing"

func TestFibonacciFor(t *testing.T){
	target := 55
	got := fibonacciFor(10)
	if got != target {
		t.Errorf("Se esperaba %d se obtuvo %d", target, got)
	}
}

func TestFibonacciRec(t *testing.T) {
	target := 55
	got := fibonacciRec(10)
	if got != target {
		t.Errorf("Se esperaba %d se obtuvo %d", target, got)
	}
}

func TestFibonacciRecMem(t *testing.T) {
	target := 55
	got := fibonacciRecMem(10)
	if got != target {
		t.Errorf("Se esperaba %d se obtuvo %d", target, got)
	}
}

func BenchmarkFibonacciFor(b *testing.B) {
	for n:=0; n < b.N; n++ {
		fibonacciFor(30)
	}
}

func BenchmarkFibonacciRec(b *testing.B) {
	for n:=0; n < b.N; n++ {
		fibonacciRec(30)
	}
}

func BenchmarkFibonacciRecMem(b *testing.B) {
	for n:=0; n < b.N; n++ {
		fibonacciRecMem(30)
	}
}