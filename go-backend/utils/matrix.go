package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-backend/types"
	"gonum.org/v1/gonum/mat"
	"io"
	"net/http"
	"os"
)

func RotateMatrix(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])
	rotated := make([][]float64, cols)

	for i := 0; i < cols; i++ {
		rotated[i] = make([]float64, rows)
		for j := 0; j < rows; j++ {
			rotated[i][j] = matrix[rows-j-1][i]
		}
	}

	return rotated
}

func PostMatrixAnalysis(payload types.QrPayload, token string) (*types.AnalyzeMatrixResponse, error) {
	payloadJSON, err := json.Marshal(payload)

	if err != nil {
		return nil, fmt.Errorf("error serializando JSON: %w", err)
	}

	req, err := http.NewRequest("POST", os.Getenv("EXPRESS_URL")+"/api/matrix/analyze", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, fmt.Errorf("error al crear solicitud HTTP: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en POST HTTP: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta HTTP: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("respuesta HTTP con código %d: %s", resp.StatusCode, string(body))
	}

	var result types.AnalyzeMatrixResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error deserializando JSON: %w", err)
	}

	return &result, nil
}

func ValidateMatrix(matrix [][]float64) error {
	numMatrix := len(matrix)
	if numMatrix == 0 {
		return errors.New("la matriz no puede estar vacía")
	}

	numCols := len(matrix[0])
	if numCols == 0 {
		return errors.New("la matriz no puede tener columnas vacías")
	}

	for i, row := range matrix {
		if len(row) != numCols {
			return fmt.Errorf("fila %d tiene longitud %d diferente a %d", i, len(row), numCols)
		}
	}

	if numMatrix < numCols {
		return fmt.Errorf("la matriz debe tener filas >= columnas, filas=%d, columnas=%d", numMatrix, numCols)
	}

	return nil
}

func MatrixToSlice(m *mat.Dense) [][]float64 {
	rows, cols := m.Dims()
	slice := make([][]float64, rows)
	for i := range slice {
		slice[i] = make([]float64, cols)
		for j := range slice[i] {
			slice[i][j] = m.At(i, j)
		}
	}
	return slice
}
