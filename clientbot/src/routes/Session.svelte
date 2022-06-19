<script lang="ts">
    import { messenger, unread } from "./stores.js"; 
    import { onMount } from 'svelte';
    import { fade, fly } from 'svelte/transition';
    import { createEventDispatcher } from 'svelte';

    interface Data{
        name: String, 
        surname: String,
        email: String,
    }
    const postData = {
      name: '',
      surname: '',
      email: ''
    }

    const dispatch = createEventDispatcher();

    let session = async (etc:any) => {
        const formBody = Object.keys(postData).map(key => encodeURIComponent(key) + '=' + encodeURIComponent(postData[key as keyof Data])).join('&');
        fetch(`http://${import.meta.env.VITE_HOSTNAME}:${import.meta.env.VITE_HOST_PORT}/session`,{ 
            method:"POST", 
            mode:"no-cors", 
            credentials: "include",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            body: formBody, 
        })
        .then(res => {
            console.log(res.status)
            if(res.status===0)
                dispatch("sessiondone")
            else
                throw new Error('Network response was not OK');
            })
        .catch(err=> console.log(err))
    }
</script>

<div class="flex flex-col h-full overflow-x-auto">
    <div class="lg:flex lg:flex-wrap g-0">
      <div class="lg:w-full 2 px-0 md:px-0">
        <div class="md:p-5 md:mx-5">
          <form>
            <p class="mb-4">Please fill in your name/email:</p>
            <div class="mb-4">
              <input
                type="text"
                class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                id="exampleFormControlInput1"
                placeholder="name"
                bind:value={postData.name}
              />
            </div>
            <div class="mb-4">
              <input
                type="text"
                class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                id="exampleFormControlInput1"
                placeholder="surname"
                bind:value={postData.surname}
              />
            </div>
            <div class="mb-4">
              <input
                type="text"
                class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                id="exampleFormControlInput1"
                placeholder="email"
                bind:value={postData.email}
              />
            </div>
            <div class="text-center pt-1 mb-12 pb-1">
             <button on:click={session}
                class="inline-block px-6 py-2.5 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:shadow-lg transition duration-150 ease-in-out w-full mb-3"
                type="button"
                data-mdb-ripple="true"
                data-mdb-ripple-color="light"
                style="
                  background: linear-gradient(
                    to right,
                    #ee7724,
                    #d8363a,
                    #dd3675,
                    #b44593
                  );
                "
              >
                open chat
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
</div>

