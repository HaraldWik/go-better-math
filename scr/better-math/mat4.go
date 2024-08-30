package bm

import (
	"fmt"
)

// Mat4 represents a 4x4 matrix with elements of type T.
type Mat4[T Numeric] [4][4]T

/**
 * IdentityMat4 returns a 4x4 identity matrix.
 * For example:
 *   IdentityMat4[float64]() returns Mat4[float64]{ {1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1} }
 */
func IdentityMat4[T Numeric]() Mat4[T] {
	return Mat4[T]{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

/**
 * Add adds another 4x4 matrix to the current matrix and returns the result.
 * For example:
 *   Mat4[int]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }.Add(Mat4[int]{ {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1} })
 *   returns Mat4[int]{ {2, 3, 4, 5}, {6, 7, 8, 9}, {10, 11, 12, 13}, {14, 15, 16, 17} }
 */
func (m Mat4[T]) Add(n Mat4[T]) Mat4[T] {
	return Mat4[T]{
		{m[0][0] + n[0][0], m[0][1] + n[0][1], m[0][2] + n[0][2], m[0][3] + n[0][3]},
		{m[1][0] + n[1][0], m[1][1] + n[1][1], m[1][2] + n[1][2], m[1][3] + n[1][3]},
		{m[2][0] + n[2][0], m[2][1] + n[2][1], m[2][2] + n[2][2], m[2][3] + n[2][3]},
		{m[3][0] + n[3][0], m[3][1] + n[3][1], m[3][2] + n[3][2], m[3][3] + n[3][3]},
	}
}

/**
 * Sub subtracts another 4x4 matrix from the current matrix and returns the result.
 * For example:
 *   Mat4[int]{ {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20} }.Sub(Mat4[int]{ {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1} })
 *   returns Mat4[int]{ {4, 5, 6, 7}, {8, 9, 10, 11}, {12, 13, 14, 15}, {16, 17, 18, 19} }
 */
func (m Mat4[T]) Sub(n Mat4[T]) Mat4[T] {
	return Mat4[T]{
		{m[0][0] - n[0][0], m[0][1] - n[0][1], m[0][2] - n[0][2], m[0][3] - n[0][3]},
		{m[1][0] - n[1][0], m[1][1] - n[1][1], m[1][2] - n[1][2], m[1][3] - n[1][3]},
		{m[2][0] - n[2][0], m[2][1] - n[2][1], m[2][2] - n[2][2], m[2][3] - n[2][3]},
		{m[3][0] - n[3][0], m[3][1] - n[3][1], m[3][2] - n[3][2], m[3][3] - n[3][3]},
	}
}

/**
 * Mul multiplies the current matrix by another 4x4 matrix and returns the result.
 * For example:
 *   Mat4[int]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }.Mul(Mat4[int]{ {1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1} })
 *   returns Mat4[int]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }
 */
func (m Mat4[T]) Mul(n Mat4[T]) Mat4[T] {
	var result Mat4[T]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			var sum T
			for k := 0; k < 4; k++ {
				sum += m[i][k] * n[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

/**
 * ScalarMul multiplies the matrix by a scalar and returns the result.
 * For example:
 *   Mat4[int]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }.ScalarMul(2)
 *   returns Mat4[int]{ {2, 4, 6, 8}, {10, 12, 14, 16}, {18, 20, 22, 24}, {26, 28, 30, 32} }
 */
func (m Mat4[T]) ScalarMul(scalar T) Mat4[T] {
	return Mat4[T]{
		{m[0][0] * scalar, m[0][1] * scalar, m[0][2] * scalar, m[0][3] * scalar},
		{m[1][0] * scalar, m[1][1] * scalar, m[1][2] * scalar, m[1][3] * scalar},
		{m[2][0] * scalar, m[2][1] * scalar, m[2][2] * scalar, m[2][3] * scalar},
		{m[3][0] * scalar, m[3][1] * scalar, m[3][2] * scalar, m[3][3] * scalar},
	}
}

/**
 * Transpose returns the transpose of the matrix.
 * For example:
 *   Mat4[int]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }.Transpose()
 *   returns Mat4[int]{ {1, 5, 9, 13}, {2, 6, 10, 14}, {3, 7, 11, 15}, {4, 8, 12, 16} }
 */
func (m Mat4[T]) Transpose() Mat4[T] {
	return Mat4[T]{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}
}

/**
 * Determinant calculates the determinant of the matrix.
 * For example:
 *   Mat4[float64]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }.Determinant()
 *   returns 0 (as this matrix is singular)
 */
func (m Mat4[T]) Determinant() T {
	var a T
	for i := 0; i < 4; i++ {
		a += m[0][i] * m.Cofactor(0, i)
	}
	return a
}

/**
 * Cofactor calculates the cofactor of an element in the matrix.
 * For example:
 *   Mat4[int]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }.Cofactor(0, 0)
 *   returns -216 (the cofactor of the element at row 0, column 0)
 */
func (m Mat4[T]) Cofactor(row, col int) T {
	minor := m.Minor(row, col)
	if (row+col)%2 == 0 {
		return minor
	}
	return -minor
}

/**
 * Minor calculates the minor of an element in the matrix.
 * For example:
 *   Mat4[int]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }.Minor(0, 0)
 *   returns 216 (the minor of the element at row 0, column 0)
 */
func (m Mat4[T]) Minor(row, col int) T {
	var subMat [3][3]T
	r, c := 0, 0
	for i := 0; i < 4; i++ {
		if i == row {
			continue
		}
		c = 0
		for j := 0; j < 4; j++ {
			if j == col {
				continue
			}
			subMat[r][c] = m[i][j]
			c++
		}
		r++
	}
	return Mat3[T](subMat).Determinant()
}

/**
 * Inverse returns the inverse of the matrix and a boolean indicating success.
 * For example:
 *   Mat4[float64]{ {1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1} }.Inverse()
 *   returns (Mat4[float64]{ {1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1} }, true)
 *   Mat4[float64]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }.Inverse()
 *   returns (Mat4[float64]{}, false) (as this matrix is singular)
 */
func (m Mat4[T]) Inverse() (Mat4[T], bool) {
	det := m.Determinant()
	if det == 0 {
		return Mat4[T]{}, false
	}
	var adjoint Mat4[T]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			adjoint[j][i] = m.Cofactor(i, j)
		}
	}
	invDet := 1 / det
	return adjoint.ScalarMul(invDet), true
}

/**
 * String returns a string representation of the matrix.
 * For example:
 *   Mat4[int]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }.String()
 *   returns "[[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12], [13, 14, 15, 16]]"
 */
func (m Mat4[T]) String() string {
	return fmt.Sprintf("[[%v, %v, %v, %v], [%v, %v, %v, %v], [%v, %v, %v, %v], [%v, %v, %v, %v]]",
		m[0][0], m[0][1], m[0][2], m[0][3],
		m[1][0], m[1][1], m[1][2], m[1][3],
		m[2][0], m[2][1], m[2][2], m[2][3],
		m[3][0], m[3][1], m[3][2], m[3][3])
}

/**
 * Norm calculates the Frobenius norm of the matrix.
 * For example:
 *   Mat4[float64]{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16} }.Norm()
 *   returns 30.594117 (approximate value)
 */
func (m Mat4[T]) Norm() T {
	var sum T
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum += m[i][j] * m[i][j]
		}
	}
	return Sqrt(sum)
}

/**
 * RotateX returns a rotation matrix around the X-axis for a given angle in radians.
 * For example:
 *   Mat4[float64]{}.RotateX(math.Pi / 2)
 *   returns Mat4[float64]{ {1, 0, 0, 0}, {0, 0, -1, 0}, {0, 1, 0, 0}, {0, 0, 0, 1} }
 */
func (m Mat4[T]) RotateX(angle T) Mat4[T] {
	cosA := Cos(angle)
	sinA := Sin(angle)
	return Mat4[T]{
		{1, 0, 0, 0},
		{0, cosA, -sinA, 0},
		{0, sinA, cosA, 0},
		{0, 0, 0, 1},
	}
}

/**
 * RotateY returns a rotation matrix around the Y-axis for a given angle in radians.
 * For example:
 *   Mat4[float64]{}.RotateY(math.Pi / 2)
 *   returns Mat4[float64]{ {0, 0, 1, 0}, {0, 1, 0, 0}, {-1, 0, 0, 0}, {0, 0, 0, 1} }
 */
func (m Mat4[T]) RotateY(angle T) Mat4[T] {
	cosA := Cos(angle)
	sinA := Sin(angle)
	return Mat4[T]{
		{cosA, 0, sinA, 0},
		{0, 1, 0, 0},
		{-sinA, 0, cosA, 0},
		{0, 0, 0, 1},
	}
}

/**
 * RotateZ returns a rotation matrix around the Z-axis for a given angle in radians.
 * For example:
 *   Mat4[float64]{}.RotateZ(math.Pi / 2)
 *   returns Mat4[float64]{ {0, -1, 0, 0}, {1, 0, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1} }
 */
func (m Mat4[T]) RotateZ(angle T) Mat4[T] {
	cosA := Cos(angle)
	sinA := Sin(angle)
	return Mat4[T]{
		{cosA, -sinA, 0, 0},
		{sinA, cosA, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}
