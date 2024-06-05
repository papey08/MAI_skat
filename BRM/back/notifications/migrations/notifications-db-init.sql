CREATE TABLE "notifications" (
    "id" SERIAL PRIMARY KEY,
    "company_id" INTEGER NOT NULL,
    "date" DATE NOT NULL,
    "viewed" BOOLEAN NOT NULL,
    "type" VARCHAR(100) NOT NULL
);

CREATE TABLE "new_lead_notifications" (
    "id" INTEGER NOT NULL,
    "lead_id" INTEGER NOT NULL,
    "client_company" INTEGER NOT NULL
);

CREATE TABLE "closed_lead_notifications" (
    "id" INTEGER NOT NULL,
    "ad_id" INTEGER NOT NULL,
    "producer_company" INTEGER NOT NULL,
    "answered" BOOLEAN NOT NULL
);
