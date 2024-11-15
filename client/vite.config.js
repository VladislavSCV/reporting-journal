import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'https://reporting-journal-2.onrender.com', // Ваш сервер
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''), // Убирает "/api" из пути
      },
    },
  },
});
