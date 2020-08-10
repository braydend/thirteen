import { PrismaClient } from "@prisma/client";

const suitNames = [
    { name: 'Spades', value: 1 },
    { name: 'Clubs', value: 2 },
    { name: 'Diamonds', value: 3 },
    { name: 'Hearts', value: 4 },
];

const cardValues = [
    { name: 'Three', value: 3 },
    { name: 'Four', value: 4 },
    { name: 'Five', value: 5 },
    { name: 'Six', value: 6 },
    { name: 'Seven', value: 7 },
    { name: 'Eight', value: 8 },
    { name: 'Nine', value: 9 },
    { name: 'Ten', value: 10 },
    { name: 'Jack', value: 11 },
    { name: 'Queen', value: 12 },
    { name: 'King', value: 13 },
    { name: 'Ace', value: 14 },
    { name: 'Two', value: 15 },
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
    await suitNames.forEach(async ({ name, value }) => {
        await client.suit.create({ data: { name, value } });
    });

    console.log('--- Seeded Card Suits ---');
};

const seedValues = async (client: PrismaClient): Promise<void> => {
    await cardValues.forEach(async ({ name, value}) => {
        await client.value.create({ data: { name, value } });
    });

    console.log('--- Seeded Card Values ---');
};

export default seedDeck;