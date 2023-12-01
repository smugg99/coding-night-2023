<head>
	<!-- SMUI Styles -->
	<link rel="stylesheet" href="/build/smui.css" media="(prefers-color-scheme: light)" />
	<link
		rel="stylesheet"
		href="/build/smui-dark.css"
		media="screen and (prefers-color-scheme: dark)"
	/>
</head>

<main>
	<h1>Haseus</h1>

	<p>
	Help The Others During Crises
	</p>

	<button type="button" >login</button>
	<button type="button" >register</button>

</main>

<br>

<script>
    import 'leaflet/dist/leaflet.css'; //Don't forget to declare leaflet css
    import {Map, TileLayer, Marker, Popup} from 'svelte-map-leaflet'

    const mapOptions = { center: [49.609841,20.703781], zoom: 16};
    const tileUrl = "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png";
    const markerLatLng = [49.609841,20.703781];
    const popupMessage = "Statue of Liberty National Monument";
</script>

<div class="map">
    <Map options={mapOptions}>
        <TileLayer url={tileUrl}></TileLayer>
        <Marker latLng={markerLatLng}>
            <Popup>{popupMessage}</Popup>
        </Marker>
    </Map>
</div>

<style>
    .map{
        height: 400px;
        width: 400px;
    }
</style>
