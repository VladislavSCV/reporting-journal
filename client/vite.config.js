import { defineConfig } from 'vite';
import React from "react";


export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'https://reporting-journal-1.onrender.com', // Адрес вашего бэкенда на Go
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
