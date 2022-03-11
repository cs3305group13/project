// Waits for html and css to load before executing init()
window.addEventListener('DOMContentLoaded', init, false);

let suits = ["s", "h", "d", "c"];
let facevalues = ["2", "3", "4", "5", "6", "7", "8", "9", "10", "j", "q", "k", "a"];
let deck = new Array();
let dealerCards = new Array();
let dealerScore = 0;
let playerCards = new Array();
let playerScore = 0;
function init() {

    // find by id #btnStart button and attach it an event listener
    let btnStart = document.querySelector("#btnStart");
    btnStart.addEventListener( "click", startGame, false );
    btnHit.addEventListener( "click", Hit, false );
    btnStand.addEventListener( "click", Stand, false );
}

function startGame() {
    // function to initialize the game  
    createDeck();
    shuffle(); 
    deal();
}

function getRandomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function createDeck() {
    // function that creates and assignes weights to cards in the deck
    deck = new Array(); // reset deck

    for (let i = 0; i < facevalues.length; i++) {
        for (let x = 0; x < suits.length; x++) {
            let weight = parseInt(facevalues[i]);
            if (facevalues[i] == "j" || facevalues[i] == "q" || facevalues[i] == "k")
                weight = 10;
            if (facevalues[i] == "a")
                weight = 11;
            let card = { Value: facevalues[i], Suit: suits[x], Weight: weight };
            deck.push(card);
        }
    }
}

function shuffle() {
    // function that shuffles the created deck by swapping a card at a given 
    // index with a random card in the deck
    for(let i = 0; i < 52; i++) {
        let rndNo = getRandomInt(1,52);
        let card = deck[i];
        
        deck[i] = deck[rndNo];
        deck[rndNo] = card;
    }
}

function deal() {
    // function that initially deals the player 2 cards, and the dealer one.
    playerCards = [nextCard(), nextCard()]
    dealerCards = [nextCard()]
    Score();
}

function nextCard() {
    // function that returns the next available card from the deck and removes it from the deck
    return deck.shift();
}

function Hit() {
    // function that deals the player another card and recalculates their score
    playerScore = 0;
    playerCards.push(deck.shift());
    Score();
}

function Stand() {
    // function that deals the dealer another card
    if (dealerScore < 21)
        dealerCards.push(deck.shift());
        Score();
} 

function Score() {
    // logic of blackjack implemented 
    playerScore = 0;
    for(i = 0; i < playerCards.length; i ++) {
        playerScore += playerCards[i].Weight;
        console.log(playerCards);
        console.log("Player Score: " + playerScore);
    }

    dealerScore = 0;
    for(i = 0; i < dealerCards.length; i ++) {
        dealerScore += dealerCards[i].Weight;
        console.log(dealerCards);
        console.log("Dealer Score: " + dealerScore);
    }
    
    if (playerScore > 21) {
        // change the value of "A" from 11 if the score is over 21 to 1.
        for (i = 0; i < playerCards.length; i ++) {
            if (playerCards[i].Value == "a" && playerCards[i].Weight == 11) {
                playerCards[i].Weight == 1;
                playerScore += 1
            } else {
                console.log("Dealer wins!")
            }
        }
    } if (dealerScore > 21) {
        console.log("Player wins")
    } if (playerScore == dealerScore) {
        console.log("Tie!")
    } if (dealerScore == 21 && playerScore != 21) {
        console.log("Dealer wins")
    } if (playerScore == 21 && dealerScore != 21) {
        console.log("Player wins")
    }
}