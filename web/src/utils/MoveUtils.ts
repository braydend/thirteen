import { Card, Move } from "../types";
import MoveEnum from "../enum/Move";
import { min, max, sortCards, toString } from "./CardUtils";
import ValueEnum from "../enum/Value";
import SuitEnum from "../enum/Suit";

export const createMoveFromCards = (cards: Card[]): Move => {
    const newMove: Move = { cards };

    return { ...newMove, type: getMoveType(newMove) };
};

export const addCardsToMove = (move: Move, cards: Card[]): Move => {
    const newCards = [ ...move.cards, ...cards ];
    const newMove: Move = { ...move,  cards: newCards };
    
    validateMove(newMove);

    return { ...newMove, type: getMoveType(newMove) };
};

export const removeCardFromMove = (move: Move, cardToMove: Card): Move => {
    if (!move.cards.find(card => toString(card) === toString(cardToMove))) {
        throw new Error(`Error: Card ${toString(cardToMove)} is not in the move`);
    }

    const newMove: Move = { ...move, cards: move.cards.filter(cardToFilter => toString(cardToFilter) !== toString(cardToMove)) };

    return { ...newMove, type: getMoveType(newMove) };
};

export const getMoveType = (move: Move): MoveEnum | undefined => {
    if (isSingleCard(move)) return MoveEnum.SingleCard;

    if (isMatchingValue(move)) return MoveEnum.MatchingValue;

    if (isPotentialStraight(move)) return MoveEnum.PotentialStraight;

    if (isStraight(move)) return MoveEnum.Straight;

    if (isChop(move)) return MoveEnum.ThreeConsecutivePairs;

    if (isPotentialChop(move)) return MoveEnum.PotentialThreeConsecutivePairs;

    return undefined;
};

export const isChop = (move: Move): boolean => {
    if (isMatchingValue(move) && move.cards.length === 4) return true;

    if (move.cards.length !== 6) return false;

    const firstSeedValue = move.cards[0].value;
    const secondSeedValue = move.cards.filter(card => card.value !== firstSeedValue)[0].value;
    const thirdSeedValue = move.cards.filter(card => card.value !== firstSeedValue && card.value !== secondSeedValue)[0].value;

    const firstPair = move.cards.filter(card => card.value === firstSeedValue);
    const secondPair = move.cards.filter(card => card.value === secondSeedValue);
    const thirdPair = move.cards.filter(card => card.value === thirdSeedValue);

    if (firstPair.length !== 2 || secondPair.length !== 2 || thirdPair.length !== 2) return false;

    const minCard = min(move.cards);
    const maxCard = max(move.cards);

    return (minCard.value + 2) === maxCard.value;
}

export const isBeatenBy = (currentMove: Move, newMove: Move): boolean => {
    const currentHighestCard = max(currentMove.cards);
    const newHighestCard = max(newMove.cards);

    if (isChop(newMove) && !isChop(currentMove)) return true;
    if (!isChop(newMove) && isChop(currentMove)) return false;
    if (isChop(newMove) && isChop(currentMove)) return max([currentHighestCard, newHighestCard]) === newHighestCard;
    if (newMove.cards.length !== currentMove.cards.length || 
        newMove.type !== currentMove.type ) return false;

    return max([currentHighestCard, newHighestCard]) === newHighestCard;
}

export const validateMove = (move: Move): void => {
    if (!isValid(move)) {
        throw new Error(`Invalid Move! The following cards cannot make a valid move:${move.cards.map(card => ` ${ValueEnum[card.value]} of ${SuitEnum[card.suit]}`)}`);
    }
}

const isPotentialChop = ({ cards }: Move): boolean => {
    if (cards.length <= 1) return false;
    const lowestCard = min(cards);
    const highestCard = max(cards);

    return highestCard.value <= (lowestCard.value + 2);
}

const isValid = (move: Move): boolean => {
    if (move.cards.length === 0) return false;

    return (
        isStraight(move) || 
        isMatchingValue(move) || 
        isPotentialStraight(move) ||
        isSingleCard(move) || 
        isPotentialChop(move)
    );
}

const isPotentialStraight = ({ cards }: Move) => {
    if (cards.length <= 1) return false;

    const lowestCard = min(cards);

    if (cards.every(card => card.suit !== lowestCard.suit)) return false;

    return sortCards(cards).every((current, index, cards) => {
        if (index === 0) return true;
        const previous = cards[index - 1];

        if(current.suit === previous.suit){
            if (current.value === (previous.value + 1)) return true;
        }
        return false;
    }); 
};

const isStraight = ({ cards }: Move) => {
    if (cards.length < 3) return false;

    const lowestCard = min(cards);

    if (cards.every(card => card.suit !== lowestCard.suit)) return false;

    return sortCards(cards).every((current, index, cards) => {
        if (index === 0) return true;
        const previous = cards[index - 1];

        if(current.suit === previous.suit){
            if (current.value === (previous.value + 1)) return true;
        }
        return false;
    });        
}

const isSingleCard = ({ cards }: Move): boolean => cards.length === 1;

const isMatchingValue = ({ cards }: Move) => {
    if (cards.length <= 1) return false;
    
    const firstValue = cards[0].value;

    return cards.every(card => card.value === firstValue);
}