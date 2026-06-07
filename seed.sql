TRUNCATE TABLE replay_troops, replay_defenses, battle_replays, buildings_village_mappings, buildings, troop_village_mappings, village_def_mappings, troops, defenses, villages, users RESTART IDENTITY CASCADE;

INSERT INTO users (id, username, name, hashed_password, fights_won) VALUES
(1, 'dragon_lord', 'Aiden Cross',  '$2a$10$X5rXjnZNSSDRxUccwRNfV.SmAh2RkinHhGec35bN4pjZp9ymVFF/6', 42),
(2, 'iron_witch',  'Priya Sharma', '$2a$10$EehrDTRrzmVXQmwgAaXbTOzlzBmdcaCgnNQWRJ4e6arNa6qqEIysC', 17),
(3, 'ghost_rider', 'Marcus Webb',  '$2a$10$hteS7N6MntegtKekvO.Vr.jLj4PJEfdEmrWvyXtg0yswSBKItYmbu',  5);

INSERT INTO villages (id, user_id, village_level, gold, oil, money, farm_land, mines, level_constraints) VALUES
(1, 1, 8, 95000, 4200, 12000, 15, 6, '{"max_buildings":30,"max_troops":200}'),
(2, 2, 5, 40000, 1800,  6500, 10, 3, '{"max_buildings":20,"max_troops":150}'),
(3, 3, 2,  8000,  300,   900,  5, 1, '{"max_buildings":10,"max_troops":50}');

INSERT INTO defenses (id, type, defensive_power, capabilities, attack_power, cost) VALUES
(1, 'cannon',       120, '{"range":5,"splash":false}',   80, 500),
(2, 'archer_tower',  95, '{"range":8,"splash":false}',   60, 300),
(3, 'mortar',       200, '{"range":6,"splash":true}',   100, 800),
(4, 'air_defense',  150, '{"range":7,"anti_air":true}', 120, 700),
(5, 'wall',          30, '{"blocks_ground":true}',         0,  50);

INSERT INTO village_def_mappings (id, defenses_id, amount, village_id) VALUES
(1,  1,  4, 1),
(2,  2,  6, 1),
(3,  3,  2, 1),
(4,  4,  3, 1),
(5,  5, 40, 1),
(6,  1,  2, 2),
(7,  2,  3, 2),
(8,  3,  1, 2),
(9,  5, 20, 2),
(10, 1,  1, 3),
(11, 2,  2, 3),
(12, 5, 10, 3);

INSERT INTO troops (id, type, health, offensive_power, capabilities, level, cost) VALUES
(1, 'barbarian',  300,  80, '["sword_slash","charge"]',    5, 100),
(2, 'archer',     180, 120, '["arrow_shot","rapid_fire"]', 4, 150),
(3, 'giant',      900,  60, '["stomp","shield_bash"]',     4, 500),
(4, 'dragon',    1200, 200, '["fire_breath","dive_bomb"]', 3, 2000),
(5, 'goblin',     100, 100, '["loot_raid","stab"]',        2,  80);

INSERT INTO troop_village_mappings (id, village_id, troops_id, quantity) VALUES
(1, 1, 1, 50),
(2, 1, 2, 40),
(3, 1, 3, 10),
(4, 1, 4,  5),
(5, 2, 1, 30),
(6, 2, 2, 20),
(7, 2, 5, 25),
(8, 3, 1, 15),
(9, 3, 2, 10);

INSERT INTO buildings (id, width, height, type, cost) VALUES
(1, 3, 3, 'town_hall',  5000),
(2, 2, 2, 'barracks',   1000),
(3, 2, 2, 'mine',        800),
(4, 2, 2, 'farm',        600),
(5, 1, 1, 'wall',        100),
(6, 2, 2, 'cannon_base', 900);

INSERT INTO buildings_village_mappings (id, village_id, x, y, buildings_id) VALUES
(1,  1,  2,  2, 1),
(2,  1,  6,  2, 2),
(3,  1, 10,  2, 2),
(4,  1,  2,  6, 3),
(5,  1,  6,  6, 4),
(6,  1, 14,  3, 6),
(7,  2,  3,  3, 2),
(8,  2,  7,  3, 2),
(9,  2,  3,  8, 3),
(10, 2, 10,  7, 4),
(11, 3,  4,  4, 2),
(12, 3,  8,  4, 2);

INSERT INTO battle_replays (id, attacker_id, defender_id, attacker_loot, defender_loot, winner) VALUES
(1, 1, 2, 5000, -5000, true);

INSERT INTO replay_defenses (id, attacks_used, battle_replay_id, defenses_id, num_deployed) VALUES
(1, 3, 1, 1,  2),
(2, 5, 1, 2,  3),
(3, 2, 1, 3,  1),
(4, 0, 1, 5, 20);

INSERT INTO replay_troops (id, amount_deployed, attacks_allocated, battle_replay_id, troops_id) VALUES
(1, 30, 60, 1, 1),
(2, 20, 40, 1, 2),
(3,  5, 10, 1, 3),
(4,  2,  4, 1, 4);

SELECT setval('users_id_seq',                    (SELECT MAX(id) FROM users));
SELECT setval('villages_id_seq',                 (SELECT MAX(id) FROM villages));
SELECT setval('defenses_id_seq',                 (SELECT MAX(id) FROM defenses));
SELECT setval('village_def_mappings_id_seq',     (SELECT MAX(id) FROM village_def_mappings));
SELECT setval('troops_id_seq',                   (SELECT MAX(id) FROM troops));
SELECT setval('troop_village_mappings_id_seq',   (SELECT MAX(id) FROM troop_village_mappings));
SELECT setval('buildings_id_seq',                (SELECT MAX(id) FROM buildings));
SELECT setval('buildings_village_mappings_id_seq',(SELECT MAX(id) FROM buildings_village_mappings));
SELECT setval('battle_replays_id_seq',           (SELECT MAX(id) FROM battle_replays));
SELECT setval('replay_defenses_id_seq',          (SELECT MAX(id) FROM replay_defenses));
SELECT setval('replay_troops_id_seq',            (SELECT MAX(id) FROM replay_troops));
