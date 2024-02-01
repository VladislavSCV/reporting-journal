import React from 'react'
import './groups.scss'
import GroupCard from '../GroupCard/GroupCard'
import {objectGroupCard} from '../../helpers/objectGroupCard'
const CuratorGroup = () => {
  return (
    
    <div className='groups'>

      <div className="groups__container">
        {objectGroupCard.map((obj, index)=>{
        return <GroupCard group={obj.group} key={index}/>})}
      
      </div>
    </div>
  )
}

export default CuratorGroup