import React from 'react';

import ModalStudentAdd from './common/ModalStudent/ModalStudentAdd/ModalStudentAdd';
import ModalStudentDelete from './common/ModalStudent/ModalStudentDelete/ModalStudentDelete';
import ModalStudentSettings from './common/ModalStudent/ModalStudentSettings/ModalStudentSettings';
import ModalScheduleAdd from './common/ModalSchedule/ModalScheduleAdd/ModalScheduleAdd'
import ModalGroupLinks from './common/ModalGroup/ModalGroupLinks/ModalGroupLinks'
import ModalGroupSettings from './common/ModalGroup/ModalGroupSettings/ModalGroupSettings';
import ModalMainPageInfo from './common/ModalMainPage/ModalMainPageInfo/ModalMainPageInfo'
import ModalScheduleLinks from './common/ModalSchedule/ModalScheduleLinks/ModalScheduleLinks'
import ModalGroupAdd from './common/ModalGroup/ModalGroupAdd/ModalGroupAdd'
import ModalCuratorGroupAdd from './common/ModalCuratorGroup/ModalCuratorGroupAdd/ModalCuratorGroupAdd'
const ModalManager = ({ closeFn, modal = '' }) => {
  return (
    <>
      <ModalStudentAdd closeFn={closeFn} open={modal === 'ModalStudentAdd'} />
      <ModalStudentDelete closeFn={closeFn} open={modal === 'ModalStudentDelete'} />
      <ModalStudentSettings closeFn={closeFn} open={modal === 'ModalStudentSettings'} />
      <ModalScheduleAdd closeFn={closeFn} open={modal === 'ModalScheduleAdd'}/>
      <ModalGroupLinks closeFn={closeFn} open={modal === 'modalGroupLinks'}/>
      <ModalGroupSettings closeFn={closeFn} open={modal === 'modalGroupSettings'}/>
      <ModalMainPageInfo closeFn={closeFn} open={modal === 'modalMainPageInfo'}/>
      <ModalScheduleLinks closeFn={closeFn} open={modal === 'modalScheduleLinks'}/>
      <ModalGroupAdd closeFn={closeFn} open={modal === 'modalGroupAdd'}/>
      <ModalCuratorGroupAdd closeFn={closeFn} open={modal === 'modalCuratorGroupAdd'}/>

    </>
  );
};

export default ModalManager;
