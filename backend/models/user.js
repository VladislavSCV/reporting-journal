const { Sequelize, DataTypes } = require("sequelize");

const sequelize = require("../db");

const User = sequelize.define("user", {
  id: { type: DataTypes.INTEGER, primaryKey: true, autoIncrement: true },
  name: { type: DataTypes.STRING },
  role: { type: DataTypes.STRING },
  group: { type: DataTypes.STRING },
  login: { type: DataTypes.STRING },
  password: { type: DataTypes.STRING },
});

module.exports = { User };
