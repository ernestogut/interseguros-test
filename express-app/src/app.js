const express = require('express');
const cors = require('cors');
const morgan = require('morgan');
const errorHandler = require('./middlewares/ErrorHandler');
const routes = require('./routes/MatrixRoutes');
const authenticateJWT = require('./middlewares/AuthMiddleware');
require('dotenv').config();

const app = express();

app.use(cors());
app.use(express.json());
app.use(morgan('dev'));

app.get('/express/health', (req, res) => {
  res.status(200).send('OK');
});

app.use('/express', authenticateJWT, routes);

app.use(errorHandler);

module.exports = app;