# livematrix
Embedded live chat integration for your website, using matrix. 

Live chat with your website's visitors, using your Matrix.org client to communicate.


## What is livechat?

An oversimplified embedded live chat widget that allows your website's visitors to send you messages seamlessly to your Matrix.org account.


# How to use

##  :computer: client

Svelte is a vite framework, so everything goes in one bundle, its fast and awesome to embed. No need for external libraries or big framework.js files.

To embed the widget is very straightforward, copy the following and paste on your index.html (or other):

```
    <script type="module" crossorigin src="/static/js/index-03328893.js"></script>
    <link rel="stylesheet" href="/static/css/index-0e3c6aef.css">
    <div id="app" class="tw-fixed"></div>
```

### Configure paths on client-side

On the "client" folder (Not the submodule), download the config.json file, and put it inside a /static/config.json folder relative to your index.html:

```
root/ .
      |_ index.html
      |_ static/.  
                |_ config.json 
```             

Edit the config.json to set your server's addresses and ports;

#### ❔ 
If you don't wish to compile things yourself using Svelte and Vite's .env variables, this is the way. 



## 🛰️ Server

On the "server" folder (Not the submodule), download the binary server and the .env file; Edit the .env file and add your personal details. 

You must create a 


