

insert into bildings (id_bild, build_name, cost, milit_boost, med_boost, manag_boost, people_boost) values 
    (1, 'main bild', 0,10,10,10,10),
    (2, 'main milit', 10,10,10,10,10),
    (3, 'main med', 10,10,10,10,10),
    (4, 'main manag', 10,10,10,10,10)
    on conflict (id_bild) do update set
    id_bild = excluded.id_bild,
    build_name = excluded.build_name,
    cost = excluded.cost,
    milit_boost = excluded.milit_boost,
    med_boost = excluded.med_boost,
    manag_boost = excluded.manag_boost,
    people_boost = excluded.people_boost;

insert into events (id_event, ev_name, ev_comment, ev_money, ev_people, ev_milit, ev_med, ev_manag, ev_bild) values
    (1, 'бандиты', 'razboi', -20, -5, -3,-1,-1, 0),
    (2, 'хворь', 'хворь', -20, -5, -1,-3,-1, 0)
    on conflict (id_event) do update set
    id_event = excluded.id_event,
    ev_name = excluded.ev_name,
    ev_comment = excluded.ev_comment,
    ev_money = excluded.ev_money,
    ev_people = excluded.ev_people,
    ev_milit = excluded.ev_milit,
    ev_med = excluded.ev_med,
    ev_manag = excluded.ev_manag,
    ev_bild = excluded.ev_bild;    

 with w_add_user as (
     insert into users (user_name, score,money_sum,money_per_turn,people_all,people_per_turn,people_free,people_milit,people_med,people_manag,milit_sum_boost,med_sum_boost,manag_sum_boost) 
     values ('testuser', 10, 100, 0, 10,0,10,0,0,0,0,0,0) 
     returning id_user
 )

 insert into maps (user_id, bilding_id, xCoord, yCoord) values
    ((select id_user from w_add_user), 1, 50, 50);




