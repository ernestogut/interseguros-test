const express = require('express');
const router = express.Router();
const matrixController = require('../controllers/MatrixController');

/**
 * @swagger
 * /api/stats:
 *   post:
 *     summary: Procesa matrices y devuelve estadísticas
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               q:
 *                 type: array
 *                 items:
 *                   type: array
 *                   items:
 *                     type: number
 *               r:
 *                 type: array
 *                 items:
 *                   type: array
 *                   items:
 *                     type: number
 *     responses:
 *       200:
 *         description: Estadísticas calculadas
 */
router.post('/stats', matrixController.calculateStats);

module.exports = router;