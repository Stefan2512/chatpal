// add listeners when the app is loading
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

// get the text value from the ui and make request for the response message
function sendMessage() {
    var comment = document.getElementById("comment");
    composeMessage("sender", comment.value);
    requestBotMessage();
    comment.value = "";
    comment.focus()
}

// compose and append messages to be displayed on page
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

    actorText.innerText = receiverText;

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

// ajax request for message
function requestBotMessage() {
    $.ajax({
        type: "POST",
        url: "http://localhost:4419/message",
        dataType: "json",
        data: JSON.stringify({"Text": comment.value})
    }).done(function (response) {
        // after the response message from server is received, compose and render it on page
        composeMessage("receiver", response.Text);

        var comment = document.getElementById("comment");
        comment.value = ""
    })
}

document.addEventListener("DOMContentLoaded", onLoad, false);

// format the time to be displayed under the comment texts
function formatMessageTime() {
    var days = ['Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday'];

    var now = new Date();
    var day = days[now.getDay()];
    var hours = now.getHours();
    var minutes = now.getMinutes();

    return day + " " + hours + ":" + minutes;
}