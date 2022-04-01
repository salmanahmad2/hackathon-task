package pointer_test

import (
	"testing"
	"time"

	"hackathon/pkg/utils/pointer"
)

func TestBool_WhenBoolPassedIn_ShouldReturnBoolPointer(t *testing.T) {
	// Arrange
	b := false

	// Act
	result := pointer.Bool(b)

	// Assert
	if *result != b {
		t.Errorf("Expected %v; got %v", b, *result)
	}
}

func TestByte_WhenBytePassedIn_ShouldReturnBytePointer(t *testing.T) {
	// Arrange
	b := []byte{'a'}

	// Act
	result := pointer.Byte(b[0])

	// Assert
	if *result != b[0] {
		t.Errorf("Expected %v; got %v", b, *result)
	}
}

func TestComplex64_WhenComplex64PassedIn_ShouldReturnComplex64Pointer(t *testing.T) {
	// Arrange
	var c complex64 = complex(1, 2)

	// Act
	result := pointer.Complex64(c)

	// Assert
	if *result != c {
		t.Errorf("Expected %v; got %v", c, *result)
	}
}

func TestComplex128_WhenComplex128PassedIn_ShouldReturnComplex128Pointer(t *testing.T) {
	// Arrange
	var c complex128 = complex(1, 2)

	// Act
	result := pointer.Complex128(c)

	// Assert
	if *result != c {
		t.Errorf("Expected %v; got %v", c, *result)
	}
}

func TestFloat32_WhenFloat32PassedIn_ShouldReturnFloat32Pointer(t *testing.T) {
	// Arrange
	var f float32 = 123.45

	// Act
	result := pointer.Float32(f)

	// Assert
	if *result != f {
		t.Errorf("Expected %v; got %v", f, *result)
	}
}

func TestFloat64_WhenFloat64PassedIn_ShouldReturnFloat64Pointer(t *testing.T) {
	// Arrange
	var f float64 = 123.45

	// Act
	result := pointer.Float64(f)

	// Assert
	if *result != f {
		t.Errorf("Expected %v; got %v", f, *result)
	}
}

func TestInt_WhenIntPassedIn_ShouldReturnIntPointer(t *testing.T) {
	// Arrange
	i := 0

	// Act
	result := pointer.Int(i)

	// Assert
	if *result != i {
		t.Errorf("Expected %v; got %v", i, *result)
	}
}

func TestInt8_WhenInt8PassedIn_ShouldReturnInt8Pointer(t *testing.T) {
	// Arrange
	i := int8(0)

	// Act
	result := pointer.Int8(i)

	// Assert
	if *result != i {
		t.Errorf("Expected %v; got %v", i, *result)
	}
}

func TestInt16_WhenInt16PassedIn_ShouldReturnInt16Pointer(t *testing.T) {
	// Arrange
	i := int16(0)

	// Act
	result := pointer.Int16(i)

	// Assert
	if *result != i {
		t.Errorf("Expected %v; got %v", i, *result)
	}
}

func TestInt32_WhenInt32PassedIn_ShouldReturnInt32Pointer(t *testing.T) {
	// Arrange
	i := int32(0)

	// Act
	result := pointer.Int32(i)

	// Assert
	if *result != i {
		t.Errorf("Expected %v; got %v", i, *result)
	}
}

func TestInt64_WhenInt64PassedIn_ShouldReturnInt64Pointer(t *testing.T) {
	// Arrange
	i := int64(0)

	// Act
	result := pointer.Int64(i)

	// Assert
	if *result != i {
		t.Errorf("Expected %v; got %v", i, *result)
	}
}

func TestRune_WhenRunePassedIn_ShouldReturnRunePointer(t *testing.T) {
	// Arrange
	var r rune = 1

	// Act
	result := pointer.Rune(r)

	// Assert
	if *result != r {
		t.Errorf("Expected %v; got %v", r, *result)
	}
}

func TestString_WhenStringPassedIn_ShouldReturnStringPointer(t *testing.T) {
	// Arrange
	str := "String"

	// Act
	result := pointer.String(str)

	// Assert
	if *result != str {
		t.Errorf("Expected %v; got %v", str, *result)
	}
}

func TestTime_WhenTimePassedIn_ShouldReturnTimePointer(t *testing.T) {
	// Arrange
	tm := time.Date(2006, 1, 2, 3, 4, 5, 7, time.UTC)

	// Act
	result := pointer.Time(tm)

	// Assert
	if *result != tm {
		t.Errorf("Expected %v; got %v", tm, *result)
	}
}

func TestUInt_WhenUIntPassedIn_ShouldReturnUIntPointer(t *testing.T) {
	// Arrange
	ui := uint(0)

	// Act
	result := pointer.UInt(ui)

	// Assert
	if *result != ui {
		t.Errorf("Expected %v; got %v", ui, *result)
	}
}

func TestUInt8_WhenUInt8PassedIn_ShouldReturnUInt8Pointer(t *testing.T) {
	// Arrange
	ui := uint8(0)

	// Act
	result := pointer.UInt8(ui)

	// Assert
	if *result != ui {
		t.Errorf("Expected %v; got %v", ui, *result)
	}
}

func TestUInt16_WhenUInt16PassedIn_ShouldReturnUInt16Pointer(t *testing.T) {
	// Arrange
	ui := uint16(0)

	// Act
	result := pointer.UInt16(ui)

	// Assert
	if *result != ui {
		t.Errorf("Expected %v; got %v", ui, *result)
	}
}

func TestUInt32_WhenUInt32PassedIn_ShouldReturnUInt32Pointer(t *testing.T) {
	// Arrange
	ui := uint32(0)

	// Act
	result := pointer.UInt32(ui)

	// Assert
	if *result != ui {
		t.Errorf("Expected %v; got %v", ui, *result)
	}
}

func TestUInt64_WhenUInt64PassedIn_ShouldReturnUInt64Pointer(t *testing.T) {
	// Arrange
	ui := uint64(0)

	// Act
	result := pointer.UInt64(ui)

	// Assert
	if *result != ui {
		t.Errorf("Expected %v; got %v", ui, *result)
	}
}
