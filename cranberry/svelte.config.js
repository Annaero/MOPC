import { vitePreprocess } from '@sveltejs/kit/vite';
import adapter from "@sveltejs/adapter-vercel"

/** @type {import('@sveltejs/kit').Config} */
const config = {
  preprocess: vitePreprocess(),
  vitePlugin: {
    // This enables compile-time warnings to be
    // visible in the learn.svelte.dev editor
    onwarn: (warning, defaultHandler) => {
      console.log('svelte:warnings:%s', JSON.stringify(warning));
      defaultHandler(warning);
    }
  },
  kit: {
    adapter: adapter({
      // adapter options go here
    })
  }
};

export default config;
