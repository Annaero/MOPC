import { defineConfig } from 'vite'
import { sveltekit } from '@sveltejs/kit/vite';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [sveltekit()],
  server: {
    proxy: {
      "/events": {
        target: "http://0.0.0.0:8000",
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
