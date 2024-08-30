package bm

type Vec2[T Numeric] struct {
	X, Y T
}

func NewVec2[T Numeric](x, y T) Vec2[T] {
	return Vec2[T]{X: x, Y: y}
}

func (v Vec2[T]) Add(other Vec2[T]) Vec2[T] {
	return Vec2[T]{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vec2[T]) Sub(other Vec2[T]) Vec2[T] {
	return Vec2[T]{X: v.X - other.X, Y: v.Y - other.Y}
}

func (v Vec2[T]) Mul(other Vec2[T]) Vec2[T] {
	return Vec2[T]{X: v.X * other.X, Y: v.Y * other.Y}
}

func (v Vec2[T]) Div(other Vec2[T]) Vec2[T] {
	return Vec2[T]{X: v.X / other.X, Y: v.Y / other.Y}
}

func (v Vec2[T]) Scale(scalar T) Vec2[T] {
	return Vec2[T]{X: v.X * scalar, Y: v.Y * scalar}
}

func (v Vec2[T]) Dot(other Vec2[T]) T {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2[T]) Mag() T {
	return Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2[T]) Norm() Vec2[T] {
	mag := v.Mag()
	if mag == 0 {
		return Vec2[T]{0, 0}
	}
	return Vec2[T]{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}

func (v Vec2[T]) Angle(other Vec2[T]) T {
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

func (v Vec2[T]) Neg() Vec2[T] {
	return Vec2[T]{X: -v.X, Y: -v.Y}
}

func (v Vec2[T]) Abs() Vec2[T] {
	return Vec2[T]{
		X: Abs(v.X),
		Y: Abs(v.Y),
	}
}

func (v Vec2[T]) Clamp(min, max T) Vec2[T] {
	return Vec2[T]{
		X: Clamp(v.X, min, max),
		Y: Clamp(v.Y, min, max),
	}
}

func (v Vec2[T]) Dist(other Vec2[T]) T {
	return v.Sub(other).Mag()
}

func (v Vec2[T]) Lerp(other Vec2[T], t T) Vec2[T] {
	return Vec2[T]{
		X: v.X + (other.X-v.X)*t,
		Y: v.Y + (other.Y-v.Y)*t,
	}
}

func (v Vec2[T]) Reflect(normal Vec2[T]) Vec2[T] {
	dot := v.Dot(normal)
	return v.Sub(normal.Scale(dot * 2))
}
