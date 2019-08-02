var loc = window.location
var uri = "ws:"
if (loc.protocol === "https:") {uri = "wss:"}
uri += "//" + loc.host
uri += loc.pathname + "ws"
ws = new WebSocket(uri)

/*
function samePlayer(e) {
	if (e["player"]["name"] == e["previously"]["player"]["name"]){
		return true
	} else {
		return false
	}
}
*/

function doubleKill(e) {
	if (e["player"]["state"]["round_kills"] == 2 && e["previously"]["player"]["state"]["round_kills"] == 1){
		return true
	} else {
		return false
	}
}

$(document).ready(function () {
	$("#badge").hide();

	ws.onopen = function() {
		console.log("Connected")
		ws.send("Connected")
	}

	ws.onmessage = function(evt) {
		event = JSON.parse(evt.data)
		if(event.hasOwnProperty("previously")){
			console.log(event["previously"]["player"]["state"])
			// evaluate state changes
			if (event["round"]["phase"] == "live"){
				if (doubleKill(event)){
					$("#cheer")[0].play()
					$("#badge").slideDown("fast").delay(2000).fadeOut("fast")
				}
			}
		}
		last_event = event
	}
});