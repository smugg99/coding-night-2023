<script lang="ts">
	import 'leaflet/dist/leaflet.css';
    import {Map, TileLayer, Marker, Popup} from 'svelte-map-leaflet'

    const mapOptions = { center: [49.609841,20.703781], zoom: 16};
    const tileUrl = "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png";
    // const markerLatLng = [49.609841,20.703781];
    // const popupMessage = "Statue of Liberty National Monument";
</script>

<div class="map">
	<Map options={mapOptions}>
		<TileLayer url={tileUrl}></TileLayer>
		<!-- <Marker latLng={markerLatLng}>
			<Popup>{popupMessage}</Popup>
		</Marker> -->
	</Map>
</div>