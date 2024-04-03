/*
  Warnings:

  - A unique constraint covering the columns `[owner_id]` on the table `events` will be added. If there are existing duplicate values, this will fail.
  - Changed the type of `owner_id` on the `events` table. No cast exists, the column would be dropped and recreated, which cannot be done if there is data, since the column is required.

*/
-- DropForeignKey
ALTER TABLE "events" DROP CONSTRAINT "events_owner_id_fkey";

-- AlterTable
ALTER TABLE "events" ADD COLUMN     "profileId" INTEGER,
DROP COLUMN "owner_id",
ADD COLUMN     "owner_id" UUID NOT NULL;

-- AlterTable
ALTER TABLE "profiles" ADD COLUMN     "avatarUUID" UUID;

-- CreateIndex
CREATE UNIQUE INDEX "events_owner_id_key" ON "events"("owner_id");

-- AddForeignKey
ALTER TABLE "events" ADD CONSTRAINT "events_profileId_fkey" FOREIGN KEY ("profileId") REFERENCES "profiles"("id") ON DELETE SET NULL ON UPDATE CASCADE;
