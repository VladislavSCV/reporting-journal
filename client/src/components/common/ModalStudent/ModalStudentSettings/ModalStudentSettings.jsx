import React, { useState } from "react";
import styles from "./ModalStudentSettings.module.scss";
import Modal from "../../modal/Modal";
import axios from "axios";

const ModalStudentSetting = ({ closeFn = () => null, open = false, id }) => {
  const [name, setName] = useState("");
  const [role, setRole] = useState("");
  const putStudent = async (key) => {
    try {
      await axios.put(`http://localhost:5001/api/students/${key}`, {
        name,
        role,
      });
      setName("");
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
          <header className="modal__header">
            <h1 className="modal__title">Изменение студента</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <label htmlFor="studentname" className={styles.modal__label}>
              Введите новое ФИО студента:
            </label>
            <input
              type="text"
              onChange={(e) => setName(e.target.value)}
              className={styles.modal__input}
              placeholder="Иванов Иван Иванович"
              id="studentname"
            />
            <label htmlFor="studentrole" className={styles.modal__label}>
              Введите новую роль студента:
            </label>
            <input
              type="text"
              onChange={(e) => setRole(e.target.value)}
              className={styles.modal__input}
              placeholder="Староста"
              id="studentrole"
            />
            <button
              className={styles.modal__button}
              onClick={() => putStudent(id)}
            >
              Изменить
            </button>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentSetting;
