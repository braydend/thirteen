import React from "react";
import { render, screen } from "@testing-library/react";
import Game from ".";

const { getAllByLabelText } = screen;

describe('<Game />', () => {
    test('renders correctly', async () => {
        render(<Game />);

        expect(getAllByLabelText('player')).toHaveLength(4);
    });
});