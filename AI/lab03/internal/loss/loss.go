package loss

type Loss int

const (
	None Loss = iota
	CrossEntropy
	BinCrossEntropy
	MeanSquared
)

func GetLoss(loss Loss) (f func(e, i [][]float64) float64, df func(e, i, a float64) float64) {
	switch loss {
	case CrossEntropy:
		return CrossEntropyF, CrossEntropyDf
	case MeanSquared:
		return MeanSquaredF, MeanSquaredDf
	case BinCrossEntropy:
		return BinCrossEntropyF, BinCrossEntropyDf
	default:
		return CrossEntropyF, CrossEntropyDf
	}
}

func (l Loss) String() string {
	switch l {
	case CrossEntropy:
		return "CE"
	case BinCrossEntropy:
		return "BinCE"
	case MeanSquared:
		return "MSE"
	}
	return "N/A"
}
