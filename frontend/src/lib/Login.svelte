<script lang="ts">
	import Paper, { Title, Subtitle, Content } from '@smui/paper';
	
	export let open = false;

	let paper: Paper;
</script>

<div class="paper-container">
	<Paper bind:this={paper} variant="modal" bind:open>
		<Title>Login</Title>
		<Subtitle>This is a sheet of paper.</Subtitle>
		<Content>Paper is used to build an elevated surface.</Content>

	</Paper>
</div>
