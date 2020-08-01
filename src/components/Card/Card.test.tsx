import React from 'react';
import { render, screen } from "@testing-library/react";
import Card from '.';
import { Card as CardType } from '../../types';
import SuitEnum from '../../enum/Suit';
import ValueEnum from '../../enum/Value';
import { getSuitIcon } from '../../utils/CardUtils';

const { getByText } = screen;

describe('<Card />', () => {
    test('renders correctly', () => {
        const card: CardType = { suit: SuitEnum.Clubs, value: ValueEnum.Ace };
        render(<Card card={card} onClick={jest.fn()} />);

        expect(getByText(getSuitIcon(SuitEnum.Clubs))).toBeInTheDocument();
        expect(getByText(ValueEnum[ValueEnum.Ace])).toBeInTheDocument();
    });
});