package mathutils

/*
	Package mathutils: exported numeric utility functions.

	Second subpackage to reinforce the module structure pattern.
	Same exported/unexported rules apply — Clamp and Percentage are visible
	from main, the internal clamp helper is not.

	Notice that both stringutils and mathutils are separate packages
	under the same module (learngo/modules-practice).
	A module can contain as many packages as needed, each in its own subdirectory.
*/

// Clamp returns value v constrained to the range [min, max].
// Exported — useful utility in trading systems for capping order sizes,
// price ranges, or any value that must stay within bounds.
func Clamp(v, min, max int) int {
	// Delegating to unexported helper
	// The caller (main) doesn't need to know clamp() exists — just uses Clamp()
	return clamp(v, min, max)
}

// Percentage returns what percent `part` is of `total` as a float64.
// Exported — handles divide-by-zero by returning 0 as a sensible default.
func Percentage(part, total int) float64 {
	if total == 0 {
		// returning a sensible zero value instead of panicking
		// this is the Go way: handle expected edge cases, don't panic
		return 0
	}
	return (float64(part) / float64(total)) * 100
}

// clamp is the unexported internal implementation.
// Exported Clamp() delegates here — this separation lets us change the
// internal logic later without affecting the public API signature.
// main package cannot call mathutils.clamp() — only mathutils.Clamp().
func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
