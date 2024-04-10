const { Note } = require("../models/note");

class NoteController {
  async addGroup(req, res) {
    const { title, body } = req.body;
    const group = await Group.create({ title, body });
    return res.json(group);
  }

  async getGroup(req, res) {
    const note = await Note.findAll();
    return res.json(note);
  }
  async deleteGroup(req, res) {
    const note = await Note.destroy({
      where: {
        id: req.params.id,
      },
    });
    return res.json(note);
  }
}

module.exports = new NoteController();
