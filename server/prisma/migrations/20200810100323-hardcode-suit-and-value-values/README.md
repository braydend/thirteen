# Migration `20200810100323-hardcode-suit-and-value-values`

This migration has been generated by Brayden at 8/10/2020, 8:03:23 PM.
You can check out the [state of the schema](./schema.prisma) after the migration.

## Database Steps

```sql
PRAGMA foreign_keys=OFF;

CREATE TABLE "new_Suit" (
"value" INTEGER NOT NULL,
"name" TEXT NOT NULL,
PRIMARY KEY ("value"))

INSERT INTO "new_Suit" ("value", "name") SELECT "value", "name" FROM "Suit"

PRAGMA foreign_keys=off;
DROP TABLE "Suit";;
PRAGMA foreign_keys=on

ALTER TABLE "new_Suit" RENAME TO "Suit";

CREATE UNIQUE INDEX "Suit.name_unique" ON "Suit"("name")

CREATE TABLE "new_Value" (
"value" INTEGER NOT NULL,
"name" TEXT NOT NULL,
PRIMARY KEY ("value"))

INSERT INTO "new_Value" ("value", "name") SELECT "value", "name" FROM "Value"

PRAGMA foreign_keys=off;
DROP TABLE "Value";;
PRAGMA foreign_keys=on

ALTER TABLE "new_Value" RENAME TO "Value";

CREATE UNIQUE INDEX "Value.name_unique" ON "Value"("name")

PRAGMA foreign_key_check;

PRAGMA foreign_keys=ON;
```

## Changes

```diff
diff --git schema.prisma schema.prisma
migration 20200809113826-card-suit-and-value-unique-constraint..20200810100323-hardcode-suit-and-value-values
--- datamodel.dml
+++ datamodel.dml
@@ -1,7 +1,7 @@
 datasource db {
     provider = "sqlite" 
-    url = "***"
+    url = "***"
 }
 generator client {
     provider = "prisma-client-js"
@@ -47,12 +47,12 @@
     games       Game[]      @relation(references: [id])
 }
 model Suit {
-    value       Int     @id @default(autoincrement())
+    value       Int     @id
     name        String  @unique
 }
 model Value {
-    value       Int     @id @default(autoincrement())
+    value       Int     @id
     name        String  @unique
 }
```


