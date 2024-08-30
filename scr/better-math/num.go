package bm

import (
	"log"
	"math"
)

// Numeric is a generic type that includes various numeric types.
type Numeric interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Constants

/**
 * Pi is the mathematical constant π, the ratio of a circle's circumference to its diameter.
 * For example:
 *   Pi returns approximately 3.141592653589793
 */
const Pi = math.Pi

/**
 * E is the mathematical constant e, the base of the natural logarithm.
 * For example:
 *   E returns approximately 2.718281828459045
 */
const E = math.E

/**
 * Phi is the golden ratio, approximately 1.618033988749895.
 * For example:
 *   Phi returns approximately 1.618033988749895
 */
const Phi = math.Phi

// Basic Arithmetic and Utility Functions

/**
 * Abs returns the absolute value of x.
 * For example:
 *   Abs(-5) returns 5
 *   Abs(3) returns 3
 */
func Abs[T Numeric](x T) T {
	return T(math.Abs(float64(x)))
}

/**
 * Lerp performs linear interpolation between a and b with the given ratio t.
 * If t is 0, it returns a. If t is 1, it returns b. If t is between 0 and 1,
 * it returns a value between a and b.
 * For example:
 *   Lerp(0, 10, 0.5) returns 5
 */
func Lerp[T Numeric](a, b T, t float64) T {
	return T(float64(a) + (float64(b)-float64(a))*t)
}

/**
 * Clamp restricts the value of x to be within the range [min, max].
 * If x is less than min, it returns min.
 * If x is greater than max, it returns max.
 * Otherwise, it returns x.
 * For example:
 *   Clamp(5, 1, 10) returns 5
 *   Clamp(-3, 0, 10) returns 0
 *   Clamp(15, 1, 10) returns 10
 */
