<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();
    let bubble:Element;
    let bubbleicon:Element;
    let toggleShow = () => {
        ['icon'].map((val)=> bubbleicon.classList.toggle(val));
        ['buuubble'].map((val)=> bubble.classList.toggle(val));
        setTimeout(()=>{ 
            ['rounded-full','rounded-t-2xl'].map((val)=> bubble.classList.toggle(val));
        }, 200)
        dispatch("bubbling");
    }
</script>

<div class="max-w-sm flex flex-col flex-auto h-full p-6">
    <div bind:this="{bubble}" on:click={toggleShow} class="bubble grid grid-cols-5 gap-3 bg-slate-200 p-4 rounded-full place-items-center">
        <div bind:this={bubbleicon} class="bubbleicon" style="">
            <img alt="bubble_icon" src="/images/bubbleicon.svg"/> 
        </div>
    </div>
</div>

<div class="buuubble icon">ok</div>

<style>

.icon{
    opacity: 0;
}

.bubbleicon{
    position:absolute;
    z-index: -1;
    color: #afafaf;
    transition: 0.25s all;
}

.bubble:hover .bubbleicon {
    transform: scale(1.3);
}

.bubble {
    z-index:999;
    position:fixed;
    bottom: 20px;
    right: 45px;
    width: 60px;
    height: 60px;
    transition: 0.2s ease-in-out;
}

.buuubble{
    opacity: 0;
    width: 336px;
    bottom: 415px;
    -webkit-transition: bottom 0.25s ease-in-out 0.35s, width 0.4s ease-in-out;
    -moz-transition: bottom 0.25s ease-in-out 0.35s, width 0.4s ease-in-out;
    -o-transition: bottom 0.25s ease-in-out 0.35s, width 0.4s ease-in-out;
    transition: bottom 0.25s ease-in-out 0.25s, width 0.25s ease-in-out, opacity 0.4s ease-in-out 0.35s;
}

</style>
