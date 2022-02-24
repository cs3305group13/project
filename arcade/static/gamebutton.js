window.addEventListener('DOMContentLoaded', init, false);

let readyButton;

let currentPlayerMakingMove;

let foldButton;
let callButton;
let checkButton;
let raiseButton;

let chipsTextField;

// let usernameDiv;
// let seatNumberDiv;

function init() {

    foldButton = document.querySelector("#fold_button");
    callButton = document.querySelector("#call_button");
    checkButton = document.querySelector("#check_button");
    raiseButton = document.querySelector("#raise_button");

    chipsTextField = document.querySelector("chips_text_field");

    // usernameDiv = document.querySelector("#username_div");
    currentPlayerMakingMove = document.querySelector("#current_player_making_move");

    readyButton = document.querySelector("#ready_button");
    readyButton.addEventListener( "click", ListenReadyUpButton, false );
}

function ListenReadyUpButton() {
    sendUserRequest( "ready_button=true" )
}

let userRequestURL;
let userRequest;
// Function sends request to url `/user_request` with the attached request form fields. example: "ready_button=true"
function sendUserRequest( requestFormFields ) {
    userRequestURL = '/user_request?'+requestFormFields;
    userRequest = new XMLHttpRequest();
    userRequest.addEventListener('readystatechange', handleUserRequestResponse, false);
    userRequest.open('GET', userRequestURL, true);
    userRequest.send();
}

let generalMessageSpan;
let problemMessageSpan;
function handleUserRequestResponse() {
    if (this.readyState == 4 && this.status == 200) {

        let response = userRequest.responseText;
        message = response.split("\n");
        if (response.includes("PROBLEM:")) {
            problemMessageSpan = document.querySelector("#problem_message_span");
            problemMessageSpan.innerHTML = message[1];
        } 
        else if (response.includes("MESSAGE:")) {
            generalMessageSpan = document.querySelector("#general_message_span");
            generalMessageSpan.innerHTML = message[1];
        }
            
    }
}