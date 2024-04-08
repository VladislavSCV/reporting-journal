const Router = require("express");

const router = new Router();

const noteRouter = require("./noteRouter");

router.use("/notes", noteRouter);

module.exports = router;
