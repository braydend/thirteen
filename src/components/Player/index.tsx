import React, { useState } from 'react';
import { Hand as HandType, Move as MoveType, Card as CardType } from '../../types';
import Hand from '../Hand';
import Move from '../Move';
import { addCardsToHand, removeCardFromHand } from '../../utils/HandUtils';
import { addCardsToMove, removeCardFromMove } from '../../utils/MoveUtils';
import './player.css';

export type Props = {
    name: string;
    hand: HandType;
    onPlay: (move: MoveType) => void,
};

const Player: React.FC<Props> = ({ name, hand: initialHand, onPlay }) => {
    const [hand, setHand] = useState<HandType>(initialHand);
    const [move, setMove] = useState<MoveType>({ cards: [] });
    const [error, setError] = useState<Error>();

    const moveCardFromHandToMove = (cardToMove: CardType) => {
        setError(undefined);
        try {
            setMove(addCardsToMove(move, [cardToMove]));
            setHand(removeCardFromHand(hand, cardToMove));
        } catch (e) {
            setError(e);
        }
    };    
    
    const moveCardFromMoveToHand = (cardToMove: CardType) => {
        setError(undefined);
        try {
            setHand(addCardsToHand(hand, [cardToMove]));
            setMove(removeCardFromMove(move, cardToMove));
        } catch (e) {
            setError(e);
        }
    };

    const resetMove = ({ cards }: MoveType) => {
        setHand(addCardsToHand(hand, cards));
        setMove({ cards: [] });
    };

    const playMove = (move: MoveType) => {
        onPlay(move);
        setMove({ cards: [] });
    };

    const hasMadeMove = move.cards.length !== 0;
    const hasError = error !== undefined;

    return (
        <div aria-label="player">
            {hasError && <div className="error">Hold up! {error?.message}</div>}
            {name}
            {hasMadeMove && (
                <Move 
                    move={move} 
                    onCardClick={moveCardFromMoveToHand} 
                    onReset={resetMove} 
                    onPlay={playMove} 
                />
            )}
            <Hand hand={hand} onCardClick={moveCardFromHandToMove} />
        </div>
    );
};

export default Player;