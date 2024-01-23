import React from 'react'

const ScheduleModal = () => {
  return (
    <div className='studentList__modal'>
       
        <input className='studentList__modal-input' type="text" placeholder='Имя'/>
        <input className='studentList__modal-input' type="text" placeholder='Фамилия'/>
        <input className='studentList__modal-input' type="text" placeholder='Роль'/>
        <button className="studentList__modal-button">Добавить</button>
    </div>
  )
}

export default ScheduleModal