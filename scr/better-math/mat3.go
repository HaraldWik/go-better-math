package bm

import "fmt"

// Mat3 represents a 3x3 matrix with elements of type T.
type Mat3[T Numeric] [3][3]T

/**
 * IdentityMat3 returns a 3x3 identity matrix. The identity matrix has ones on the diagonal and zeros elsewhere.
 * For example:
 *   IdentityMat3[int]() returns Mat3[int]{ {1, 0, 0}, {0, 1, 0}, {0, 0, 1} }
 *   IdentityMat3[float64]() returns Mat3[float64]{ {1, 0, 0}, {0, 1, 0}, {0, 0, 1} }
 */
func IdentityMat3[T Numeric]() Mat3[T] {
	return Mat3[T]{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

/**
 * Add adds another 3x3 matrix to the current matrix and returns the result.
 * For example:
 *   Mat3[int]{ {1, 2, 3}, {4, 5, 6}, {7, 8, 9} }.Add(Mat3[int]{ {9, 8, 7}, {6, 5, 4}, {3, 2, 1} })
 *   returns Mat3[int]{ {10, 10, 10}, {10, 10, 10}, {10, 10, 10} }
 */
func (m Mat3[T]) Add(n Mat3[T]) Mat3[T] {
	return Mat3[T]{
		{m[0][0] + n[0][0], m[0][1] + n[0][1], m[0][2] + n[0][2]},
		{m[1][0] + n[1][0], m[1][1] + n[1][1], m[1][2] + n[1][2]},
		{m[2][0] + n[2][0], m[2][1] + n[2][1], m[2][2] + n[2][2]},
	}
}

/**
 * Sub subtracts another 3x3 matrix from the current matrix and returns the result.
 * For example:
 *   Mat3[int]{ {9, 8, 7}, {6, 5, 4}, {3, 2, 1} }.Sub(Mat3[int]{ {1, 2, 3}, {4, 5, 6}, {7, 8, 9} })
 *   returns Mat3[int]{ {8, 6, 4}, {2, 0, -2}, {-4, -6, -8} }
 */
func (m Mat3[T]) Sub(n Mat3[T]) Mat3[T] {
	return Mat3[T]{
		{m[0][0] - n[0][0], m[0][1] - n[0][1], m[0][2] - n[0][2]},
		{m[1][0] - n[1][0], m[1][1] - n[1][1], m[1][2] - n[1][2]},
		{m[2][0] - n[2][0], m[2][1] - n[2][1], m[2][2] - n[2][2]},
	}
}

/**
 * Mul multiplies the current matrix by another 3x3 matrix and returns the result.
 * For example:
 *   Mat3[int]{ {1, 2, 3}, {4, 5, 6}, {7, 8, 9} }.Mul(Mat3[int]{ {9, 8, 7}, {6, 5, 4}, {3, 2, 1} })
 *   returns Mat3[int]{ {30, 24, 18}, {84, 69, 54}, {138, 114, 90} }
 */
func (m Mat3[T]) Mul(n Mat3[T]) Mat3[T] {
	return Mat3[T]{
		{
			m[0][0]*n[0][0] + m[0][1]*n[1][0] + m[0][2]*n[2][0],
			m[0][0]*n[0][1] + m[0][1]*n[1][1] + m[0][2]*n[2][1],
			m[0][0]*n[0][2] + m[0][1]*n[1][2] + m[0][2]*n[2][2],
		},
		{
			m[1][0]*n[0][0] + m[1][1]*n[1][0] + m[1][2]*n[2][0],
			m[1][0]*n[0][1] + m[1][1]*n[1][1] + m[1][2]*n[2][1],
			m[1][0]*n[0][2] + m[1][1]*n[1][2] + m[1][2]*n[2][2],
		},
		{
			m[2][0]*n[0][0] + m[2][1]*n[1][0] + m[2][2]*n[2][0],
			m[2][0]*n[0][1] + m[2][1]*n[1][1] + m[2][2]*n[2][1],
			m[2][0]*n[0][2] + m[2][1]*n[1][2] + m[2][2]*n[2][2],
		},
	}
}

/**
 * ScalarMul multiplies the matrix by a scalar and returns the result.
 * For example:
 *   Mat3[int]{ {1, 2, 3}, {4, 5, 6}, {7, 8, 9} }.ScalarMul(2)
 *   returns Mat3[int]{ {2, 4, 6}, {8, 10, 12}, {14, 16, 18} }
 */
func (m Mat3[T]) ScalarMul(scalar T) Mat3[T] {
	return Mat3[T]{
		{m[0][0] * scalar, m[0][1] * scalar, m[0][2] * scalar},
		{m[1][0] * scalar, m[1][1] * scalar, m[1][2] * scalar},
		{m[2][0] * scalar, m[2][1] * scalar, m[2][2] * scalar},
	}
}

/**
 * Transpose returns the transpose of the matrix. The transpose is obtained by swapping rows with columns.
 * For example:
 *   Mat3[int]{ {1, 2, 3}, {4, 5, 6}, {7, 8, 9} }.Transpose()
 *   returns Mat3[int]{ {1, 4, 7}, {2, 5, 8}, {3, 6, 9} }
 */
func (m Mat3[T]) Transpose() Mat3[T] {
	return Mat3[T]{
		{m[0][0], m[1][0], m[2][0]},
		{m[0][1], m[1][1], m[2][1]},
		{m[0][2], m[1][2], m[2][2]},
	}
}

/**
 * Determinant calculates the determinant of the matrix. The determinant is a scalar value that provides information about the matrix.
 * For example:
 *   Mat3[int]{ {1, 2, 3}, {4, 5, 6}, {7, 8, 9} }.Determinant()
 *   returns 0 (indicating the matrix is singular)
 */
func (m Mat3[T]) Determinant() T {
	return m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) -
		m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) +
		m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])
}

