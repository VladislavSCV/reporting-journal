const Router = require("express");
const { check } = require("express-validator");
const router = new Router();

const authController = require("../controllers/authController");
const authMiddleware = require("../middleware/authMiddleware");
const roleMiddleware = require("../middleware/roleMiddleware");

router.post(
  "/registration",
  [
    check("login").isLength({ min: 3, max: 10 }),
    check("password").isLength({ min: 3, max: 15 }),
  ],
  authController.registration
);
router.post("/login", authController.login);
router.get(
  "/users",
 
  authController.getUsers
);

module.exports = router;
