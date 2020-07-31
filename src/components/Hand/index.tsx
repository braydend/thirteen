import React from 'react';
import { Card as CardType, Hand as HandType } from '../../types';
import Card from '../Card';
import './hand.css';
import ValueEnum from '../../enum/Value';
import SuitEnum from '../../enum/Suit';
import { sortCards } from '../../utils/CardUtils';

export type Props = { 
    hand: HandType,
    onCardClick: (card: CardType) => void,
};

const Hand: React.FC<Props> = ({ hand, onCardClick }) => {
    const cards = hand.cards;

    return (
        <div>
            <div aria-label="hand" className="hand">
                {sortCards(cards).map(card => <Card key={`${ValueEnum[card.value]}-${SuitEnum[card.suit]}`} card={card} onClick={() => onCardClick(card)} />)}
            </div>
        </div>
    );
};

export default Hand;