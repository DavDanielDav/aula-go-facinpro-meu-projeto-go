package calculoimc

func CalculoIMC(peso, altura float64) float64 {
	if altura == 0 {
		return 0
	}
	return peso / (altura * altura)
}
