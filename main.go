package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	v := 42
	b := new(Buffer)
	b.Send(v)
	var x, y int
	b.Recv(&x)
	b.RecvReflect(&y)
	z := b.RecvReturn(nil)

	fmt.Println("Receiver v:", v, "x:", x, "y:", y, "z:", z)
}

type Buffer struct {
	ptr unsafe.Pointer
}

// Send is a generic sending function that sends val through the Buffer b.
func (b *Buffer) Send(val interface{}) {
	b.ptr = unsafe.Pointer(&val)
}

// Recv is a generic receiving function that receives from Buffer b and writes
// to the memory location pointed to by ptr. Note that ptr is a pointer to an
// internal location and is volatile.
func (b *Buffer) Recv(ptr interface{}) error {
	var word uint

	// ptrConcrete points to the concrete value of the ptr interface at runtime.
	// interface{} is stored as 2 words, and we only look at the 2nd word.
	ptrConcrete := unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))

	// bPtrConcrete points to the real value of the ptr field of the buffer.
	// interface{} is stored as 2 words, and we only look at the 2nd word.
	bPtrConcrete := unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word))

	// ptrAddr is the memory location storing *ptr.
	// This is the memory location to write into.
	ptrAddr := *(**unsafe.Pointer)(ptrConcrete)

	*ptrAddr = **(**unsafe.Pointer)(bPtrConcrete)
	return nil
}

// RecvReflect is a receive function that uses reflection to inspect the runtime
// type of ptr and sets its value appropriately. The use of reflect means this
// is slow.
func (b *Buffer) RecvReflect(ptr interface{}) error {
	v := *(*interface{})(b.ptr)

	ptrValue := reflect.ValueOf(ptr)
	val := reflect.Indirect(ptrValue)
	val.Set(reflect.ValueOf(v))
	return nil
}

// RecvReturn is a receive function that discards the parameter (only use as
// type hint for deserialisation) and return the actual value (not pointer).
// The advantage is that it is simple but does not give safer guarantee than
// Recv (caller needs to make a copy).
func (b *Buffer) RecvReturn(ptr interface{}) interface{} {
	v := *(*interface{})(b.ptr)
	return v
}
