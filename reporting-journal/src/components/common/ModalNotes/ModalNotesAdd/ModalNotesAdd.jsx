import React from "react";
import styles from "./ModalNotesAdd.module.scss";
import Modal from "../../modal/Modal";

import addNote from "../../../../js/addNote";

const ModalNotesAdd = ({ closeFn = () => null, open = false }) => {
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
            <label htmlFor="noteName" className={styles.modal__label}>
              Введите название заметки:
            </label>
            <input
              type="text"
              className={styles.modal__input}
              placeholder="Супер пупер заметка"
              id="noteName"
            />
            <label htmlFor="noteDescription" className={styles.modal__label}>
              Введите описание заметки:
            </label>
            <input
              type="text"
              className={styles.modal__input}
              placeholder="Это супер пупер заметка, если вы ее прочитали скиньте 100р, пж."
              id="noteDescription"
            />
            <button className={styles.modal__button} onClick={addNote}>
              Добавить
            </button>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalNotesAdd;
