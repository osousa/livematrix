import { writable } from 'svelte/store';

export let unread = writable(0);
export let messenger = writable([{ author: "0", body: "hello guest, how are you doing?" }, { author: "1", body: "i'm doing fine, thank you very much! " }]);

//type State = {
//  requests: Array<Request>;
//};
//export const state = writable<State>({
//  requests: [],
//});

















