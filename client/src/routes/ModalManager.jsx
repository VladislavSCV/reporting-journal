import React from "react";

import ModalMainInfo from "../components/common/ModalMain/ModalMainInfo/ModalMainInfo";

import ModalStudentAdd from "../components/common/ModalStudent/ModalStudentAdd/ModalStudentAdd";
import ModalStudentDelete from "../components/common/ModalStudent/ModalStudentDelete/ModalStudentDelete";
import ModalStudentSettings from "../components/common/ModalStudent/ModalStudentSettings/ModalStudentSettings";

import ModalScheduleAdd from "../components/common/ModalSchedule/ModalScheduleAdd/ModalScheduleAdd";
import ModalScheduleLinks from "../components/common/ModalSchedule/ModalScheduleLinks/ModalScheduleLinks";
import ModalScheduleSettings from "../components/common/ModalSchedule/ModalScheduleSettings/ModalScheduleSettings";
import ModalGroupAdd from "../components/common/ModalGroup/ModalGroupAdd/ModalGroupAdd";
import ModalGroupLinks from "../components/common/ModalGroup/ModalGroupLinks/ModalGroupLinks";
import ModalGroupSettings from "../components/common/ModalGroup/ModalGroupSettings/ModalGroupSettings";
import ModalGroupDelete from "../components/common/ModalGroup/ModalGroupDelete/ModalGroupDelete";

import ModalCuratorGroupAdd from "../components/common/ModalCuratorGroup/ModalCuratorGroupAdd/ModalCuratorGroupAdd";

import ModalNotesAdd from "../components/common/ModalNotes/ModalNotesAdd/ModalNotesAdd";
import ModalNotesDelete from "../components/common/ModalNotes/ModalNotesDelete/ModalNotesDelete";
import ModalNotesSettings from "../components/common/ModalNotes/ModalNotesSettings/ModalNotesSettings";

const ModalManager = ({ closeFn, modal = "", id, day }) => {
  return (
    <>
      <ModalMainInfo closeFn={closeFn} open={modal === "modalMainInfo"} />

      <ModalStudentAdd
        closeFn={closeFn}
        open={modal === "ModalStudentAdd"}
        id={id}
      />
      <ModalStudentDelete
        closeFn={closeFn}
        open={modal === "ModalStudentDelete"}
      />
      <ModalStudentSettings
        closeFn={closeFn}
        open={modal === "ModalStudentSettings"}
        id={id}
      />

      <ModalScheduleAdd
        closeFn={closeFn}
        open={modal === "ModalScheduleAdd"}
        day={day}
        id={id}
      />
      <ModalScheduleLinks
        closeFn={closeFn}
        open={modal === "modalScheduleLinks"}
        id={id}
      />
      <ModalScheduleSettings
        closeFn={closeFn}
        open={modal === "modalScheduleSettings"}
        id={id}
      />

      <ModalGroupAdd closeFn={closeFn} open={modal === "modalGroupAdd"} />
      <ModalGroupLinks closeFn={closeFn} open={modal === "modalGroupLinks"} />
      <ModalGroupSettings
        closeFn={closeFn}
        open={modal === "modalGroupSettings"}
        id={id}
      />
      <ModalGroupDelete
        closeFn={closeFn}
        open={modal === "modalGroupDelete"}
        id={id}
      />

      <ModalCuratorGroupAdd
        closeFn={closeFn}
        open={modal === "modalCuratorGroupAdd"}
      />

      <ModalNotesAdd
        closeFn={closeFn}
        open={modal === "modalNotesAdd"}
        id={id}
      />
      <ModalNotesDelete closeFn={closeFn} open={modal === "modalNotesDelete"} />
      <ModalNotesSettings
        closeFn={closeFn}
        open={modal === "modalNotesSettings"}
        id={id}
      />
    </>
  );
};

export default ModalManager;
