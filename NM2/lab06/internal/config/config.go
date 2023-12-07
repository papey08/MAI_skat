package config

type Config struct {
	XBegin float64
	XEnd   float64

	TBegin float64
	TEnd   float64

	H     float64
	Sigma float64

	A float64

	Phi0 func(t, a float64) float64
	Phi1 func(t, a float64) float64

	Psi0 func(x, a float64) float64
	Psi1 func(x, a float64) float64
}
