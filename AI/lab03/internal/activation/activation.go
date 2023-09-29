package activation

type Mode int

const (
	Default Mode = iota
	MultiClass
	Regression
	Binary
	MultiLabel
)

func OutputActivation(c Mode) Activation {
	switch c {
	case MultiClass:
		return Softmax
	case Regression:
		return Linear
	case Binary, MultiLabel:
		return Sigmoid
	default:
		return None
	}
}

func GetActivation(act Activation) (f, df func(x float64) float64) {
	switch act {
	case Sigmoid:
		return SigmoidF, SigmoidDf
	case Tanh:
		return TanhF, TanhDf
	case ReLU:
		return ReLUF, ReLUDf
	case Linear:
		return LinearF, LinearDf
	default:
		return LinearF, LinearDf
	}
}

type Activation int

const (
	None Activation = iota
	Sigmoid
	Tanh
	ReLU
	Linear
	Softmax
)
