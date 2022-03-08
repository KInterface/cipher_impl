import legacy from '@vitejs/plugin-legacy'
import wasmPack from 'vite-plugin-wasm-pack';
import { defineConfig } from 'vite'

export default defineConfig({
   plugins: [
      legacy({
         targets: ['defaults', 'not IE 11']
      }), wasmPack('./wasm')
   ]
})