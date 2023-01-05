package main

func CalcularSalarioXCategoria(minutostrabajados int, categoria string)float64{
	var horas = minutostrabajados / 60
	switch categoria{
	case "A":
		return 3000 * float64(horas) + (1500)
	case "B":
		return 1500 * float64(horas) + (300)
	case "C":
		return 1000 * float64(horas)
	}
	return 0
}