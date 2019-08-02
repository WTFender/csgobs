# CSGOBS
CSGO gamestate integration for OBS.

## Setup
- Save [gamestate_integration_csgobs.cfg](gamestate_integration_csgobs.cfg) to your cfg directory:
  - Default:&nbsp;&nbsp;&nbsp;`C:\Program Files (x86)\Steam\steamapps\common\Counter-Strike Global Offensive\csgo\cfg`
- Run server.exe.
- Add an OBS browser source:
  - URL:&nbsp;&nbsp;&nbsp;`http://127.0.0.1:3000`

## Multi-PC
- Run server.exe on your gaming PC.
- Set your OBS browser source URL to the private IP address of your gaming PC.  
  - Ex:&nbsp;&nbsp;&nbsp;`http://192.168.0.3:3000`

## Demo
[![](https://imgur.com/bFf8eyF.png)](https://gfycat.com/aggressivesafechrysomelid)

## Notes
A golang webserver accepts CSGO's gamestate payloads and forwards the JSON via websocket to a client browser. Javascript handles the event logic and presentation. Tried using AJAX to poll the webserver for gamestate data, but you are stuck on an interval which leads to an obvious on-screen delay. You could try polling at 10ms, but that doesn't scale well to even a home network. If events arrive too quickly, they need to be queued and then batch processed, which just leads to more delay. Using websockets was much better; I had difficulty even measuring the time it took an event to be received.