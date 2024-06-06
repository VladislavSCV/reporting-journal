import React, { useState } from "react";
import styles from "./ModalGroupAdd.module.scss";
import Modal from "../../modal/Modal";
import axios from "axios";
const ModalStudentAdd = ({ closeFn = () => null, open = false }) => {
  const [name, setName] = useState("");

  const addGroup = async () => {
    try {
      await axios.post("http://localhost:5001/api/groups", { name });
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
          <header className="modal__header">
            <h1 className="modal__title">Добавление группы</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <form>
              <label htmlFor="groupName" className={styles.modal__label}>
                Введите название группы:
              </label>
              <input
                required
                type="text"
                className={styles.modal__input}
                onChange={(e) => setName(e.target.value)}
                placeholder="21ИС3-4Д"
                id="groupName"
              />

              <button onClick={addGroup} className={styles.modal__button}>
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
