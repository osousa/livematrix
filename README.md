# livematrix
Embedded live chat integration for your website, using matrix. 

Live chat with your website's visitors, using your Matrix client to communicate.


# 📌 What is livechat?

An *oversimplified* embedded live chat widget that allows your website's visitors to send you messages seamlessly to your Matrix account.


## 📺 Working demo 


### Visitor's first interaction, creates session:
![demo_1](https://github.com/livematrix/.github/blob/main/images/demo01.gif?raw=true)



### Chatting with visitor, through Element client:
![demo_2](https://github.com/livematrix/.github/blob/main/images/demo02.gif?raw=true)



## 🛠️ How does it work? 

There would be many ways of doing a live chat, for example, we could allow each visitor to create a Matrix account, but people might not want to go through that. You could also use your own homeserver and [allow guests](https://spec.matrix.org/latest/client-server-api/#guest-access) . Why your own? Because most homeservers disallow guests. 

The next best bet is for you to create a new Matrix account and all visitor's chats will be mediated by it. Each new conversation is a new room. When a visitor starts a chat he/she needs to introduce their name, surname and email. If Melissa Brandon starts a chat, you'll receive an invitation on your personal account to join a room of the name `#Melissa_Brandon4212353:Matrix.org`


# 📗 How to use

##  :computer: client

Built using Svelte, so everything goes in one nice bundle, its fast and awesome to embed. No need for external libraries or big framework.js files.
Just import into your website's structure one JS and one CSS files. 

To embed the widget is very straightforward, copy the following and paste on your index.html (or other):

```
    <script type="module" crossorigin src="/static/js/index-03328893.js"></script>
    <link rel="stylesheet" href="/static/css/index-0e3c6aef.css">
    <div id="app" class="tw-fixed"></div>
```

### Configure paths on client-side

On the repo's folder **"_client"**  (Not the submodule), download the files, and put them inside folders relative to your index.html, like this:

```
root/ .
      |_ index.html
      |_ static/.  
         |_ config.json 
         |_ images/
         |   |_ bubbleicon.svg
         |_ js   
         |   |_ index-xxxxxxx.js
         |_ css
             |_ index-xxxxxxx.css

```

Edit the config.json to set your server's addresses and ports;

#### ❔
If you can't use this directories' sctructure, change the corresponding paths on the Svelte's App and compile it again;
Follow the **submodule** named client@xxxx


## 🛰️ Server

On the repo's folder **"_server"**  (Not the submodule), download the binary server and the .env file; Edit the .env file and add your personal details. 

You must create a Matrix user specifically to mediate conversations with your visitors. The .env looks like this:

```
DATABASE_NAME=livechat
DATABASE_USER=osousa
DATABASE_PASSWORD=password

# Your personal Matrix.org account
MATRIX_RECIPIENT=@osousa:matrix.org

# Account used ONLY for mediation
MATRIX_USERNAME=@osousa:privex.io
MATRIX_PASSWORD=password

SERVER_PORT=8000
```

DO NOT use your personal account's password and username. CREATE a new one for this purpose only.


🔥 Just fire up the server and you're done. Your live chat is working:

```
$ chmod 755 livematrix
$ ./livematrix &
```

#### ❔ 
If you do not want to use the compiled binary "livechat" you can compile your own. 
Follow the **submodule** named server@xxxx



#### :bulb: Disclaimer 
- The server uses an ORM from [wish](https://www.wish.com/) , that i wrote to learn Go (This is my first Go project). 
- Its not performant, as uses Go's reflection a lot and i had no time to write tests, but it should avoid SQLi attacks 
- No unit tests? In my defence, i have very little spare time.
- Use at your own risk.

