import SuitEnum from "../enum/Suit";
import ValueEnum from "../enum/Value";
import { Deck, Game, Player } from "../types";

export const createDeck = (): Deck => {
    return [
        { suit: SuitEnum.Clubs, value: ValueEnum.Three },
        { suit: SuitEnum.Clubs, value: ValueEnum.Four },
        { suit: SuitEnum.Clubs, value: ValueEnum.Five },
        { suit: SuitEnum.Clubs, value: ValueEnum.Six },
        { suit: SuitEnum.Clubs, value: ValueEnum.Seven },
        { suit: SuitEnum.Clubs, value: ValueEnum.Eight },
        { suit: SuitEnum.Clubs, value: ValueEnum.Nine },
        { suit: SuitEnum.Clubs, value: ValueEnum.Ten },
        { suit: SuitEnum.Clubs, value: ValueEnum.Jack },
        { suit: SuitEnum.Clubs, value: ValueEnum.Queen },
        { suit: SuitEnum.Clubs, value: ValueEnum.King },
        { suit: SuitEnum.Clubs, value: ValueEnum.Ace },
        { suit: SuitEnum.Clubs, value: ValueEnum.Two },
        { suit: SuitEnum.Spades, value: ValueEnum.Three },
        { suit: SuitEnum.Spades, value: ValueEnum.Four },
        { suit: SuitEnum.Spades, value: ValueEnum.Five },
        { suit: SuitEnum.Spades, value: ValueEnum.Six },
        { suit: SuitEnum.Spades, value: ValueEnum.Seven },
        { suit: SuitEnum.Spades, value: ValueEnum.Eight },
        { suit: SuitEnum.Spades, value: ValueEnum.Nine },
        { suit: SuitEnum.Spades, value: ValueEnum.Ten },
        { suit: SuitEnum.Spades, value: ValueEnum.Jack },
        { suit: SuitEnum.Spades, value: ValueEnum.Queen },
        { suit: SuitEnum.Spades, value: ValueEnum.King },
        { suit: SuitEnum.Spades, value: ValueEnum.Ace },
        { suit: SuitEnum.Spades, value: ValueEnum.Two },
        { suit: SuitEnum.Hearts, value: ValueEnum.Three },
        { suit: SuitEnum.Hearts, value: ValueEnum.Four },
        { suit: SuitEnum.Hearts, value: ValueEnum.Five },
        { suit: SuitEnum.Hearts, value: ValueEnum.Six },
        { suit: SuitEnum.Hearts, value: ValueEnum.Seven },
        { suit: SuitEnum.Hearts, value: ValueEnum.Eight },
        { suit: SuitEnum.Hearts, value: ValueEnum.Nine },
        { suit: SuitEnum.Hearts, value: ValueEnum.Ten },
        { suit: SuitEnum.Hearts, value: ValueEnum.Jack },
        { suit: SuitEnum.Hearts, value: ValueEnum.Queen },
        { suit: SuitEnum.Hearts, value: ValueEnum.King },
        { suit: SuitEnum.Hearts, value: ValueEnum.Ace },
        { suit: SuitEnum.Hearts, value: ValueEnum.Two },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Three },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Four },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Five },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Six },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Seven },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Eight },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Nine },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Ten },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Jack },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Queen },
        { suit: SuitEnum.Diamonds, value: ValueEnum.King },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Ace },
        { suit: SuitEnum.Diamonds, value: ValueEnum.Two },
    ];
};

export const shuffleDeck = (deck: Deck): Deck => {

    // @see: https://stackoverflow.com/a/12646864
    function shuffleArray(array: any[]) {
        for (let i = array.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            [array[i], array[j]] = [array[j], array[i]];
        }

        return array;
    }

    return shuffleArray(deck);
};

export const dealDeck = (deck: Deck): Game => 
{
    const player1: Player = { hand: { cards: []}, move: { cards: [] } };
    const player2: Player = { hand: { cards: []}, move: { cards: [] } };
    const player3: Player = { hand: { cards: []}, move: { cards: [] } };
    const player4: Player = { hand: { cards: []}, move: { cards: [] } };

    deck.forEach((card, index) => {
        switch (index % 4){
            case 0: 
                player1.hand.cards.push(card);
                break;

            case 1: 
                player2.hand.cards.push(card);
                break;

            case 2: 
                player3.hand.cards.push(card);
                break;

            case 3: 
                player4.hand.cards.push(card);
                break;
        }
    });

    return [
        player1,
        player2,
        player3,
        player4,
    ];
}