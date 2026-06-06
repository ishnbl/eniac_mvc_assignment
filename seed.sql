TRUNCATE TABLE replay_troops, replay_defenses, battle_replays, buildings, troops, defenses, villages, users RESTART IDENTITY CASCADE;

INSERT INTO users (id, username, name, hashed_password, fights_won) VALUES
(1, 'dragon_lord', 'Aiden Cross',  '$2a$10$X5rXjnZNSSDRxUccwRNfV.SmAh2RkinHhGec35bN4pjZp9ymVFF/6', 42),
(2, 'iron_witch',  'Priya Sharma', '$2a$10$EehrDTRrzmVXQmwgAaXbTOzlzBmdcaCgnNQWRJ4e6arNa6qqEIysC', 17),
(3, 'ghost_rider', 'Marcus Webb',  '$2a$10$hteS7N6MntegtKekvO.Vr.jLj4PJEfdEmrWvyXtg0yswSBKItYmbu',  5);

INSERT INTO villages (id, user_id, village_level, gold, oil, money, farm_land, mines, map) VALUES
(1, 1, 8, 95000, 4200, 12000, 15, 6, '{"width":20,"height":20,"terrain":"mixed"}'),
(2, 2, 5, 40000, 1800,  6500, 10, 3, '{"width":20,"height":20,"terrain":"forest"}'),
(3, 3, 2,  8000,  300,   900,  5, 1, '{"width":20,"height":20,"terrain":"plains"}');

INSERT INTO defenses (id, village_id, type, defensive_power, amount) VALUES
(1,  1, 'cannon',       120, 4),
(2,  1, 'archer_tower',  95, 6),
(3,  1, 'mortar',       200, 2),
(4,  1, 'air_defense',  150, 3),
(5,  1, 'wall',         300, 40),
(6,  2, 'cannon',        80, 2),
(7,  2, 'archer_tower',  65, 3),
(8,  2, 'mortar',       140, 1),
(9,  2, 'wall',         180, 20),
(10, 3, 'cannon',        40, 1),
(11, 3, 'archer_tower',  30, 2),
(12, 3, 'wall',          60, 10);

INSERT INTO troops (id, village_id, type, health, offensive_power, level, quantity, attacks) VALUES
(1, 1, 'barbarian', 300,  80, 5, 50, '["sword_slash","charge"]'),
(2, 1, 'archer',    180, 120, 4, 40, '["arrow_shot","rapid_fire"]'),
(3, 1, 'giant',     900,  60, 4, 10, '["stomp","shield_bash"]'),
(4, 1, 'dragon',   1200, 200, 3,  5, '["fire_breath","dive_bomb"]'),
(5, 2, 'barbarian', 200,  60, 3, 30, '["sword_slash"]'),
(6, 2, 'archer',    120,  90, 3, 20, '["arrow_shot"]'),
(7, 2, 'goblin',    100, 100, 2, 25, '["loot_raid","stab"]'),
(8, 3, 'barbarian', 100,  40, 1, 15, '["sword_slash"]'),
(9, 3, 'archer',     80,  50, 1, 10, '["arrow_shot"]');

INSERT INTO buildings (id, village_id, building_level, x, y, width, height) VALUES
(1,  1, 10,  2,  2, 2, 2),
(2,  1,  8,  5,  2, 2, 2),
(3,  1,  8,  8,  2, 2, 2),
(4,  1,  6,  2,  6, 3, 3),
(5,  1,  5,  8,  8, 2, 2),
(6,  1,  7, 14,  3, 2, 2),
(7,  2,  6,  3,  3, 2, 2),
(8,  2,  5,  7,  3, 2, 2),
(9,  2,  4,  3,  8, 2, 2),
(10, 2,  3, 10,  7, 2, 2),
(11, 3,  2,  4,  4, 2, 2),
(12, 3,  1,  8,  4, 2, 2);

INSERT INTO battle_replays (id, winner_village_id, loser_village_id) VALUES
(1, 1, 2);

INSERT INTO replay_defenses (id, battle_replay_id, village_id, type, defensive_power, amount_deployed, attacks_used) VALUES
(1, 1, 2, 'cannon',       80,  2, 3),
(2, 1, 2, 'archer_tower', 65,  3, 5),
(3, 1, 2, 'mortar',      140,  1, 2),
(4, 1, 2, 'wall',        180, 20, 0);

INSERT INTO replay_troops (id, battle_replay_id, village_id, type, health, offensive_power, level, amount_deployed, attacks_allocated) VALUES
(1, 1, 1, 'barbarian', 300,  80, 5, 30, 60),
(2, 1, 1, 'archer',    180, 120, 4, 20, 40),
(3, 1, 1, 'giant',     900,  60, 4,  5, 10),
(4, 1, 1, 'dragon',   1200, 200, 3,  2,  4);

SELECT setval('users_id_seq',          (SELECT MAX(id) FROM users));
SELECT setval('villages_id_seq',        (SELECT MAX(id) FROM villages));
SELECT setval('defenses_id_seq',        (SELECT MAX(id) FROM defenses));
SELECT setval('troops_id_seq',          (SELECT MAX(id) FROM troops));
SELECT setval('buildings_id_seq',       (SELECT MAX(id) FROM buildings));
SELECT setval('battle_replays_id_seq',  (SELECT MAX(id) FROM battle_replays));
SELECT setval('replay_defenses_id_seq', (SELECT MAX(id) FROM replay_defenses));
SELECT setval('replay_troops_id_seq',   (SELECT MAX(id) FROM replay_troops));
