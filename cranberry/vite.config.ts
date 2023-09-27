import { defineConfig } from 'vite'
import { sveltekit } from '@sveltejs/kit/vite';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [sveltekit()],
  server: {
    proxy: {
      "/api": {
        target: "http://0.0.0.0:3000",
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