/**
 * Inverse returns the inverse of the matrix and a boolean indicating success. If the matrix is singular (determinant is zero), it returns false indicating that the inverse does not exist.
 * For example:
 *   Mat3[int]{ {1, 2, 3}, {0, 1, 4}, {5, 6, 0} }.Inverse()
 *   returns (Mat3[int]{ { -24, 18, 5 }, { 20, -15, -4 }, { -5, 4, 1 } }, true)
 *   Mat3[int]{ {1, 2, 3}, {4, 5, 6}, {7, 8, 9} }.Inverse()
 *   returns (Mat3[int]{}, false) since the matrix is singular.
 */
func (m Mat3[T]) Inverse() (Mat3[T], bool) {
	det := m.Determinant()
	if det == 0 {
		return Mat3[T]{}, false
	}
	invDet := 1 / det
	return Mat3[T]{
		{
			(m[1][1]*m[2][2] - m[1][2]*m[2][1]) * invDet,
			(m[0][2]*m[2][1] - m[0][1]*m[2][2]) * invDet,
			(m[0][1]*m[1][2] - m[0][2]*m[1][1]) * invDet,
		},
		{
			(m[1][2]*m[2][0] - m[1][0]*m[2][2]) * invDet,
			(m[0][0]*m[2][2] - m[0][2]*m[2][0]) * invDet,
			(m[0][2]*m[1][0] - m[0][0]*m[1][2]) * invDet,
		},
		{
			(m[1][0]*m[2][1] - m[1][1]*m[2][0]) * invDet,
			(m[0][1]*m[2][0] - m[0][0]*m[2][1]) * invDet,
			(m[0][0]*m[1][1] - m[0][1]*m[1][0]) * invDet,
		},
	}, true
}

/**
 * String returns a string representation of the matrix.
 * For example:
 *   Mat3[int]{ {1, 2, 3}, {4, 5, 6}, {7, 8, 9} }.String()
 *   returns "[[1, 2, 3], [4, 5, 6], [7, 8, 9]]"
 */
func (m Mat3[T]) String() string {
	return fmt.Sprintf("[[%v, %v, %v], [%v, %v, %v], [%v, %v, %v]]",
		m[0][0], m[0][1], m[0][2],
		m[1][0], m[1][1], m[1][2],
		m[2][0], m[2][1], m[2][2])
}

/**
 * Norm calculates the Frobenius norm of the matrix, which is the square root of the sum of the absolute squares of its elements.
 * For example:
 *   Mat3[int]{ {1, 2, 3}, {4, 5, 6}, {7, 8, 9} }.Norm()
 *   returns 16.881943 (which is sqrt(1^2 + 2^2 + ... + 9^2))
 */
