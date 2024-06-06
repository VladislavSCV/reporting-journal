import React, { useState } from "react";
import styles from "./ModalScheduleSettings.module.scss";
import Modal from "../../modal/Modal";
import axios from "axios";

const ModalStudentAdd = ({ closeFn = () => null, open = false, id }) => {
  const [subject, setSubject] = useState("");
  const [teacher, setTeacher] = useState("");
  const putSchedule = async (key) => {
    try {
      await axios.put(`http://localhost:5001/api/schedule/${key}`, {
        subject,
        teacher,
      });
    } catch (error) {
      console.error(error);
    }
  };
  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
          <header className="modal__header">
            <h1 className="modal__title">Изменить предмет</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <label htmlFor="lesson" className={styles.modal__label}>
              Выберите предмет:
            </label>
            <div className={styles.modal__customSelect}>
              <select
                id="chooseLesson"
                className={styles.modal__select}
                onChange={(e) => setSubject(e.target.value)}
              >
                <option value="Выберите предмет">Выберите предмет</option>
                <option value="Граф. дизайн">Граф. дизайн</option>
                <option value="Программирование">Программирование</option>
                <option value="Тестирование">Тестирование</option>
                <option value="Веб-разработка">Веб-разработка</option>
                <option value="???">???</option>
                <option value="хз">хз</option>
              </select>
            </div>
            <label htmlFor="lessonTeacher" className={styles.modal__label}>
              Выберите имя преподавателя:
            </label>
            <div className={styles.modal__customSelect}>
              <select
                id="chooseTeacher"
                className={styles.modal__select}
                onChange={(e) => setTeacher(e.target.value)}
              >
                <option value="Выберите преподавателя">
                  Выберите преподавателя
                </option>
                <option value="Ньютон">Ньютон</option>
                <option value="Стэтхем">Стэтхем</option>
                <option value="Дуейн Джонсон">Дуейн Джонсон</option>
                <option value="Сэм Сулек">Сэм Сулек</option>
                <option value="???">???</option>
                <option value="хз">хз</option>
              </select>
            </div>
            <button
              className={styles.modal__button}
              onClick={() => putSchedule(id)}
            >
              Добавить
            </button>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentAdd;
