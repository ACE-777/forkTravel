CREATE TABLE IF NOT EXISTS projects.countries
(
    countries Nullable(VARCHAR),
    mountain Nullable(VARCHAR),
    sea Nullable(VARCHAR),
    excursion Nullable(VARCHAR),
    health Nullable(VARCHAR)
) ENGINE = Memory;

INSERT INTO projects.countries (countries, mountain, sea, excursion, health) VALUES ('Russia', 'R', 'R', 'R', 'R');
INSERT INTO projects.countries (countries, mountain, sea, excursion, health) VALUES ('Argentina', 'A', 'A', 'A', 'A');
