import React from "react";
import { render, RenderResult, screen } from "@testing-library/react";
import Game from ".";

describe('<Game />', () => {
    const setUp = (): RenderResult => render(<Game />);

    test('renders correctly', async () => {
        setUp();

        expect(screen.getAllByLabelText('player')).toHaveLength(4);
    });
});