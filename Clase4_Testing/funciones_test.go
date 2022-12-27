package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestCalcularSalario_Caso1(t *testing.T){
	Resultado := CalcularSalario(20000)
	var ResultadoEsperado float64 = 20000

	assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
}

func TestCalcularSalario_Caso2(t *testing.T){
	Resultado := CalcularSalario(55000)
	var ResultadoEsperado float64 = 45650

	assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
}

func TestCalcularSalario_Caso3(t *testing.T){
	Resultado := CalcularSalario(160000)
	var ResultadoEsperado float64 = 116800

	assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
}

