import React from "react";
import "./../../common/ModalStudentAdd/modalStudentAdd.scss";
import "./../../common/modal.scss";
import Modal from "../modal/Modal";

const ModalStudentAdd = ({ closeFn = () => null, open = false }) => {
  return (
    <Modal open={open}>
      <div className="modal--mask">
        <div className="modal-window">
          <header className="modal-header">
            <h1 className="modal-title">Добавление студента</h1>
            <button className="modal-close" type="button" onClick={closeFn}>
              X
            </button>
          </header>
          <div className="modal-body">
            <label htmlFor="studentname" className="modal-label">Введите ФИО студента:</label>
            <input type="text" className="modal-input" placeholder="Иванов Иван Иванович" id="studentname"/>
            <label htmlFor="studentrole" className="modal-label">Введите роль студента:</label>
            <input type="text" className="modal-input" placeholder="Староста" id="studentrole"/>
            <button className="modal-button">Добавить</button>
          </div>
    
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentAdd;
