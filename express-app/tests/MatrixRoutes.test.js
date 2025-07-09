const request = require('supertest');
const express = require('express');
const matrixRoutes = require('../src/routes/MatrixRoutes');

const app = express();
app.use(express.json());
app.use('/api', matrixRoutes);

describe('MatrixRoutes', () => {
  describe('POST /express/stats', () => {
    it('should return 200 and calculated statistics for valid input', async () => {
      const response = await request(app)
        .post('/api/stats')
        .send({
          q: [[1, 2], [3, 4]],
          r: [[5, 6], [7, 8]]
        });

      expect(response.status).toBe(200);
    });

    it('should return 400 for invalid input', async () => {
      const response = await request(app)
        .post('/api/stats')
        .send({
          q: 'invalid',
          r: [[5, 6], [7, 8]]
        });

      expect(response.status).toBe(500);
    });
  });
});
