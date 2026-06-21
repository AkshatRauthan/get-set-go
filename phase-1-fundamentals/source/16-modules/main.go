package main

// RUN COMMAND (from inside 16-modules/ folder):
// go run main.go

// SETUP COMMAND (only needed once, already done if go.mod exists):
// go mod init learngo/modules-practice

import (
	"fmt"

	// Importing our own subpackages using the module path declared in go.mod
	// Module path (learngo/modules-practice) + subfolder name = import path
	"github.com/developer-akshat/16-modules/mathutils"
	"github.com/developer-akshat/16-modules/stringutils"
)

func main() {

	// 01. Using exported functions from stringutils package
	// Only Reverse() and CountVowels() are accessible here — isVowel() is unexported
	fmt.Print("\n01. String Utils Package\n")
	StringUtilsDemo()

	// 02. Using exported functions from mathutils package
	fmt.Print("\n\n02. Math Utils Package\n")
	MathUtilsDemo()
}

func StringUtilsDemo() {

	// Reverse is exported (capital R) — accessible from main package
	reversed := stringutils.Reverse("golang")
	fmt.Println("Reverse of 'golang':", reversed)

	// CountVowels is exported — accessible from main package
	count := stringutils.CountVowels("backend developer")
	fmt.Println("Vowel count in 'backend developer':", count)

	// isVowel is unexported (lowercase i) — lives inside stringutils only
	// Trying to call stringutils.isVowel() here would be a compile error:
	// "cannot refer to unexported name stringutils.isVowel"
}

func MathUtilsDemo() {

	// Clamp is exported — keeps a value between a min and max bound
	// Common in trading systems: clamp order quantity between allowed limits
	clamped := mathutils.Clamp(150, 0, 100)
	fmt.Println("Clamp(150, min=0, max=100):", clamped) // => 100

	clamped = mathutils.Clamp(50, 0, 100)
	fmt.Println("Clamp(50, min=0, max=100):", clamped) // => 50

	// Percentage is exported — calculates what percent `part` is of `total`
	pct := mathutils.Percentage(25, 200)
	fmt.Printf("Percentage(25 of 200): %.2f%%\n", pct) // => 12.50%

	// clamp is unexported — internal helper used only inside mathutils
	// mathutils.clamp() from here => compile error
}
