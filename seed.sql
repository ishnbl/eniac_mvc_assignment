CREATE TABLE IF NOT EXISTS users (
  id               BIGSERIAL PRIMARY KEY,
  username         TEXT UNIQUE NOT NULL,
  name             TEXT,
  hashed_password  TEXT,
  fights_won       INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS villages (
  id                BIGSERIAL PRIMARY KEY,
  user_id           BIGINT UNIQUE REFERENCES users(id),
  village_level     INTEGER DEFAULT 1,
  gold              INTEGER DEFAULT 0,
  oil               INTEGER DEFAULT 0,
  money             INTEGER DEFAULT 0,
  farm_land         INTEGER DEFAULT 0,
  mines             INTEGER DEFAULT 0,
  level_constraints JSONB
);

CREATE TABLE IF NOT EXISTS defenses (
  id               BIGSERIAL PRIMARY KEY,
  type             TEXT UNIQUE NOT NULL,
  defensive_power  INTEGER DEFAULT 0,
  capabilities     JSONB,
  attack_power     INTEGER DEFAULT 0,
  cost             INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS troops (
  id               BIGSERIAL PRIMARY KEY,
  type             TEXT,
  health           INTEGER DEFAULT 0,
  offensive_power  INTEGER DEFAULT 0,
  capabilities     JSONB,
  level            INTEGER DEFAULT 1,
  cost             INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS battle_replays (
  id             BIGSERIAL PRIMARY KEY,
  attacker_id    BIGINT REFERENCES villages(id),
  defender_id    BIGINT REFERENCES villages(id),
  attacker_loot  INTEGER DEFAULT 0,
  defender_loot  INTEGER DEFAULT 0,
  winner         BOOLEAN DEFAULT false
);

CREATE TABLE IF NOT EXISTS replay_defenses (
  id                BIGSERIAL PRIMARY KEY,
  attacks_used      INTEGER DEFAULT 0,
  battle_replay_id  BIGINT REFERENCES battle_replays(id),
  defenses_id       BIGINT REFERENCES defenses(id),
  num_deployed      INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS replay_troops (
  id                  BIGSERIAL PRIMARY KEY,
  amount_deployed     INTEGER DEFAULT 0,
  attacks_allocated   INTEGER DEFAULT 0,
  battle_replay_id    BIGINT REFERENCES battle_replays(id),
  troops_id           BIGINT REFERENCES troops(id)
);

CREATE TABLE IF NOT EXISTS village_def_mappings (
  id           BIGSERIAL PRIMARY KEY,
  defenses_id  BIGINT REFERENCES defenses(id),
  amount       INTEGER DEFAULT 0,
  village_id   BIGINT REFERENCES villages(id)
);

CREATE TABLE IF NOT EXISTS troop_village_mappings (
  id          BIGSERIAL PRIMARY KEY,
  village_id  BIGINT REFERENCES villages(id),
  troops_id   BIGINT REFERENCES troops(id),
  quantity    INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS buildings (
  id                  BIGSERIAL PRIMARY KEY,
  x                   INTEGER DEFAULT 0,
  y                   INTEGER DEFAULT 0,
  building_type       TEXT,
  level               INTEGER DEFAULT 1,
  village_id          BIGINT REFERENCES villages(id),
  resource_collected  TIMESTAMP
);

CREATE TABLE IF NOT EXISTS level_specifics (
  level           INTEGER PRIMARY KEY,
  min_vill_level  INTEGER DEFAULT 1,
  upg_cost        INTEGER DEFAULT 0,
  h               INTEGER DEFAULT 1,
  w               INTEGER DEFAULT 1
);

CREATE TABLE IF NOT EXISTS resource_colls (
  level           INTEGER,
  type            TEXT,
  resource_yield  INTEGER DEFAULT 0,
  time_refill     INTEGER DEFAULT 0,
  PRIMARY KEY (level, type)
);

CREATE TABLE IF NOT EXISTS storages (
  level  INTEGER,
  type   TEXT,
  store  INTEGER DEFAULT 0,
  PRIMARY KEY (level, type)
);

INSERT INTO users (id, username, name, hashed_password, fights_won) VALUES
(1, 'dragon_lord', 'Aiden Cross',  '$2a$10$a7dJf6zWuFT3zwRJdFYusumESQlqFd5g3eoksCiFk9GGy0SAojL4u', 0),
(2, 'iron_witch',  'Priya Sharma', '$2a$10$a7dJf6zWuFT3zwRJdFYusumESQlqFd5g3eoksCiFk9GGy0SAojL4u', 0),
(3, 'ghost_rider', 'Marcus Webb',  '$2a$10$a7dJf6zWuFT3zwRJdFYusumESQlqFd5g3eoksCiFk9GGy0SAojL4u', 0);

INSERT INTO villages (id, user_id, village_level, gold, oil, money, farm_land, mines, level_constraints) VALUES
(1, 1, 1, 200, 100, 2000, 5, 0, '{"MaxBuildings":10}'),
(2, 2, 1, 150,  80, 1500, 5, 0, '{"MaxBuildings":10}'),
(3, 3, 1, 100,  50, 1000, 5, 0, '{"MaxBuildings":10}');

INSERT INTO defenses (id, type, defensive_power, capabilities, attack_power, cost) VALUES
(1, 'cannon',       120, '[]', 80, 500),
(2, 'archer_tower',  95, '[]', 60, 300),
(3, 'wall',          30, '[]',  0,  50);

INSERT INTO troops (id, type, health, offensive_power, capabilities, level, cost) VALUES
(1,  'recruit', 120, 40,  '[{"CapType":"jab","RelativeDamage":1.0},{"CapType":"guard","RelativeDamage":0.5}]', 1,  50),
(2,  'recruit', 160, 55,  '[{"CapType":"jab","RelativeDamage":1.1},{"CapType":"guard","RelativeDamage":0.6}]', 2,  80),
(3,  'recruit', 200, 70,  '[{"CapType":"jab","RelativeDamage":1.2},{"CapType":"guard","RelativeDamage":0.7}]', 3, 120),
(4,  'recruit', 250, 90,  '[{"CapType":"jab","RelativeDamage":1.3},{"CapType":"guard","RelativeDamage":0.8}]', 4, 180),
(5,  'goblin',  100, 100, '[{"CapType":"stab","RelativeDamage":1.0},{"CapType":"loot_raid","RelativeDamage":0.8}]', 1,  80),
(6,  'goblin',  130, 130, '[{"CapType":"stab","RelativeDamage":1.1},{"CapType":"loot_raid","RelativeDamage":0.9}]', 2, 120),
(7,  'goblin',  160, 160, '[{"CapType":"stab","RelativeDamage":1.2},{"CapType":"loot_raid","RelativeDamage":1.0}]', 3, 180),
(8,  'goblin',  200, 200, '[{"CapType":"stab","RelativeDamage":1.3},{"CapType":"loot_raid","RelativeDamage":1.1}]', 4, 260),
(9,  'archer',  140, 110, '[{"CapType":"arrow_shot","RelativeDamage":1.0},{"CapType":"rapid_fire","RelativeDamage":1.2}]', 1, 130),
(10, 'archer',  180, 145, '[{"CapType":"arrow_shot","RelativeDamage":1.1},{"CapType":"rapid_fire","RelativeDamage":1.3}]', 2, 190),
(11, 'archer',  220, 180, '[{"CapType":"arrow_shot","RelativeDamage":1.2},{"CapType":"rapid_fire","RelativeDamage":1.4}]', 3, 270),
(12, 'archer',  270, 220, '[{"CapType":"arrow_shot","RelativeDamage":1.3},{"CapType":"rapid_fire","RelativeDamage":1.5}]', 4, 370),
(13, 'barbarian', 250, 70,  '[{"CapType":"sword_slash","RelativeDamage":1.0},{"CapType":"charge","RelativeDamage":1.3}]', 1, 100),
(14, 'barbarian', 330, 90,  '[{"CapType":"sword_slash","RelativeDamage":1.1},{"CapType":"charge","RelativeDamage":1.4}]', 2, 150),
(15, 'barbarian', 420, 115, '[{"CapType":"sword_slash","RelativeDamage":1.2},{"CapType":"charge","RelativeDamage":1.5}]', 3, 220),
(16, 'barbarian', 520, 145, '[{"CapType":"sword_slash","RelativeDamage":1.3},{"CapType":"charge","RelativeDamage":1.6}]', 4, 310);

INSERT INTO troop_village_mappings (id, village_id, troops_id, quantity) VALUES
(1, 1, 1, 4),
(2, 2, 5, 4),
(3, 3, 9, 4);

INSERT INTO village_def_mappings (id, defenses_id, amount, village_id) VALUES
(1, 1, 2, 1),
(2, 3, 5, 1),
(3, 2, 1, 2),
(4, 3, 3, 2),
(5, 1, 1, 3),
(6, 3, 2, 3);

INSERT INTO buildings (id, x, y, building_type, level, village_id, resource_collected) VALUES
(1, 1, 1, 'Barrack', 1, 1, '2020-01-01 00:00:00'),
(2, 4, 1, 'Gold',    1, 1, '2020-01-01 00:00:00'),
(3, 1, 1, 'Barrack', 1, 2, '2020-01-01 00:00:00'),
(4, 4, 1, 'Gold',    1, 2, '2020-01-01 00:00:00'),
(5, 1, 1, 'Barrack', 1, 3, '2020-01-01 00:00:00'),
(6, 4, 1, 'Gold',    1, 3, '2020-01-01 00:00:00');

INSERT INTO level_specifics (level, min_vill_level, upg_cost, h, w) VALUES
(1, 1,  500, 2, 2),
(2, 2, 1000, 2, 2),
(3, 3, 2000, 3, 3),
(4, 4, 4000, 3, 3),
(5, 5, 8000, 4, 4);

INSERT INTO resource_colls (level, type, resource_yield, time_refill) VALUES
(1, 'Gold',  100, 60),
(2, 'Gold',  250, 60),
(3, 'Gold',  500, 60),
(4, 'Gold', 1000, 60),
(5, 'Gold', 2000, 60),
(1, 'Oil',    50, 90),
(2, 'Oil',   120, 90),
(3, 'Oil',   250, 90),
(4, 'Oil',   500, 90),
(5, 'Oil',  1000, 90),
(1, 'Farm',  200, 45),
(2, 'Farm',  450, 45),
(3, 'Farm',  900, 45),
(4, 'Farm', 1800, 45),
(5, 'Farm', 3500, 45);

INSERT INTO storages (level, type, store) VALUES
(1, 'Gold',  10000),
(2, 'Gold',  25000),
(3, 'Gold',  50000),
(4, 'Gold', 100000),
(5, 'Gold', 200000),
(1, 'Oil',    5000),
(2, 'Oil',   12000),
(3, 'Oil',   25000),
(4, 'Oil',   50000),
(5, 'Oil',  100000),
(1, 'Farm',  20000),
(2, 'Farm',  45000),
(3, 'Farm',  90000),
(4, 'Farm', 180000),
(5, 'Farm', 350000);

SELECT setval('users_id_seq',                  (SELECT MAX(id) FROM users));
SELECT setval('villages_id_seq',               (SELECT MAX(id) FROM villages));
SELECT setval('defenses_id_seq',               (SELECT MAX(id) FROM defenses));
SELECT setval('village_def_mappings_id_seq',   (SELECT MAX(id) FROM village_def_mappings));
SELECT setval('troops_id_seq',                 (SELECT MAX(id) FROM troops));
SELECT setval('troop_village_mappings_id_seq', (SELECT MAX(id) FROM troop_village_mappings));
SELECT setval('buildings_id_seq',              (SELECT MAX(id) FROM buildings));
