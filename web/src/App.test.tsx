import React from 'react';
import { render } from '@testing-library/react';
import App from './App';

test('renders game title', () => {
  const { getByLabelText } = render(<App />);
  const gameTitle = getByLabelText('game-title');
  
  expect(gameTitle).toBeInTheDocument();
  expect(gameTitle).toHaveTextContent('Thirteen');
});
