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

const Role = sequelize.define("role", {
  value: {
    type: DataTypes.STRING,
    unique: true,
    allowNull: false,
    defaultValue: "user",
  },
});

const Note = sequelize.define("note", {
  id: { type: DataTypes.INTEGER, primaryKey: true, autoIncrement: true },
  title: { type: DataTypes.STRING },
  body: { type: DataTypes.STRING },
  group: { type: DataTypes.STRING },
});

const Group = sequelize.define("group", {
  id: { type: DataTypes.INTEGER, primaryKey: true, autoIncrement: true },
  name: { type: DataTypes.STRING },
  body: { type: DataTypes.STRING },
});

User.belongsToMany(Role, { through: "userRoles" });
Role.belongsToMany(User, { through: "userRoles" });

module.exports = { User, Role, Note, Group };
