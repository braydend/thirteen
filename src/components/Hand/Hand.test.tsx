import { render, fireEvent, screen, within, RenderResult, prettyDOM } from "@testing-library/react";
import Hand, { Props } from ".";
import React from "react";
import { Card, Hand as HandType } from "../../types";
import SuitEnum from "../../enum/Suit";
import ValueEnum from "../../enum/Value";
import { getSuitIcon } from "../../utils/CardUtils";


describe('<Hand />', () => {
    const setUp = (customProps?: Partial<Props>): RenderResult => {
        const card: Card = { suit: SuitEnum.Clubs, value: ValueEnum.Ace };
        const hand: HandType = { cards: [card] };
        const defaultProps: Props = {
            hand
        };
        const props = { ...defaultProps, ...customProps };

        return render(<Hand {...props} />);
    };

    describe('renders correctly', () => {
        test('correctly maps cards to components', async () => {
            const kingOfSpades: Card = { suit: SuitEnum.Spades, value: ValueEnum.King };
            const aceOfSpades: Card = { suit: SuitEnum.Spades, value: ValueEnum.Ace };
            const twoOfSpades: Card = { suit: SuitEnum.Spades, value: ValueEnum.Two };
            const hand: HandType = { cards: [kingOfSpades,aceOfSpades,twoOfSpades] };
            const wrapper = setUp({ hand });

            expect(wrapper.getAllByText(getSuitIcon(SuitEnum.Spades))).toHaveLength(3);
        });
    });

    describe('move', () => {
        test('clicking on card moves it to move', async () => {
            setUp();
            const hand = screen.getByLabelText(/hand/);
            // const move = screen.getByLabelText(/move/);
            const card = screen.getByLabelText(/card/);

            // expect(hand.textContent).toContain('Ace');
            // act(() => {
            // console.log(prettyDOM(hand));
// 
               await fireEvent(card, new MouseEvent('click'));
            // });

            // screen.debug();

            // await waitForElementToBeRemoved(card);
            // console.log(prettyDOM(card));
            // console.log(prettyDOM(hand));


            // const move = await screen.findByLabelText(/move/i);
            // await waitFor(() => {
                // within(hand).getByLabelText(/card/);
                // expect(within(hand).queryByLabelText(/card/)).not.toBeInTheDocument();
            // });
            // console.log(move);

        });
        test('clicking on card in move adds it to the hand', () => {
            
        });
    })
});