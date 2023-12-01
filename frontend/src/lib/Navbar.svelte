<script lang="ts">
	import TopAppBar, {
		Row,
		Section,
		Title,
		AutoAdjust,
	} from '@smui/top-app-bar';
	import IconButton from '@smui/icon-button';

	let topAppBar: TopAppBar;
</script>

<nav>
	<TopAppBar bind:this={topAppBar} variant="fixed">
	<Row>
		<Section>
		<IconButton class="material-icons">menu</IconButton>
		<Title>Haseus</Title>
		</Section>
		<Section align="end" toolbar>

		</Section>
	</Row>
	</TopAppBar>
</nav>