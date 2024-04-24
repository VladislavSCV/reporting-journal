import React, { useState } from "react";
import styles from "./ModalStudentAdd.module.scss";
import Modal from "../../modal/Modal";
import axios from "axios";
const ModalStudentAdd = ({ closeFn = () => null, open = false }) => {
  const [name, setName] = useState("");
  const [role, setRole] = useState("");

  const addStudent = async () => {
    try {
      await axios.post("http://localhost:5001/api/students", { name, role });
    } catch (error) {
      console.error(error);
    }
  };

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
            <form action="">
              <label htmlFor="studentname" className={styles.modal__label}>
                Введите ФИО студента:
              </label>
              <input
                type="text"
                className={styles.modal__input}
                placeholder="Иванов Иван Иванович"
                onChange={(e) => setName(e.target.value)}
                id="studentname"
              />
              <label htmlFor="studentrole" className={styles.modal__label}>
                Введите роль студента:
              </label>
              <input
                type="text"
                className={styles.modal__input}
                placeholder="Староста"
                id="studentrole"
                onChange={(e) => setRole(e.target.value)}
              />
              <button className={styles.modal__button} onClick={addStudent}>
                Добавить
              </button>
            </form>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentAdd;
