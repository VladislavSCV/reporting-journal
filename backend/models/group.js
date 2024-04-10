const { Sequelize, DataTypes } = require("sequelize");

const sequelize = require("../db");

const Group = sequelize.define("group", {
  id: { type: DataTypes.INTEGER, primaryKey: true, autoIncrement: true },
  title: { type: DataTypes.STRING },
  body: { type: DataTypes.STRING },
});

module.exports = { Group };
