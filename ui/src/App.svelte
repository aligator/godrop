<script lang="ts">

    import {GetNode, GetNodeQuery, GetNodeQueryVariables} from "./api/types";
    import type {Readable} from "svelte/store";
    import type {ApolloQueryResult} from "@apollo/client";
    import {ObservableQuery} from "@apollo/client";

    export let name: string;
    const path = "/"
    let node: Readable<ApolloQueryResult<GetNodeQuery> & {query: ObservableQuery<GetNodeQuery, GetNodeQueryVariables>}> = GetNode({variables: {path}})

    $: node = node

    node.subscribe((value) => {
        console.log(value?.data?.getNode)
    })
</script>

<main>
    <h1>Hello {name}!</h1>
    <p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
    <ul>
        {#each $node?.data?.getNode?.name || "" as filename}
            <li>{filename}</li>
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