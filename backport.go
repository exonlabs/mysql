package mysql

// Ordered is a constraint that allows any ordered type, i.e.,
// types that support comparison operators like <, <=, >=, and >.
// This includes various numeric types (integers and floats) and strings.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// From Go v1.24
// The min built-in function returns the smallest value of a fixed number of
// arguments of [cmp.Ordered] types. There must be at least one argument.
// If T is a floating-point type and any of the arguments are NaNs,
// min will return NaN.
func min[T Ordered](x T, y ...T) T {
	r := x
	for _, v := range y {
		if v < r {
			r = v
		}
	}
	return r
}
