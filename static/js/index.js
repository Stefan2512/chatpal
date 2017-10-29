function onLoad() {
    document.getElementsByClassName("fa-send")[0].addEventListener("click", requestBotMessage);
    document.getElementsByClassName("fa-send")[0].addEventListener("click", composeMessage.bind(null, "sender", null))
}

function composeMessage(actor, receiverText) {
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

    var text = receiverText;
    console.log('receiver text', receiverText);
    if (!receiverText) {
        text = document.getElementById("comment").value;
    }

    actorText.innerText = text;

    var time = document.createElement("SPAN");
    time.className = "message-time pull-right";
    time.innerText = "Sun";

    actorSubElement.appendChild(actorText);
    actorSubElement.appendChild(time);
    actorElement.appendChild(actorSubElement);
    messageBody.appendChild(actorElement);
    conversation.appendChild(messageBody);

    document.getElementById('conversation').lastChild.scrollIntoView(true);
}

function requestBotMessage() {
    var text = document.getElementById("comment").value;

    $.ajax({
        type: "POST",
        url: "http://localhost:4419/message",
        dataType: "json",
        data: JSON.stringify({"MessageID": "1231231", "Text": text})
    }).done(function (response) {
        composeMessage("receiver", response.Text)
    })
}

document.addEventListener("DOMContentLoaded", onLoad, false);