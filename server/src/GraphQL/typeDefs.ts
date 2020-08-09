import gql from 'graphql-tag';

export default gql`
    type Query {
        cards: [Card!]!
    }

    type Card {
        suit: Suit!
        suitValue: Int!
        value: Value!
        valueValue: Int!
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