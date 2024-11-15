import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    proxy: {
      '/api': 'https://reporting-journal-2.onrender.com'
    }
  }
});
