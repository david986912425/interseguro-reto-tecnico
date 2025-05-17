type Matrix = number[][];

interface MatrixStats {
    max: number;
    min: number;
    sum: number;
    average: number;
    isDiagonal: boolean;
}

interface AnalyzeMatrixResult {
    rotatedMatrix: Matrix;
    qStats: MatrixStats;
    rStats: MatrixStats;
    qMatrix: Matrix;
    rMatrix: Matrix;
}

export function analyzeMatrices(Q: Matrix, R: Matrix, rotated: Matrix): AnalyzeMatrixResult {
    return {
        rotatedMatrix: rotated,
        qStats: getMatrixStats(Q),
        rStats: getMatrixStats(R),
        qMatrix: Q,
        rMatrix: R
    };
}

function getMatrixStats(matrix: Matrix): MatrixStats {
    const { max, min, sum, count } = getBasicStats(matrix);
    return {
        max,
        min,
        sum,
        average: count > 0 ? sum / count : 0,
        isDiagonal: checkIfDiagonal(matrix),
    };
}

function getBasicStats(matrix: Matrix): { max: number; min: number; sum: number; count: number } {
    let max = Number.NEGATIVE_INFINITY;
    let min = Number.POSITIVE_INFINITY;
    let sum = 0;
    let count = 0;

    for (const row of matrix) {
        for (const val of row) {
            max = Math.max(max, val);
            min = Math.min(min, val);
            sum += val;
            count++;
        }
    }
    return { max, min, sum, count };
}

function checkIfDiagonal(matrix: Matrix): boolean {
    const rowCount = matrix.length;
    const colCount = matrix[0].length;

    for (let i = 0; i < rowCount; i++) {
        for (let j = 0; j < colCount; j++) {
            if (i !== j && matrix[i][j] !== 0) {
                return false;
            }
        }
    }
    return true;
}
