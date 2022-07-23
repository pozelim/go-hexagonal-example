-- CreateTable
CREATE TABLE IF NOT EXISTS estate (
                                      "estate_id" TEXT NOT NULL,
                                      "name" TEXT NOT NULL,
                                      "lat" FLOAT NOT NULL,
                                      "lon" FLOAT NOT NULL,
                                      "area" INTEGER NOT null,

                                      PRIMARY KEY ("estate_id")
);



