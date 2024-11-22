package composite

import "testing"

func TestAthlete_Train(t *testing.T) {
	athlete := Athlete{}
	athlete.Train()
}

func TestSwimmer_Swim(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: &localSwim,
	}
	swimmer.MyAthlete.Train()
	(*swimmer.MySwim)()
}
