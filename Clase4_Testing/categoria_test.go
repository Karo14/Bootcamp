package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcularSalarioXCategoria_Caso1(t *testing.T){
	Resultado := CalcularSalarioXCategoria(6000,"A")
	var ResultadoEsperado float64 = 301500

	assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
}

func TestCalcularSalarioXCategoria_Caso2(t *testing.T){
	Resultado := CalcularSalarioXCategoria(600,"B")
	var ResultadoEsperado float64 = 15300

	assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
}

func TestCalcularSalarioXCategoria_Caso3(t *testing.T){
	Resultado := CalcularSalarioXCategoria(7800,"C")
	var ResultadoEsperado float64 = 130000

	assert.Equal(t, ResultadoEsperado, Resultado, "el valor resultado y resultado esperado deberian ser iguales")
}