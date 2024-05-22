const { Note } = require("../models/models");

class NoteController {
  async addNote(req, res) {
    const { title, body, groupId } = req.body;
    const note = await Note.create({ title, body, groupId });
    return res.json(note);
  }

  async getNote(req, res) {
    const note = await Note.findAll();
    return res.json(note);
  }
  async deleteNote(req, res) {
    const note = await Note.destroy({
      where: {
        id: req.params.id,
      },
    });
    return res.json(note);
  }

  async putNote(req, res) {
    const note = Note.update(
      {
        title: req.body.title,
        body: req.body.body,
      },
      { where: { id: req.params.id } }
    );
  }
}

module.exports = new NoteController();
