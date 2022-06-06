import { writable } from 'svelte/store';

export const unread = writable(0);
export const messenger = writable([{ author: 0, body: "hello author, how are you doing?" }, { author: 1, body: "i'm doing fine, thank you very much! " }]);

//type State = {
//  requests: Array<Request>;
//};
//export const state = writable<State>({
//  requests: [],
//});

















