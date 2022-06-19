import { writable } from 'svelte/store';

export let unread = writable(0);
export let messenger = writable([{ author: "0", body: "Hi there,  i'll answer in a few minutes :)" }]);

//type State = {
//  requests: Array<Request>;
//};
//export const state = writable<State>({
//  requests: [],
//});

















