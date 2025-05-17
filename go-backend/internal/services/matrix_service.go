package services

import (
	"fmt"
	"go-backend/types"
	"go-backend/utils"
	"gonum.org/v1/gonum/mat"
)

func ProcessMatrix(matrix [][]float64, token string) (*types.AnalyzeMatrixResponse, error) {
	if err := utils.ValidateMatrix(matrix); err != nil {
		return nil, err
	}

	qData, rData, err := getQR(matrix)

	rotated := utils.RotateMatrix(matrix)

	if err != nil {
		return nil, fmt.Errorf("error al factorizar la matriz QR: %w", err)
	}

	payload := types.QrPayload{
		Q:       qData,
		R:       rData,
		Rotated: rotated,
	}

	response, err := utils.PostMatrixAnalysis(payload, token)
	if err != nil {
		return nil, fmt.Errorf("error en petición POST al API: %w", err)
	}

	return response, nil
}

func getQR(matrix [][]float64) ([][]float64, [][]float64, error) {
	rows := len(matrix)
	cols := len(matrix[0])

	flat := make([]float64, rows*cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			flat[row*cols+col] = matrix[row][col]
		}
	}

	matDense := mat.NewDense(rows, cols, flat)

	var qr mat.QR
	qr.Factorize(matDense)

	var qFull mat.Dense
	var rFull mat.Dense

	qr.QTo(&qFull)
	qr.RTo(&rFull)

	qReduced := qFull.Slice(0, rows, 0, cols).(*mat.Dense)
	rReduced := rFull.Slice(0, cols, 0, cols).(*mat.Dense)

	return utils.MatrixToSlice(qReduced), utils.MatrixToSlice(rReduced), nil
}
