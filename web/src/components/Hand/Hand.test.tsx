import { render, screen, RenderResult } from "@testing-library/react";
import Hand, { Props } from ".";
import React from "react";
import { Card, Hand as HandType } from "../../types";
import SuitEnum from "../../enum/Suit";
import ValueEnum from "../../enum/Value";
import { getSuitIcon } from "../../utils/CardUtils";

const { getAllByText } = screen;

describe('<Hand />', () => {
    const setUp = (customProps?: Partial<Props>): RenderResult => {
        const card: Card = { suit: SuitEnum.Clubs, value: ValueEnum.Ace };
        const hand: HandType = { cards: [card] };
        const defaultProps: Props = {
            hand,
            onCardClick: jest.fn(),
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
            setUp({ hand });

            expect(getAllByText(getSuitIcon(SuitEnum.Spades))).toHaveLength(3);
        });
    });
});