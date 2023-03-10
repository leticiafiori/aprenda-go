package iteracao

func Repetir(caractere string, qtdeRepeticoes int) string {
	var repeticoes string
	for i := 0; i < qtdeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
