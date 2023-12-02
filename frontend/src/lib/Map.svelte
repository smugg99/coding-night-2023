<script lang="ts">
	import 'leaflet/dist/leaflet.css';
    import {Map, TileLayer, Marker, Popup} from 'svelte-map-leaflet'

	let map:Map;


const options = {method: 'GET', headers: {'User-Agent': 'insomnium/0.2.3'}};

fetch('http://localhost:8080/incident/?lat=49.60925678752547&lon=20.697714282865814&distance=10.0', options)
  .then(response => response.json())
  .then(response => console.log(response))
  .catch(err => console.error(err));

    const mapOptions = { center: [49.60925678752547, 20.697714282865814], zoom: 16, zoomControl: false};
    const tileUrl = "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png";
    // const markerLatLng = [49.609841,20.703781];
    // const popupMessage = "Statue of Liberty National Monument";
</script>

<div class="map">
	<Map bind:this={map} options={mapOptions}>
		<TileLayer url={tileUrl}></TileLayer>
		<!-- <Marker latLng={markerLatLng}>
			<Popup>{popupMessage}</Popup>
		</Marker> -->
	</Map>
</div>