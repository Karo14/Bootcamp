package hunt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSharkHuntsSuccessfully(t *testing.T) {

	s := Shark{
		hungry: false,
		tired: true,
		speed: 100,
	}

	p := Prey{
		name: "Agria",
		speed: 90,
	}

	err := s.Hunt(&p)

	assert.NoError(t,err)
	assert.Equal(t,true,s.tired)
	assert.Equal(t,false,s.hungry)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	ErrorEsperado := errors.New("cannot hunt, i am really tired")

	s := Shark{
		hungry: false,
		tired: true,
		speed: 90,
	}

	p := Prey{
		name: "Jessi",
		speed: 100,
	}

	err := s.Hunt(&p)

	assert.EqualError(t, ErrorEsperado, err.Error())
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	ErrorEsperado := errors.New("cannot hunt, i am not hungry")

	s := Shark{
		hungry: false,
		tired: false,
		speed: 100,
	}

	p := Prey{
		name: "Cami",
		speed: 90,
	}

	err := s.Hunt(&p)

	assert.EqualError(t, ErrorEsperado, err.Error())
}

func TestSharkCannotReachThePrey(t *testing.T) {
	ErrorEsperado := errors.New("could not catch it")

	s := Shark{
		hungry: true,
		tired: false,
		speed: 90,
	}

	p := Prey{
		name: "Angie",
		speed: 100,
	}

	err := s.Hunt(&p)

	assert.EqualError(t, ErrorEsperado, err.Error())
}

func TestSharkHuntNilPrey(t *testing.T) {
	ErrorEsperado := errors.New("cannot hunt, prey nil")

	s := Shark{
		hungry: true,
		tired: false,
		speed: 90,
	}

	err := s.Hunt(nil)

	assert.EqualError(t, ErrorEsperado, err.Error())
}
