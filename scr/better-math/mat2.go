package bm

import (
	"fmt"
)

// Mat2 represents a 2x2 matrix with elements of type T.
type Mat2[T Numeric] [2][2]T

/**
 * IdentityMat2 returns a 2x2 identity matrix. The identity matrix has ones on the diagonal and zeros elsewhere.
 * For example:
 *   IdentityMat2[int]() returns Mat2[int]{ {1, 0}, {0, 1} }
 *   IdentityMat2[float64]() returns Mat2[float64]{ {1, 0}, {0, 1} }
 */
func IdentityMat2[T Numeric]() Mat2[T] {
	return Mat2[T]{
		{1, 0},
		{0, 1},
	}
}

/**
 * Add adds another 2x2 matrix to the current matrix and returns the result.
 * For example:
 *   Mat2[int]{ {1, 2}, {3, 4} }.Add(Mat2[int]{ {5, 6}, {7, 8} }) returns Mat2[int]{ {6, 8}, {10, 12} }
 */
func (m Mat2[T]) Add(n Mat2[T]) Mat2[T] {
	return Mat2[T]{
		{m[0][0] + n[0][0], m[0][1] + n[0][1]},
		{m[1][0] + n[1][0], m[1][1] + n[1][1]},
	}
}

/**
 * Sub subtracts another 2x2 matrix from the current matrix and returns the result.
 * For example:
 *   Mat2[int]{ {5, 6}, {7, 8} }.Sub(Mat2[int]{ {1, 2}, {3, 4} }) returns Mat2[int]{ {4, 4}, {4, 4} }
 */
func (m Mat2[T]) Sub(n Mat2[T]) Mat2[T] {
	return Mat2[T]{
		{m[0][0] - n[0][0], m[0][1] - n[0][1]},
		{m[1][0] - n[1][0], m[1][1] - n[1][1]},
	}
}

/**
 * Mul multiplies the current matrix by another 2x2 matrix and returns the result.
 * For example:
 *   Mat2[int]{ {1, 2}, {3, 4} }.Mul(Mat2[int]{ {5, 6}, {7, 8} }) returns Mat2[int]{ {19, 22}, {43, 50} }
 */
func (m Mat2[T]) Mul(n Mat2[T]) Mat2[T] {
	return Mat2[T]{
		{
			m[0][0]*n[0][0] + m[0][1]*n[1][0],
			m[0][0]*n[0][1] + m[0][1]*n[1][1],
		},
		{
			m[1][0]*n[0][0] + m[1][1]*n[1][0],
			m[1][0]*n[0][1] + m[1][1]*n[1][1],
		},
	}
}

/**
 * Div divides the current matrix by another 2x2 matrix (element-wise) and returns the result.
 * For example:
 *   Mat2[float64]{ {2, 4}, {6, 8} }.Div(Mat2[float64]{ {1, 2}, {3, 4} }) returns Mat2[float64]{ {2, 2}, {2, 2} }
 */
func (m Mat2[T]) Div(n Mat2[T]) Mat2[T] {
	return Mat2[T]{
		{m[0][0] / n[0][0], m[0][1] / n[0][1]},
		{m[1][0] / n[1][0], m[1][1] / n[1][1]},
	}
}

/**
 * Scale multiplies the current matrix by a scalar and returns the result.
 * For example:
 *   Mat2[int]{ {1, 2}, {3, 4} }.Scale(2) returns Mat2[int]{ {2, 4}, {6, 8} }
 */
func (m Mat2[T]) Scale(scalar T) Mat2[T] {
	return Mat2[T]{
		{m[0][0] * scalar, m[0][1] * scalar},
		{m[1][0] * scalar, m[1][1] * scalar},
	}
}

/**
 * Transpose returns the transpose of the matrix. The transpose is obtained by swapping rows with columns.
 * For example:
 *   Mat2[int]{ {1, 2}, {3, 4} }.Transpose() returns Mat2[int]{ {1, 3}, {2, 4} }
 */
func (m Mat2[T]) Transpose() Mat2[T] {
	return Mat2[T]{
		{m[0][0], m[1][0]},
		{m[0][1], m[1][1]},
	}
}

/**
 * Determinant calculates the determinant of the matrix. The determinant is a scalar value that provides information about the matrix.
 * For example:
 *   Mat2[int]{ {1, 2}, {3, 4} }.Determinant() returns -2
 */
func (m Mat2[T]) Determinant() T {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

/**
 * Inverse returns the inverse of the matrix. The inverse is a matrix such that when multiplied with the original matrix, results in the identity matrix.
 * If the matrix is singular (determinant is zero), it returns false indicating that the inverse does not exist.
 * For example:
 *   Mat2[int]{ {1, 2}, {3, 4} }.Inverse() returns (Mat2[int]{ {-2, 1}, {1.5, -0.5} }, true)
 *   Mat2[int]{ {1, 2}, {2, 4} }.Inverse() returns (Mat2[int]{}, false) // Singular matrix
 */
func (m Mat2[T]) Inverse() (Mat2[T], bool) {
	det := m.Determinant()
	if Abs(float32(det)) < 1e-10 {
		return Mat2[T]{}, false // Matrix is singular, no inverse
	}
	invDet := 1 / det
	return Mat2[T]{
		{m[1][1] * invDet, -m[0][1] * invDet},
		{-m[1][0] * invDet, m[0][0] * invDet},
	}, true
}

/**
 * String returns a string representation of the matrix in a human-readable format.
 * For example:
 *   Mat2[int]{ {1, 2}, {3, 4} }.String() returns "[[1, 2], [3, 4]]"
 */
func (m Mat2[T]) String() string {
	return fmt.Sprintf("[[%v, %v], [%v, %v]]", m[0][0], m[0][1], m[1][0], m[1][1])
}

/**
 * Norm calculates the Frobenius norm of the matrix. The Frobenius norm is the square root of the sum of the squares of all elements.
 * For example:
 *   Mat2[float64]{ {1, 2}, {3, 4} }.Norm() returns approximately 5.477
 */
func (m Mat2[T]) Norm() T {
	var sum T
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			sum += m[i][j] * m[i][j]
		}
	}
	return Sqrt(sum)
}

/**
 * Rotate returns a rotation matrix around the origin for a given angle in radians. The rotation matrix is used to rotate points in the 2D plane.
 * For example:
 *   Mat2[float64]{ {1, 0}, {0, 1} }.Rotate(math.Pi/2) returns Mat2[float64]{ {0, -1}, {1, 0} }
 */
func (m Mat2[T]) Rotate(angle T) Mat2[T] {
	var cosA, sinA T = Cos(angle), Sin(angle)
	return Mat2[T]{
		{cosA, -sinA},
		{sinA, cosA},
	}
}
