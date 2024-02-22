import React from "react";
import styles from "./ModalGroupAdd.module.scss";
import Modal from "../../modal/Modal";
import addGroup from "../../../../js/addGroup";
const ModalStudentAdd = ({ closeFn = () => null, open = false }) => {
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
            <label htmlFor="groupName" className={styles.modal__label}>
              Введите название группы:
            </label>
            <input
              type="text"
              className={styles.modal__input}
              placeholder="21ИС3-4Д"
              id="groupName"
            />

            <button
              className={styles.modal__button}
              onClick={addGroup}
              id="studentAddBtn"
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
