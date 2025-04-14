let ws;
let userName = "";
let roomID = "";

function getBackendURL(path = "", isWebSocket = false) {
    const protocol = isWebSocket ? "ws:" : "http:";
    const backendHost = window.location.hostname;
    const backendPort = "8082";
  
    return `${protocol}//${backendHost}:${backendPort}${path}`;
  }  

function createRoom() {
  userName = document.getElementById("name").value;
  if (!userName) return alert("Enter your name");

  fetch(getBackendURL("/create"), { method: "POST" })
    .then(res => res.json())
    .then(data => {
      roomID = data.roomID;
      document.getElementById("roomID").value = roomID;
      connectWebSocket();
    })
    .catch(err => {
      console.error("Error creating room:", err);
    });
}

function joinRoom() {
  userName = document.getElementById("name").value;
  roomID = document.getElementById("roomID").value;
  if (!userName || !roomID) return alert("Enter name and room ID");

  connectWebSocket();
}

function connectWebSocket() {
  ws = new WebSocket(getBackendURL(`/join?roomID=${roomID}`, true));

  ws.onopen = () => {
    ws.send(JSON.stringify({ type: "name", data: userName }));
    document.getElementById("chat").style.display = "block";
    document.getElementById("form-section").style.display = "none";
    document.getElementById("currentRoom").innerText = roomID;
  };

  ws.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    const div = document.createElement("div");
    div.textContent = `${msg.name}: ${msg.data}`;
    document.getElementById("messages").appendChild(div);
  };

  ws.onerror = (error) => {
    console.error("WebSocket error:", error);
  };

  ws.onclose = () => {
    alert("Connection closed");
  };
}

function sendMessage() {
  const input = document.getElementById("messageInput");
  const message = input.value;
  if (!message) return;

  ws.send(JSON.stringify({ data: message }));
  input.value = "";
}