func (m Mat3[T]) Norm() T {
	var sum T
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += m[i][j] * m[i][j]
		}
	}
	return Sqrt(sum)
}

/**
 * RotateX returns a 3x3 rotation matrix for a rotation around the x-axis by a given angle in radians.
 * For example:
 *   RotateX(math.Pi / 4) returns Mat3[float64]{ {1, 0, 0}, {0, 0.7071, -0.7071}, {0, 0.7071, 0.7071} }
 */
func RotateX[T Numeric](angle T) Mat3[T] {
	var cosA, sinA T = Cos(angle), Sin(angle)
	return Mat3[T]{
		{1, 0, 0},
		{0, cosA, -sinA},
		{0, sinA, cosA},
	}
}

/**
 * RotateY returns a 3x3 rotation matrix for a rotation around the y-axis by a given angle in radians.
 * For example:
 *   RotateY(math.Pi / 4) returns Mat3[float64]{ {0.7071, 0, 0.7071}, {0, 1, 0}, {-0.7071, 0, 0.7071} }
 */
func RotateY[T Numeric](angle T) Mat3[T] {
	var cosA, sinA T = Cos(angle), Sin(angle)
	return Mat3[T]{
		{cosA, 0, sinA},
		{0, 1, 0},
		{-sinA, 0, cosA},
	}
}

/**
 * RotateZ returns a 3x3 rotation matrix for a rotation around the z-axis by a given angle in radians.
 * For example:
 *   RotateZ(math.Pi / 4) returns Mat3[float64]{ {0.7071, -0.7071, 0}, {0.7071, 0.7071, 0}, {0, 0, 1} }
 */
func RotateZ[T Numeric](angle T) Mat3[T] {
	var cosA, sinA T = Cos(angle), Sin(angle)
	return Mat3[T]{
		{cosA, -sinA, 0},
		{sinA, cosA, 0},
		{0, 0, 1},
	}
}

/**
 * RotateX returns a rotation matrix around the X-axis for a given angle in radians.
 * For example:
 *   Mat3[float64]{ {1, 0, 0}, {0, 0.7071, -0.7071}, {0, 0.7071, 0.7071} }.RotateX(math.Pi / 4)
 *   returns Mat3[float64]{ {1, 0, 0}, {0, 0.7071, -0.7071}, {0, 0.7071, 0.7071} }
 */
func (m Mat3[T]) RotateX(angle T) Mat3[T] {
	cosA := Cos(angle)
	sinA := Sin(angle)
	return Mat3[T]{
		{1, 0, 0},
		{0, cosA, -sinA},
		{0, sinA, cosA},
	}
}

/**
 * RotateY returns a rotation matrix around the Y-axis for a given angle in radians.
 * For example:
 *   Mat3[float64]{ {0.7071, 0, 0.7071}, {0, 1, 0}, {-0.7071, 0, 0.7071} }.RotateY(math.Pi / 4)
 *   returns Mat3[float64]{ {0.7071, 0, 0.7071}, {0, 1, 0}, {-0.7071, 0, 0.7071} }
 */
func (m Mat3[T]) RotateY(angle T) Mat3[T] {
	cosA := Cos(angle)
	sinA := Sin(angle)
	return Mat3[T]{
		{cosA, 0, sinA},
		{0, 1, 0},
		{-sinA, 0, cosA},
	}
}

/**
 * RotateZ returns a rotation matrix around the Z-axis for a given angle in radians.
 * For example:
 *   Mat3[float64]{ {0.7071, -0.7071, 0}, {0.7071, 0.7071, 0}, {0, 0, 1} }.RotateZ(math.Pi / 4)
 *   returns Mat3[float64]{ {0.7071, -0.7071, 0}, {0.7071, 0.7071, 0}, {0, 0, 1} }
 */
func (m Mat3[T]) RotateZ(angle T) Mat3[T] {
	cosA := Cos(angle)
	sinA := Sin(angle)
	return Mat3[T]{
		{cosA, -sinA, 0},
		{sinA, cosA, 0},
		{0, 0, 1},
	}
}
