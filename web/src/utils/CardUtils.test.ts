import { Card } from "../types";
import SuitEnum from "../enum/Suit";
import ValueEnum from "../enum/Value";
import { toString, min, max, sortCards, getSuitIcon } from "./CardUtils";

describe('Card Utils', () => {
    describe('toString', () => {
        test('ace of clubs', () => {
            const card: Card = {suit: SuitEnum.Clubs, value: ValueEnum.Ace};
            expect(toString(card)).toBe('Ace of Clubs');
        });

        test('getSuitIcon', () => {
            expect(getSuitIcon(SuitEnum.Hearts)).toBe('♥');
            expect(getSuitIcon(SuitEnum.Diamonds)).toBe('♦');
            expect(getSuitIcon(SuitEnum.Clubs)).toBe('♣');
            expect(getSuitIcon(SuitEnum.Spades)).toBe('♠');
        }); 
    });

    describe('min', () => {
        let threeOfSpades: Card;
        let fiveOfSpades: Card;
        let fiveOfHearts: Card;

        beforeEach(() => {
            threeOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Three };
            fiveOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Five };
            fiveOfHearts = { suit: SuitEnum.Hearts, value: ValueEnum.Five };
        });

        test('single card', () => {
            const cards = [threeOfSpades];

            expect(min(cards)).toBe(threeOfSpades);
        });

        test('same suit, different value', () => {
            const cards = [threeOfSpades, fiveOfSpades];

            expect(min(cards)).toBe(threeOfSpades);
        });

        test('different suit, same value', () => {
            const cards = [fiveOfHearts, fiveOfSpades];

            expect(min(cards)).toBe(fiveOfSpades);
        });

        test('different suit, different value', () => {
            const cards = [threeOfSpades, fiveOfHearts];

            expect(min(cards)).toBe(threeOfSpades);
        });
    });

    describe('max', () => {
        let threeOfSpades: Card;
        let fourOfSpades: Card;
        let fiveOfSpades: Card;
        let sixOfSpades: Card;
        let sevenOfSpades: Card;
        let fiveOfHearts: Card;
        let fiveOfClubs: Card;
        let fiveOfDiamonds: Card;

        beforeEach(() => {
            threeOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Three};
            fourOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Four};
            fiveOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Five};
            sixOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Six};
            sevenOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Seven};
            fiveOfHearts = { suit: SuitEnum.Hearts, value: ValueEnum.Five};
            fiveOfClubs = { suit: SuitEnum.Clubs, value: ValueEnum.Five};
            fiveOfDiamonds = { suit: SuitEnum.Diamonds, value: ValueEnum.Five};
        });

        test('single card', () => {
            const cards = [threeOfSpades];

            expect(max(cards)).toBe(threeOfSpades);
        });
        
        describe('matching value', () => {
            test('pair', () => {
                const cards = [fiveOfSpades, fiveOfClubs];

                expect(max(cards)).toBe(fiveOfClubs);
            });
            test('three of a kind', () => {
                const cards = [fiveOfSpades, fiveOfClubs, fiveOfDiamonds];

                expect(max(cards)).toBe(fiveOfDiamonds);
            });
            test('four of a kind (CHOP)', () => {
                const cards = [fiveOfSpades, fiveOfClubs, fiveOfDiamonds, fiveOfHearts];

                expect(max(cards)).toBe(fiveOfHearts);
            });
        });

        describe('straight', () => {
            test('three cards', () => {
                const cards = [threeOfSpades, fourOfSpades, fiveOfSpades];

                expect(max(cards)).toBe(fiveOfSpades);
            });
            test('four cards', () => {
                const cards = [threeOfSpades, fourOfSpades, fiveOfSpades, sixOfSpades];

                expect(max(cards)).toBe(sixOfSpades);
            });
            test('five cards', () => {
                const cards = [threeOfSpades, fourOfSpades, fiveOfSpades, sevenOfSpades];

                expect(max(cards)).toBe(sevenOfSpades);
            });
        })
    });

    describe('sortCards', () => {
        let cards: Card[];
        const aceOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Ace};
        const aceOfHearts = { suit: SuitEnum.Hearts, value: ValueEnum.Ace};
        const threeOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Three};
        const twoOfDiamonds = { suit: SuitEnum.Diamonds, value: ValueEnum.Two};

        beforeEach(() => {
            cards = [aceOfSpades, aceOfHearts, threeOfSpades, twoOfDiamonds];
        });

        test('asc', () => {
            expect(sortCards(cards)).toStrictEqual([
                threeOfSpades, 
                aceOfSpades, 
                aceOfHearts, 
                twoOfDiamonds,
            ]);
        });

        test('desc', () => {
            expect(sortCards(cards, 'desc')).toStrictEqual([
                twoOfDiamonds, 
                aceOfHearts, 
                aceOfSpades, 
                threeOfSpades,
            ]);
        });
    });
});
