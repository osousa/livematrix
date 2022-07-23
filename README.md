# livematrix
Chat with your website's visitors, using your Matrix client of choice.

If you need help, you can ask here: **#livematrix:matrix.org**


# 📌 What is livematrix?

An **oversimplified**, easy to implement, embedded live chat widget that allows your website's visitors to send you messages seamlessly to your Matrix account.

#### :bust_in_silhouette: Why livematrix?
livematrix is intended for personal use, very lightweight, stripped down and easy to install. You have a personal blog/website and really enjoy using Matrix? Well, give it a try :) 


#### :heavy_check_mark: Todo
- [ ] Add reCaptcha V3 or hcaptcha
- [ ] Allow for chat history to be fetched
- [ ] OLM encryption (not e2ee, but close if you host this yourself)

If you wish to clone this repo, remember to sync submodules after cloning it with : 
```
git submodule update --init
```
A quick note about encrytion being implemented: If you host the server on a VPS of your own, and use TLS on your website, it should enable for a fairly robust private conversation with your visitors:

[Browser] <---TLS---> [livematrix] <---e2ee---> [Matrix homeserver]  

## 📺 Working demo 


### Visitor's first interaction, creates session:
(Wait for the GIF to load...)

![demo_1](https://github.com/livematrix/.github/blob/main/images/demo01.gif?raw=true)



### Chatting with visitor, through Element client:
(Wait for the GIF to load...)

![demo_2](https://github.com/livematrix/.github/blob/main/images/demo02.gif?raw=true)



## 🛠️ How does it work? 

There would be many ways of doing a live chat, for example, we could create a Matrix account for each visitor using an SDK that allows this on the browser. You could also use your own homeserver and [allow guests](https://spec.matrix.org/latest/client-server-api/#guest-access) . Why your own? Because most homeservers disallow guests. 

The next best bet, and subjectively less cumbersome, is for you to create a new Matrix account and all visitor's chats will be mediated by it. Each new conversation is a new room. When a visitor starts a chat he/she needs to introduce their name, surname and email. If Melissa Brandon starts a chat, you'll receive an invitation on your personal account to join a room of the name `#Melissa_Brandon4212353:Matrix.org`


# 📗 How to use

##  :computer: client

Built using Svelte, so everything goes in one nice bundle, its fast and awesome to embed. No need for external libraries or big framework.js files.
Just import into your website's structure **one** JS and **one** CSS files. 

To embed the widget is very straightforward, copy the following and paste on your index.html (or other):

```html
    <script type="module" crossorigin src="/static/js/index-e73fa2f7.js"></script>
    <link rel="stylesheet" href="/static/css/index-de5f8656.css">
    <div id="app" class="tw-fixed"></div>
```

### Configure paths on client-side

On the repo's folder [**_client**](https://github.com/osousa/livematrix/tree/main/_client)  (Not the submodule), download the files, and put them inside folders relative to your index.html, like this:

```
root/ .
      |_ index.html
      |_ static/
         |_ config.json 
         |_ images/
         |   |_ bubbleicon.svg
         |_ js/
         |   |_ index-xxxxxxx.js
         |_ css/
             |_ index-xxxxxxx.css

```

Edit the config.json to set your server's addresses and ports;

#### ❔
If you can't use this directories' sctructure, change the corresponding paths on the Svelte's App and compile it again;
Follow the **submodule** named client@xxxx


## 🛰️ Server

On the repo's folder  [**_server**](https://github.com/osousa/livematrix/tree/main/_server)  (Not the submodule), download the binary "livematrix" (Only 2.1MB, small for Go standards) and the .env file; Edit the .env file and add your personal details. 
You can download the same binary and .env file from [**here**](https://github.com/livematrix/server) on *Releases* if you don't wish to clone this repo.

You must create a Matrix user specifically to mediate conversations with your visitors. The .env looks like this:

```python
# MySQL/MariaDB credentials
DATABASE_NAME=livematrix
DATABASE_USER=username
DATABASE_PASSWORD=password
DATABASE_IPADDR="127.0.0.1"
DATABASE_PORT="3306"

# Your personal Matrix.org account
MATRIX_RECIPIENT=@username:matrix.org

# Account used ONLY for mediation
MATRIX_SERVER=matrix.privex.io
MATRIX_USERNAME=@ousername:privex.io
MATRIX_PASSWORD=password

# Time to wait until next login (in days)
# Defautl: Access Token is kept for 7 days
MATRIX_TIMEOUT=7

# Leave empty if you want it to bind
# to every available interfaces. 
SERVER_IFACE="127.0.0.1"
SERVER_PORT=8000
```

```diff
- DO NOT use your personal account's password and username. CREATE one for this purpose only
```

🔥 Just fire up the server and you're done. Your live chat is working:

```
$ chmod 755 livematrix
$ ./livematrix &
```
or if you wish to suppress logs from terminal:

```
$ chmod 755 livematrix
$ ./livematrix 2>/dev/null &
```


#### ❔ 
If you do not want to use the compiled binary "livematrix" you can compile your own. 
Follow the **submodule** named server@xxxx



#### :bulb: Disclaimer 
- The server uses an ORM written by me to minimize imports (Should avoid SQLi).
- No unit tests? Skipping this as i shouldn't. Next, adding tests.
- Encryption? Nope (Currently implementing it), you can implement it if you want to. It shouldn't be hard. See [Mautrix](https://github.com/mautrix/go):
    - Use TLS and encryption. Why?  
        - [browser] <---> [your_server] <---> [matrix]

