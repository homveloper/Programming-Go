<svelte:head>
    <script src="https://cdn.jsdelivr.net/npm/handsontable/dist/handsontable.full.min.js" 
        on:load={scriptLoaded}></script>
    <link href="https://cdn.jsdelivr.net/npm/handsontable/dist/handsontable.full.min.css" rel="stylesheet" 
        on:load={styleLoaded}>
</svelte:head>

<script>
    import { onMount } from 'svelte';
    import Handsontable  from 'handsontable'

    export let data;

    // string[] 
    export let header;

    let gridElement;

    let gridStatus = {
        isScriptLoaded: false,
        isStyleLoaded: false,
        isMounted: false,
        isInited: false
    }
    
    onMount(() => {
        gridStatus.isMounted = true;
        if (gridStatus.isScriptLoaded && gridStatus.isStyleLoaded) gridInit()
    })

    function scriptLoaded() {
        gridStatus.isScriptLoaded = true;
        if (gridStatus.isMounted && gridStatus.isStyleLoaded) gridInit()
    }

    function styleLoaded() {
        gridStatus.isStyleLoaded = true;
        if (gridStatus.isScriptLoaded && gridStatus.isMounted) gridInit()
    }

    function gridInit() {
        if (!gridStatus.isInited) {
            gridStatus.isInited = true;
            new Handsontable(gridElement,{
                data:data,
                rowHeaders:true,
                colHeaders:header,
                licenseKey: 'non-commercial-and-evaluation',
            });
        } 
    }
</script>

<div bind:this={gridElement}></div>