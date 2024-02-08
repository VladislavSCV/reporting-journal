import React from "react";
import styles from "./ModalNotesSettings.module.scss";
import Modal from "../../modal/Modal";

const ModalNotesSettings = ({ closeFn = () => null, open = false }) => {
  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
          <header className="modal__header">
            <h1 className="modal__title">Изменение заметки</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <label htmlFor="noteName" className={styles.modal__label}>
              Введите новое название заметки:
            </label>
            <input
              type="text"
              className={styles.modal__input}
              placeholder="Пожать 100кг"
              id="noteName"
            />
            <label htmlFor="noteDescription" className={styles.modal__label}>
              Введите новое описание заметки:
            </label>
            <input
              type="text"
              className={styles.modal__input}
              placeholder="Пожать на жиме лежа 100кг до 20 лет"
              id="noteDescription"
            />
            <button className={styles.modal__button}>Изменить</button>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalNotesSettings;
