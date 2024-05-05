const Router = require("express");

const router = new Router();

const noteRouter = require("./noteRouter");
const groupRouter = require("./groupRouter");
const studentRouter = require("./studentRouter");
const authRouter = require("./authRouter");

router.use("/notes", noteRouter);
router.use("/groups", groupRouter);
router.use("/students", studentRouter);
router.use("/auth", authRouter);

module.exports = router;
