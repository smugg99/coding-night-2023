import './build/smui.css'
import App from './App.svelte'
import 'leaflet/dist/leaflet.css';

const app = new App({
  target: document.getElementById('app')!,
})

export default app