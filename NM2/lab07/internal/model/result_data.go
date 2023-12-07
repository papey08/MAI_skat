package model

type ResultData struct {
	A       [][]float64 `json:"A"`
	B       []float64   `json:"b"`
	Mapping [][]int     `json:"mapping"`
	Res     [][]float64 `json:"res"`
	X       []float64   `json:"x"`
	Y       []float64   `json:"y"`
}
