import MoveEnum from "./enum/Move";
import ValueEnum from "./enum/Value";
import SuitEnum from "./enum/Suit";

export type Card = {
    value: ValueEnum,
    suit: SuitEnum,
};

export type Deck = Card[];

export type Game = [Player, Player, Player, Player];

export type Hand = {
    cards: Card[],
};

export type Move = {
    cards: Card[],
    type?: MoveEnum,
};

export type Player = {
    hand: Hand,
    move: Move,
};

export type Table = {
    pile: Move[],
    isReset: boolean,
};
