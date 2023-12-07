package config

type Config struct {
	XBegin float64
	XEnd   float64

	YBegin float64
	YEnd   float64

	HX float64
	HY float64

	Phi0 func(y float64) float64
	Phi1 func(y float64) float64

	Psi0 func(x float64) float64
	Psi1 func(x float64) float64
}
