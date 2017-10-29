function onLoad() {
    document.getElementsByClassName("fa-send")[0].addEventListener("click", composeMessage.bind(this, "receiver", "this"))
}

function composeMessage(actor, text) {
    var conversation = document.getElementById("conversation");

    var messageBody = document.createElement("DIV");
    messageBody.className = "row message-body";

    var actorElement = document.createElement("DIV");
    actorElement.className = "col-sm-12 message-main-receiver";
    if (actor === "sender") {
        actorElement.className = "col-sm-12 message-main-sender";
    }

    var actorSubElement = document.createElement("DIV");
    actorSubElement.className = "receiver";
    if (actor === "sender") {
        actorSubElement.className = "sender";
    }

    var actorText = document.createElement("DIV");
    actorText.className = "message-text";
    actorText.innerText = text;

    var time = document.createElement("SPAN");
    time.className = "message-time pull-right";
    time.innerText = "Sun";

    actorSubElement.appendChild(actorText);
    actorSubElement.appendChild(time);
    actorElement.appendChild(actorSubElement);
    messageBody.appendChild(actorElement);
    conversation.appendChild(messageBody);
}

document.addEventListener("DOMContentLoaded", onLoad, false);