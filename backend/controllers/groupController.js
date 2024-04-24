const { Group } = require("../models/group");
const { Note } = require("../models/note");

class GroupController {
  async addGroup(req, res) {
    const { name } = req.body;
    const group = await Group.create({ name });
    return res.json(group);
  }

  async getGroup(req, res) {
    const group = await Group.findAll();
    return res.json(group);
  }
  async deleteGroup(req, res) {
    const group = await Group.destroy({
      where: {
        id: req.params.id,
      },
    });
    return res.json(group);
  }
}

module.exports = new GroupController();
