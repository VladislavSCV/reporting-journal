import React from "react";
import styles from './ModalScheduleAdd.module.scss'
import Modal from "../../modal/Modal";
import addLesson from "../../../../js/addLesson";
const ModalStudentAdd = ({ closeFn = () => null, open = false }) => {
  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
          <header className="modal__header">
            <h1 className="modal__title">Добавление студента</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <label htmlFor="lesson" className={styles.modal__label}>
              Выберите предмет:
            </label>
            <div className={styles.modal__customSelect}>
              <select  id="chooseLesson">
                <option value="0">Выберите предмет</option>
                <option value="1">Граф. дизайн</option>
                <option value="2">Программирование</option>
                <option value="3">Тестирование</option>
                <option value="4">Веб-разработка</option>
                <option value="5">???</option>
                <option value="6">хз</option>
              </select>
            </div>
            <label htmlFor="lessonTeacher" className={styles.modal__label}>
              Выберите имя преподавателя:
            </label>
            <div className={styles.modal__customSelect}>
              <select  id="chooseTeacher">
                <option value="0">Выберите преподавателя</option>
                <option value="1">Ньютон</option>
                <option value="2">Стэтхем</option>
                <option value="3">Дуейн Джонсон</option>
                <option value="4">Сэм Сулек</option>
                <option value="5">???</option>
                <option value="6">хз</option>
              </select>
            </div>
            <button className={styles.modal__button} onClick={addLesson}>Добавить</button>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentAdd;
