<script lang="ts">
	import Drawer, {
		AppContent,
		Content,
		Header,
		Title,
		Subtitle,
		Scrim,
	} from '@smui/drawer';
	import Button, { Label } from '@smui/button';
	import List, { Item, Text, Graphic, Separator, Subheader } from '@smui/list';
	import { getContext } from 'svelte';
	
	let drawer: Drawer;

	export let open = false;
	let active: string;

	let toggleLogin: CallableFunction = getContext('toggleLogin');
	let toggleRegister: CallableFunction = getContext('toggleRegister');

	function setActive(value: string) {
		active = value;
		open = false;

		switch (active) {
			case 'Login':
				toggleLogin();
				break;
		
			case 'Register':
				toggleRegister();
				break;

			default:
				break;
		}
	}
</script>


<div class="drawer-container">
	<Drawer bind:this={drawer} variant="modal" bind:open>
		<Header>
			<Title>Menu</Title>
			<Subtitle></Subtitle>
		</Header>
		<Content>
			<Item
				on:click={() => setActive('Login')}
				activated={active === 'Login'}>
				<Graphic class="material-icons" aria-hidden="true">login</Graphic>
				<Text>Login</Text>
			</Item>
			<Item
				on:click={() => setActive('Register')}
				activated={active === 'Register'}>
				<Graphic class="material-icons" aria-hidden="true">assignment_ind</Graphic>
				<Text>Register</Text>
			</Item>

			<Item
				on:click={() => setActive('Inbox')}
				activated={active === 'Inbox'}>
				<Graphic class="material-icons" aria-hidden="true">help</Graphic>
				<Text>About</Text>
			</Item>
			<Item
				on:click={() => setActive('Inbox')}
				activated={active === 'Inbox'}>
				<Graphic class="material-icons" aria-hidden="true">logout</Graphic>
				<Text>Logout</Text>
			</Item>
		</Content>
	</Drawer>

	<Scrim/>
	<AppContent class="app-content">

	</AppContent>
</div>