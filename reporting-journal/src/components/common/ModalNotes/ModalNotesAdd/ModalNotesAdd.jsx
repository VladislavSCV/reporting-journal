import React, { useState } from "react";
import styles from "./ModalNotesAdd.module.scss";
import Modal from "../../modal/Modal";
import axios from "axios";
import addNote from "../../../../js/addNote";

const ModalNotesAdd = ({ closeFn = () => null, open = false }) => {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");

  const addNote = async () => {
    try {
      await axios.post("http://localhost:5001/api/notes", { title, body });
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
          <header className="modal__header">
            <h1 className="modal__title">Добавление заметки</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <form onChange={addNote}>
              <label htmlFor="noteName" className={styles.modal__label}>
                Введите название заметки:
              </label>
              <input
                type="text"
                className={styles.modal__input}
                placeholder="Супер пупер заметка"
                onChange={(e) => setTitle(e.target.value)}
                id="noteName"
              />
              <label htmlFor="noteDescription" className={styles.modal__label}>
                Введите описание заметки:
              </label>
              <input
                type="text"
                className={styles.modal__input}
                onChange={(e) => setBody(e.target.value)}
                placeholder="Это супер пупер заметка, если вы ее прочитали скиньте 100р, пж."
                id="noteDescription"
              />
              <button
                type="button"
                className={styles.modal__button}
                onClick={console.log(1)}
              >
                Добавить
              </button>
            </form>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalNotesAdd;
