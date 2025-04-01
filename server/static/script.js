function sendRequest(endpoint) {
    fetch(endpoint)
}

document.getElementById("alert").addEventListener("click", function() {
    sendRequest("/alert");
});
document.getElementById("suction").addEventListener("click", function() {
    sendRequest("/suction");
});
document.getElementById("readjust").addEventListener("click", function() {
    sendRequest("/readjust");
});
document.getElementById("pee").addEventListener("click", function() {
    sendRequest("/pee");
});
document.getElementById("bedtime").addEventListener("click", function() {
    sendRequest("/bedtime");
});
document.getElementById("chair").addEventListener("click", function() {
    sendRequest("/chair");
});
document.getElementById("bathroom").addEventListener("click", function() {
    sendRequest("/bathroom");
});
document.getElementById("getup").addEventListener("click", function() {
    sendRequest("/getup");
});
document.getElementById("bed").addEventListener("click", function() {
    sendRequest("/bed");
});
document.getElementById("cancel").addEventListener("click", function() {
    sendRequest("/shutoff");
});
