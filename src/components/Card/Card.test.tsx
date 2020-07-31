import React from 'react';
import { render } from "@testing-library/react";
import Card from '.';
import { Card as CardType } from '../../types';
import SuitEnum from '../../enum/Suit';
import ValueEnum from '../../enum/Value';
import { getSuitIcon } from '../../utils/CardUtils';

describe('<Card />', () => {
    const card: CardType = { suit: SuitEnum.Clubs, value: ValueEnum.Ace};
    const setUp = () => render(<Card card={card} onClick={jest.fn()} />);
    test('renders correctly', () => {
        const { getByText } = setUp();

        expect(getByText(getSuitIcon(SuitEnum.Clubs))).toBeInTheDocument();
        expect(getByText(ValueEnum[ValueEnum.Ace])).toBeInTheDocument();
    });
});