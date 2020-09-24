var app = new Vue({
    el: '#app',
    data: {
      ws: null,
      serverUrl: "ws://localhost:8080/ws",
      messages: [],
      newMessage: ""
    },
    mounted: function() {
      this.connectToWebsocket();
    },
    methods: {
      connectToWebsocket() {
        this.ws = new WebSocket( this.serverUrl );
        this.ws.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
        this.ws.addEventListener('message', (event) => { this.handleNewMessage(event) });
      },
      onWebsocketOpen() {
        console.log("connected to WS!");        
      },
      handleNewMessage(event) {
        let data = event.data;
        data = data.split(/\r?\n/);

        for (let i = 0; i < data.length; i++) {
            let msg = JSON.parse(data[i]);
            this.messages.push(msg);

        }   
      },
      sendMessage() {
        if(this.newMessage !== "") {
          this.ws.send(JSON.stringify({message: this.newMessage}));
          this.newMessage = "";
        }
      }

    }
  })
