const Router = require("express");

const router = new Router();

const noteController = require("../controllers/noteController");

router.post("/", noteController.addNote);
router.delete("/:id", noteController.deleteNote);
router.get("/", noteController.getNote);
router.put("/:id", noteController.putNote);

module.exports = router;
