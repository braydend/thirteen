import { addCardsToHand, removeCardFromHand } from './HandUtils';
import { Hand, Card } from "../types";
import SuitEnum from '../enum/Suit';
import ValueEnum from '../enum/Value';

describe('HandUtils', () => {
    test('addCardsToHand', () => {
        const hand: Hand = { cards: [] };
        const cards: Card[] = [
            { suit: SuitEnum.Spades, value: ValueEnum.Three },
            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
        ];

        const result = addCardsToHand(hand, cards);

        expect(result.cards).toHaveLength(2);
    });

    test('throws error when trying to add card to hand that aready contains card', () => {
        const hand: Hand = { cards: [{ suit: SuitEnum.Spades, value: ValueEnum.Three }] };
        const cards: Card[] = [
            { suit: SuitEnum.Spades, value: ValueEnum.Three },
            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
        ];

        expect(() => addCardsToHand(hand, cards)).toThrowError('Invalid Deck: Three of Spades cannot be added to hand as it already exists in this hand.')
    });

    test('removeCardFromHand', () => {
        const hand: Hand = { cards: [
            { suit: SuitEnum.Spades, value: ValueEnum.Three },
            { suit: SuitEnum.Spades, value: ValueEnum.Four },
            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
        ] };

        const result = removeCardFromHand(hand, { suit: SuitEnum.Spades, value: ValueEnum.Three });

        expect(result.cards).toHaveLength(2);
    });    
    
    test('removeCardFromHand throws error if card is not in hand', () => {
        const hand: Hand = { cards: [] };
        const card: Card = { suit: SuitEnum.Spades, value: ValueEnum.Three };

        expect(() => removeCardFromHand(hand, card)).toThrowError('Error: Cannot remove Three of Spades from hand as it does not exist in the hand.');
    });
});