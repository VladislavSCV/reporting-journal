require("dotenv").config();
const express = require("express");
const sequelize = require("./db");
const bodyParser = require("body-parser");
const cors = require("cors");
const coockieParser = require("cookie-parser");
const router = require("./routes/index");
const corsMiddleware = require("./middleware/cors.middleware");
const PORT = process.env.PORT || 5002;

const app = express();
app.use(corsMiddleware);
app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.json());
app.use("/api", router);

const start = async () => {
  try {
    await sequelize.authenticate();
    await sequelize.sync();
    app.listen(PORT, () => console.log(`Server started on port ${PORT}`));
  } catch (e) {
    console.log(e);
  }
};

start();
