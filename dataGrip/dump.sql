CREATE TABLE IF NOT EXISTS projects.countries
(
    countries Nullable(VARCHAR),
    mountain Nullable(VARCHAR),
    sea Nullable(VARCHAR),
    excursion Nullable(VARCHAR),
    health Nullable(VARCHAR),
    visa Nullable(VARCHAR),
    continent Nullable(VARCHAR),
    info Nullable(VARCHAR)

) ENGINE = Memory;

INSERT INTO projects.countries (countries, mountain, visa, continent, info) VALUES ('Аргентниа', 'Серро-Катадраль', 'Безвизовая','Южная Америка','Ушуая, Эль-Калафата');
INSERT INTO projects.countries (countries, sea,  visa, continent, info) VALUES ('Аргентниа', 'Мар-Дель-Плато', 'Безвизовая','Южная Америка','Ушуая, Эль-Калафата');
INSERT INTO projects.countries (countries, excursion, visa, continent, info) VALUES ('Аргентниа','Боэносариес', 'Безвизовая','Южная Америка','Ушуая, Эль-Калафата');
INSERT INTO projects.countries (countries, health, visa, continent, info) VALUES ('Аргентниа','Бонэносариес', 'Безвизовая','Южная Америка','Ушуая, Эль-Калафата');

INSERT INTO projects.countries (countries, sea, visa, continent, info) VALUES ('Бразилия', 'Бузиос','Безвизовая','Южная Америка','Амазония, Игуасу, карнавал');
INSERT INTO projects.countries (countries, excursion, visa, continent, info) VALUES ('Бразилия', 'Рио де Жанейро','Безвизовая','Южная Америка','Амазония, Игуасу, карнавал');
-- INSERT INTO projects.countries (countries, mountain, sea, excursion, health, visa, continent, info) VALUES ('Аргентниа', 'Серро-Катадраль', 'Мар-Дель-Плато', 'Боэносариес','Бонэносариес', 'Безвизовая','Южная Америка','Ушуая, Эль-Калафата');

-- INSERT INTO projects.countries (countries, sea, excursion, visa, continent, info) VALUES ('Бразилия', 'Бузиос', 'Рио де Жанейро','Безвизовая','Южная Америка','Амазония, Игуасу, карнавал');
