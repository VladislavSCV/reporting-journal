import React from 'react';

import ModalStudentAdd from './common/ModalStudentAdd/ModalStudentAdd';
import ModalStudentDelete from './common/ModalStudentDelete/ModalStudentDelete';
import ModalStudentSetting from './common/ModalStudentSetting/ModalStudentSetting';

const ModalManager = ({ closeFn, modal = '' }) => {
  return (
    <>
      <ModalStudentAdd closeFn={closeFn} open={modal === 'ModalStudentAdd'} />
      <ModalStudentDelete closeFn={closeFn} open={modal === 'ModalStudentDelete'} />
      <ModalStudentSetting closeFn={closeFn} open={modal === 'ModalStudentSettings'} />
    </>
  );
};

export default ModalManager;
