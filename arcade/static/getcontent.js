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


function insertPlayersIntoHTML( players ) {
    for (let i=1; i<=players.length; i++) {
        let usernameTAG = document.querySelector("#username_" + i);
        let fundsTAG = document.querySelector("#funds_" + i);
        let stateTAG = document.querySelector("#state_" + i);
        let moneyInPotTAG = document.querySelector("#money_in_pot_" + i);
        let cardsTAG = document.querySelector("#cards_" + i);

        let username = players[i-1].Username;
        let funds = players[i-1].Funds;
        // let seatNumber = players[i-1].SeatNumber;
        let playerState = players[i-1].PlayerState;
        let moneyInPot = players[i-1].MoneyInPot;
        let cards = players[i-1].Cards;

        if ( usernameTAG.innerHTML === username ) {
            continue;
        }
        usernameTAG.innerHTML = username;
        fundsTAG.innerHTML = funds;
        stateTAG.innerHTML = playerState;
        moneyInPotTAG.innerHTML = moneyInPot;
        cardsTAG.innerHTML = cards
    }
}

function insertDetailsIntoHTML( tableDetails ) {
    communityCardsTAG = document.querySelector("#community_cards");
    currentPlayerMakingMoveTAG = document.querySelector("#current_player_making_move");
    gameStateTAG = document.querySelector("#game_state");

    communityCards = tableDetails.CommunityCards;
    currentPlayerMakingMove = tableDetails.CurrentPlayerMakingMove;
    gameState = tableDetails.GameState;

    communityCardsTAG.innerHTML = communityCards;
    currentPlayerMakingMoveTAG.innerHTML = currentPlayerMakingMove;
    gameState.innerHTML = gameState;
}