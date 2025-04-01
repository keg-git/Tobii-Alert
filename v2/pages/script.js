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
document.getElementById("bathroom").addEventListener("click", function() {
    sendRequest("/bathroom");
});
document.getElementById("bedtime").addEventListener("click", function() {
    sendRequest("/bedtime");
});
document.getElementById("getup").addEventListener("click", function() {
    sendRequest("/getup");
});
