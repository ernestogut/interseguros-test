const matrixService = require('../services/MatrixService');

/**
 * Calcula estadísticas de matrices Q y R
 * @param {Object} req - Request de Express
 * @param {Object} res - Response de Express
 * @param {Function} next - Next middleware
 */
exports.calculateStats = async (req, res, next) => {
  try {
    const { q, r } = req.body;
    
    if (!q || !r) {
      return res.status(400).json({ error: 'Q and R matrices are required' });
    }

    const stats = await matrixService.computeMatrixStats(q, r);
    res.json(stats);
  } catch (error) {
    next(error);
  }
};