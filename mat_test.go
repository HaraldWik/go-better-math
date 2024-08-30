package bm

import (
	"testing"
)

// TestMat2Rotation tests the rotation functionality for Mat2.
func TestMat2Rotation(t *testing.T) {
	angle := Pi / 4 // 45 degrees
	expected := Mat2[float64]{
		{0.7071067811865476, -0.7071067811865475},
		{0.7071067811865475, 0.7071067811865476},
	}
	m := IdentityMat2[float64]()
	rotated := m.Rotate(angle)

	if rotated != expected {
		t.Errorf("Mat2 Rotate() = %v, want %v", rotated, expected)
	}
}

// TestMat2Identity tests the creation of the identity matrix for Mat2.
func TestMat2Identity(t *testing.T) {
	expected := Mat2[float64]{
		{1, 0},
		{0, 1},
	}
	identity := IdentityMat2[float64]()

	if identity != expected {
		t.Errorf("Mat2 IdentityMat2() = %v, want %v", identity, expected)
	}
}

// TestMat3RotationX tests the X-axis rotation functionality for Mat3.
func TestMat3RotationX(t *testing.T) {
	angle := Pi / 4 // 45 degrees
	expected := Mat3[float64]{
		{1, 0, 0},
		{0, 0.7071067811865476, -0.7071067811865475},
		{0, 0.7071067811865475, 0.7071067811865476},
	}
	m := IdentityMat3[float64]()
	rotated := m.RotateX(angle)

	if rotated != expected {
		t.Errorf("Mat3 RotateX() = %v, want %v", rotated, expected)
	}
}

// TestMat3RotationY tests the Y-axis rotation functionality for Mat3.
func TestMat3RotationY(t *testing.T) {
	angle := Pi / 4 // 45 degrees
	expected := Mat3[float64]{
		{0.7071067811865476, 0, 0.7071067811865475},
		{0, 1, 0},
		{-0.7071067811865475, 0, 0.7071067811865476},
	}
	m := IdentityMat3[float64]()
	rotated := m.RotateY(angle)

	if rotated != expected {
		t.Errorf("Mat3 RotateY() = %v, want %v", rotated, expected)
	}
}

// TestMat3RotationZ tests the Z-axis rotation functionality for Mat3.
func TestMat3RotationZ(t *testing.T) {
	angle := Pi / 4 // 45 degrees
	expected := Mat3[float64]{
		{0.7071067811865476, -0.7071067811865475, 0},
		{0.7071067811865475, 0.7071067811865476, 0},
		{0, 0, 1},
	}
	m := IdentityMat3[float64]()
	rotated := m.RotateZ(angle)

	if rotated != expected {
		t.Errorf("Mat3 RotateZ() = %v, want %v", rotated, expected)
	}
}

// TestMat4RotationX tests the X-axis rotation functionality for Mat4.
func TestMat4RotationX(t *testing.T) {
	angle := Pi / 4 // 45 degrees
	expected := Mat4[float64]{
		{1, 0, 0, 0},
		{0, 0.7071067811865476, -0.7071067811865475, 0},
		{0, 0.7071067811865475, 0.7071067811865476, 0},
		{0, 0, 0, 1},
	}
	m := IdentityMat4[float64]()
	rotated := m.RotateX(angle)

	if rotated != expected {
		t.Errorf("Mat4 RotateX() = %v, want %v", rotated, expected)
	}
}

// TestMat4RotationY tests the Y-axis rotation functionality for Mat4.
func TestMat4RotationY(t *testing.T) {
	angle := Pi / 4 // 45 degrees
	expected := Mat4[float64]{
		{0.7071067811865476, 0, 0.7071067811865475, 0},
		{0, 1, 0, 0},
		{-0.7071067811865475, 0, 0.7071067811865476, 0},
		{0, 0, 0, 1},
	}
	m := IdentityMat4[float64]()
	rotated := m.RotateY(angle)

	if rotated != expected {
		t.Errorf("Mat4 RotateY() = %v, want %v", rotated, expected)
	}
}

// TestMat4RotationZ tests the Z-axis rotation functionality for Mat4.
func TestMat4RotationZ(t *testing.T) {
	angle := Pi / 4 // 45 degrees
	expected := Mat4[float64]{
		{0.7071067811865476, -0.7071067811865475, 0, 0},
		{0.7071067811865475, 0.7071067811865476, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	m := IdentityMat4[float64]()
	rotated := m.RotateZ(angle)

	if rotated != expected {
		t.Errorf("Mat4 RotateZ() = %v, want %v", rotated, expected)
	}
}
