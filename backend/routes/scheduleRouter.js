const Router = require("express");

const router = new Router();

const scheduleController = require("../controllers/scheduleController");

router.get("/", scheduleController.getSchedule);
router.post("/", scheduleController.addSchedule);
router.delete("/:id", scheduleController.deleteSchedule);
router.put("/:id", scheduleController.putSchedule);
module.exports = router;
