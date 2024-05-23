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
router.get("/auth", authMiddleware, authController.auth);

router.get("/user", authController.getUsers);
router.get("/user", authController.getUser);
router.delete("/user/:id", authController.deleteUser);

router.post("/role", authController.postRole);
router.get("/role", authController.getRoles);
router.delete("/role/:id", authController.deleteRole);

module.exports = router;
