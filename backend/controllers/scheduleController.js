const { Schedule } = require("../models/models");
const { Sequelize } = require("sequelize");
class ScheduleController {
  async getSchedule(req, res) {
    const schedule = await Schedule.findAll({
      order: [
        [
          Sequelize.literal(`
          CASE
            WHEN "dayOfWeek" = 'Monday' THEN 1
            WHEN "dayOfWeek" = 'Tuesday' THEN 2
            WHEN "dayOfWeek" = 'Wednesday' THEN 3
            WHEN "dayOfWeek" = 'Thursday' THEN 4
            WHEN "dayOfWeek" = 'Friday' THEN 5
            WHEN "dayOfWeek" = 'Saturday' THEN 6
            WHEN "dayOfWeek" = 'Sunday' THEN 7
          END
      `),
        ],
      ],
    });

    return res.json(schedule);
  }

  async addSchedule(req, res) {
    const { subject, teacher, dayOfWeek, groupId } = req.body;
    const schedule = await Schedule.create({
      subject,
      teacher,
      dayOfWeek,
      groupId,
    });
    return res.json(schedule);
  }

  async deleteSchedule(req, res) {
    const schedule = await Schedule.destroy({
      where: { id: req.params.id },
    });
    return res.json(schedule);
  }

  async putSchedule(req, res) {
    const { subject, teacher } = req.body;
    const schedule = await Schedule.update(
      {
        subject,
        teacher,
      },
      {
        where: { id: req.params.id },
      }
    );
    return res.json(schedule);
  }
}

module.exports = new ScheduleController();
