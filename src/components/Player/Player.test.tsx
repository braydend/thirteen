import React from 'react';
import { render, screen, RenderResult, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import Player, { Props } from ".";
import SuitEnum from '../../enum/Suit';
import ValueEnum from '../../enum/Value';

const { getByText, getByLabelText, queryByLabelText } = screen;

describe('<Player />',() => {
    const setUp = (customProps?: Partial<Props>): RenderResult => {
        const defaultProps: Props = {
            name: "Player",
            hand: { cards:[] },
            onPlay: jest.fn(),
        };
        const props = { ...defaultProps, ...customProps };

        return render(<Player {...props} />);
    };
    test('renders correctly', () => {
        setUp({ name: 'Blue Player' });

        expect(getByText('Blue Player')).toBeInTheDocument();
        expect(getByLabelText('hand')).toBeInTheDocument();
        expect(queryByLabelText('move')).not.toBeInTheDocument();
    });

    test("player's cards start in hand", () => {
        setUp({ hand: { cards: [ { suit: SuitEnum.Spades, value: ValueEnum.Three } ] } });
        const hand = getByLabelText('hand');
        
        expect(within(hand).getByLabelText('card')).toBeInTheDocument();
    });

    test("move renders when cards are added", () => {
        setUp({ hand: { cards: [{ suit: SuitEnum.Spades, value: ValueEnum.Three }] } });

        userEvent.click(screen.getByLabelText('card'));
        expect(getByLabelText('move')).toBeInTheDocument();
    });

    test("card moves between <Hand /> and <Move /> when clicked", () => {
        setUp({ hand: { cards: [ { suit: SuitEnum.Spades, value: ValueEnum.Three } ] } });
        const hand = getByLabelText('hand');
        
        expect(within(hand).getByLabelText('card')).toBeInTheDocument();

        userEvent.click(within(hand).getByLabelText('card'));

        expect(within(getByLabelText('move')).getByLabelText('card')).toBeInTheDocument();
        expect(within(hand).queryByLabelText('card')).not.toBeInTheDocument();

        userEvent.click(within(screen.getByLabelText('move')).getByLabelText('card'));

        expect(within(hand).getByLabelText('card')).toBeInTheDocument();
    });

    test('cards are moved back to hand on reset', () => {
        setUp({ hand: { cards: [ 
            { suit: SuitEnum.Spades, value: ValueEnum.Three },
            { suit: SuitEnum.Spades, value: ValueEnum.Four },
            { suit: SuitEnum.Spades, value: ValueEnum.Five },
        ] } });
        const hand = screen.getByLabelText('hand');

        within(hand).getAllByLabelText('card').forEach(node => userEvent.click(node));
        userEvent.click(screen.getByText('Reset'));

        expect(within(hand).getAllByLabelText('card')).toHaveLength(3);
    });
});