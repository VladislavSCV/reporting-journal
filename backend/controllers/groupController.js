const { Group } = require("../models/models");

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

  async putGroup(req, res) {
    const posts = Group.update(
      {
        name: req.body.name,
      },
      {
        where: {
          id: req.params.id,
        },
      }
    );
  }
}

module.exports = new GroupController();
