import React from 'react'
import './userNav.scss'
const UserNav = () => {
  return (
    <div className='userNav'>
      <nav className='userNav__navigation'>
        <h1 className='userNav__text'></h1>
        <div className="userNav__user">
          <div className="userNav__user-avatar"></div>
          <div className="userNav__user-container">
          <div className="userNav__user-name">Гилоян Роман</div>
          <div className="userNav__user-role">Преподаватель</div>
          </div>
        </div>
      </nav>
    </div>
  )
}

export default UserNav