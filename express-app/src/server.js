const app = require('./app');
const config = require('./config');

app.listen(config.PORT, () => {
  console.log(`Matrix API running on port ${config.PORT}`);
});