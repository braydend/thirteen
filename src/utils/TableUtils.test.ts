import { Card, Move, Table } from "../types";
import SuitEnum from "../enum/Suit";
import ValueEnum from "../enum/Value";
import { playMove } from "./TableUtils";
import MoveEnum from "../enum/Move";

describe('tableUtils', () => {
    describe('playMove', () => {
        describe('table is reset', () => {
            let table: Table;;
            
            beforeEach(() => {
                table = {deck: [], isReset: true, pile: []};
            });

            test('single card', () => {
                const aceOfSpades: Card = { suit: SuitEnum.Spades, value: ValueEnum.Ace};
                const move: Move = { cards: [aceOfSpades], type: MoveEnum.SingleCard };

                const result = playMove(table, move);

                expect(result).toStrictEqual(table);
                expect(table.pile).toHaveLength(1);
            });            
            
            test('matching value', () => {
                const aceOfSpades: Card = { suit: SuitEnum.Spades, value: ValueEnum.Ace};
                const aceOfHearts: Card = { suit: SuitEnum.Hearts, value: ValueEnum.Ace};
                const move: Move = { cards: [aceOfSpades, aceOfHearts], type: MoveEnum.MatchingValue };

                const result = playMove(table, move);

                expect(result).toStrictEqual(table);
                expect(table.pile).toHaveLength(1);
            });            
            
            test('straight', () => {
                const twoOfSpades: Card = { suit: SuitEnum.Spades, value: ValueEnum.Two};
                const aceOfSpades: Card = { suit: SuitEnum.Spades, value: ValueEnum.Ace};
                const kingOfSpades: Card = { suit: SuitEnum.Hearts, value: ValueEnum.King};
                const move: Move = { cards: [kingOfSpades, aceOfSpades, twoOfSpades], type: MoveEnum.Straight };
                
                const result = playMove(table, move);

                expect(result).toStrictEqual(table);
                expect(table.pile).toHaveLength(1);
            });
        });
    });
});