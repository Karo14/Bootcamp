package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcularPromedio_Caso1(t *testing.T){
	Notas := []float32{4.5,4.6,3.5,4.0,1.0}
	Resultado := CalcularPromedio(Notas)
	var ResultadoEsperado float32 = 3.52

	assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
}