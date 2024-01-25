import React from "react";
import "./modalStudentAdd.scss";
import Modal from "../../modal/Modal";

import addStudent from "../../../../js/addStudent";

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
            <label htmlFor="studentsurename" className="modal__label">Введите фамилию студента:</label>
            <input type="text" className="modal__input" placeholder="Иванов" id="studentsurename"/>
            <label htmlFor="studentname" className="modal__label">Введите имя студента:</label>
            <input type="text" className="modal__input" placeholder="Иван" id="studentname"/>
            <label htmlFor="studentpatronymic" className="modal__label">Введите отчество студента:</label>
            <input type="text" className="modal__input" placeholder="Иванович" id="studentpatronymic"/>
            <label htmlFor="studentrole" className="modal__label">Введите роль студента:</label>
            <input type="text" className="modal__input" placeholder="Староста" id="studentrole"/>
            <button className="modal__button" onClick={addStudent} id="studentAddBtn">Добавить</button>
          </div>
    
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentAdd;
