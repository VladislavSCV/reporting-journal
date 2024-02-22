import React from "react";
import styles from "./ModalScheduleLinks.module.scss";
import Modal from "../../modal/Modal";
import { Link } from "react-router-dom";

const ModalStudentAdd = ({ closeFn = () => null, open = false }) => {
  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
          <header className="modal__header">
            <h1 className="modal__title">Что вы хотите сделать?</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <div className={styles.modal__links}>
              <Link to="/studentAttendance" className={styles.modal__linkBtn}>
                <button className={styles.modal__linkBtn}>
                  Отметить отсутствующих
                </button>
              </Link>
              <button className={styles.modal__linkBtn}>Удалить предмет</button>
            </div>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentAdd;
