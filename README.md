# livematrix
Embedded live chat integration for your website, using matrix. 

Live chat with your website's visitors, using your Matrix.org client to communicate.


## What is livechat?

An oversimplified embedded live chat widget that allows your website's visitors to send you messages seamlessly to your Matrix.org account.

## How does it work? 

There would be many ways of doing a live chat, for example, we could allow each visitor to create a Matrix.org account, but people might not want to go through that. You could also use your own homeserver and [allow guests](https://spec.matrix.org/latest/client-server-api/#guest-access) . Why your own? Because most homeservers disallow guests. 

The next best bet is for you to create a new Matrix.org account and all visitor's chats will be mediated by it. Each new conversation is a new room. When a visitor starts a chat he/she needs to introduce their name, surname and email. If Melissa Brandon starts a chat, you'll receive an invitation on your personal account to join a room of the name `#Melissa_Brandon4212353:Matrix.org`


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

You must create a Matrix.org user specifically to mediate conversations with your visitors. The .env looks like this:

```
DATABASE_NAME=plexerbot
DATABASE_USER=osousa
DATABASE_PASSWORD=password

# Your personal Matrix.org account
MATRIX_RECIPIENT=@osousa:matrix.io

# Account used ONLY for mediation
MATRIX_USERNAME=@osousa:privex.io
MATRIX_PASSWORD=6VrdT8DCsa1xDvyaOghxT

SERVER_PORT=8000
```

DO NOT use your personal account's password and username. CREATE a new one for this purpose only.

