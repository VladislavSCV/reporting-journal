const { Student } = require("../models/student");

class StudentController {
  async addStudent(req, res) {
    const { name, role } = req.body;
    const student = await Student.create({ name, role });
    return res.json(student);
  }

  async getStudent(req, res) {
    const student = await Student.findAll();
    return res.json(student);
  }
  async deleteStudent(req, res) {
    const student = await Student.destroy({
      where: {
        id: req.params.id,
      },
    });
    return res.json(student);
  }
}

module.exports = new StudentController();
