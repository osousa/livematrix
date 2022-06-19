<script lang="ts">
    import { createEventDispatcher, SvelteComponent } from 'svelte';
	import { messenger, unread } from "./stores.js";
    import { tick } from 'svelte';
    import { onMount } from 'svelte';
    import Messages from "./Messages.svelte"
    import Session from "./Session.svelte"
    import { Socket } from "./websockets"
    
    export let chatboxOpen:boolean;
    const dispatch = createEventDispatcher();
    let socket:Socket;  
    let MessagesComponent:SvelteComponent; 
    let msg:string = "";
    let node:Element;
    let session:any 
	let countValue:number;

	unread.subscribe( (value:number) => {
		countValue = value;
	});

    let notify = async () => {
        await tick().then(()=>{
            if(chatboxOpen){
                unread.set(0)
            }else{
                unread.update((n:number) => n + 1);
            }
            console.log(`unread: ${countValue}`)
        })
    }

    let scrolld = async () => {
        await tick().then(MessagesComponent.scrollDown);
        console.log("works....");
    }
    let watchSession = () => {
        if(session && !socket)
            socket = new Socket({url:`ws://${import.meta.env.VITE_HOSTNAME}:${import.meta.env.VITE_HOST_PORT}/entry`, store:messenger, callback: [scrolld, notify]});
        else
            console.log("Socket either connected or no session yet")
    } 

    const addMsg = async () => {
        if (msg==="")
            return
        let message:string = JSON.stringify({author:"1", body: msg});
        messenger.update((messenger : any) => [...messenger, JSON.parse(message)]);
        socket.sendData(message)
        await tick().then(scrolld);
        dispatch("newmessage");
        msg = "";
    }

    //$: msg, addMsg()

    let watchMe = (node:Element) => {
        if(node===undefined || node===null)
            return;
        if(chatboxOpen){
            node.classList.add('grow');
            node.classList.remove('chatHidden');
            unread.set(0);
        }else{
            node.classList.remove('grow');
            node.classList.add('chatHidden');
        }
        dispatch("togglechat", {
            variable: "empty" 
        });
    }

    onMount(() => {
        session = document.cookie.match(/^(.*;)?\s*session_id\s*=\s*[^;]+(.*)?$/)
        console.log(session)
        watchSession()
    });

    $: chatboxOpen, watchMe(node);

</script>

<div class="fixed msgbox">
    <div bind:this={node}  class="max-w-sm flex flex-col flex-auto h-full p-6 chatHidden" >
      <div class="grid grid-cols-5 gap-3 bg-slate-200 p-4 rounded-t-2xl place-items-center">
        <div class="">...</div>
        <div class="col-start-2 col-span-3 ">chat with me</div>
        <div class="">x</div>
      </div>
      <div class="flex flex-col flex-auto flex-shrink-0 rounded-b-2xl bg-gray-100 h-full px-4 pb-4">
          {#if !session}
            <Session on:sessiondone={()=> (session = !session) && watchSession() } />
          {:else}
            <Messages bind:this={MessagesComponent} /> 
            <div class="flex flex-row items-center h-16 rounded-xl bg-white w-full px-4">
              <div>
              </div>
              <div class="flex-grow">
                <div class="relative w-full">
                    <input bind:value={msg} on:keydown={e => e.key === 'Enter' && addMsg() } type="text" class="flex w-full border rounded-xl focus:outline-none focus:border-indigo-300 pl-4 pr-10 h-10">
                  <button class="absolute flex items-center justify-center h-full w-12 right-0 top-0 text-gray-400 hover:text-gray-600">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </button>
                </div>
              </div>
              <div class="ml-4">
                  <button on:click={() => msg!=="" && addMsg()} class="flex items-center justify-center bg-indigo-500 hover:bg-indigo-600 rounded-xl text-white px-4 py-1 flex-shrink-0">
                  <span class="p-5px">
                    <svg class="w-4 h-4 transform rotate-45 -mt-px" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
                    </svg>
                  </span>
                </button>
              </div>
            </div>
          {/if}
      </div>
    </div>
</div>

<div class="grow"></div>

<style>

.msgbox{
    bottom:0;
    right: 20px;
}

.chatHidden{
    opacity: 0;
    padding-bottom: 5rem;
    height:0px;
    width:0px;
    transition: height 0.2s ease-in-out, opacity 0.2s ease-in-out, width 0.2s ease-in-out;
}
.grow{
    opacity: 1;
    width:100%;
    padding-bottom: 5rem;
    height: 500px;
    -webkit-transition:  height 1s ease-in-out;
    -moz-transition:  height 1s ease-in-out;
    -o-transition:  height 1s ease-in-out;
    transition: height 0.25s ease-in-out 0.25s, opacity 0.5s ease-in-out 0.5s;
}

</style>
