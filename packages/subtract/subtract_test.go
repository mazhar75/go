package subtract
import "testing"
func TestSubtract(t *testing.T){
	got := Subtract(2,5)
	want := -3
	if got != want {
        t.Errorf("Subtract(2, 5) = %d; want %d", got, want)
    }
}
