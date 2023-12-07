package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards []card
	bid   int
	score int
}

type card struct {
	rank  rune
	value int
}

const (
	FiveOfAKind  int = 7
	FourOfAKind  int = 6
	FullHouse    int = 5
	ThreeOfAKind int = 4
	TwoPair      int = 3
	OnePair      int = 2
	HighCard     int = 1
)

var rankValues = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

func main() {
	fmt.Println("Hello 2023 day 7!")

	hands := readHands("input.txt", false)

	for i, v := range hands {
		fmt.Println("Hand", i+1)
		fmt.Print("Cards: ")
		for _, c := range v.cards {
			fmt.Printf("%v", string(c.rank))
		}
		fmt.Printf("\nBid: %v, Score:%v\n", v.bid, v.score)
	}

	fmt.Printf("The solution for part 1 is %v\n", calculateWinnings(hands))

	hands2 := readHands("input.txt", true)

	fmt.Printf("The solution for part 2 is %v\n", calculateWinnings(hands2))
}

func calculateWinnings(hands []hand) int {
	rankedHands := []hand{}
	returnTotal := 0

	// rank the hands by strength
	for _, h := range hands {
		if len(rankedHands) == 0 {
			rankedHands = append(rankedHands, h)
			continue
		}
		inserted := false
		for j, v := range rankedHands {
			if h.score > v.score {
				rankedHands = insertHand(rankedHands, j, h)
				inserted = true
			} else if h.score == v.score && playWar(h, v) {
				rankedHands = insertHand(rankedHands, j, h)
				inserted = true
			}
			if inserted {
				break
			}
		}
		if inserted {
			inserted = false
			continue
		}
		rankedHands = append(rankedHands, h)
	}

	for i, h := range rankedHands {
		returnTotal += (len(rankedHands) - i) * h.bid
	}

	return returnTotal
}

// Returns true if a is stronger than b
func playWar(a hand, b hand) bool {
	for i, card := range a.cards {
		if card.value > b.cards[i].value {
			return true
		} else if card.value < b.cards[i].value {
			return false
		}
	}
	return true
}

func readHands(file string, jokers bool) []hand {
	returnHands := make([]hand, 0)
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Couldn't read %v, %v\n", file, err)
	}

	lines := strings.Split(trim(string(f)), "\n")

	for _, l := range lines {
		returnHands = append(returnHands, parseHand(l, jokers))
	}

	return returnHands
}

func insertHand(hands []hand, index int, h hand) []hand {
	if len(hands) == index {
		return append(hands, h)
	}
	hands = append(hands[:index+1], hands[index:]...)
	hands[index] = h

	return hands
}

func parseHand(line string, jokers bool) hand {
	returnHand := hand{}
	handAndBid := strings.Split(line, " ")

	bid, err := strconv.Atoi(handAndBid[1])
	if err != nil {
		log.Fatalf("Couldn't convert bid value %v, %v\n", handAndBid[1], err)
	}
	returnHand.bid = bid

	for _, r := range handAndBid[0] {
		if jokers {
			returnHand.cards = append(returnHand.cards, card{rank: r, value: cardValueWithJokers(rankValues[r])})
		} else {
			returnHand.cards = append(returnHand.cards, card{rank: r, value: rankValues[r]})
		}
	}

	returnHand.score = scoreHand(returnHand, jokers)

	return returnHand
}

func cardValueWithJokers(v int) int {
	switch v {
	case 10:
		return 1
	case 11, 12, 13:
		return v
	default:
		return v + 1
	}
}

type cardCount struct {
	Key   card
	Value int
}

func scoreHand(h hand, joker bool) int {
	countsMap := make(map[card]int, 0)

	for _, c := range h.cards {
		countsMap[c]++
	}

	var countsSlice []cardCount

	for k, v := range countsMap {
		countsSlice = append(countsSlice, cardCount{Key: k, Value: v})
	}

	sort.Slice(countsSlice, func(i, j int) bool { return countsSlice[i].Value > countsSlice[j].Value })
	jokers := countJokers(countsSlice)

	switch countsSlice[0].Value {
	case 5:
		return FiveOfAKind
	case 4:
		if joker && jokers > 0 {
			return FiveOfAKind
		}
		return FourOfAKind
	case 3:
		if joker {
			if jokers >= 2 {
				return FiveOfAKind
			} else if jokers == 1 {
				return FourOfAKind
			} else if countsSlice[1].Value == 2 {
				return FullHouse
			} else {
				return ThreeOfAKind
			}
		}
		if countsSlice[1].Value == 2 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	case 2:
		if joker {
			if jokers == 2 {
				if countsSlice[1].Value == 2 {
					return FourOfAKind
				}
			} else if jokers == 1 && countsSlice[1].Value == 2 {
				return FullHouse
			}
			return ThreeOfAKind
		}
		if countsSlice[1].Value == 2 {
			return TwoPair
		} else {
			return OnePair
		}
	default:
		if joker {
			return OnePair
		}
		return HighCard
	}
}

func countJokers(handSlice []cardCount) int {
	for _, v := range handSlice {
		if v.Key.rank == 'J' {
			return v.Value
		}
	}
	return 0
}

func (h hand) getHand() string {
	returnString := ""
	for _, v := range h.cards {
		returnString += fmt.Sprintf("%v", string(v.rank))
	}
	return returnString
}

func trim(text string) string {
	return strings.TrimSpace(text)
}
