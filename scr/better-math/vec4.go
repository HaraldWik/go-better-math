package bm

type Vec4[T Numeric] struct {
	X, Y, Z, W T
}

func NewVec4[T Numeric](x, y, z, w T) Vec4[T] {
	return Vec4[T]{X: x, Y: y, Z: z, W: w}
}

func (v Vec4[T]) Add(other Vec4[T]) Vec4[T] {
	return Vec4[T]{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z, W: v.W + other.W}
}

func (v Vec4[T]) Sub(other Vec4[T]) Vec4[T] {
	return Vec4[T]{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z, W: v.W - other.W}
}

func (v Vec4[T]) Mul(other Vec4[T]) Vec4[T] {
	return Vec4[T]{X: v.X * other.X, Y: v.Y * other.Y, Z: v.Z * other.Z, W: v.W * other.W}
}

func (v Vec4[T]) Div(other Vec4[T]) Vec4[T] {
	return Vec4[T]{X: v.X / other.X, Y: v.Y / other.Y, Z: v.Z / other.Z, W: v.W / other.W}
}

func (v Vec4[T]) Scale(scalar T) Vec4[T] {
	return Vec4[T]{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar, W: v.W * scalar}
}

func (v Vec4[T]) Dot(other Vec4[T]) T {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z + v.W*other.W
}

func (v Vec4[T]) Mag() T {
	return Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

func (v Vec4[T]) Norm() Vec4[T] {
	mag := v.Mag()
	if mag == 0 {
		return Vec4[T]{0, 0, 0, 0}
	}
	return Vec4[T]{
		X: v.X / mag,
		Y: v.Y / mag,
		Z: v.Z / mag,
		W: v.W / mag,
	}
}

func (v Vec4[T]) Angle(other Vec4[T]) T {
	dot := v.Dot(other)
	magA := v.Mag()
	magB := other.Mag()

	cosTheta := float32(dot / (magA * magB))

	if cosTheta > 1.0 {
		cosTheta = 1.0
	} else if cosTheta < -1.0 {
		cosTheta = -1.0
	}

	angle := Acos(cosTheta)

	return T(angle)
}

func (v Vec4[T]) Neg() Vec4[T] {
	return Vec4[T]{X: -v.X, Y: -v.Y, Z: -v.Z, W: -v.W}
}

func (v Vec4[T]) Abs() Vec4[T] {
	return Vec4[T]{
		X: Abs(v.X),
		Y: Abs(v.Y),
		Z: Abs(v.Z),
		W: Abs(v.W),
	}
}

func (v Vec4[T]) Clamp(min, max T) Vec4[T] {
	return Vec4[T]{
		X: Clamp(v.X, min, max),
		Y: Clamp(v.Y, min, max),
		Z: Clamp(v.Z, min, max),
		W: Clamp(v.W, min, max),
	}
}

func (v Vec4[T]) Dist(other Vec4[T]) T {
	return v.Sub(other).Mag()
}

func (v Vec4[T]) Lerp(other Vec4[T], t T) Vec4[T] {
	return Vec4[T]{
		X: v.X + (other.X-v.X)*t,
		Y: v.Y + (other.Y-v.Y)*t,
		Z: v.Z + (other.Z-v.Z)*t,
		W: v.W + (other.W-v.W)*t,
	}
}

func (v Vec4[T]) Reflect(normal Vec4[T]) Vec4[T] {
	dot := v.Dot(normal)
	return v.Sub(normal.Scale(dot * 2))
}
