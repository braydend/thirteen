import React from "react";
import { RenderResult, render, screen } from "@testing-library/react";
import Move, { Props } from "./index";
import SuitEnum from "../../enum/Suit";
import ValueEnum from "../../enum/Value";
import { Card, Move as MoveType } from "../../types";
import MoveEnum from "../../enum/Move";

const setUp = (customProps?: Partial<Props>): RenderResult => {
    const defaultProps: Props = { move: { cards: [] }, onCardClick: jest.fn() };
    const props = { ...defaultProps, ...customProps };

    return render(<Move {...props} />);
};

describe('<Move />', () => {
    test('renders correctly', () => {
        const { getByLabelText } = setUp();
        const moveContainer = getByLabelText('move');

        expect(moveContainer).toBeInTheDocument();
    });

    describe('renders move type', () => {
        test('single card', () => {
            const move: MoveType = { cards: [], type: MoveEnum.SingleCard };
            setUp({ move });

            screen.getByText(MoveEnum.SingleCard);
        });
        test('matching value', () => {
            const move: MoveType = { cards: [], type: MoveEnum.MatchingValue };
            setUp({ move });

            screen.getByText(MoveEnum.MatchingValue);
        });
        test('straight', () => {
            const move: MoveType = { cards: [], type: MoveEnum.Straight };
            setUp({ move });

            screen.getByText(MoveEnum.Straight);
        });
        test('potential three consecutive pairs (CHOP)', () => {
            const move: MoveType = { cards: [], type: MoveEnum.ThreeConsecutivePairs };
            setUp({ move });

            screen.getByText(MoveEnum.ThreeConsecutivePairs);
        });
        test('three consecutive pairs (CHOP)', () => {
            const move: MoveType = { cards: [], type: MoveEnum.ThreeConsecutivePairs };
            setUp({ move });

            screen.getByText(MoveEnum.ThreeConsecutivePairs);
        });
    });

    test('renders cards correctly', () => {
        const cards: Card[] = [
            { suit: SuitEnum.Diamonds, value: ValueEnum.Ace },
            { suit: SuitEnum.Diamonds, value: ValueEnum.Two },
        ];
        const { getAllByLabelText } = setUp({ move: { cards } });
        const cardComponents = getAllByLabelText('card');

        expect(cardComponents).toHaveLength(2);
    });
});