package helpers

func ToPtr[T any](input T) *T {
	return &input
}
