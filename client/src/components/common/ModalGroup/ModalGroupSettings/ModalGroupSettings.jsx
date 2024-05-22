import React, { useState } from "react";
import styles from "./ModalGroupSettings.module.scss";
import Modal from "../../modal/Modal";
import { Link } from "react-router-dom";
import axios from "axios";

const ModalGroupSettings = ({ closeFn = () => null, open = false, id }) => {
  const [name, setName] = useState("");

  const putGroup = async (key) => {
    try {
      await axios.put(`http://localhost:5001/api/groups/${key}`, {
        name,
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
            <h1 className="modal__title">Настройка группы</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <form action="">
              <label htmlFor="studentsurename" className={styles.modal__label}>
                Введите новое название:
              </label>
              <input
                onChange={(e) => setName(e.target.value)}
                type="text"
                className={styles.modal__input}
                placeholder="21ИС3-4Д"
                id="studentsurename"
              />
              <button
                className={styles.modal__button}
                onClick={() => putGroup(id)}
              >
                Изменить
              </button>
            </form>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalGroupSettings;
