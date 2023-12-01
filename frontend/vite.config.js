import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import ViteSassPlugin from 'vite-plugin-sass';


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte(), ViteSassPlugin()],
})