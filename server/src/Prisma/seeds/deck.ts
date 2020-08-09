import { PrismaClient } from "@prisma/client";

const suitNames = [
    'Spades',
    'Clubs',
    'Diamonds',
    'Hearts',
];

const cardValues = [
    'Three',
    'Four',
    'Five',
    'Six',
    'Seven',
    'Eight',
    'Nine',
    'Ten',
    'Jack',
    'Queen',
    'King',
    'Ace',
    'Two',
];

const seedDeck = async (client: PrismaClient): Promise<void> => {
    await seedSuits(client);
    await seedValues(client);
};

const seedSuits = async (client: PrismaClient): Promise<void> => {
    await suitNames.forEach(async (name) => {
        await client.suit.create({ data: { name } });
    });

    console.log('--- Seeded Card Suits ---');
};

const seedValues = async (client: PrismaClient): Promise<void> => {
    await cardValues.forEach(async (value) => {
        await client.value.create({ data: { name: value } });
    });

    console.log('--- Seeded Card Values ---');
};

export default seedDeck;