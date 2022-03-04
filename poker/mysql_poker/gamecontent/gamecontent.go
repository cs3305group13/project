window.addEventListener('DOMContentLoaded', init, false);

// Returns a Promise that resolves after "ms" Milliseconds
const timer = ms => new Promise(res => setTimeout(res, ms));

let contentURL;
let gameContentRequest;

let MAX_NUMBER_OF_PLAYERS = 8;

function init() {
    contentURL = 'content_request';
    gameContentRequest = new XMLHttpRequest();
    gameContentRequest.addEventListener('readystatechange', handleContentResponse, false);
    sendRequests();
}

async function sendRequests() {
    while ( true ) {
        gameContentRequest.open('GET', contentURL, true);
        gameContentRequest.send(null);
        await timer(2000); // then the created Promise can be awaited
    }
}

function handleContentResponse() {
    if (this.readyState == 4 && this.status == 200) {
        let gameContent = gameContentRequest.responseText;
        insertGameContent(gameContent);
    }
}


function insertGameContent( gameContent ) {
    var content = JSON.parse(gameContent);

    insertPlayersIntoHTML(content.Players);
    insertDetailsIntoHTML(content.TableDetails);
}


let usernameTAG;
let fundsTAG;
let stateTAG;
let moneyInPotTAG;
let cardsTAG;

let username;
let funds;
// let seatNumber;
let playerState;
let moneyInPot;
let cards;

function insertPlayersIntoHTML( players ) {
    for (let i=1; i<=players.length; i++) {
        seatTAG = document.querySelector("#seat_" + i);

        usernameTAG = document.querySelector("#username_" + i);
        fundsTAG = document.querySelector("#funds_" + i);
        stateTAG = document.querySelector("#state_" + i);
        moneyInPotTAG = document.querySelector("#money_in_pot_" + i);
        cardsTAG = document.querySelector("#cards_" + i);

        username = players[i-1].Username;
        funds = players[i-1].Funds;
        // seatNumber = players[i-1].SeatNumber;
        playerState = players[i-1].PlayerState;
        moneyInPot = players[i-1].MoneyInPot;
        cards = players[i-1].Cards;

        currentPlayerMakingMoveTAG = document.querySelector("#current_player_making_move");

        if ( detectRefresh() ) {
            continue;
        }

        usernameTAG.innerHTML = username;
        fundsTAG.innerHTML = funds;
        stateTAG.innerHTML = playerState;
        moneyInPotTAG.innerHTML = moneyInPot;
        cardsTAG.innerHTML = cards;
    }
}

// function checks if player list html should refresh
function detectRefresh() {
    if ( usernameTAG.innerHTML !== username ) {
        return false;
    }
    if ( fundsTAG.innerHTML !== funds ) {
        return false;
    }
    if ( stateTAG.innerHTML !== playerState ) {
        return false;
    }
    if ( moneyInPotTAG.innerHTML !== moneyInPot ) {
        return false;
    }
    if ( cardsTAG.innerHTML !== cards ) {
        return false;
    }
    
    return true;
}

function insertDetailsIntoHTML( tableDetails ) {
    communityCardsTAG = document.querySelector("#community_cards");
    currentPlayerMakingMoveTAG = document.querySelector("#current_player_making_move");
    moneyInPotTAG = document.querySelector("#money_in_pot")

    communityCards = tableDetails.CommunityCards;
    currentPlayerMakingMove = tableDetails.CurrentPlayerMakingMove;
    moneyInPot = tableDetails.MoneyInPot;

    communityCardsTAG.innerHTML = communityCards;
    currentPlayerMakingMoveTAG.innerHTML = currentPlayerMakingMove;
    moneyInPotTAG.innerHTML = moneyInPot;

    hiddenUsernameTAG = document.querySelector("#hidden_username_tag");
    hiddenSeatNumberTAG = document.querySelector("#hidden_seatnumber_tag");

    SeatTAG = document.querySelector("#seat_" + hiddenSeatNumberTAG.innerHTML);


    let gameButtonsFormTAG = document.querySelector("#game_buttons_form");
    if ( hiddenUsernameTAG.innerHTML === currentPlayerMakingMoveTAG.innerHTML ) {
        SeatTAG.style.backgroundColor = "cadetblue";
        gameButtonsFormTAG.style.display = "block";
        
    } else {
        SeatTAG.style.backgroundColor = "blue";
        gameButtonsFormTAG.style.display = "none";
    }

    
    gameState = tableDetails.GameState;
    let readyButtonFormTAG = document.querySelector("#player_state_button_form");
    if (gameState == "1" ) { // aka. game is in progress
        readyButtonFormTAG.style.display = "none";

    } else {
        readyButtonFormTAG.style.display = "block";
    }
}
