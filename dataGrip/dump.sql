CREATE DATABASE tour;

CREATE TABLE tour.users_info
(
    id_of_tours Nullable(Int32),
    countries Nullable(String)
) ENGINE = Memory;