func Clamp[T Numeric](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

/**
 * Max returns the greater of two values a and b.
 * For example:
 *   Max(3, 5) returns 5
 *   Max(7.2, 7.1) returns 7.2
 */
func Max[T Numeric](a, b T) T {
	if a > b {
		return a
	}
	return b
}

/**
 * Min returns the lesser of two values a and b.
 * For example:
 *   Min(3, 5) returns 3
 *   Min(7.2, 7.1) returns 7.1
 */
func Min[T Numeric](a, b T) T {
	if a < b {
		return a
	}
	return b
}

/**
 * Mod returns the remainder of x divided by y.
 * For example:
 *   Mod(7, 3) returns 1
 *   Mod(-7, 3) returns 2
 */
func Mod[T Numeric](x, y T) T {
	return T(math.Mod(float64(x), float64(y)))
}

/**
 * Pow returns x raised to the power of y.
 * For example:
 *   Pow(2, 3) returns 8
 *   Pow(9, 0.5) returns 3
 */
func Pow[T Numeric](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

/**
 * Round returns the nearest integer, rounding half away from zero.
 * For example:
 *   Round(1.5) returns 2
 *   Round(-1.5) returns -2
 */
func Round[T Numeric](x T) T {
	return T(math.Round(float64(x)))
}

/**
 * Trunc returns the integer value of x, removing any fractional part.
 * For example:
 *   Trunc(3.9) returns 3
 *   Trunc(-3.9) returns -3
 */
func Trunc[T Numeric](x T) T {
	return T(math.Trunc(float64(x)))
}

/**
 * Sqrt returns the square root of x.
 * For example:
 *   Sqrt(9) returns 3
 * If x is negative, it returns 0.
 */
func Sqrt[T Numeric](x T) T {
	if x < 0 {
		return 0 // Default value for undefined square root
	}
	return T(math.Sqrt(float64(x)))
}

/**
 * Cbrt returns the cube root of x.
 * For example:
 *   Cbrt(27) returns 3
 */
func Cbrt[T Numeric](x T) T {
	return T(math.Cbrt(float64(x)))
}

/**
 * Ldexp returns x * (2^exp).
 * For example:
 *   Ldexp(1, 3) returns 8
 */
func Ldexp[T Numeric](x T, exp int) T {
	return T(math.Ldexp(float64(x), exp))
}

/**
 * Hypot returns the hypotenuse of x and y.
 * For example:
 *   Hypot(3, 4) returns 5
 */
func Hypot[T Numeric](x, y T) T {
	return T(math.Hypot(float64(x), float64(y)))
}

/**
 * Signbit returns true if x is negative.
 * For example:
 *   Signbit(-1) returns true
 *   Signbit(1) returns false
 */
func Signbit[T Numeric](x T) bool {
	return math.Signbit(float64(x))
}

/**
 * Copysign returns a value with the magnitude of x and the sign of sign.
 * For example:
 *   Copysign(1, -1) returns -1
 */
func Copysign[T Numeric](x, sign T) T {
	return T(math.Copysign(float64(x), float64(sign)))
}

/**
 * Dim returns the positive difference between x and y.
 * For example:
 *   Dim(5, 3) returns 2
 *   Dim(3, 5) returns 0
 */
func Dim[T Numeric](x, y T) T {
	return T(math.Dim(float64(x), float64(y)))
}

/**
 * Ceil returns the smallest integer greater than or equal to x.
 * For example:
 *   Ceil(1.2) returns 2
 *   Ceil(-1.2) returns -1
 */
func Ceil[T Numeric](x T) T {
	return T(math.Ceil(float64(x)))
}

/**
 * Floor returns the largest integer less than or equal to x.
 * For example:
 *   Floor(1.2) returns 1
 *   Floor(-1.2) returns -2
 */
func Floor[T Numeric](x T) T {
	return T(math.Floor(float64(x)))
}

// Trigonometric Functions

/**
 * Sin returns the sine of x (in radians).
 * For example:
 *   Sin(0) returns 0
 *   Sin(π/2) returns 1
 */
func Sin[T Numeric](x T) T {
	return T(math.Sin(float64(x)))
}

/**
 * Cos returns the cosine of x (in radians).
 * For example:
 *   Cos(0) returns 1
 *   Cos(π/2) returns 0
 */
func Cos[T Numeric](x T) T {
	return T(math.Cos(float64(x)))
}

/**
 * Tan returns the tangent of x (in radians).
 * For example:
 *   Tan(0) returns 0
 *   Tan(π/4) returns 1
 */
func Tan[T Numeric](x T) T {
	return T(math.Tan(float64(x)))
}

/**
 * Asin returns the arc sine (inverse sine) of x.
 * The result is in radians, in the range [-π/2, π/2].
 * For example:
 *   Asin(0) returns 0
 *   Asin(1) returns π/2
 */
func Asin[T Numeric](x T) T {
	return T(math.Asin(float64(x)))
}

/**
 * Acos returns the arc cosine (inverse cosine) of x.
 * The result is in radians, in the range [0, π].
 * For example:
 *   Acos(1) returns 0
 *   Acos(0) returns π/2
 */
func Acos[T Numeric](x T) T {
	return T(math.Acos(float64(x)))
}

/**
 * Atan returns the arc tangent (inverse tangent) of x.
 * The result is in radians, in the range [-π/2, π/2].
 * For example:
 *   Atan(0) returns 0
 *   Atan(1) returns π/4
 */
func Atan[T Numeric](x T) T {
	return T(math.Atan(float64(x)))
}

/**
 * Atan2 returns the arc tangent of y/x, using the signs of both arguments to determine the quadrant.
 * The result is in radians, in the range [-π, π].
 * For example:
 *   Atan2(1, 1) returns π/4
 *   Atan2(-1, -1) returns -3π/4
 */
func Atan2[T Numeric](y, x T) T {
	return T(math.Atan2(float64(y), float64(x)))
}

// Hyperbolic Functions

/**
 * Sinh returns the hyperbolic sine of x.
 * For example:
 *   Sinh(0) returns 0
 *   Sinh(1) returns 1.1752011936438014
 */
func Sinh[T Numeric](x T) T {
	return T(math.Sinh(float64(x)))
}

/**
 * Cosh returns the hyperbolic cosine of x.
 * For example:
 *   Cosh(0) returns 1
 *   Cosh(1) returns 1.5430806348152437
 */
func Cosh[T Numeric](x T) T {
	return T(math.Cosh(float64(x)))
}

/**
 * Tanh returns the hyperbolic tangent of x.
 * For example:
 *   Tanh(0) returns 0
 *   Tanh(1) returns 0.7615941559557649
 */
func Tanh[T Numeric](x T) T {
	return T(math.Tanh(float64(x)))
}

/**
 * Asinh returns the inverse hyperbolic sine of x.
 * For example:
 *   Asinh(0) returns 0
 *   Asinh(1) returns 0.881373587019543
 */
func Asinh[T Numeric](x T) T {
	return T(math.Asinh(float64(x)))
}

/**
 * Acosh returns the inverse hyperbolic cosine of x.
 * For example:
 *   Acosh(1) returns 0
 *   Acosh(2) returns 1.3169578969248166
 */
func Acosh[T Numeric](x T) T {
	return T(math.Acosh(float64(x)))
}

/**
 * Atanh returns the inverse hyperbolic tangent of x.
 * For example:
 *   Atanh(0) returns 0
 *   Atanh(0.5) returns 0.5493061443340549
 */
func Atanh[T Numeric](x T) T {
	return T(math.Atanh(float64(x)))
}

// Exponential and Logarithmic Functions

/**
 * Exp returns e raised to the power of x.
 * For example:
 *   Exp(1) returns 2.718281828459045
 *   Exp(0) returns 1
 */
func Exp[T Numeric](x T) T {
	return T(math.Exp(float64(x)))
}

/**
 * Exp2 returns 2 raised to the power of x.
 * For example:
 *   Exp2(3) returns 8
 *   Exp2(0) returns 1
 */
func Exp2[T Numeric](x T) T {
	return T(math.Exp2(float64(x)))
}

/**
 * Log returns the natural logarithm (base e) of x.
 * If x is less than or equal to 0, it returns 0.
 * For example:
 *   Log(1) returns 0
 *   Log(e) returns 1
 */
func Log[T Numeric](x T) T {
	if x <= 0 {
		return 0 // Default value for undefined logarithm
	}
	return T(math.Log(float64(x)))
}

/**
 * Log2 returns the base-2 logarithm of x.
 * If x is less than or equal to 0, it returns 0.
 * For example:
 *   Log2(1) returns 0
 *   Log2(8) returns 3
 */
func Log2[T Numeric](x T) T {
	if x <= 0 {
		return 0 // Default value for undefined logarithm
	}
	return T(math.Log2(float64(x)))
}

/**
 * Log10 returns the base-10 logarithm of x.
 * If x is less than or equal to 0, it returns 0.
 * For example:
 *   Log10(1) returns 0
 *   Log10(100) returns 2
 */
func Log10[T Numeric](x T) T {
	if x <= 0 {
		return 0 // Default value for undefined logarithm
	}
	return T(math.Log10(float64(x)))
}

/**
 * Gamma returns the gamma function of x.
 * For example:
 *   Gamma(1) returns 1
 *   Gamma(2) returns 1
 */
func Gamma[T Numeric](x T) T {
	return T(math.Gamma(float64(x)))
}

/**
 * Lgamma returns the natural logarithm of the absolute value of the gamma function of x and a flag indicating
 * whether the result is valid.
 * For example:
 *   Lgamma(1) returns 0
 *   Lgamma(2) returns 0
 */
func Lgamma[T Numeric](x T) (T, int) {
	a, b := math.Lgamma(float64(x))
	return T(a), b
}

/**
 * Ilogb returns the exponent of the binary representation of x.
 * For example:
 *   Ilogb(8) returns 3
 *   Ilogb(0.5) returns -1
 */
func Ilogb[T Numeric](x T) int {
	return int(math.Ilogb(float64(x)))
}

// Bezier Functions

/**
 * Bezier2 computes a point on a quadratic Bézier curve defined by points p0, p1, and p2 with parameter t.
 * For example:
 *   Bezier2(0, 10, 20, 0.5) returns 10
 */
func Bezier2[T Numeric](p0, p1, p2 T, t float64) T {
	oneMinusT := 1 - t
	return T(float64(p0)*oneMinusT*oneMinusT + float64(p1)*2*oneMinusT*t + float64(p2)*t*t)
}

/**
 * Bezier3 computes a point on a cubic Bézier curve defined by points p0, p1, p2, and p3 with parameter t.
 * For example:
 *   Bezier3(0, 5, 15, 20, 0.5) returns 10
 */
func Bezier3[T Numeric](p0, p1, p2, p3 T, t float64) T {
	oneMinusT := 1 - t
	t2 := t * t
	oneMinusT2 := oneMinusT * oneMinusT
	return T(float64(p0)*oneMinusT2*oneMinusT + float64(p1)*3*oneMinusT2*t + float64(p2)*3*oneMinusT*t2 + float64(p3)*t*t2)
}

// Cubic Spline

/**
 * CubicSpline computes the value of a cubic spline interpolation for a given x value, based on the provided
 * spline coefficients a, b, c, d and x values.
 * For example:
 *   CubicSpline(x, a, b, c, d, 2.5) returns the interpolated value at x = 2.5
 * If x is out of the range of the provided x values, it returns 0.
 */
func CubicSpline[T Numeric](x []T, a, b, c, d []T, xVal T) T {
	if len(x) != len(a) || len(x) != len(b) || len(x) != len(c) || len(x) != len(d) || len(x) < 2 {
		log.Println("invalid input lengths")
		return 0 // Default value for invalid input
	}

	n := len(x)
	if xVal < x[0] || xVal > x[n-1] {
		log.Println("xVal out of bounds")
		return 0 // Default value for out-of-bounds
	}

	i := 0
	for i < n-1 && x[i+1] < xVal {
		i++
	}

	h := float64(xVal) - float64(x[i])
	return T(float64(a[i]) + float64(b[i])*h + float64(c[i])*h*h + float64(d[i])*h*h*h)
}
