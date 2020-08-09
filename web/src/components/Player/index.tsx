import React, { useState } from 'react';
import { Hand as HandType, Move as MoveType, Card as CardType } from '../../types';
import Hand from '../Hand';
import Move from '../Move';
import { addCardsToHand, removeCardFromHand } from '../../utils/HandUtils';
import { addCardsToMove, removeCardFromMove } from '../../utils/MoveUtils';

export type Props = {
    name: string;
    hand: HandType;
};

const Player: React.FC<Props> = ({ name, hand: initialHand }) => {
    const [hand, setHand] = useState<HandType>(initialHand);
    const [move, setMove] = useState<MoveType>({ cards: [] });

    const moveCardFromHandToMove = (cardToMove: CardType) => {
        setHand(removeCardFromHand(hand, cardToMove));
        setMove(addCardsToMove(move, [cardToMove]));
    };    
    
    const moveCardFromMoveToHand = (cardToMove: CardType) => {
        setHand(addCardsToHand(hand, [cardToMove]));
        setMove(removeCardFromMove(move, cardToMove));
    };

    return (
        <div aria-label="player">
            {name}
            <Move move={move} onCardClick={moveCardFromMoveToHand} />
            <Hand hand={hand} onCardClick={moveCardFromHandToMove} />
        </div>
    );
};

export default Player;