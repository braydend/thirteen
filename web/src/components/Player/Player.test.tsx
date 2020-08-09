import React from 'react';
import { render, screen, RenderResult, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import Player, { Props } from ".";
import SuitEnum from '../../enum/Suit';
import ValueEnum from '../../enum/Value';

const { getByText, getByLabelText } = screen;

describe('<Player />',() => {
    const setUp = (customProps?: Partial<Props>): RenderResult => {
        const defaultProps: Props = {
            name: "Player",
            hand: { cards:[] },
        };
        const props = { ...defaultProps, ...customProps };

        return render(<Player {...props} />);
    };
    test('renders correctly', () => {
        setUp({ name: 'Blue Player' });

        getByText('Blue Player');
        getByLabelText('hand');
        getByLabelText('move');
    });

    test("player's cards start in hand", () => {
        setUp({ hand: { cards: [ { suit: SuitEnum.Spades, value: ValueEnum.Three } ] } });
        const hand = getByLabelText('hand');
        const move = getByLabelText('move');
        
        expect(within(hand).getByLabelText('card')).toBeInTheDocument();
        expect(within(move).queryByLabelText('card')).not.toBeInTheDocument();
    });

    test("card moves between <Hand /> and <Move /> when clicked", () => {
        setUp({ hand: { cards: [ { suit: SuitEnum.Spades, value: ValueEnum.Three } ] } });
        const hand = getByLabelText('hand');
        const move = getByLabelText('move');
        
        expect(within(hand).getByLabelText('card')).toBeInTheDocument();
        expect(within(move).queryByLabelText('card')).not.toBeInTheDocument();

        userEvent.click(within(hand).getByLabelText('card'));

        expect(within(move).getByLabelText('card')).toBeInTheDocument();
        expect(within(hand).queryByLabelText('card')).not.toBeInTheDocument();

        userEvent.click(within(move).getByLabelText('card'));

        expect(within(hand).getByLabelText('card')).toBeInTheDocument();
        expect(within(move).queryByLabelText('card')).not.toBeInTheDocument();
    });
});