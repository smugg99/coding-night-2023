<script lang="ts">
	import 'leaflet/dist/leaflet.css';
    import {Map, TileLayer, Marker, Popup} from 'svelte-map-leaflet'

	let map:Map;


const options = {method: 'GET', headers: {'User-Agent': 'insomnium/0.2.3'}};

let mapOptions = {}
let titleUrl = ""

if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition((pos) => {
            const [lat, lon] = [pos.coords.latitude, pos.coords.longitude]

            const incidents = fetch(`http://localhost:8080/incident/?lat=${lat}&lon=${lon}&distance=10.0`, options)
            .then(response => response.json())
            .then(response => console.log(response))
            .catch(err => console.error(err))

            mapOptions = { center: [lat, lon], zoom: 16, zoomControl: false}
            titleUrl = "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            
        },
        null,
    )
}
    // const markerLatLng = [49.609841,20.703781];
    // const popupMessage = "Statue of Liberty National Monument";
</script>

<div class="map">
	<Map bind:this={map} options={mapOptions}>
		<TileLayer url={titleUrl}></TileLayer>
		<!-- <Marker latLng={markerLatLng}>
			<Popup>{popupMessage}</Popup>
		</Marker> -->
	</Map>
</div>
