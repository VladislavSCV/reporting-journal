const Router = require("express");

const router = new Router();

const studentController = require("../controllers/studentController");

router.post("/", studentController.addStudent);
router.delete("/:id", studentController.deleteStudent);
router.get("/", studentController.getStudent);

module.exports = router;
