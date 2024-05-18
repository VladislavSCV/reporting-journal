import React from 'react';
import styles from './ModalStudentDelete.module.scss'
import Modal from '../../modal/Modal';

const ModalStudentDelete = ({ closeFn = () => null, open = false }) => {
  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
        <header className="modal__header">
            <h1 className="modal__title">Удаление студента</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <p className={styles.modal__text}>Вы точно хотите удалить студента?</p>
            <button className={styles.modal__buttonDelete}>Удалить</button>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentDelete;
