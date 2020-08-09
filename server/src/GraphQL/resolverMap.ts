import { PrismaClient } from "@prisma/client";

interface Context {
    prisma: PrismaClient,
};

export default {
    Query : {
        cards: async (parent: unknown, args: unknown, context: Context) => context.prisma.card.findMany({ include: { suit: true, value: true }}),
    },
};