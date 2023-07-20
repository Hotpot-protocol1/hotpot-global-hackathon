const express = require('express');
const app = express();
const port = process.env.PORT || 5000; // Use environment variable or default

app.get('/', (req, res) => {
  res.send('Hello, world!');
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});