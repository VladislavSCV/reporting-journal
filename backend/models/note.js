const { Sequelize, DataTypes } = require("sequelize");

const sequelize = require("../db");

const Note = sequelize.define("note", {
  id: { type: DataTypes.INTEGER, primaryKey: true, autoIncrement: true },
  title: { type: DataTypes.STRING },
  body: { type: DataTypes.STRING },
});

module.exports = { Note };
