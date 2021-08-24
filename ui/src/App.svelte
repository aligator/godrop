<script lang="ts">
    import {GetNode} from "./api/types";

    export let name: string;
    const path = "/"
    let node = GetNode({variables: {path}})
    $: node = node
</script>

<main>
    <h1>Hello {name}!</h1>
    <p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
    <ul>
        {#each $node?.data?.getNode?.children || "" as node}
            <li>{node.name}</li>
        {/each}
    </ul>
</main>

<style>
    main {
        text-align: center;
        padding: 1em;
        max-width: 240px;
        margin: 0 auto;
    }

    h1 {
        color: #ff3e00;
        text-transform: uppercase;
        font-size: 4em;
        font-weight: 100;
    }

    @media (min-width: 640px) {
        main {
            max-width: none;
        }
    }
</style>