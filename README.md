# **Arcade Game**

*A visually simple arcade lobby with free **online poker** and **side-scroller** games.*

---

## **How to run project:**
- ### Running the arcade lobby
      $ cd arcade
      arcade$ go run handler.go

    ***Reminder:*** Read & fill in 'DUMMY.env' file with necessary credentials 
---

## **Section 1 - Contained Packages**
These are discussed in more detail in later sections

- **arcade** *package*:
  * Incoming request handling
  * User login, sign up and user page validation 
  * Distributes html, css & js
  
- **poker** *package*:
  * User request handling & game logic
  
- **sidescroller** *package*:
  * User request handling & game logic
  
- **mysql_db** *package*:
  * Establishes database connections for sql queries
  * Offers useful mysql queries

- **testing** *package*:
  * **mysql_poker**
    - Contains necessary mysql exec queries for refreshing database
      tables necessary for testing
  * **utils**
    - Contains dummy http requests and responses

- **utils** *package* contains:
   * Useful helper methods
   
   * **env** *package* contains:
     - Methods which help retrieve environment variables
   * **token** *package* contains:
     - Methods which handle creating and validating json tokens

- **cards** *package* contains:
   * Implements card deck logic 

---

## **Section 2 - Words of Wisdom**

 - ### Understanding the poker package
    * **Overview**
      * Due to the large quantity of methods necessary for the game to work, we (as expected) subdivided the package into smaller sub-packages based on their usage.  
      The top level files (**ajax_request.go**, **pokertable.go**, and **user_request.go**) are used for transitioning from http request handling into sql handling. (ex. user sends http request containing form input through `"/user_request"` url this would be directed to user_request.go and then here the users form input `'action=Fold'` would be directed to a specific method in the sub package mysql_poker where all necessary sql querying and execution would take place.)  

    * **Poker game content retrieval**
      * The game content is retrieved by the users browser via javascript AJAX requests sent under the url path `'/content_request'`, this simply pings the server with a blank http.Request which must have a non-expired session token, once it's confirmed that the token is valid, then the necessary content is queried and sent back to the user as a JSON object. The JSON object is then decoded and the inclosed details would be inserted into their appropriate html locations. These would include general game details such as the community cards, the name of the current player making move, and the money in the pot. Also player specific details would include usernames, their funds and other details like their cards, seat number and the action they may have taken (ie. folded, raised, called or checked)  

    * **Poker lobby join and setup**
      * The game setup or join mechanism is triggered in the implementation by filling the necessary form found in **userpage.go** in arcade package. Once this form is filled, it's form handling is directed to **pokertable.go**, from here it is decided whether the user wants to create a game or join one. Both of these actions require multiple sql queries which depend on each other, due to this fact we had to wrap their execution in sql transactions.

---

## **Section 3 - Key Features**

* When a poker game is created its members are monitored and kicked if idle, this 
  process is active as long as their are players not idle. *more details can be found in gameobserver package*
* Poker hand evaluation is fully functioning thanks to Che-Hsun Liu's package [github.com/chehsunliu/poker](https://github.com/chehsunliu/poker)

---

## **Section 4 - Features Needing Fixing**
  * ### **Things to keep in mind**
    * Currently there is no tie game functionality in poker which would split the pot 
      accordingly
    * An All-In player still may win entire pot even if they didn't match a betting round.
    * Deadlocks still may occur in grey areas such as during start game sequence. (Basically any time a user performs an action and quickly they or someone else performs another action, this may trigger a deadlock due to a transaction still being executed, very unlikely though.)
  
  * ### **Bugs**
    * During testing it was seen users using a Safari browser would be kicked for being idle whereas browsers like Google Chrome, Brave Browser and FireFox would be unaffected.

--- 

## **Other Details**
* ### **Packages Used**
  * Poker Hand Evaluation thanks to [github.com/chehsunliu/poker](https://github.com/chehsunliu/poker)

  * JSON Tokens: [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)

  * Environment Variable Retrieval: [github.com/joho/godotenv](https://github.com/joho/godotenv)

  * Cryptographic password hashing: [golang.org/x/crypto/scrypt](https://golang.org/x/crypto/scrypt)

---

## **Example Poker Game**

* #### **Playing Poker Game Example**
  https://user-images.githubusercontent.com/78961144/156378791-dda22416-b768-41a5-91d0-56d55de20fd9.mp4

* #### **Leaving & Rejoining Game**
  

https://user-images.githubusercontent.com/78961144/156379189-d2a50502-76e7-4883-ad05-24c5a7aae25d.mp4

