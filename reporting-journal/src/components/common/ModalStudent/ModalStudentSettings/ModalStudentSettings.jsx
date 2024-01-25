import React from 'react';
import "./modalStudentSettings.scss";
import Modal from '../../modal/Modal';

const ModalStudentSetting = ({ closeFn = () => null, open = false }) => {
  return (
    <Modal open={open}>
      <div className="modal__mask">
        <div className="modal__window">
        <header className="modal__header">
            <h1 className="modal__title">Изменение студента</h1>
            <button className="modal__close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal__body">
            <label htmlFor="studentname" className="modal__label">Введите новое ФИО студента:</label>
            <input type="text" className="modal__input" placeholder="Иванов Иван Иванович" id="studentname"/>
            <label htmlFor="studentrole" className="modal__label">Введите новую роль студента:</label>
            <input type="text" className="modal__input" placeholder="Староста" id="studentrole"/>
            <button className="modal__button">Изменить</button>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentSetting;
