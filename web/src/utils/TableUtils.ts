import { Table, Move } from "../types";

export const playMove = (table: Table, move: Move): Table => {
    if (isMovePlayable(table, move)) {
        table.pile.push(move);

        return table;
    }

    throw new Error('Move is not playable');
}

const isMovePlayable = ({ isReset, pile }: Table, { type }: Move): boolean => {
    const lastType = pile[pile.length]?.type;

    if (isReset) {
        return true;
    }

    if (type === lastType){ 
        return true;
    }

    return false;
};