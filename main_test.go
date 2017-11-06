package main

import "testing"

func TestCorrect(t *testing.T) {
	v := 42
	buf := new(Buffer)
	buf.Send(v)

	var x, y int
	buf.Recv(&x)
	buf.RecvReflect(&y)
	var zhint int
	z := buf.RecvReturn(&zhint)
	if want, got := v, x; want != got {
		t.Errorf("Recv error: want %d, got %d", want, got)
	}
	if want, got := v, y; want != got {
		t.Errorf("RecvReflect error: want %d, got %d", want, got)
	}
	if want, got := v, z; want != got {
		t.Errorf("RecvReturn error: want %d, got %d", want, got)
	}
}

func BenchmarkRecv(b *testing.B) {
	v := 42
	buf := new(Buffer)
	for i := 0; i < b.N; i++ {
		buf.Send(v)
		var x int
		buf.Recv(&x)
	}
}

func BenchmarkRecvReturn(b *testing.B) {
	v := 42
	buf := new(Buffer)
	for i := 0; i < b.N; i++ {
		buf.Send(v)
		x := buf.RecvReturn(nil)
		_ = x
	}
}

func BenchmarkRecvReflect(b *testing.B) {
	v := 42
	buf := new(Buffer)
	for i := 0; i < b.N; i++ {
		buf.Send(v)
		var x int
		buf.RecvReflect(&x)
	}
}
