import { GraphQLServer } from 'graphql-yoga';
import { PrismaClient } from '@prisma/client';
import typeDefs from './GraphQL/typeDefs';
import resolvers from './GraphQL/resolverMap';

const prisma = new PrismaClient();

const server = new GraphQLServer({ typeDefs, resolvers, context: { prisma } });

server.start(() => console.log('Server running at http://localhost:4000'));