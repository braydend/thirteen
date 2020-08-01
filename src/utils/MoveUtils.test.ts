import { Card, Move } from "../types";
import SuitEnum from "../enum/Suit";
import ValueEnum from "../enum/Value";
import { 
    addCardsToMove, 
    createMoveFromCards, 
    getMoveType, 
    isChop, 
    isBeatenBy,
    removeCardFromMove, 
} from "./MoveUtils";
import MoveEnum from "../enum/Move";

describe('MoveUtils', () => {
    describe('createMoveFromCards', () => {
        test('single card', () => {
            const card: Card = { suit: SuitEnum.Hearts, value: ValueEnum.Two };

            const result = createMoveFromCards([card]);

            expect(result.cards).toHaveLength(1);
            expect(result.cards).toContain(card);
            expect(result.type).toBe(MoveEnum.SingleCard);
        });
    });    
    
    describe('removeCardFromMove', () => {
        test('remove card succesfully', () => {
            const card: Card = { suit: SuitEnum.Hearts, value: ValueEnum.Two };
            const move: Move = { cards: [card], type: MoveEnum.SingleCard };  
            
            const result = removeCardFromMove(move, card);
            
            expect(result.cards).toHaveLength(0);
            expect(result.cards).not.toContain(card);
            expect(result.type).not.toBe(MoveEnum.SingleCard);
        });
        test('throws error if card is not in move', () => {
            const missingCard: Card = { suit: SuitEnum.Spades, value: ValueEnum.Three };
            const presentCard: Card = { suit: SuitEnum.Hearts, value: ValueEnum.Two };
            const move: Move = { cards: [presentCard] };  

            expect(() => {
                removeCardFromMove(move, missingCard);
            }).toThrowError('Error: Card Three of Spades is not in the move');
        });
    });
    describe('addCardsToMove', () => {
        describe('valid moves', () => {
            test('single card', () => {
                const card: Card = { suit: SuitEnum.Hearts, value: ValueEnum.Two };
                const move: Move = { cards: [] };

                const result = addCardsToMove(move, [card]);

                expect(result.cards).toHaveLength(1);
                expect(result.cards).toContain(card);
                expect(result.type).toBe(MoveEnum.SingleCard);
            });
        });

        describe('illegal moves', () => {
            test('cannot add two unrelated cards to a move ([A,7])', () => {
                const aceOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Ace };
                const sevenOfHearts = { suit: SuitEnum.Hearts, value: ValueEnum.Seven };
                const move: Move = { cards: [] };
    
                expect(() => { 
                    addCardsToMove(move, [sevenOfHearts, aceOfSpades]);
                }).toThrow();
            });
        });
    });

    describe('getMoveType', () => {
        test('single card', () => {
            const card: Card = { suit: SuitEnum.Hearts, value: ValueEnum.Two };
            const move: Move = { cards: [card]};

            const result = getMoveType(move);

            expect(result).toBe(MoveEnum.SingleCard);
        });
    });
    
    describe('isChop', () => {
        test('four of a kind', () => {
            const fourOfAKind: Move = { cards: [
                { suit: SuitEnum.Spades, value: ValueEnum.Two },
                { suit: SuitEnum.Clubs, value: ValueEnum.Two },
                { suit: SuitEnum.Diamonds, value: ValueEnum.Two },
                { suit: SuitEnum.Hearts, value: ValueEnum.Two },
            ]};

            expect(isChop(fourOfAKind)).toBe(true);
        });
        test('three consecutive pairs', () => {
            const threeThreeFourFourFiveFive: Move = { cards: [
                { suit: SuitEnum.Spades, value: ValueEnum.Three },
                { suit: SuitEnum.Clubs, value: ValueEnum.Three },
                { suit: SuitEnum.Spades, value: ValueEnum.Four },
                { suit: SuitEnum.Clubs, value: ValueEnum.Four },
                { suit: SuitEnum.Spades, value: ValueEnum.Five },
                { suit: SuitEnum.Clubs, value: ValueEnum.Five },
            ]};

            expect(isChop(threeThreeFourFourFiveFive)).toBe(true);
        });
    });

    describe('isBeatenBy', () => {
        // Spades
        let threeOfSpades: Card;
        let fourOfSpades: Card;
        let fiveOfSpades: Card;
        let sixOfSpades: Card;
        let sevenOfSpades: Card;
        // Clubs
        let threeOfClubs: Card;
        let fiveOfClubs: Card;
        // Diamonds
        let threeOfDiamonds: Card;
        let fiveOfDiamonds: Card;
        // Hearts
        let threeOfHearts: Card;
        let fourOfHearts: Card;
        let fiveOfHearts: Card;
        let sixOfHearts: Card;
        let sevenOfHearts: Card;

        const assertABeatsB = (A: Move, B: Move): void => {
            expect(isBeatenBy(B, A)).toBe(true);  
            expect(isBeatenBy(A, B)).toBe(false);  
        };

        beforeEach(() => {
            threeOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Three };
            fourOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Four };
            fiveOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Five };
            sixOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Six };
            sevenOfSpades = { suit: SuitEnum.Spades, value: ValueEnum.Seven };
            threeOfHearts = { suit: SuitEnum.Hearts, value: ValueEnum.Three };
            fourOfHearts = { suit: SuitEnum.Hearts, value: ValueEnum.Four };
            fiveOfHearts = { suit: SuitEnum.Hearts, value: ValueEnum.Five };
            sixOfHearts = { suit: SuitEnum.Hearts, value: ValueEnum.Six };
            sevenOfHearts = { suit: SuitEnum.Hearts, value: ValueEnum.Seven };
            threeOfClubs = { suit: SuitEnum.Spades, value: ValueEnum.Three };
            threeOfDiamonds = { suit: SuitEnum.Spades, value: ValueEnum.Three };
            fiveOfClubs = { suit: SuitEnum.Clubs, value: ValueEnum.Five };
            fiveOfDiamonds = { suit: SuitEnum.Diamonds, value: ValueEnum.Five };
        });
        describe('single card', () => {
            test('different values ([3] vs [7])', () => {
                const three: Move = { cards: [threeOfSpades] };
                const seven: Move = { cards: [sevenOfSpades] };

                assertABeatsB(seven, three);
            });

            test('same value ([5] vs [5])', () => {
                const heart: Move = { cards: [fiveOfHearts] };
                const spade: Move = { cards: [fiveOfSpades] };

                assertABeatsB(heart, spade);
            });
        });
        describe('matching value', () => {
            test('pair of different value ([3,3] vs [5,5])', () => {
                const pairOfFives: Move = { cards: [
                    fiveOfHearts,
                    fiveOfClubs,
                ]};

                const pairOfThrees: Move = { cards: [
                    threeOfSpades,
                    threeOfHearts,
                ]};

                assertABeatsB(pairOfFives, pairOfThrees);
            });            
            test('pair of same value ([5,5] vs [5,5])', () => {
                const highPairOfFives: Move = { cards: [
                    fiveOfHearts,
                    fiveOfClubs,
                ]};   

                const lowPairOfFives: Move = { cards: [
                    fiveOfDiamonds,
                    fiveOfSpades,
                ]};

                assertABeatsB(highPairOfFives, lowPairOfFives);
            });
            test('three of a kind ([3,3,3] vs [4,4,4])', () => {
                const threeThrees: Move = { cards: [
                    threeOfHearts,
                    threeOfDiamonds,
                    threeOfClubs,
                ]};          
                const threeFives: Move = { cards: [
                    fiveOfClubs,
                    fiveOfDiamonds,
                    fiveOfHearts,
                ]};

                assertABeatsB(threeFives, threeThrees);
            });
            test('four of a kind (CHOP) ([3,3,3,3] vs [4,4,4,4])', () => {
                const fourThrees: Move = { cards: [
                    threeOfClubs,
                    threeOfDiamonds,
                    threeOfHearts,
                    threeOfSpades,
                ]};            
                const fourFives: Move = { cards: [
                    fiveOfSpades,
                    fiveOfHearts,
                    fiveOfDiamonds,
                    fiveOfClubs,
                ]};

                assertABeatsB(fourFives, fourThrees);
            });
        });
        describe('straight', () => {
            test('different amount of cards is not a valid move ([3,4,5] vs [3,4,5,6])', () => {
                const spades: Move = { cards: [
                    threeOfSpades,
                    fourOfSpades,
                    fiveOfSpades,
                ]};
                const hearts: Move = { cards: [
                    threeOfHearts,
                    fourOfHearts,
                    fiveOfHearts,
                    sixOfHearts,
                ]};

                expect(isBeatenBy(hearts, spades)).toBe(false);                
                expect(isBeatenBy(spades, hearts)).toBe(false); 
            });

            describe('same value', () => {
                test('three cards ([3,4,5] vs [3,4,5])', () => {
                    const spades: Move = { cards: [
                        threeOfSpades,
                        fourOfSpades,
                        fiveOfSpades,
                    ]};
                    const hearts: Move = { cards: [
                        threeOfHearts,
                        fourOfHearts,
                        fiveOfHearts,
                    ]};

                    assertABeatsB(hearts, spades);
                });
                test('four cards ([3,4,5,6] vs [3,4,5,6])', () => {
                    const spades: Move = { cards: [
                        threeOfSpades,
                        fourOfSpades,
                        fiveOfSpades,
                        sixOfSpades,
                    ]};
                    const hearts: Move = { cards: [
                        threeOfHearts,
                        fourOfHearts,
                        fiveOfHearts,
                        sixOfHearts,
                    ]};

                    assertABeatsB(hearts, spades);                
                });
                test('five cards ([3,4,5,6,7] vs [3,4,5,6,7])', () => {
                    const spades: Move = { cards: [
                        threeOfSpades,
                        fourOfSpades,
                        fiveOfSpades,
                        sixOfSpades,
                        sevenOfSpades,
                    ]};
                    const hearts: Move = { cards: [
                        threeOfHearts,
                        fourOfHearts,
                        fiveOfHearts,
                        sixOfHearts,
                        sevenOfHearts,
                    ]};

                    assertABeatsB(hearts, spades);
                });
            });
            describe('different value', () => {
                test('higher value ([3,4,5] vs [K,A,2])', () => {
                    const threeFourFive: Move = { cards: [
                        threeOfSpades,
                        fourOfSpades,
                        fiveOfSpades,
                    ]};                    
                    const fiveSixSeven: Move = { cards: [
                        fiveOfSpades,
                        sixOfSpades,
                        sevenOfSpades,
                    ]};

                    assertABeatsB(fiveSixSeven, threeFourFive);
                });
            });
        });

        describe('chop', () => {
            describe('four of a kind', () => {
                let fourThrees: Move;

                beforeEach(() => {
                    fourThrees = { cards: [
                        threeOfSpades,
                        threeOfHearts,
                        threeOfClubs,
                        threeOfDiamonds,
                    ]};
                });

                test('single card ([2] vs [3,3,3,3]', () => {
                    const twoOfHearts: Move = { cards: [{ suit: SuitEnum.Hearts, value: ValueEnum.Two }] };
                   
                    assertABeatsB(fourThrees, twoOfHearts);
                });
                describe('matching values', () => {
                    test('pair ([2,2] vs [3,3,3,3])', () => {
                        const pairOfTwos: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Two },
                        ]};

                        assertABeatsB(fourThrees, pairOfTwos);
                    });
                    test('three of a kind ([2,2,2] vs [3,3,3,3])', () => {
                        const threeTwos: Move = { cards: [
                            { suit: SuitEnum.Clubs, value: ValueEnum.Two },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Two },
                        ]};

                        assertABeatsB(fourThrees, threeTwos);
                    });
                });
                describe('straight', () => {
                    test('three card straight ([K,A,2] vs [3,3,3,3])', () => {
                        const threeCardStraight: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.King },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Ace },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
                        ]};

                        assertABeatsB(fourThrees, threeCardStraight);
                    });
                    test('four card straight ([Q,K,A,2] vs [3,3,3,3])', () => {
                        const fourCardStraight: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Queen },
                            { suit: SuitEnum.Hearts, value: ValueEnum.King },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Ace }, 
                            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
                        ]};

                        assertABeatsB(fourThrees, fourCardStraight);
                    });
                    test('five card straight([J,Q,K,A,2] vs [3,3,3,3])', () => {
                        const fiveCardStraight: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Jack },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Queen },
                            { suit: SuitEnum.Hearts, value: ValueEnum.King },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Ace },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
                        ]};

                        assertABeatsB(fourThrees, fiveCardStraight);
                    });
                });
                describe('chop', () => {
                    test('four of a kind ([3,3,3,3] vs [4,4,4,4])', () => {
                        const fourFours: Move = { cards: [
                            { suit: SuitEnum.Clubs, value: ValueEnum.Four },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Four },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Four },
                            { suit: SuitEnum.Spades, value: ValueEnum.Four },
                        ]};

                        assertABeatsB(fourFours, fourThrees);  
                    });
                    test('three consecutive pairs ([3,3,3,3] vs [3,3,4,4,5,5])', () => {
                        const threeConsecutivePairs: Move = { cards: [
                            { suit: SuitEnum.Clubs, value: ValueEnum.Three },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Three },
                            { suit: SuitEnum.Clubs, value: ValueEnum.Four },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Four },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Five },
                            { suit: SuitEnum.Spades, value: ValueEnum.Five },
                        ]};

                        assertABeatsB(threeConsecutivePairs, fourThrees);
                    });
                });
            });

            describe('three conseutive pairs', () => {
                let threeThreeFourFourFiveFive: Move;

                beforeEach(() => {
                    threeThreeFourFourFiveFive = { cards: [
                        threeOfSpades,
                        threeOfHearts,
                        fourOfHearts,
                        fourOfSpades,
                        fiveOfHearts,
                        fiveOfDiamonds,
                    ]};
                });
                test('single card ([2] vs [3,3,4,4,5,5]', () => {
                    const two: Move = { cards: [{ suit: SuitEnum.Hearts, value: ValueEnum.Two }] };

                    assertABeatsB(threeThreeFourFourFiveFive, two);
                });
                describe('matching values', () => {
                    test('pair ([2,2] vs [3,3,4,4,5,5])', () => {
                        const pairOfTwos: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Two },
                        ]};

                        assertABeatsB(threeThreeFourFourFiveFive, pairOfTwos);
                    });
                    test('three of a kind ([2,2,2] vs [3,3,4,4,5,5])', () => {
                        const threeTwos: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Two },
                            { suit: SuitEnum.Clubs, value: ValueEnum.Two },
                        ]};

                        assertABeatsB(threeThreeFourFourFiveFive, threeTwos);
                    });
                });
                describe('straight', () => {
                    test('three card straight ([K,A,2] vs [3,3,4,4,5,5])', () => {
                        const threeCardStraight: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Ace },
                            { suit: SuitEnum.Hearts, value: ValueEnum.King },
                        ]};
                        
                        assertABeatsB(threeThreeFourFourFiveFive, threeCardStraight);
                    });
                    test('four card straight ([Q,K,A,2] vs [3,3,4,4,5,5])', () => {
                        const fourCardStraight: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Ace },
                            { suit: SuitEnum.Hearts, value: ValueEnum.King },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Queen },
                        ]};

                        assertABeatsB(threeThreeFourFourFiveFive, fourCardStraight);
                    });
                    test('five card straight([J,Q,K,A,2] vs [3,3,4,4,5,5])', () => {
                        const fiveCardStraight: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Two },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Ace },
                            { suit: SuitEnum.Hearts, value: ValueEnum.King },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Queen },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Jack },
                        ]};

                        assertABeatsB(threeThreeFourFourFiveFive, fiveCardStraight); 
                    });
                });
                describe('chop', () => {
                    test('four of a kind ([6,6,6,6] vs [3,3,4,4,5,5])', () => {
                        const fourSixes: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Six },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Six },
                            { suit: SuitEnum.Clubs, value: ValueEnum.Six },
                            { suit: SuitEnum.Spades, value: ValueEnum.Six },
                        ]};

                        assertABeatsB(fourSixes, threeThreeFourFourFiveFive);  
                    });
                    
                    test('four of a kind ([4,4,4,4] vs [3,3,4,4,5,5])', () => {
                        const fourFours: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Four },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Four },
                            { suit: SuitEnum.Clubs, value: ValueEnum.Four },
                            { suit: SuitEnum.Spades, value: ValueEnum.Four },
                        ]};

                        assertABeatsB(threeThreeFourFourFiveFive, fourFours);  
                    });
                    test('three consecutive pairs ([4,4,5,5,6,6] vs [3,3,4,4,5,5])', () => {
                        const fourFourFiveFiveSixSix: Move = { cards: [
                            { suit: SuitEnum.Hearts, value: ValueEnum.Four },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Four },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Five },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Five },
                            { suit: SuitEnum.Hearts, value: ValueEnum.Six },
                            { suit: SuitEnum.Diamonds, value: ValueEnum.Six },
                        ]};

                        assertABeatsB(fourFourFiveFiveSixSix, threeThreeFourFourFiveFive);
                    });
                });
            });
        });
    });
});