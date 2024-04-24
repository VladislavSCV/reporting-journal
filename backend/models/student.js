const { Sequelize, DataTypes } = require("sequelize");

const sequelize = require("../db");

const Student = sequelize.define("student", {
  id: { type: DataTypes.INTEGER, primaryKey: true, autoIncrement: true },
  name: { type: DataTypes.STRING },
  role: { type: DataTypes.STRING },
  group: { type: DataTypes.STRING },
});

module.exports = { Student };
