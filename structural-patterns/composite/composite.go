package composite

type Athlete struct{}

func (*Athlete) Train() {}

func Swim() {}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}
