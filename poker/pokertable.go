<!DOCTYPE html>
<html>
    <head>
        <title>Protected Page</title>
        <link rel="stylesheet" type="text/css" href="/static/css/main.css" />
        <script src="/static/js/getcontent.js"></script>
        <script src="/static/js/gamebuttons.js"></script>
    </head>
    <body>

        
        <h1>Lobby Code: {{.TableID}}</h1>
        <h1 hidden id="hidden_username_tag" >{{.Username}}</h1>
        <h1 hidden id="hidden_seatnumber_tag" >{{.SeatNumber}}</h1>
        <h1 hidden id="hidden_funds_tag">{{.Funds}}</h1>

        <ul>
            <li>
                <span id="money_in_pot" ></span>
            </li>
            <li>
                <span id="side_pot" ></span>
            </li>
            <li>
                <span id="current_player_making_move" ></span>
            </li>
            <li>
                <span id="community_cards"></span>
            </li>
        </ul>

        
        <ul id="player_list">
            <li id="seat_1" >
                <ul><li> Seat 1:</li>
                    <li id="username_1"></li>
                    <li id="funds_1"></li>
                    <li id="state_1"></li>
                    <li id="money_in_pot_1"></li>
                    <li id="cards_1"></li>
                </ul>
            </li>
            <li id="seat_2" >
                <ul><li> Seat 2:</li>
                    <li id="username_2"></li>
                    <li id="funds_2"></li>
                    <li id="state_2"></li>
                    <li id="money_in_pot_2"></li>
                    <li id="cards_2"></li>
                </ul>
            </li>
            <li id="seat_3" >
                <ul><li> Seat 3:</li>
                    <li id="username_3"></li>
                    <li id="funds_3"></li>
                    <li id="state_3"></li>
                    <li id="money_in_pot_3"></li>
                    <li id="cards_3"></li>
                </ul>
            </li>
            <li id="seat_4" >
                <ul><li> Seat 4:</li>
                    <li id="username_4"></li>
                    <li id="funds_4"></li>
                    <li id="state_4"></li>
                    <li id="money_in_pot_4"></li>
                    <li id="cards_4"></li>
                </ul>
            </li>
            <li id="seat_5" >
                <ul><li> Seat 5:</li>
                    <li id="username_5"></li>
                    <li id="funds_5"></li>
                    <li id="state_5"></li>
                    <li id="money_in_pot_5"></li>
                    <li id="cards_5"></li>
                </ul>
            </li>
            <li id="seat_6" >
                <ul><li> Seat 6:</li>
                    <li id="username_6"></li>
                    <li id="funds_6"></li>
                    <li id="state_6"></li>
                    <li id="money_in_pot_6"></li>
                    <li id="cards_6"></li>
                </ul>
            </li>
            <li id="seat_7" >
                <ul><li> Seat 7:</li>
                    <li id="username_7"></li>
                    <li id="funds_7"></li>
                    <li id="state_7"></li>
                    <li id="money_in_pot_7"></li>
                    <li id="cards_7"></li>
                </ul>
            </li>
            <li id="seat_8" >
                <ul><li> Seat 8:</li>
                    <li id="username_8"></li>
                    <li id="funds_8"></li>
                    <li id="state_8"></li>
                    <li id="money_in_pot_8"></li>
                    <li id="cards_8"></li>
                </ul>
            </li>
        </ul>

        <span id="general_message_span"></span>
        <span id="problem_message_span"></span>

        <form id="player_state_button_form" >
            <fieldset>
                <input type="button" id="ready_button" value="Ready Up!" />
            </fieldset>
        </form>

        <form id="game_buttons_form" >
            <fieldset>
                <input type="button" id="fold_button" value="Fold" />
                
                <input type="button" id="check_button" value="Check" />
                <input type="button" id="call_button" value="Call" />

                <input type="button" id="raise_button" value="Raise" />
                
                <!-- value in label will alternate between Call and Raise depending on chips -->
                <label for="amount" >Amount:</label>
                <input type="text" id="amount" name="amount" />
            </fieldset>
        </form>

    </body>
</html>
