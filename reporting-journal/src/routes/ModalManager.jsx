import React from 'react';

import ModalStudentAdd from '../components/common/ModalStudent/ModalStudentAdd/ModalStudentAdd';
import ModalStudentDelete from '../components/common/ModalStudent/ModalStudentDelete/ModalStudentDelete';
import ModalStudentSettings from '../components/common/ModalStudent/ModalStudentSettings/ModalStudentSettings';
import ModalScheduleAdd from '../components/common/ModalSchedule/ModalScheduleAdd/ModalScheduleAdd'
import ModalGroupLinks from '../components/common/ModalGroup/ModalGroupLinks/ModalGroupLinks'
import ModalGroupSettings from '../components/common/ModalGroup/ModalGroupSettings/ModalGroupSettings';
import ModalMainPageInfo from '../components/common/ModalMainPage/ModalMainPageInfo/ModalMainPageInfo'
import ModalScheduleLinks from '../components/common/ModalSchedule/ModalScheduleLinks/ModalScheduleLinks'
import ModalGroupAdd from '../components/common/ModalGroup/ModalGroupAdd/ModalGroupAdd'
import ModalCuratorGroupAdd from '../components/common/ModalCuratorGroup/ModalCuratorGroupAdd/ModalCuratorGroupAdd'
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
