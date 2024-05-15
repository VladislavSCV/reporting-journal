const { User, Role } = require("../models/models");
const bcrypt = require("bcryptjs");
const { validationResult } = require("express-validator");
const jwt = require("jsonwebtoken");

const { SECRET_KEY } = process.env;

const generateAccessToken = (id, role) => {
  const payload = {
    id,
    role,
  };

  return jwt.sign(payload, SECRET_KEY, { expiresIn: "24h" });
};

class authController {
  async registration(req, res) {
    try {
      const errors = validationResult(req);
      if (!errors.isEmpty()) {
        return res.status(400).json({ message: "Некорректные данные" });
      }
      const { login, password } = req.body;

      const candidate = await User.findOne({ where: { login } });
      if (candidate) {
        return res
          .status(400)
          .json({ message: "Такой пользователь уже существует" });
      }
      const hashPassword = bcrypt.hashSync(password, 7);
      const userRole = await Role.findOne({ where: { value: "user" } });

      const user = await User.create({
        name: "Гилоян Роман",
        login,
        password: hashPassword,
        role: userRole.value,
        group: "21ис3-4д",
      });

      await user.save();
      return res.json("Пользователь создан");
    } catch (e) {
      console.error(e);
      res.status(500).json({ error: "Произошла ошибка при регистрации" });
    }
  }
  async login(req, res) {
    try {
      const { login, password } = req.body;
      const user = await User.findOne({ where: { login } });
      if (!user) {
        return res
          .status(400)
          .json({ message: `Пользователь ${login} не найден` });
      }
      const validPassword = bcrypt.compareSync(password, user.password);
      if (!validPassword) {
        return res.status(400).json({ message: "Неверный пароль" });
      }
      const token = generateAccessToken(user.id, user.role);
      return res.json({ token });
    } catch (e) {}
  }
  async getUsers(req, res) {
    try {
      const adminRole = new Role({ value: "admin" });

     
      await adminRole.save();
      res.json({ adminRole });
      // const users = await User.findAll();
      // res.json(users);
    } catch (e) {
      console.error(e);
      res.status(500).json({ error: "Произошла ошибка при создании роли" });
    }
  }
}

module.exports = new authController();
