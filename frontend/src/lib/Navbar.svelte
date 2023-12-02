<script lang="ts">
	import TopAppBar, {
		Row,
		Section,
		Title,
		AutoAdjust,
	} from '@smui/top-app-bar';
	import Wrapper from '@smui/touch-target';
	import Button from '@smui/button';
	import { Icon } from '@smui/button';
	import IconButton from '@smui/icon-button';
	import Drawer from './_Drawer.svelte';
	import { writable } from 'svelte/store';
	import { getContext } from 'svelte';

	let topAppBar: TopAppBar;
	let toggleHamburger: CallableFunction = getContext('toggleHamburger');
</script>

<nav>
	<TopAppBar bind:this={topAppBar} variant="standard">
	<Row>
		<Section>
		<Wrapper>
			<IconButton on:click={() => { toggleHamburger(); }} touch>
				<Icon class="material-icons">menu</Icon>
			</IconButton>
		</Wrapper>
		<Title>Haseus</Title>
		</Section>
		<Section align="end" toolbar>
		</Section>
	</Row>
	</TopAppBar>
</nav>
