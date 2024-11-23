import { defineConfig } from 'vite';
import React from "react";


export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8000', // Адрес вашего бэкенда на Go
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
