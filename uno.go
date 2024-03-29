package main

// Import used packages
import (
	"fmt"
	"math/rand"
	"time"
)

// Rules
/*
var (
	O7        bool = false
	JumpIn    bool = false
	Challenge bool = false
)
*/

// This is a struct that shows the schematics of what a card should be built on.
type card struct {
	color       string
	name        string
	amountCards int
}

type handCard struct {
	color string
	name  string
}

// This is a struct that individual player data
type playerChar struct {
	name         string
	amountInHand int
	cardsInHand  [30]handCard
	isUNO        bool
}

// Create players
var player playerChar
var computer1 playerChar
var computer2 playerChar
var computer3 playerChar

// This is our deck, here's where all the cards in the UNO game is stored.
var deckSchematic = []card{

	/*
		Cards in a UNO deck

		To explain the structure

		{"Color", "Type", amount}
	*/

	// Black/Universal cards
	{"Black", "Wildcard", 4}, {"Black", "Plus4", 4},

	// Zero cards
	{"Red", "Zero", 1}, {"Yellow", "Zero", 1}, {"Green", "Zero", 1}, {"Blue", "Zero", 1},

	// +2 cards
	{"Red", "Plus2", 2}, {"Yellow", "Plus2", 2}, {"Green", "Plus2", 2}, {"Blue", "Plus2", 2},

	// Reverse cards
	{"Red", "Reverse", 2}, {"Yellow", "Reverse", 2}, {"Green", "Reverse", 2}, {"Blue", "Reverse", 2},

	// Block cards
	{"Red", "Block", 2}, {"Yellow", "Block", 2}, {"Green", "Block", 2}, {"Blue", "Block", 2},

	// Red cards
	{"Red", "One", 2}, {"Red", "Two", 2}, {"Red", "Three", 2}, {"Red", "Four", 2}, {"Red", "Five", 2}, {"Red", "Six", 2}, {"Red", "Seven", 2}, {"Red", "Eight", 2}, {"Red", "Nine", 2},

	// Yellow cards
	{"Yellow", "One", 2}, {"Yellow", "Two", 2}, {"Yellow", "Three", 2}, {"Yellow", "Four", 2}, {"Yellow", "Five", 2}, {"Yellow", "Six", 2}, {"Yellow", "Seven", 2}, {"Yellow", "Eight", 2}, {"Yellow", "Nine", 2},

	// Green cards
	{"Green", "One", 2}, {"Green", "One", 2}, {"Green", "One", 2}, {"Green", "One", 2}, {"Green", "Five", 2}, {"Green", "Six", 2}, {"Green", "Seven", 2}, {"Green", "Eight", 2}, {"Green", "Nine", 2},

	// Blue cards
	{"Red", "One", 2}, {"Red", "Two", 2}, {"Red", "Three", 2}, {"Red", "Four", 2}, {"Red", "Five", 2}, {"Red", "Six", 2}, {"Red", "Seven", 2}, {"Red", "Eight", 2}, {"Red", "Nine", 2},
}

// Copies the deckSchematic to deck so we can vary the contents during the progress of the game
var deck = deckSchematic

// Global variables
var (
	playerTurn int = 0

	// 0 = Left, 1 = Right
	gameDirection bool = true

	currentCard handCard
)

// Get total cards
func getCardTotal() int {
	a := 0
	for i := 0; i < len(deck); i++ {
		a = a + deck[i].amountCards
	}
	return a
}

// Generate an integer
func randomInt(a int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(a)
}

// Draw a card
func drawCard(a string) {
	for true {
		var b int = randomInt(len(deckSchematic) - 1)

		// Checks that we dont have a full hand
		switch a {
		case "player":
			if player.amountInHand == 30 {
				break
			}
		case "computer1":
			if computer1.amountInHand == 30 {
				break
			}
		case "computer2":
			if computer2.amountInHand == 30 {
				break
			}
		case "computer3":
			if computer3.amountInHand == 30 {
				break
			}
		}

		// Checks that we have the card in the deck
		if deck[b].amountCards != 0 {

			// Selects player and gives them the card
			switch a {
			case "player":
				player.cardsInHand[player.amountInHand] = handCard{deck[b].color, deck[b].name}
				deck[b].amountCards--
				player.amountInHand++
			case "computer1":
				computer1.cardsInHand[computer1.amountInHand] = handCard{deck[b].color, deck[b].name}
				deck[b].amountCards--
				computer1.amountInHand++
			case "computer2":
				computer2.cardsInHand[computer2.amountInHand] = handCard{deck[b].color, deck[b].name}
				deck[b].amountCards--
				computer2.amountInHand++
			case "computer3":
				computer3.cardsInHand[computer3.amountInHand] = handCard{deck[b].color, deck[b].name}
				deck[b].amountCards--
				computer3.amountInHand++
			}
			break
		}
	}
}

