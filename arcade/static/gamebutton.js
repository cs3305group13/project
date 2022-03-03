window.addEventListener('DOMContentLoaded', init, false);

let readyButton;

let foldButton;
let callButton;
let checkButton;
let raiseButton;

let chipsTextField;

function init() {

    foldButton = document.querySelector("#fold_button");
    callButton = document.querySelector("#call_button");
    checkButton = document.querySelector("#check_button");
    raiseButton = document.querySelector("#raise_button");

    chipsTextField = document.querySelector("#amount");

    // usernameDiv = document.querySelector("#username_div");

    readyButton = document.querySelector("#ready_button");
    readyButton.addEventListener( "click", ListenReadyUpButton, false );


    foldButton = document.querySelector("#fold_button");
    foldButton.addEventListener( "click", ListenActionButton, false );

    checkButton = document.querySelector("#check_button");
    checkButton.addEventListener( "click", ListenActionButton, false );
    
    callButton = document.querySelector("#call_button");
    callButton.addEventListener( "click", ListenActionButton, false );
    
    raiseButton = document.querySelector("#raise_button");
    raiseButton.addEventListener( "click", ListenActionButton, false );

}


function ListenActionButton(event) {
    let request = "action=" + event.target.value + "&amount=" + chipsTextField.value;
    sendUserRequest(request);
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