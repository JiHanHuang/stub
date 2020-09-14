package v1

import (
	"html/template"
	"log"
	"strconv"
	"strings"

	"github.com/JiHanHuang/stub/pkg/logging"
	"github.com/JiHanHuang/stub/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func Echo(c *gin.Context) {
	client, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer client.Close()
	for {
		mt, message, err := client.ReadMessage()
		if err != nil {
			logging.Error("Read. ", err.Error())
			break
		}
		logging.Debug("Recv:", string(message))
		err = client.WriteMessage(mt, message)
		if err != nil {
			logging.Error("Write. ", err.Error())
			break
		}
	}
}

func Home(c *gin.Context) {
	logging.Debug("ws:", c.Request.URL.Host, c.Request.URL.Port(), "xxxx")
	isHttps := false
	if strings.Contains(c.Request.Host, strconv.Itoa(setting.ServerSetting.HttpsPort)) {
		isHttps = true
	}
	if isHttps {
		homeTemplate.Execute(c.Writer, "wss://"+c.Request.Host+"/api/v1/websocket/echo")
	} else {
		homeTemplate.Execute(c.Writer, "ws://"+c.Request.Host+"/api/v1/websocket/echo")
	}
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
