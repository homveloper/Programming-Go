<script>
	import {Sheet,Toolbar} from "Svelte-sheets"
	let sheetNames;
	let sheets;
	let active;
	let data;
	let columns;
	let mergeCells;
	let style;

	$: {
	data = sheets[active].data;
	columns = sheets[active].columns;
	mergeCells = sheets[active].mergeCells;
	style = sheets[active].style;
	}
</script>

<main>
	<Toolbar bind:sheetNames bind:sheets bind:active />
	<Sheet style={style} mergeCells={mergeCells} columns={columns} data={data} />
</main>

<style>
</style>
