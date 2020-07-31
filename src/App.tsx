import React from 'react';
import './App.css';
import Game from './components/Game';

function App() {
  return (
    <div className="App">
      <header aria-label="game-title">
        Thirteen
      </header>
      <Game />
    </div>
  );
}

export default App;
