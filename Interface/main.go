package main 
import "fmt"
type Stringer interface {
    String() string
}

type Human struct {
    Firstname string
    Lastname string
    Age int
    Country string
}

func (h Human) String() string {
    return fmt.Sprintf("human named %s %s of age %d living in %s",h.Firstname,h.Lastname,h.Age,h.Country)
}

func main() {
    var john Human
    john.Firstname = "John"
    john.Lastname = "Doe"
    john.Country = "USA"
    john.Age = 45

    fmt.Println(john)
}