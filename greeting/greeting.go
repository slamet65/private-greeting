// Package greeting menyediakan contoh API kecil untuk demo versioning module.
package greeting

import "fmt"

// Version menunjukkan versi perilaku library pada commit saat ini.
const Version = "v2.0.0"

// Hello menghasilkan sapaan dari versi pertama library.
func Hello(name string) string {
	return fmt.Sprintf("Halo, %s! Salam dari private-greeting %s.", name, Version)
}
