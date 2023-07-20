const express = require('express');
const app = express();
const port = process.env.PORT || 5000; // Use environment variable or default
var router = express.Router();
const { User } = require('./models')


const db = require("./models");
db.sequelize.sync()
  .then(() => {
    console.log("Synced db.");
  })
  .catch((err) => {
    console.log("Failed to sync db: " + err.message);
  });



app.get('/', (req, res) => {
  console.log("GET /");
  res.send('Hello, world!');
});

app.get('/users/:id', async (req, res) => {
  console.log("GET /users/:id");
  const id = req.params.id;
  try {
    User.findByPk(id)
    .then(data => {
      if (data) {
        //log data
        console.log(data);

        res.send(data);
      } else {
        res.status(404).send({
          message: 'cant find that ish'
        })
      }
    })
  } catch (error) {
    res.status(500).send(error.message);
  }
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});

module.exports = router;
