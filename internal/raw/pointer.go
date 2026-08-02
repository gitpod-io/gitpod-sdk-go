package client

func Pointer[T any](v T) *T {
	return &v
}

func ToString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
