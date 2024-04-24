const Router = require("express");

const router = new Router();

const noteRouter = require("./noteRouter");
const groupRouter = require("./groupRouter");
const studentRouter = require("./studentRouter");

router.use("/notes", noteRouter);
router.use("/groups", groupRouter);
router.use("/students", studentRouter);

module.exports = router;
