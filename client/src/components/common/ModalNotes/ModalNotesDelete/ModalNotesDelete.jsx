import React from "react";
import styles from ".//ModalNotesDelete.module.scss";
import Modal from "../../modal/Modal";

const ModalNotesDelete = ({ closeFn = () => null, open = false }) => {
  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
          <header className="modal__header">
            <h1 className="modal__title">Удаление заметки</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <p className={styles.modal__text}>
              Вы точно хотите удалить заметку?
            </p>
            <button className={styles.modal__buttonDelete}>Удалить</button>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalNotesDelete;
