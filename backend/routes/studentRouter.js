const Router = require("express");

const router = new Router();

const studentController = require("../controllers/studentController");

router.post("/", studentController.addStudent);
router.delete("/:id", studentController.deleteStudent);
router.get("/", studentController.getStudent);
router.put("/:id", studentController.putStudent);
module.exports = router;
