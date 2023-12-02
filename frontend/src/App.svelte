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

	<link rel="stylesheet" href="https://unpkg.com/leaflet@1.6.0/dist/leaflet.css"
		integrity="sha512-xwE/Az9zrjBIphAcBb3F6JVqxf46+CDLwfLMHloNu6KEQCAWi6HcDUbeOfBIptF7tcCzusKFjFw2yuvEpDL9wQ=="
		crossorigin=""
	/>
</head>

<script lang="ts">
	import Navbar from './lib/Navbar.svelte';
	import Map from './lib/Map.svelte';
	import Drawer from './lib/Drawer.svelte';
	import Login from './lib/Login.svelte';
	import { setContext } from 'svelte';

	let hamburgerOpen: boolean = false;
	let loginOpen: boolean = false;

	setContext('toggleHamburger', () => {
		hamburgerOpen = !hamburgerOpen;
	});
	
	setContext('toggleLogin', () => {
		loginOpen = !loginOpen;
	});
</script>


<body>
	<Navbar/>
	<Drawer open={hamburgerOpen}/>
	<Login open={loginOpen}/>
	<main>
		<Map/>
	</main>
</body>