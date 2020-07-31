import { Table, Move } from "../types";
import { isBeatenBy } from "./MoveUtils";

export const playMove = (table: Table, move: Move): Table => {
    if (isMovePlayable(table, move)) {
        table.pile.push(move);

        return { ...table, pile: [...table.pile, move], isReset: false };
    }

    throw new Error('Move is not playable');
}

const isMovePlayable = ({ isReset, pile }: Table, move: Move): boolean => {
    const lastMove = pile[pile.length - 1];

    if (isReset) {
        return true;
    }

    return isBeatenBy(lastMove, move);
};