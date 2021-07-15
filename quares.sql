-- список пользователей по рейтингу
select user_name, score, money_sum, people_all from users ORDER BY score DESC;

-- получить списки зданий пользователей по имени и с названием зданий и их координат
SELECT user_name, build_name, xCoord,yCoord FROM users, bildings, maps WHERE maps.user_id = users.id_user and maps.bilding_id=bildings.id_bild;

-- получить статистику пользователя по имени
select user_name, score, money_sum, money_per_turn, people_all, people_per_turn, people_free,people_milit,people_med,people_manag,milit_sum_boost,med_sum_boost, manag_sum_boost from users where user_name = 'testuser';

-- посмотреть какие ивенты случились с конкретным пользователем
SELECT user_name, ev_name, milit_lost,med_lost,manag_lost,money_lost,people_lost,idBild_lost FROM users, eventUsers, events WHERE eventUsers.user_id = users.id_user and eventUsers.event_id=events.id_event and user_name = 'testuser';

