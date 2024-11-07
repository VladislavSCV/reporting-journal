import { configureStore } from "@reduxjs/toolkit";
import { thunk } from 'redux-thunk';  // Именованный импорт
import userReducer from "./userReducer";

export const store = configureStore({
  reducer: {
    user: userReducer,
  },
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(thunk),
});

export default store;
