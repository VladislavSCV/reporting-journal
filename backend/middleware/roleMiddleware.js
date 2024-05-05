const jwt = require("jsonwebtoken");

module.exports = function (roles) {
  return function (req, res, next) {
    if (req.method === "OPTIONS") {
      next();
    }

    try {
      const token = req.headers.authorization.split(" ")[1];
      if (!token) {
        return res
          .status(403)
          .json({ message: "Пользователь212 не авторизован" });
      }

      const { role: userRoles } = jwt.verify(token, process.env.SECRET_KEY);
      console.log(roles);
      console.log(userRoles);

      let hasRole = false;
      if (userRoles == "admin") {
        hasRole = true;
      }

      if (!hasRole) {
        return res.status(403).json({ message: "У вас нет доступа" });
      }
      next();
    } catch (e) {
      console.log(e);
      return res.status(403).json({ message: "Пользовател1ь не авторизован" });
    }
  };
};
