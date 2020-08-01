import { createDeck, shuffleDeck, dealDeck } from "./GameUtils";
import SuitEnum from "../enum/Suit";
import ValueEnum from "../enum/Value";
import { Deck, Hand } from "../types";
import { toString } from "./CardUtils";

const validFullDeck: Deck = [
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

describe('GameUtils', () => {
    describe('createDeck', () => {
        test('correctly creates deck', () => {
            const result = createDeck();

            expect(result).toHaveLength(52);
            expect(result).toStrictEqual(validFullDeck);
        });
    });

    describe('shuffleDeck', () => {
        test('deck is not identical to generated deck', () => {
            const result = shuffleDeck(createDeck());

            expect(result).toHaveLength(52);
            expect(result).not.toStrictEqual(validFullDeck);
        });

        test('generates random decks sequentially', async () => {
            const resultOne = shuffleDeck(createDeck());
            const resultTwo = shuffleDeck(createDeck());
            const resultThree = shuffleDeck(createDeck());

            expect(resultOne).not.toStrictEqual(resultTwo);
            expect(resultTwo).not.toStrictEqual(resultThree);
        });
    });

    describe('dealDeck', () => {
        test('all players are dealt 13 cards', () => {
            const [player1, player2, player3, player4] = dealDeck(createDeck());

            expect(player1.hand.cards.length).toBe(13);
            expect(player2.hand.cards.length).toBe(13);
            expect(player3.hand.cards.length).toBe(13);
            expect(player4.hand.cards.length).toBe(13);
        });

        test('all hands dealt are unique', () => {
            const [one,two,three,four] = dealDeck(createDeck());

            const doesHandContainSameCard = (handA: Hand) => (handB: Hand) => {
                return handA.cards.find(cardOne => (
                    handB.cards.forEach(cardTwo => {
                        return (toString(cardOne) === toString(cardTwo));
                    })
                ));
            };

            expect(doesHandContainSameCard(one.hand)(two.hand)).toBeFalsy();
            expect(doesHandContainSameCard(two.hand)(three.hand)).toBeFalsy();
            expect(doesHandContainSameCard(three.hand)(four.hand)).toBeFalsy();
        });
    });
});