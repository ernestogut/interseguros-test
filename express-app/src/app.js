const express = require('express');
const cors = require('cors');
const morgan = require('morgan');
const errorHandler = require('./middlewares/ErrorHandler');
const routes = require('./routes/MatrixRoutes');
const authenticateJWT = require('./middlewares/AuthMiddleware');
require('dotenv').config();

const app = express();

// Middlewares globales
app.use(cors());
app.use(express.json());
app.use(morgan('dev')); // Logging de requests


app.use('/api', authenticateJWT, routes);

// Manejador de errores centralizado
app.use(errorHandler);

module.exports = app;