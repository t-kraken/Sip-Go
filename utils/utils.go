package utils

// Check two string pointers for equality as follows:
// - If neither pointer is nil, check equality of the underlying strings.
// - If either pointer is nil, return true if and only if they both are.
func StrPtrEq(a *string, b *string) bool {
	if a == nil || b == nil {
		return a == b
	}

	return *a == *b
}

// Check two uint16 pointers for equality as follows:
// - If neither pointer is nil, check equality of the underlying uint16s.
// - If either pointer is nil, return true if and only if they both are.
func Uint16PtrEq(a *uint16, b *uint16) bool {
	if a == nil || b == nil {
		return a == b
	}

	return *a == *b
}
