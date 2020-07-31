import React from 'react';
import { Card as CardType } from '../../types';
import ValueEnum from '../../enum/Value';
import './card.css';
import { getSuitIcon } from '../../utils/CardUtils';
import SuitEnum from '../../enum/Suit';

export type Props = {
    card: CardType,
    onClick?: () => void,
};

const Card: React.FC<Props> = ({ card: { value, suit }, onClick = () => {} }) => {
    const isRed = suit === SuitEnum.Diamonds || suit === SuitEnum.Hearts;

    return (
    <div className={`card ${isRed ? 'red' : 'black'}`} aria-label="card" onClick={onClick}>
        <div>{getSuitIcon(suit)}</div>
        <div>{ValueEnum[value]}</div>
    </div>
);
    };

export default Card;