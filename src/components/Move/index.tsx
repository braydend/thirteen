import React from 'react';
import { Card as CardType, Move as MoveType } from '../../types';
import Card from '../Card';
import './move.css';

export type Props = {
    move: MoveType;
    onCardClick: (card: CardType) => void;
};

const Move: React.FC<Props> = ({ move: { cards, type }, onCardClick }) => (
    <div aria-label="move" className="move">
        {type && <p>{type}</p>}
        {cards.map((card, index) => <Card key={`card-${index}`} onClick={() => onCardClick(card)} card={card} />)}
    </div>
);

export default Move;