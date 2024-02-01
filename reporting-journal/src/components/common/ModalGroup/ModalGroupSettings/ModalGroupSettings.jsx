import React from "react";
import styles from "./ModalGroupSettings.module.scss";
import Modal from "../../modal/Modal";
import {Link} from "react-router-dom";

const ModalGroupSettings = ({ closeFn = () => null, open = false }) => {
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
                
                <label htmlFor="studentsurename" className={styles.modal__label}>Введите новое название:</label>
                <input type="text" className={styles.modal__input} placeholder="21ИС3-4Д" id="studentsurename"/>
                <button className={styles.modal__button}> Изменить </button>

          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalGroupSettings;
