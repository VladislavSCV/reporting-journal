import React from "react";
import styles from "./ModalCuratorGroupAdd.module.scss";
import Modal from "../../modal/Modal";
import addCuratorGroup from "../../../../js/addCuratorGroup";
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
            <label htmlFor="curatorGroupName" className={styles.modal__label}>
              Введите название группы:
            </label>
            <input
              type="text"
              className={styles.modal__input}
              placeholder="21ИС3-4Д"
              id="curatorGroupName"
            />

            <button className={styles.modal__button} onClick={addCuratorGroup}>
              Добавить
            </button>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentAdd;
