package mod
import "testing"
func TestMod(t *testing.T){
	got,_:=Mod(100,0)
	want :=-1
	if got != want {
        t.Errorf("Mod(100, -1) = %d; want %d", got, want)
    }
}