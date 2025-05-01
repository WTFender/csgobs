# CSGOBS
CSGO gamestate integration for OBS.

## Setup
- Save [gamestate_integration_csgobs.cfg](gamestate_integration_csgobs.cfg) to your cfg directory:
  - Default:&nbsp;&nbsp;&nbsp;`C:\Program Files (x86)\Steam\steamapps\common\Counter-Strike Global Offensive\csgo\cfg`
- Run `server.exe`.
  - (optional) Run `server_dynamic.exe` if you want to modify or add files to `static/`.
- Add an OBS browser source:
  - URL:&nbsp;&nbsp;&nbsp;`http://127.0.0.1:3000`

## Multi-PC
- Run server.exe on your gaming PC.
- Set your OBS browser source URL to the private IP address of your gaming PC.  
  - Ex:&nbsp;&nbsp;&nbsp;`http://192.168.0.3:3000`

## Demo
In-game event (double kill) triggers animation and sound.  

https://github.com/user-attachments/assets/b10b4a4c-9018-4b51-ba2b-9fd834f28345
