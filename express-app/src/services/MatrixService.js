/**
 * Calcula estadísticas de matrices
 * @param {Array} q - Matriz Q
 * @param {Array} r - Matriz R
 * @returns {Object} Estadísticas calculadas
 */
exports.computeMatrixStats = (q, r) => {
  const allValues = [...flattenMatrix(q), ...flattenMatrix(r)];
  
  if (allValues.length === 0) {
    throw new Error('Matrices cannot be empty');
  }

  return {
    max: Math.max(...allValues),
    min: Math.min(...allValues),
    average: calculateAverage(allValues),
    sum: calculateSum(allValues),
    isQDiagonal: isDiagonalMatrix(q),
    isRDiagonal: isDiagonalMatrix(r),
  };
};

// Helper: Aplana matriz 2D a array 1D
const flattenMatrix = (matrix) => matrix.flat();

// Helper: Calcula promedio
const calculateAverage = (values) => {
  const sum = values.reduce((acc, val) => acc + val, 0);
  return sum / values.length;
};

// Helper: Calcula suma total
const calculateSum = (values) => values.reduce((acc, val) => acc + val, 0);

// Helper: Verifica si matriz es diagonal
const isDiagonalMatrix = (matrix) => {
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      if (i !== j && matrix[i][j] !== 0) {
        return false;
      }
    }
  }
  return true;
};