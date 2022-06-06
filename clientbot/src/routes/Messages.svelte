<script lang="ts">
    import { messenger, unread } from "./stores.js"; 
    import { onMount } from 'svelte';
    import { fade, fly } from 'svelte/transition';

    let box:Element; 

    export function scrollDown(){
        box.scrollTo({top: box.scrollHeight, behavior: 'smooth'});
    } 

	onMount(async () => {
        scrollDown();
    });
</script>

<div class="flex flex-col h-full overflow-x-auto">
    <div class="flex flex-col overflow-y-auto max-h-80"  bind:this={box}>
    <div class="grid grid-cols-12 gap-y-2" >
     {#each $messenger as message} 
      {#if message.author==0}
       <div class="col-start-1 col-end-13 p-3 rounded-lg" transition:fade>
         <div class="flex flex-row items-center">
           <div class="flex items-center justify-center h-10 w-10 rounded-full bg-indigo-500 flex-shrink-0">
             S
           </div>
           <div class="relative ml-3 text-sm bg-white py-2 px-4 shadow rounded-xl">
             <div>{message.body}</div>
           </div>
         </div>
       </div>
      {:else}
       <div class="col-start-1 col-end-13 p-3 rounded-lg" transition:fade>
         <div class="flex items-center justify-start flex-row-reverse">
           <div class="flex items-center justify-center h-10 w-10 rounded-full bg-indigo-500 flex-shrink-0">
             G
           </div>
           <div class="relative mr-3 text-sm bg-indigo-100 py-2 px-4 shadow rounded-xl">
               <div>{ message.body }</div>
           </div>
         </div>
       </div>
      {/if}
     {/each}
    </div>
  </div>
</div>

