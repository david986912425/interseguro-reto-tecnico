package types

type MatrixStats struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Sum        float64 `json:"sum"`
	Average    float64 `json:"average"`
	IsDiagonal bool    `json:"isDiagonal"`
}

type AnalyzeMatrixResponse struct {
	RotatedMatrix [][]float64 `json:"rotatedMatrix"`
	QStats        MatrixStats `json:"qStats"`
	RStats        MatrixStats `json:"rStats"`
	QMatrix       [][]float64 `json:"qMatrix"`
	RMatrix       [][]float64 `json:"rMatrix"`
}

type QrPayload struct {
	Q       [][]float64 `json:"Q"`
	R       [][]float64 `json:"R"`
	Rotated [][]float64 `json:"rotated"`
}
