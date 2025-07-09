const { computeMatrixStats } = require('../src/services/MatrixService');

describe('MatrixService', () => {
  describe('computeMatrixStats', () => {
    it('should calculate correct statistics for valid matrices', () => {
      const q = [
        [1, 0],
        [0, 1]
      ];
      const r = [
        [2, 3],
        [4, 5]
      ];

      const result = computeMatrixStats(q, r);

      expect(result).toEqual({
        max: 5,
        min: 0,
        average: 2,
        sum: 16,
        isQDiagonal: true,
        isRDiagonal: false
      });
    });

    it('should throw an error if matrices are empty', () => {
      expect(() => computeMatrixStats([], [])).toThrow('Matrices cannot be empty');
    });

    it('should handle matrices with negative numbers', () => {
      const q = [
        [-1, 0],
        [0, -2]
      ];
      const r = [
        [3, 4],
        [5, 6]
      ];

      const result = computeMatrixStats(q, r);

      expect(result).toEqual({
        max: 6,
        min: -2,
        average: 1.875,
        sum: 15,
        isQDiagonal: true,
        isRDiagonal: false
      });
    });
  });
});
