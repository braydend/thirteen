import { PrismaClient } from '@prisma/client';
import seedDeck from './seeds/deck';

// For seeding a fresh database
const seed = async (client: PrismaClient) => {
    await seedDeck(client);
};

const seedDatabase = async () => {
    const prisma = new PrismaClient();

    try {
        await seed(prisma);
        await prisma.$disconnect();

        console.log('--- Seeding Complete ---');
        return;
    } catch (e){
        throw e;
    }
};

seedDatabase();
