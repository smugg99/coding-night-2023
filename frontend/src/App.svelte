<!-- App.svelte -->
<head>
	<!-- SMUI Styles -->
	<link rel="stylesheet" href="/build/smui.css" media="(prefers-color-scheme: light)" />
	<link
		rel="stylesheet"
		href="/build/smui.css"
		media="screen and (prefers-color-scheme: dark)"
	/>

	<!-- SMUI Styles -->
	<!-- <link rel="stylesheet" href="node_modules/svelte-material-ui/bare.css" /> -->

	<!-- Hint where we get fonts from. -->
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />

	<!-- Material Icons, Roboto, and Roboto Mono fonts -->
	<link
		href="https://fonts.googleapis.com/css2?family=Material+Icons&Roboto+Mono:ital@0;1&family=Roboto:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900&display=swap"
		rel="stylesheet"
	/>
</head>

<script lang="ts">
	import Navbar from './lib/Navbar.svelte';
</script>

<main>
	<Navbar>
		
	</Navbar>
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
