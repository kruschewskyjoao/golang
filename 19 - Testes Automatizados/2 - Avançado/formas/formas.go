package formas

import (
	"math"
)

// so tem assinatura de metodos.
// n pode passar campos
type Forma interface {
	area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {
	return math.Pow(c.raio, 2) * math.Pi
}
