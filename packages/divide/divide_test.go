package divide
import "testing"

func TestDivide(t *testing.T){
	got,_:= Divide(10,2)
	want :=5
	if got != want {
        t.Errorf("Divide(10, 5) = %d; want %d", got, want)
    }
}