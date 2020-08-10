import { PrismaClient } from "@prisma/client";

interface Context {
    prisma: PrismaClient,
};

interface CreateGameInput {
    input: {
        playerId: string,
    },
};

export default {
    Query : {
        cards: async (parent: unknown, args: unknown, context: Context) => context.prisma.card.findMany({ include: { suit: true, value: true }}),
    },
    Mutation: {
        createGame: async (parent: unknown, args: CreateGameInput, context: Context) => {
            console.log(args.input.playerId);
            return await context.prisma.game.create({ data: {players:{connect: { id: args.input.playerId}}  } });
        },
        createPlayer: async (parent: unknown, args: unknown, context: Context) => context.prisma.player.create({ data: {} }),
    },
};