// Initialize with a card
func initCard() {
	for true {
		var b int = randomInt(len(deckSchematic) - 1)

		// Checks that we have the card in the deck
		if deck[b].amountCards != 0 && deck[b].color != "Black" {
			currentCard = handCard{deck[b].color, deck[b].name}
			deck[b].amountCards--
			break
		}
	}
}

// Checks if a card is playable
func checkPossibility(cardColor, cardName string) bool {
	if cardColor == "Black" {
		return true
	} else {
		if cardColor == currentCard.color {
			return true
		} else {
			if cardName == currentCard.name {
				return true
			} else {
				return false
			}
		}
	}
}

// Compares two cards and returns a boolean lazily
func compareLazyCards(card1, card2 handCard) bool {
	if card1.name == card2.name || card1.color == card2.color {
		return true
	} else {
		return false
	}
}

// Compares two cards and returns a boolean strictly
func compareStrictCards(card1, card2 handCard) bool {
	if card1.name == card2.name && card1.color == card2.color {
		return true
	} else {
		return false
	}
}

// Check if card is owned by player
func ifCardInPlayer(cardToCheck handCard, playerName string) bool {
	switch playerName {
	case "player":
		for x := 0; x < cap(player.cardsInHand); x++ {
			if cardToCheck == player.cardsInHand[x] {
				return true
			}
		}
		return false
	case "computer1":
		for x := 0; x < cap(computer1.cardsInHand); x++ {
			if cardToCheck == computer1.cardsInHand[x] {
				return true
			}
		}
		return false
	case "computer2":
		for x := 0; x < cap(computer2.cardsInHand); x++ {
			if cardToCheck == computer2.cardsInHand[x] {
				return true
			}
		}
		return false
	case "computer3":
		for x := 0; x < cap(computer3.cardsInHand); x++ {
			if cardToCheck == computer3.cardsInHand[x] {
				return true
			}
		}
		return false
	}
	return false
}

// Places a card and reallocates other cards
func placeCard(selectedCard handCard, playerName string) {
	switch playerName {
	case "player":
		for x := 0; x < cap(player.cardsInHand); x++ {
			if selectedCard == player.cardsInHand[x] {
				cardNum := x
				for i := cardNum + 1; cardNum < player.amountInHand; i++ {
					if i <= 29 {
						player.cardsInHand[i-1] = player.cardsInHand[i]
						player.cardsInHand[player.amountInHand] = handCard{"", ""}
					} else {
						break
					}
				}
				currentCard = selectedCard
				break
			}
		}
		player.amountInHand = player.amountInHand - 1
	case "computer1":
		for x := 0; x < cap(computer1.cardsInHand); x++ {
			if selectedCard == computer1.cardsInHand[x] {
				cardNum := x
				for i := cardNum + 1; cardNum < computer1.amountInHand; i++ {
					if i <= 29 {
						computer1.cardsInHand[i-1] = computer1.cardsInHand[i]
						computer1.cardsInHand[computer1.amountInHand] = handCard{"", ""}
					} else {
						break
					}
				}
				currentCard = selectedCard
				break
			}
		}
		computer1.amountInHand = computer1.amountInHand - 1
	case "computer2":
		for x := 0; x < cap(computer2.cardsInHand); x++ {
			if selectedCard == computer2.cardsInHand[x] {
				cardNum := x
				for i := cardNum + 1; cardNum < computer2.amountInHand; i++ {
					if i <= 29 {
						computer2.cardsInHand[i-1] = computer2.cardsInHand[i]
						computer2.cardsInHand[computer2.amountInHand] = handCard{"", ""}
					} else {
						break
					}
				}
				currentCard = selectedCard
				break
			}
		}
		computer2.amountInHand = computer2.amountInHand - 1
	case "computer3":
		for x := 0; x < cap(computer3.cardsInHand); x++ {
			if selectedCard == computer3.cardsInHand[x] {
				cardNum := x
				for i := cardNum + 1; cardNum < computer3.amountInHand; i++ {
					if i <= 29 {
						computer3.cardsInHand[i-1] = computer3.cardsInHand[i]
						computer3.cardsInHand[computer3.amountInHand] = handCard{"", ""}
					} else {
						break
					}
				}
				currentCard = selectedCard
				break
			}
		}
		computer3.amountInHand = computer3.amountInHand - 1
	}
}

