const { Note } = require("../models/models");

class NoteController {
  async addNote(req, res) {
    const { title, body } = req.body;
    const note = await Note.create({ title, body });
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
    const note = Student.update(
      {
        title: req.body.title,
        body: req.body.body,
      },
      { where: { id: req.params.id } }
    );
  }
}

module.exports = new NoteController();
