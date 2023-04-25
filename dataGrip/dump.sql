CREATE DATABASE tour;

CREATE TABLE tour.countries
(
    countries Nullable(VARCHAR),
    mountain Nullable(VARCHAR),
    sea Nullable(VARCHAR),
    excursion Nullable(VARCHAR),
    health Nullable(VARCHAR)
) ENGINE = Memory;

INSERT INTO tour.countries (countries, mountain, sea, excursion, health) VALUES ('Russia', 'R', 'R', 'R', 'S');
INSERT INTO tour.countries (countries, mountain, sea, excursion, health) VALUES ('Argentina', 'A', 'A', 'A', 'A');
