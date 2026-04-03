// test temconv.go
package tempconv

import "testing"

func TestTempconv(t *testing.T) {
	t.Run("test conversation from Fahrenheit to Celsius", func(t *testing.T) {
		got := CToF(100)
		want := Fahrenheit(76)
		if got != want {
			t.Errorf("Expected conversation of 100 Celsius is %g, received %g", want, got)
		}
	})
}
