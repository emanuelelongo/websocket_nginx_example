window.onload = function () {
    const output = document.getElementById("output");
    const conn = new WebSocket(`ws://${document.location.host}/ws`);

    conn.onclose = function () {
        var item = document.createElement("div");

        if (output.innerText == "") {
            item.innerHTML = "<b>Can't connect 🤔</b>";
        }
        else {
            item.innerHTML = "<b>Bye 😀</b>";
        }

        output.appendChild(item);
    };

    conn.onmessage = function (event) {
        var item = document.createElement("div");
        item.innerText = event.data;
        output.appendChild(item);
    };
};