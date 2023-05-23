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

INSERT INTO projects.countries (countries, mountain, sea, excursion, health, visa, continent, info) VALUES ('Аргентниа', 'Серро-Катадраль', 'Мар-Дель-Плато', 'Боэносариес','Бонэносариес', 'Безвизовая','Южная Америка','Ушуая, Эль-Калафата');
INSERT INTO projects.countries (countries, mountain, sea, excursion, health, visa, continent, info) VALUES ('Бразилия', 'Нет', 'Бузиос', 'Рио де Жанейро', 'Нет','Безвизовая','Южная Америка','Амазония, Игуасу, карнавал');
