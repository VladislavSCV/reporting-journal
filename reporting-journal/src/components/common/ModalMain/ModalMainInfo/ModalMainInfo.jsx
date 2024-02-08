import React from "react";
import styles from "./ModalMainInfo.module.scss";
import Modal from "../../modal/Modal";

const ModalStudentDelete = ({ closeFn = () => null, open = false }) => {
  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
          <header className="modal__header">
            <h1 className="modal__title">Как поулчить доступ?</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <p className={styles.modal__text}>
              Для получения доступа обратитесь к администратору сайта, он выдаст
              вам логин и пароль для входа в аккаунт
            </p>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentDelete;
