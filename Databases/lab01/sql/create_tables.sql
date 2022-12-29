CREATE TABLE "Driver" (
    "driver_id" SERIAL PRIMARY KEY NOT NULL,
    "driver_second_name" VARCHAR(30) NOT NULL,
    "driver_name" VARCHAR(30) NOT NULL,
    "driver_third_name" VARCHAR(30) NOT NULL,
    "driver_class" VARCHAR(30) NOT NULL,
    "vehicle_sigh" VARCHAR(9) NOT NULL
);

CREATE TABLE "Types" (
    "type_id" SERIAL PRIMARY KEY NOT NULL,
    "type_name" VARCHAR(30) NOT NULL,
    "class" VARCHAR(30) NOT NULL,
    "capacity" INTEGER NOT NULL,
    "price" FLOAT NOT NULL
);

CREATE TABLE "Vehicle" (
    "vehicle_sigh" VARCHAR(9) PRIMARY KEY NOT NULL,
    "model" VARCHAR(30) NOT NULL,
    "type_id" INTEGER NOT NULL,
    "price_coeff" FLOAT NOT NULL
);

CREATE TABLE "Voyage" (
    "voyage_id" SERIAL PRIMARY KEY NOT NULL,
    "driver_id" INTEGER NOT NULL,
    "point_begin" VARCHAR(50) NOT NULL,
    "point_end" VARCHAR(50) NOT NULL,
    "date_begin" DATE NOT NULL,
    "date_end" DATE NOT NULL
);

ALTER TABLE "Voyage"
    ADD FOREIGN KEY ("driver_id")
        REFERENCES "Driver"("driver_id");

ALTER TABLE "Driver"
    ADD FOREIGN KEY ("vehicle_sigh")
        REFERENCES "Vehicle"("vehicle_sigh");

ALTER TABLE "Vehicle"
    ADD FOREIGN KEY ("type_id")
        REFERENCES "Types"("type_id");
