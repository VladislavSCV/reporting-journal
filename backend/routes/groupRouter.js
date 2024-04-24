const Router = require("express");

const router = new Router();

const groupController = require("../controllers/groupController");

router.post("/", groupController.addGroup);
router.delete("/:id", groupController.deleteGroup);
router.get("/", groupController.getGroup);

module.exports = router;
