interface wsConfig {
    store:any
    url:string
    callback?: { (): void; } []
}

export class Socket {
    ws: any;
    store: any;
    constructor(config:wsConfig){
        this.store = config.store;
        this.callback = config.callback;
        this.ws = new WebSocket(config.url);
        this.ws.onopen = this.onOpen;
        this.ws.onmessage = (e)=>this.onMessage(e);
        this.ws.onerror = (e)=>this.onError(e);
        this.ws.onclose = (e)=>this.onClose(e);
    }
    onOpen(event:any):any{
        return;
    }
    onMessage(event:any, store?:any):any{
        const storer = store ? store : this.store;
        const data = JSON.parse(event.data);
        storer.update((messenger:any) => [...messenger, data]);
        if(this.callback)
            for(let func of this.callback)
                func();
    }
    onError(event:any):any{
        return;
    }
    onClose(event:any):any{

    }
    sendData(data:string){
        this.ws.send(data)    
    }
}
