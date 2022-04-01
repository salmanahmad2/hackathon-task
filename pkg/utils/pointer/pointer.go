/*
Package pointer facilitates working with pointers in Go.
It is intended to make it easy to create pointers to things.
For example, this:

    a := "a string"
    p := &s

becomes:

    p := pointer.String("a string")
*/
package pointer

import (
	"time"
)

// Bool returns a pointer to a boolean.
func Bool(b bool) *bool {
	return &b
}

// Byte returns a pointer to a byte.
func Byte(b byte) *byte {
	return &b
}

// Complex64 returns a pointer to a complex64.
func Complex64(c complex64) *complex64 {
	return &c
}

// Complex128 returns a pointer to a complex128.
func Complex128(c complex128) *complex128 {
	return &c
}

// Float32 returns a pointer to a float32.
func Float32(f float32) *float32 {
	return &f
}

// Float64 returns a pointer to a float64.
func Float64(f float64) *float64 {
	return &f
}

// Int returns a pointer to an int.
func Int(i int) *int {
	return &i
}

// Int8 returns a pointer to an Int8.
func Int8(i int8) *int8 {
	return &i
}

// Int16 returns a pointer to an int16.
func Int16(i int16) *int16 {
	return &i
}

// Int32 returns a pointer to an int32.
func Int32(i int32) *int32 {
	return &i
}

// Int64 returns a pointer to an int64.
func Int64(i int64) *int64 {
	return &i
}

// Rune returns a pointer to a rune.
func Rune(r rune) *rune {
	return &r
}

// String returns a pointer to a string.
func String(s string) *string {
	return &s
}

// Time returns a pointer to a time.Time.
func Time(t time.Time) *time.Time {
	return &t
}

// UInt returns a pointer to an uint.
func UInt(ui uint) *uint {
	return &ui
}

// UInt8 returns a pointer to an uint8.
func UInt8(ui uint8) *uint8 {
	return &ui
}

// UInt16 returns a pointer to an uint16.
func UInt16(ui uint16) *uint16 {
	return &ui
}

// UInt32 returns a pointer to an uint32.
func UInt32(ui uint32) *uint32 {
	return &ui
}

// UInt64 returns a pointer to an uint64.
func UInt64(ui uint64) *uint64 {
	return &ui
}
