# JSBin Code Version 1

=============================================

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

=============================================

__Note: It may be necessary to click 'Run' two
or more times in order to engage messaging.__