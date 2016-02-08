# JSBin Code Version 2

============================================

let ws = new WebSocket('ws://localhost:4000')

let message = {
  name: 'channel add',
  data: {
    name: 'Hardware Support'
  }
};

ws.onopen = () => {
  ws.send(JSON.stringify(message))
}

ws.onmessage = (e) => {
  console.log(JSON.parse(e.data));
}

============================================

__Note: It may be necessary to click 'Run' two
or more times in order to engage messaging.__