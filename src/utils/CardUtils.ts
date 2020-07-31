import { Card } from "../types";
import ValueEnum from "../enum/Value";
import SuitEnum from "../enum/Suit";

export const toString = (card: Card): string => {
    const valueString = ValueEnum[card.value];
    const suitString = SuitEnum[card.suit];

    return `${valueString} of ${suitString}`
}

export const getSuitIcon = (suit: SuitEnum): string => {
    switch(suit){
        case SuitEnum.Hearts: return "♥";
        case SuitEnum.Diamonds: return "♦";
        case SuitEnum.Clubs: return "♣";
        case SuitEnum.Spades: return "♠";
    }
};

export const min = (cards: Card[]): Card => {
    if (cards.length === 0) throw new Error('Cannot get minimum of empty list of cards');

    return cards.reduce((lowest, current) => {
        if (current.suit === lowest.suit){
            return current.value < lowest.value ? current : lowest;
        }

        if (current.value === lowest.value){
            return current.suit < lowest.suit ? current : lowest;
        }

        return current.value < lowest.value ? current : lowest;
    }, cards[0]);
}

export const max = (cards: Card[]): Card => {
    if (cards.length === 0) throw new Error('Cannot get maximum of empty list of cards');

    return cards.reduce((highest, current) => {
        if (current.suit === highest.suit){
            return current.value > highest.value ? current : highest;
        }

        if (current.value === highest.value){
            return current.suit > highest.suit ? current : highest;
        }

        return current.value > highest.value ? current : highest;
    }, cards[0]);
}

export const sortCards = (cards: Card[], order: 'asc' | 'desc' = 'asc'): Card[] => {
    return cards.sort((a,b) => {
        if (a.value === b.value) {
            return order === 'asc' ? a.suit - b.suit : b.suit - a.suit;
        }
        return order === 'asc' ? a.value - b.value : b.value - a.value;
    });
}