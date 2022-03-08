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

}

function startGame() {
    createDeck();

    let shuffleTimes = 52;
    shuffle(shuffleTimes);
    players();
    // randomCard();
    dealHands();
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

function shuffle(timesShufflfed) {
    for(let i = 0; i < timesShufflfed; i++) {
        let rndNo = getRandomInt(1,52);
        let card = deck[i];

        deck[i] = deck[rndNo];
        deck[rndNo] = card;
    }
}


playersList = new Array();

// function players() {
//     let player1 =  document.getElementById('user1').value;
//     let player2 =  document.getElementById('user2').value;
//     let player3 =  document.getElementById('user3').value;
//     let player4 =  document.getElementById('user4').value;
//     if( player1 != "" || player2 != "" || player3 != "" || player4 != "") {
//         playersList.push(player1)
//         playersList.push(player2)
//         playersList.push(player3)
//         playersList.push(player4)
//     }
//     console.log("User 1: " + player1)
//     console.log("User 2: " + player2)
//     console.log("User 3: " + player3)
//     console.log("User 4: " + player4)
// }

function players() {
    let player1 =  document.getElementById('user1').value;
    if( player1 != "") {
        playersList.push(player1)
    console.log("User 1: " + player1)
    }
}



// function deal() {
//     let hand_div = document.createElement('div');
//     let points_div = document.createElement('div');
//     points_div.className = "points";
//     hand_div.id = "hand";
//     let playerHand = [randomCard(deck), randomCard(deck)]
//     let dealerHand = [randomCard(deck), randomCard(deck)];
//     console.log(playerHand)
//     console.log(dealerHand)
// }
    
