const Router = require("express");

const router = new Router();

const noteController = require("../controllers/noteController");

router.post("/", noteController.addNote);
router.delete("/:id", noteController.deleteNote);
router.get("/", noteController.getNote);

module.exports = router;
