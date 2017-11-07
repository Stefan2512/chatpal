function onLoad() {
    document.getElementsByClassName("fa-send")[0].addEventListener("click", handleClick);
    document.getElementById("comment").addEventListener("keypress", handleEnter)
}

function handleEnter(e) {
    if (e.key === "Enter") {
        sendMessage()
    }
}

function handleClick(e) {
    sendMessage()
}

function sendMessage() {
    var comment = document.getElementById("comment");
    composeMessage("sender", comment.value);
    requestBotMessage();
    comment.value = "";
    comment.focus()
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
    if (!receiverText) {
        text = document.getElementById("comment").value;
    }

    actorText.innerText = text;

    var time = document.createElement("SPAN");
    time.className = "message-time pull-right";

    time.innerText = formatMessageTime();

    actorSubElement.appendChild(actorText);
    actorSubElement.appendChild(time);
    actorElement.appendChild(actorSubElement);
    messageBody.appendChild(actorElement);
    conversation.appendChild(messageBody);

    document.getElementById('conversation').lastChild.scrollIntoView(true);
}

function requestBotMessage() {
    var comment = document.getElementById("comment");

    $.ajax({
        type: "POST",
        url: "http://localhost:4419/message",
        dataType: "json",
        data: JSON.stringify({"Text": comment.value})
    }).done(function (response) {
        composeMessage("receiver", response.Text);
        comment.value = ""
    })
}

document.addEventListener("DOMContentLoaded", onLoad, false);

// --------------------------- helpers -------------------------------------

function formatMessageTime() {
    var days = ['Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday'];

    var now = new Date();
    var day = days[now.getDay()];
    var hours = now.getHours();
    var minutes = now.getMinutes();

    return day + " " + hours + ":" + minutes;
}