import React, { useState } from "react";
import styles from "./ModalNotesAdd.module.scss";
import Modal from "../../modal/Modal";
import axios from "axios";

const ModalNotesAdd = ({ closeFn = () => null, open = false, id }) => {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");

  const addNote = async () => {
    try {
      await axios.post("http://localhost:5001/api/notes", {
        title,
        body,
        groupId: id,
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
            <h1 className="modal__title">Добавление заметки</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <form>
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
                onClick={addNote}
                type="submit"
                className={styles.modal__button}
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
