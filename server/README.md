# Thirteen Server

## Database
Follow the following steps to get the databse up and running: 
- Make sure the dependencies are installed with `npm i`
- Run `npx prisma migrate up --experimental` to create a new database file named `thirteen.db`
- Run `npx prisma generate` to create the Prisma Client
- Run `npm run seed` to seed the database

### Making changes to the database
- Change `prisma/schema.prisma` file to meet your needs
- Run `npx prisma migrate save --experimental` to create a new migration with the changes
- Run `npx prisma migrate up --experimental` to run the mutation
- Run `npx prisma generate` to update the Prisma Client to match the database state
  