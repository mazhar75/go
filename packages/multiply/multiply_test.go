package multiply
import "testing"

func TestMultiply(t *testing.T){
	got :=Multiply(10,5)
	want:=50
	if got != want {
        t.Errorf("Multiply(10, 5) = %d; want %d", got, want)
    }
}