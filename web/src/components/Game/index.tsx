import React from 'react';
import { dealDeck, shuffleDeck, createDeck } from '../../utils/GameUtils';
import { Table, Game as GameType } from '../../types';
import Player from '../Player';

const Game: React.FC = () => {
    const table: Table = { deck: shuffleDeck(createDeck()), isReset: true, pile: [] }
    const players: GameType = dealDeck(table.deck);

    return <div>{players.map((player, index) => <Player key={`player-${index}`} hand={player.hand} name={`Player ${index}`} />)}</div>;
};

export default Game;