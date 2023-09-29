package dataset

import (
	"ai_lab3/internal/training"
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func toPair(in []string) (training.Pair, error) {
	res, err := strconv.ParseFloat(in[0], 64)
	if err != nil {
		return training.Pair{}, err
	}
	resEncoded := make([]float64, 3)
	resEncoded[int(res)-1] = 1
	var features []float64
	for i := 1; i < len(in); i++ {
		res, err = strconv.ParseFloat(in[i], 64)
		if err != nil {
			return training.Pair{}, err
		}
		features = append(features, res)
	}
	return training.Pair{
		Response: resEncoded,
		Input:    features,
	}, nil
}

func Load(path string) (training.Pairs, error) {
	f, err := os.Open(path)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bufio.NewReader(f))
	var pairs training.Pairs
	for {
		record, readErr := r.Read()
		if readErr == io.EOF {
			break
		}
		p, pairErr := toPair(record)
		if pairErr != nil {
			return nil, pairErr
		}
		pairs = append(pairs, p)
	}
	return pairs, nil
}
