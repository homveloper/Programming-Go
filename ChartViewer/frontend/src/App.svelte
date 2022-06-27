<svelte:head>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/xlsx/0.15.5/xlsx.full.min.js"></script>

</svelte:head>

<script lang="ts">
	import {OpenFile} from '../wailsjs/go/main/App'
	import Handsontable from './Handsontable.svelte'

	let data = [
		['','Tesla','Volvo','Toyota','Ford'],
		['2019','Tesla','Volvo','Toyota','Ford'],
		['2020','Tesla','Volvo','Toyota','Ford'],
		['2021','Tesla','Volvo','Toyota','Ford'],
	]

	let header = [
		'Year', 'A','B','C','D'
	]

	let file : FileList;

	$: if (file) {
		// Note that `files` is of type `FileList`, not an Array:
		// https://developer.mozilla.org/en-US/docs/Web/API/FileList



		OpenFile(file.item.name)
	}

</script>

<main>
	<label>Grid is here:</label>
	<Handsontable data={data} header={header}/>
	<input bind:files={file} id="avatar" name="avatar" type="file"/>
</main>
