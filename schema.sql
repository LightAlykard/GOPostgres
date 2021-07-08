-- база данных браузерной игры, на развитие города, город сам развивается,
-- пользователь только только распределяет жителей по специальности, жители сами строят здания,
-- в зависимости от соотношения специальностей в городе они тратят ресурсы
-- каждый ход в зависимости от развития и сооношения типов население город лучше или хуже справляется с определенными событиями
-- 3 таблицы, первая это типы строений в игре и их характеристики
-- вторая пользователи с их прогрессом, то есть их ресурсы, население, уровень развития, прирост средсв за ход,
-- % распределения средст населения на направление постройки
-- третья с координатами строений ползователей на их личных картах

create table bildings (
id bigint generated always as identity,
name text not null,
cost int not null,
militaryboost int not null,
medicineboost int not null,
citymanegmentboost int not null,
);

create table users (
id bigint generated always as identity,
name text not null,
progress integer not null,
money int not null,
moneyperturn int not null,
people int not null,
freepeaple not null,
summilitarybild int not null,
summedicineboost int not null,
sumcityboost int not null,
militPeople int not null,
medPeople int not null,
manegmPeople int not null,

genre integer not null,
);

create table maps (
id bigint generated always as identity,
userId int not null,
bildingId int not null,
xCoord int not null,
yCoord int not null,
);

create table songs (
    id bigint generated always as identity,
    author text,
    name text,
    genre integer not null,
    license boolean,
)