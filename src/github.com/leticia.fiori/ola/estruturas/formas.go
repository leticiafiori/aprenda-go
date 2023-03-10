package estruturas

import "math"

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

//Em Go a resolução de interface é implícita.
//Se o tipo que você passar combinar com o que a interface está esperando, o código será compilado.
//Como nesse caso circulos e retangulos tem o método Area() eles se encaixam na interface Forma

type Forma interface {
	Area() float64
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

//Esse é uma função

// func Area(retangulo Retangulo) float64 {
// 	return retangulo.Largura * retangulo.Altura
// }

//Esses são métodos

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (t Triangulo) Area() float64 {
	return (t.Altura * t.Base) / 2
}
