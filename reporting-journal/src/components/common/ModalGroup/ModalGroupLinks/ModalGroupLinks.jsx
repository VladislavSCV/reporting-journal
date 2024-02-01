import React from "react";
import styles from './ModalGroupLinks.module.scss'
import Modal from "../../modal/Modal";
import {Link} from "react-router-dom";

const ModalStudentAdd = ({ closeFn = () => null, open = false }) => {
  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
          <header className="modal__header">
            <h1 className="modal__title">Куда вы хотите перейти?</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
                <div className={styles.modal__links}>
                  <Link to="/schedule">
                    <button className={styles.modal__linkBtn}>Расписание</button>
                  </Link>
                  <button className={styles.modal__linkBtn}>Вложения</button>
                  <Link to="/studentsList">
                    <button className={styles.modal__linkBtn}>Студенты</button>
                  </Link>
            </div>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentAdd;
