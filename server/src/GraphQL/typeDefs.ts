import gql from 'graphql-tag';

export default gql`
    type Query {
        cards: [Card!]!
    }

    type Mutation {
        createPlayer: Player!
        createGame(input: CreateGameInput): Game!
    }

    ## Inputs
    input CreateGameInput {
        playerId: String!
    }

    ## Types
    type Card {
        suit: Suit!
        suitValue: Int!
        value: Value!
        valueValue: Int!
    }

    type Game {
        id: String!
        players: [Player!]!
    }

    type Player {
        id: String!
        # createdAt: DateTime!
        # hand: Hand
        # games: [Game!]!
    }

    type Value {
        name: String!
        value: Int!
    }

    type Suit {
        name: String!
        value: Int!
    }
`;