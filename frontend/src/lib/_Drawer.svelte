<script lang="ts">
	import Drawer, {
		AppContent,
		Content,
		Header,
		Title,
		Subtitle,
	} from '@smui/drawer';

	import Button, { Label } from '@smui/button';
	import List, { Item, Text } from '@smui/list';
	import { getContext } from 'svelte';

	let drawer: Drawer;
	export let open = false;
</script>

<div class="drawer-container">
	<Drawer bind:this={drawer} variant="dismissible" bind:open>
		<Header>
			<Title>Super Drawer</Title>
			<Subtitle>It's the best drawer.</Subtitle>
		</Header>
		<Content>
			<List>
				<Item
				href="javascript:void(0)"
				on:click={() => {}}
				>
				<Text>Gray Kittens</Text>
				</Item>
				<Item
				href="javascript:void(0)"
				on:click={() => {}}
				>
				<Text>A Space Rocket</Text>
				</Item>
			</List>
		</Content>
	</Drawer>

	<AppContent class="app-content">
		<main class="main-content">
	</AppContent>
</div>