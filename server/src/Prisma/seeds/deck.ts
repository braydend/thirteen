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
    await seedCards(client);
};

const seedCards = async (client: PrismaClient): Promise<void> => {
    const suits = await client.suit.findMany();
    const values = await client.value.findMany();

    for (const suit of suits) {
        for (const value of values) {
            await client.card.create({ data: { 
                suit: { connect: { name: suit.name } },
                value: { connect: { name: value.name } },
            }}); 
        }
    };

    console.log('--- Seeded Deck ---')
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