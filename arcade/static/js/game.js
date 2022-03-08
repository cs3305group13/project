// Waits for html and css to load before executing startGame()
window.addEventListener('DOMContentLoaded', init, false);

let suits = ["s", "h", "d", "c"];
let facevalues = ["2", "3", "4", "5", "6", "7", "8", "9", "10", "j", "q", "k", "a"];
let deck = new Array();
let dealerCards = new Array();
let dealerScore = new Array();
let playerCards = new Array();
let playerScore = new Array();

function init() {

    // find by id #btnStart button and attach it an event listener
    let btnStart = document.querySelector("#btnStart");
    btnStart.addEventListener( "click", startGame, false );
    btnHit.addEventListener( "click", Hit, false );
    btnStand.addEventListener( "click", Stand, false );
}

function startGame() {  
    createDeck();
    shuffle();
    players();
    deal();
}

function getRandomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function createDeck() {
    deck = new Array(); // reset deck

    for (let i = 0; i < facevalues.length; i++) {
        for (let x = 0; x < suits.length; x++) {
            let weight = parseInt(facevalues[i]);
            if (facevalues[i] == "j" || facevalues[i] == "q" || facevalues[i] == "k")
                weight = 10;
            if (facevalues[i] == "a")
                weight = 11 || 1;
            let card = { Value: facevalues[i], Suit: suits[x], Weight: weight };
            deck.push(card);
        }
    }
}

function shuffle() {
    for(let i = 0; i < 52; i++) {
        let rndNo = getRandomInt(1,52);
        let card = deck[i];

        deck[i] = deck[rndNo];
        deck[rndNo] = card;
    }
    console.log(deck)
}

playersList = new Array();


function players() {
    let player1 =  document.getElementById('user1').value;
    if( player1 != "") {
        playersList.push(player1)
    console.log("User 1: " + player1)
    }
}

function deal() {
    playerCards = [nextCard(), nextCard()]
    dealerCards = [nextCard(), nextCard()]
}

function nextCard() {
    return deck.shift();
}

function Hit() {

}

function Stand() {

}