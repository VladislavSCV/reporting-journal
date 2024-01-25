import React from 'react';

import ModalStudentAdd from './common/ModalStudent/ModalStudentAdd/ModalStudentAdd';
import ModalStudentDelete from './common/ModalStudent/ModalStudentDelete/ModalStudentDelete';
import ModalStudentSettings from './common/ModalStudent/ModalStudentSettings/ModalStudentSettings';
import ModalScheduleAdd from './common/ModalSchedule/ModalScheduleAdd/ModalScheduleAdd'
const ModalManager = ({ closeFn, modal = '' }) => {
  return (
    <>
      <ModalStudentAdd closeFn={closeFn} open={modal === 'ModalStudentAdd'} />
      <ModalStudentDelete closeFn={closeFn} open={modal === 'ModalStudentDelete'} />
      <ModalStudentSettings closeFn={closeFn} open={modal === 'ModalStudentSettings'} />
      <ModalScheduleAdd closeFn={closeFn} open={modal === 'ModalScheduleAdd'}/>
    </>
  );
};

export default ModalManager;
