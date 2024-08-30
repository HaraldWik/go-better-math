package bm

type Vec3[T Numeric] struct {
	X, Y, Z T
}

func NewVec3[T Numeric](x, y, z T) Vec3[T] {
	return Vec3[T]{X: x, Y: y, Z: z}
}

func (v Vec3[T]) Add(other Vec3[T]) Vec3[T] {
	return Vec3[T]{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

func (v Vec3[T]) Sub(other Vec3[T]) Vec3[T] {
	return Vec3[T]{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

func (v Vec3[T]) Mul(other Vec3[T]) Vec3[T] {
	return Vec3[T]{X: v.X * other.X, Y: v.Y * other.Y, Z: v.Z * other.Z}
}

func (v Vec3[T]) Div(other Vec3[T]) Vec3[T] {
	return Vec3[T]{X: v.X / other.X, Y: v.Y / other.Y, Z: v.Z / other.Z}
}

func (v Vec3[T]) Scale(scalar T) Vec3[T] {
	return Vec3[T]{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

func (v Vec3[T]) Dot(other Vec3[T]) T {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3[T]) Cross(other Vec3[T]) Vec3[T] {
	return Vec3[T]{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v Vec3[T]) Mag() T {
	return Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3[T]) Norm() Vec3[T] {
	mag := v.Mag()
	if mag == 0 {
		return Vec3[T]{0, 0, 0}
	}
	return Vec3[T]{
		X: v.X / mag,
		Y: v.Y / mag,
		Z: v.Z / mag,
	}
}

func (v Vec3[T]) Angle(other Vec3[T]) T {
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

func (v Vec3[T]) Neg() Vec3[T] {
	return Vec3[T]{X: -v.X, Y: -v.Y, Z: -v.Z}
}

func (v Vec3[T]) Abs() Vec3[T] {
	return Vec3[T]{
		X: Abs(v.X),
		Y: Abs(v.Y),
		Z: Abs(v.Z),
	}
}

func (v Vec3[T]) Clamp(min, max T) Vec3[T] {
	return Vec3[T]{
		X: Clamp(v.X, min, max),
		Y: Clamp(v.Y, min, max),
		Z: Clamp(v.Z, min, max),
	}
}

func (v Vec3[T]) Dist(other Vec3[T]) T {
	return v.Sub(other).Mag()
}

func (v Vec3[T]) Lerp(other Vec3[T], t T) Vec3[T] {
	return Vec3[T]{
		X: v.X + (other.X-v.X)*t,
		Y: v.Y + (other.Y-v.Y)*t,
		Z: v.Z + (other.Z-v.Z)*t,
	}
}

func (v Vec3[T]) Reflect(normal Vec3[T]) Vec3[T] {
	dot := v.Dot(normal)
	return v.Sub(normal.Scale(dot * 2))
}
