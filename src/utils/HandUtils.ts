import { Card, Hand } from "../types";
import { toString } from "./CardUtils";

export const addCardsToHand = (hand: Hand, cards: Card[]): Hand => {
    cards.forEach(newCard => {
        hand.cards.forEach(card => {
            // Stringify and compare cards
            // Simplest way to provide object equality checks
            if (toString(card) === toString(newCard)){
                throw new Error(`Invalid Deck: ${toString(newCard)} cannot be added to hand as it already exists in this hand.`);
            }
        });
    });

    return { ...hand, cards: hand.cards.concat(cards) };
};

export const removeCardFromHand = (hand: Hand, cardToRemove: Card): Hand => {
    if (!hand.cards.find(card => 
        // Stringify and compare cards
        // Simplest way to provide object equality checks
        toString(card) === toString(cardToRemove)
            
    )) {
        throw new Error(`Error: Cannot remove ${toString(cardToRemove)} from hand as it does not exist in the hand.`);
    }

    return { ...hand, cards: hand.cards.filter(card => toString(card) !== toString(cardToRemove)) };
};