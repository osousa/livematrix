<script lang="ts">
    import Bubble from "./Bubble.svelte"
    import Chatbox from "./Chatbox.svelte"
    let openchat = false;
</script>

<h1 class="text-3xl font-bold underline">
  Hello world!
</h1>
<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<Bubble on:bubbling={()=> {openchat=!openchat}}/>
<Chatbox bind:chatboxOpen={openchat} />
