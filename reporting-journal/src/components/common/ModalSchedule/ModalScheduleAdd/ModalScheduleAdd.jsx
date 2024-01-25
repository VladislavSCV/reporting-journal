import React from "react";
import "./modalScheduleAdd.scss";
import Modal from "../../modal/Modal";

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
            <label htmlFor="lesson" className="modal__label">
              Выберите предмет:
            </label>
            <div className="custom-select">
              <select className="custom-select1">
                <option value="0">Выберите предмет</option>
                <option value="1">sfdfsd</option>
                <option value="2">fsdf</option>
                <option value="3">fdsfds</option>
                <option value="4">sdfsdf</option>
                <option value="5">fdsffd</option>
                <option value="6">fsdf</option>
              </select>
            </div>
            <label htmlFor="lessonTeacher" className="modal__label">
              Выберите имя преподавателя:
            </label>
            <div className="custom-select">
              <select className="custom-select1">
                <option value="0">Выберите преподавателя</option>
                <option value="1">sfdfsd</option>
                <option value="2">fsdf</option>
                <option value="3">fdsfds</option>
                <option value="4">sdfsdf</option>
                <option value="5">fdsffd</option>
                <option value="6">fsdf</option>
              </select>
            </div>
            <button className="modal__button">Добавить</button>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default ModalStudentAdd;
