import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    proxy: {
      "/events": {
        target: "http://0.0.0.0:8080",
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
