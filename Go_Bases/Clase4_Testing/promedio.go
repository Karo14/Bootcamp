package main

func CalcularPromedio(notas []float32) (promedio float32){
	var SumaTotal float32 = 0
	for _, num := range notas{
		if(num < 0){
			return 0
		}
		SumaTotal += num
	}
	promedio = float32(SumaTotal) / float32(len(notas))
	return promedio
}