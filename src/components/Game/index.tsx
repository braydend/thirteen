import React, { useState } from 'react';
import { dealDeck, shuffleDeck, createDeck } from '../../utils/GameUtils';
import { Table, Game as GameType, Move } from '../../types';
import Player from '../Player';
import Card from '../Card';
import { toString } from '../../utils/CardUtils';
import './game.css';
import { playMove } from '../../utils/TableUtils';

const initialTable: Table = { 
    isReset: true, 
    pile: [],
};

const Game: React.FC = () => {
    const [table, setTable] = useState<Table>(initialTable);

    const players: GameType = dealDeck(shuffleDeck(createDeck()));

    const onPlay = (move: Move): void => {
        setTable(playMove(table, move));
    };

    const lastMove = table.pile.length !== 0 ? table.pile[table.pile.length - 1] : undefined;

    return (
        <div>
            <div aria-label="pile" className="pile">
                {lastMove && lastMove.cards.map(card => <Card key={toString(card)} card={card} />)}
            </div>
            {players.map((player, index) => 
                <Player 
                    key={`player-${index}`} 
                    hand={player.hand} 
                    name={`Player ${index}`}
                    onPlay={onPlay} />
            )}
        </div>
    );
};

export default Game;