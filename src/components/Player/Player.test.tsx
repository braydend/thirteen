import React from 'react';
import { render, screen, RenderResult, within, prettyDOM } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import Player, { Props } from ".";
import SuitEnum from '../../enum/Suit';
import ValueEnum from '../../enum/Value';

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

        screen.getByText('Blue Player');
        screen.getByLabelText('hand');
        screen.getByLabelText('move');
    });

    test("player's cards start in hand", () => {
        setUp({ hand: { cards: [ { suit: SuitEnum.Spades, value: ValueEnum.Three } ] } });
        const hand = screen.getByLabelText('hand');
        const move = screen.getByLabelText('move');
        
        within(hand).getByLabelText('card');
        expect(() => prettyDOM(within(move).getByLabelText('card'))).toThrow();
    });

    test("card moves between <Hand /> and <Move /> when clicked", () => {
        setUp({ hand: { cards: [ { suit: SuitEnum.Spades, value: ValueEnum.Three } ] } });
        const hand = screen.getByLabelText('hand');
        const move = screen.getByLabelText('move');
        
        within(hand).getByLabelText('card');
        expect(() => within(move).getByLabelText('card')).toThrow();

        userEvent.click(within(hand).getByLabelText('card'));

        within(move).queryByLabelText('card');
        expect(() => within(hand).getByLabelText('card')).toThrow();

        userEvent.click(within(move).getByLabelText('card'));

        within(hand).getByLabelText('card');
        expect(() => within(move).getByLabelText('card')).toThrow();    
    });
});