// Changes the turn
func turnChange() {
	if gameDirection == true {
		switch playerTurn {
		case 0:
			playerTurn = 1
		case 1:
			playerTurn = 2
		case 2:
			playerTurn = 3
		case 3:
			playerTurn = 0
		}
	} else {
		switch playerTurn {
		case 0:
			playerTurn = 3
		case 1:
			playerTurn = 0
		case 2:
			playerTurn = 1
		case 3:
			playerTurn = 2
		}
	}
}

func checkWin() bool {
	switch 0 {
	case player.amountInHand:
		fmt.Println("\n\nYou win!")
		return true
	case computer1.amountInHand:
		fmt.Println("\n\nCom1 win!")
		return true
	case computer2.amountInHand:
		fmt.Println("\n\nCom2 win!")
		return true
	case computer3.amountInHand:
		fmt.Println("\n\nCom3 win!")
		return true
	default:
		return false
	}
}

func main() {
	// Hands everyone 7 cards
	for i := 0; i < 7; i++ {
		drawCard("player")
		drawCard("computer1")
		drawCard("computer2")
		drawCard("computer3")
	}
	// Plays out a card in the middle
	initCard()

	for true {
		// Check for Win
		winVar := checkWin()
		if winVar == true {
			break
		}

		// Do stuff
		switch playerTurn {

		// Players turn
		case 0:
			player.isUNO = false

			// Show General information about the game
			fmt.Printf("\nCom1 has %v cards. ", computer1.amountInHand)
			if computer1.isUNO == true {
				fmt.Printf("Com1 has called UNO.")
			}
			fmt.Printf("\n")
			fmt.Printf("Com2 has %v cards. ", computer2.amountInHand)
			if computer2.isUNO == true {
				fmt.Printf("Com2 has called UNO.")
			}
			fmt.Printf("\n")
			fmt.Printf("Com3 has %v cards. ", computer3.amountInHand)
			if computer3.isUNO == true {
				fmt.Printf("Com3 has called UNO.")
			}
			fmt.Printf("\n")
			if player.isUNO == true {
				fmt.Printf("You have called UNO.\n")
			}
			fmt.Printf("The current card played is a %v %v\n", currentCard.color, currentCard.name)
			if gameDirection == true {
				fmt.Printf("The game rotation is: You > Com1 > Com2 > Com3\n")
			} else {
				fmt.Printf("The game rotation is: You < Com1 < Com2 < Com3\n")
			}

			fmt.Printf("\nYou have the cards:\n")
			for i := 0; i <= player.amountInHand-1; i++ {
				fmt.Printf("%v %v\n", player.cardsInHand[i].color, player.cardsInHand[i].name)
			}
			for true {

				// Give Information to Player
				fmt.Printf("\nWhat do you want to play? Write Draw to draw a card. ")
				if player.amountInHand == 2 {
					fmt.Printf("Write UNO to call UNO")
				}
				userFirstArgument := ""
				userSecondArgument := ""

				// Reads information
				fmt.Printf("\n> ")
				fmt.Scanln(&userFirstArgument, &userSecondArgument)
				fmt.Printf("\n")

				// Checks for player action
				switch userFirstArgument {

				// Draws a card
				case "Draw":
					drawCard("player")
					fmt.Printf("You drew a %v %v!\n", player.cardsInHand[player.amountInHand-1].color, player.cardsInHand[player.amountInHand-1].name)
					turnChange()
					break

				// UNO
				case "UNO":
					if player.amountInHand <= 2 {
						fmt.Printf("You called UNO.\n")
						player.isUNO = true
					} else {
						fmt.Printf("Checking if someone hasn't called UNO...\n")
						if computer1.isUNO == false && computer1.amountInHand == 1 {
							fmt.Printf("Com1 hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("computer1")
							}
						} else if computer2.isUNO == false && computer2.amountInHand == 1 {
							fmt.Printf("Com2 hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("computer2")
							}
						} else if computer3.isUNO == false && computer3.amountInHand == 1 {
							fmt.Printf("Com3 hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("computer3")
							}
						} else {
							fmt.Printf("No one hasn't or needs to call UNO. Please choose another option.")
						}
					}

				// Places card
				default:
					if checkPossibility(userFirstArgument, userSecondArgument) && ifCardInPlayer(handCard{userFirstArgument, userSecondArgument}, "player") {
						if (player.isUNO == false && player.amountInHand >= 2) || (player.isUNO == true && player.amountInHand == 1) {
							placeCard(handCard{userFirstArgument, userSecondArgument}, "player")
							fmt.Printf("Card was placed!\n\n")

							// Special card
							switch userSecondArgument {
							case "Plus2":
								if gameDirection == true {
									for i := 0; i < 2; i++ {
										drawCard("computer1")
									}
								} else {
									for i := 0; i < 2; i++ {
										drawCard("computer3")
									}
								}
								turnChange()
								break
							case "Plus4":
								wildcardResponse := ""
								for true {
									fmt.Printf("Choose a new color\n")
									fmt.Printf("> ")
									fmt.Scanf("%s", &wildcardResponse)
									fmt.Printf("")
									if wildcardResponse == "Green" || wildcardResponse == "Red" || wildcardResponse == "Blue" || wildcardResponse == "Yellow" {
										break
									}
								}
								currentCard = handCard{wildcardResponse, currentCard.name}
								if gameDirection == true {
									for i := 0; i < 4; i++ {
										drawCard("computer1")
									}
								} else {
									for i := 0; i < 4; i++ {
										drawCard("computer3")
									}
								}
								turnChange()
								break
							case "Block":
								turnChange()
								turnChange()
								break
							case "Wildcard":
								wildcardResponse := ""
								for true {
									fmt.Printf("Choose a new color\n")
									fmt.Printf("> ")
									fmt.Scanf("%s", &wildcardResponse)
									fmt.Printf("")
									if wildcardResponse == "Green" || wildcardResponse == "Red" || wildcardResponse == "Blue" || wildcardResponse == "Yellow" {
										break
									}
								}
								currentCard = handCard{wildcardResponse, currentCard.name}
								turnChange()
								break
							case "Reverse":
								if gameDirection == true {
									gameDirection = false
								} else {
									gameDirection = true
								}
								turnChange()
								break
							default:
								player.isUNO = false
								turnChange()
								break
							}

						}
					} else {
						fmt.Printf("\nSorry, that card cannot be placed.\n")
					}
				}
				if playerTurn != 0 {
					break
				}

			}

		// Com1s turn
		case 1:
			time.Sleep(time.Duration(randomInt(3)+2) * time.Second)
			possibleCardAmount := 0
			var possibleCards [30]handCard
			canCallUNO := false

			// Stores playable cards
			for i := 0; i < 30; i++ {
				if checkPossibility(computer1.cardsInHand[i].color, computer1.cardsInHand[i].name) {
					possibleCards[possibleCardAmount] = computer1.cardsInHand[i]
					possibleCardAmount++
				}
			}

			// Checks if they can call UNO
			if player.isUNO == false && player.amountInHand == 1 {
				canCallUNO = true
			} else if computer2.isUNO == false && computer2.amountInHand == 1 {
				canCallUNO = true
			} else if computer3.isUNO == false && computer3.amountInHand == 1 {
				canCallUNO = true
			} else if computer1.amountInHand <= 2 {
				canCallUNO = true
			}

			// Does a move
			if possibleCardAmount == 0 {
				fmt.Printf("\nCom1 draws a card!\n")
				drawCard("computer1")
				possibleCardAmount++
				turnChange()
			} else {
				if canCallUNO == true && randomInt(2) == 2 {
					fmt.Printf("\nCom1 is calling UNO!\n")
					if computer1.amountInHand <= 2 {
						computer1.isUNO = true
					} else {
						if player.isUNO == false && player.amountInHand == 1 {
							fmt.Printf("Player hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("player")
							}
						} else if computer2.isUNO == false && computer2.amountInHand == 1 {
							fmt.Printf("Com2 hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("computer2")
							}
						} else if computer3.isUNO == false && computer3.amountInHand == 1 {
							fmt.Printf("Com3 hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("computer3")
							}
						}
					}

				} else {
					userArgument := possibleCards[(randomInt(possibleCardAmount))]
					userFirstArgument := userArgument.color
					userSecondArgument := userArgument.name
					placeCard(handCard{userFirstArgument, userSecondArgument}, "computer1")
					fmt.Printf("\nCom1 places a %v %v!\n", userFirstArgument, userSecondArgument)
					switch userSecondArgument {
					case "Plus2":
						if gameDirection == true {
							for i := 0; i < 2; i++ {
								drawCard("computer2")
							}
						} else {
							for i := 0; i < 2; i++ {
								drawCard("player")
							}
						}
						turnChange()
					case "Plus4":
						time.Sleep(time.Duration(randomInt(3)+2) * time.Second)
						wildcardResponse := ""
						switch randomInt(3) {
						case 0:
							fmt.Printf("Com1 changes to Blue!\n")
							wildcardResponse = "Blue"
						case 1:
							fmt.Printf("Com1 changes to Red!\n")
							wildcardResponse = "Red"
						case 2:
							fmt.Printf("Com1 changes to Yellow!\n")
							wildcardResponse = "Yellow"
						case 3:
							fmt.Printf("Com1 changes to Green!\n")
							wildcardResponse = "Green"
						}
						if wildcardResponse == "Blue" || wildcardResponse == "Red" || wildcardResponse == "Yellow" || wildcardResponse == "Green" {
							currentCard = handCard{wildcardResponse, currentCard.name}
						} else {
							currentCard = handCard{"Red", currentCard.name}
						}
						if gameDirection == true {
							for i := 0; i < 4; i++ {
								drawCard("computer2")
							}
						} else {
							for i := 0; i < 4; i++ {
								drawCard("player")
							}
						}
						turnChange()
					case "Block":
						turnChange()
						turnChange()
					case "Wildcard":
						time.Sleep(time.Duration(randomInt(3)+2) * time.Second)
						wildcardResponse := ""
						switch randomInt(3) {
						case 0:
							fmt.Printf("Com1 changes to Blue!\n")
							wildcardResponse = "Blue"
						case 1:
							fmt.Printf("Com1 changes to Red!\n")
							wildcardResponse = "Red"
						case 2:
							fmt.Printf("Com1 changes to Yellow!\n")
							wildcardResponse = "Yellow"
						case 3:
							fmt.Printf("Com1 changes to Green!\n")
							wildcardResponse = "Green"
						}
						if wildcardResponse == "Blue" || wildcardResponse == "Red" || wildcardResponse == "Yellow" || wildcardResponse == "Green" {
							currentCard = handCard{wildcardResponse, currentCard.name}
						} else {
							currentCard = handCard{"Red", currentCard.name}
						}
						turnChange()
					case "Reverse":
						fmt.Printf("Com1 changes the direction!\n")
						if gameDirection == true {
							gameDirection = false
						} else {
							gameDirection = true
						}
						turnChange()
					default:
						computer1.isUNO = false
						turnChange()
					}
				}
			}
			break

		// Com2s turn
		case 2:
			time.Sleep(time.Duration(randomInt(3)+2) * time.Second)
			possibleCardAmount := 0
			var possibleCards [30]handCard
			canCallUNO := false

			// Stores playable cards
			for i := 0; i < 30; i++ {
				if checkPossibility(computer2.cardsInHand[i].color, computer2.cardsInHand[i].name) {
					possibleCards[possibleCardAmount] = computer2.cardsInHand[i]
					possibleCardAmount++
				}
			}

			// Checks if they can call UNO
			if player.isUNO == false && player.amountInHand == 1 {
				canCallUNO = true
			} else if computer1.isUNO == false && computer1.amountInHand == 1 {
				canCallUNO = true
			} else if computer3.isUNO == false && computer3.amountInHand == 1 {
				canCallUNO = true
			} else if computer2.amountInHand <= 2 {
				canCallUNO = true
			}

			// Does a move
			if possibleCardAmount == 0 {
				fmt.Printf("\nCom2 draws a card!\n")
				drawCard("computer2")
				possibleCardAmount++
				turnChange()
			} else {
				if canCallUNO == true && randomInt(2) == 2 {
					fmt.Printf("\nCom2 is calling UNO!\n")
					if computer2.amountInHand <= 2 {
						computer2.isUNO = true
					} else {
						if player.isUNO == false && player.amountInHand == 1 {
							fmt.Printf("Player hasn't called UNO! They to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("player")
							}
						} else if computer1.isUNO == false && computer1.amountInHand == 1 {
							fmt.Printf("Com1 hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("computer1")
							}
						} else if computer3.isUNO == false && computer3.amountInHand == 1 {
							fmt.Printf("Com3 hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("computer3")
							}
						}
					}

				} else {
					userArgument := possibleCards[(randomInt(possibleCardAmount))]
					userFirstArgument := userArgument.color
					userSecondArgument := userArgument.name
					placeCard(handCard{userFirstArgument, userSecondArgument}, "computer2")
					fmt.Printf("\nCom2 places a %v %v!\n", userFirstArgument, userSecondArgument)
					switch userSecondArgument {
					case "Plus2":
						if gameDirection == true {
							for i := 0; i < 2; i++ {
								drawCard("computer3")
							}
						} else {
							for i := 0; i < 2; i++ {
								drawCard("computer1")
							}
						}
						turnChange()
					case "Plus4":
						time.Sleep(time.Duration(randomInt(3)+2) * time.Second)
						wildcardResponse := ""
						switch randomInt(3) {
						case 0:
							fmt.Printf("Com2 changes to Blue!\n")
							wildcardResponse = "Blue"
						case 1:
							fmt.Printf("Com2 changes to Red!\n")
							wildcardResponse = "Red"
						case 2:
							fmt.Printf("Com2 changes to Yellow!\n")
							wildcardResponse = "Yellow"
						case 3:
							fmt.Printf("Com2 changes to Green!\n")
							wildcardResponse = "Green"
						}
						if wildcardResponse == "Blue" || wildcardResponse == "Red" || wildcardResponse == "Yellow" || wildcardResponse == "Green" {
							currentCard = handCard{wildcardResponse, currentCard.name}
						} else {
							currentCard = handCard{"Red", currentCard.name}
						}
						if gameDirection == true {
							for i := 0; i < 4; i++ {
								drawCard("computer3")
							}
						} else {
							for i := 0; i < 4; i++ {
								drawCard("computer1")
							}
						}
						turnChange()
					case "Block":
						turnChange()
						turnChange()
					case "Wildcard":
						time.Sleep(time.Duration(randomInt(3)+2) * time.Second)
						wildcardResponse := ""
						switch randomInt(3) {
						case 0:
							fmt.Printf("Com2 changes to Blue!\n")
							wildcardResponse = "Blue"
						case 1:
							fmt.Printf("Com2 changes to Red!\n")
							wildcardResponse = "Red"
						case 2:
							fmt.Printf("Com2 changes to Yellow!\n")
							wildcardResponse = "Yellow"
						case 3:
							fmt.Printf("Com2 changes to Green!\n")
							wildcardResponse = "Green"
						}
						if wildcardResponse == "Blue" || wildcardResponse == "Red" || wildcardResponse == "Yellow" || wildcardResponse == "Green" {
							currentCard = handCard{wildcardResponse, currentCard.name}
						} else {
							currentCard = handCard{"Red", currentCard.name}
						}
						turnChange()
					case "Reverse":
						fmt.Printf("Com2 changes the direction!\n")
						if gameDirection == true {
							gameDirection = false
						} else {
							gameDirection = true
						}
						turnChange()
					default:
						computer2.isUNO = false
						turnChange()
					}
				}
			}
		case 3:
			time.Sleep(time.Duration(randomInt(3)+2) * time.Second)
			possibleCardAmount := 0
			var possibleCards [30]handCard
			canCallUNO := false

			// Stores playable cards
			for i := 0; i < 30; i++ {
				if checkPossibility(computer3.cardsInHand[i].color, computer3.cardsInHand[i].name) {
					possibleCards[possibleCardAmount] = computer3.cardsInHand[i]
					possibleCardAmount++
				}
			}

			// Checks if they can call UNO
			if player.isUNO == false && player.amountInHand == 1 {
				canCallUNO = true
			} else if computer1.isUNO == false && computer1.amountInHand == 1 {
				canCallUNO = true
			} else if computer2.isUNO == false && computer2.amountInHand == 1 {
				canCallUNO = true
			} else if computer3.amountInHand <= 2 {
				canCallUNO = true
			}

			// Does a move
			if possibleCardAmount == 0 {
				fmt.Printf("\nCom3 draws a card!\n")
				drawCard("computer3")
				possibleCardAmount++
				turnChange()
			} else {
				if canCallUNO == true && randomInt(2) == 2 {
					fmt.Printf("\nCom3 is calling UNO!\n")
					if computer3.amountInHand <= 2 {
						computer3.isUNO = true
					} else {
						if player.isUNO == false && player.amountInHand == 1 {
							fmt.Printf("Player hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("player")
							}
						} else if computer1.isUNO == false && computer1.amountInHand == 1 {
							fmt.Printf("Com1 hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("computer1")
							}
						} else if computer2.isUNO == false && computer2.amountInHand == 1 {
							fmt.Printf("Com2 hasn't called UNO! They have to draw 2 cards.\n")
							for i := 0; i < 2; i++ {
								drawCard("computer2")
							}
						}
					}

				} else {
					userArgument := possibleCards[(randomInt(possibleCardAmount))]
					userFirstArgument := userArgument.color
					userSecondArgument := userArgument.name
					placeCard(handCard{userFirstArgument, userSecondArgument}, "computer3")
					fmt.Printf("\nCom3 places a %v %v!\n", userFirstArgument, userSecondArgument)
					switch userSecondArgument {
					case "Plus2":
						if gameDirection == true {
							for i := 0; i < 2; i++ {
								drawCard("player")
							}
						} else {
							for i := 0; i < 2; i++ {
								drawCard("computer2")
							}
						}
						turnChange()
					case "Plus4":
						time.Sleep(time.Duration(randomInt(3)+2) * time.Second)
						wildcardResponse := ""
						switch randomInt(3) {
						case 0:
							fmt.Printf("Com3 changes to Blue!\n")
							wildcardResponse = "Blue"
						case 1:
							fmt.Printf("Com3 changes to Red!\n")
							wildcardResponse = "Red"
						case 2:
							fmt.Printf("Com3 changes to Yellow!\n")
							wildcardResponse = "Yellow"
						case 3:
							fmt.Printf("Com3 changes to Green!\n")
							wildcardResponse = "Green"
						}
						if wildcardResponse == "Blue" || wildcardResponse == "Red" || wildcardResponse == "Yellow" || wildcardResponse == "Green" {
							currentCard = handCard{wildcardResponse, currentCard.name}
						} else {
							currentCard = handCard{"Red", currentCard.name}
						}
						if gameDirection == true {
							for i := 0; i < 4; i++ {
								drawCard("player")
							}
						} else {
							for i := 0; i < 4; i++ {
								drawCard("computer2")
							}
						}
						turnChange()
					case "Block":
						turnChange()
						turnChange()
					case "Wildcard":
						time.Sleep(time.Duration(randomInt(3)+2) * time.Second)
						wildcardResponse := ""
						switch randomInt(3) {
						case 0:
							fmt.Printf("Com3 changes to Blue!\n")
							wildcardResponse = "Blue"
						case 1:
							fmt.Printf("Com3 changes to Red!\n")
							wildcardResponse = "Red"
						case 2:
							fmt.Printf("Com3 changes to Yellow!\n")
							wildcardResponse = "Yellow"
						case 3:
							fmt.Printf("Com3 changes to Green!\n")
							wildcardResponse = "Green"
						}
						if wildcardResponse == "Blue" || wildcardResponse == "Red" || wildcardResponse == "Yellow" || wildcardResponse == "Green" {
							currentCard = handCard{wildcardResponse, currentCard.name}
						} else {
							currentCard = handCard{"Red", currentCard.name}
						}
						turnChange()
					case "Reverse":
						fmt.Printf("Com3 changes the direction!\n")
						if gameDirection == true {
							gameDirection = false
						} else {
							gameDirection = true
						}
						turnChange()
					default:
						computer3.isUNO = false
						turnChange()
					}
				}
			}
			time.Sleep(time.Duration(randomInt(3)+2) * time.Second)
		}
	}
}
