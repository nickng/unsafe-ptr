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
// to the memory location pointed to by ptr.
func (b *Buffer) Recv(ptr interface{}) error {
	var word uint

	switch ptr.(type) {
	case *bool:
		**(**bool)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**bool)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *float32:
		**(**float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**float32)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *float64:
		**(**float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**float64)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *int:
		**(**int)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**int)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *int8:
		**(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**int8)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *int16:
		**(**int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**int16)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *int32:
		**(**int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**int32)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *int64:
		**(**int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**int64)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *uint:
		**(**uint)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**uint)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *uint8:
		**(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**uint8)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *uint16:
		**(**uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**uint16)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *uint32:
		**(**uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**uint32)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *uint64:
		**(**uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**uint64)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *uintptr:
		**(**uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**uintptr)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	case *string:
		**(**string)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**string)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	default:
		**(**unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptr)) + unsafe.Sizeof(word))) =
			**(**unsafe.Pointer)(unsafe.Pointer(uintptr(b.ptr) + unsafe.Sizeof(word)))
	}
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
