import React from 'react';
import { Card as CardType, Move as MoveType } from '../../types';
import Card from '../Card';
import './move.css';

export type Props = {
    move: MoveType;
    onCardClick: (card: CardType) => void;
    onReset: (move: MoveType) => void;
    onPlay: (move: MoveType) => void;
};

const Move: React.FC<Props> = ({ move, onCardClick, onReset, onPlay }) => {
    const { cards, type } = move;
    
    return(
        <div aria-label="move" className="move">
            {type && <p>{type}</p>}
            <div className="cards">
                {cards.map((card, index) => <Card key={`card-${index}`} onClick={() => onCardClick(card)} card={card} />)}
            </div>
            <div>
                <button onClick={() => onPlay(move)}>Play</button>
                <button onClick={() => onReset(move)}>Reset</button>
            </div>
        </div>
    );
};

export default Move;