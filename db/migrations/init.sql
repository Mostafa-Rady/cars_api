create table if not exists car_types
(
    id   int     not null primary key,
    name varchar not null unique
);

-- adding initial car types
INSERT INTO car_types(id, name)
VALUES (1, 'Sedan'),
       (2, 'Van'),
       (3, 'Suv'),
       (4, 'Motor-bike');


create table if not exists car_colors
(
    id   int     not null primary key,
    name varchar not null unique
);
-- adding initial car colors
INSERT INTO car_colors(id, name)
VALUES (1, 'Red'),
       (2, 'Green'),
       (3, 'Blue');


create table if not exists car_features
(
    id   int     not null primary key,
    name varchar not null unique
);

-- adding initial car features
INSERT INTO car_features(id, name)
VALUES (1, 'Sunroof'),
       (2, 'Panorama'),
       (3, 'Auto-parking'),
       (4, 'Surround-system');

CREATE DOMAIN speed_range AS INT CHECK (VALUE BETWEEN 0 AND 240);
create table if not exists cars
(
    id             serial primary key,
    created_at     timestamp with time zone,
    updated_at     timestamp with time zone,
    deleted_at     timestamp with time zone,
    name           varchar     not null,
    type           int         not null references car_types (id),
    color          int         not null references car_colors (id),
    speed_range_km speed_range not null
);

create index if not exists idx_cars_deleted_at
    on cars (deleted_at);

create index if not exists idx_cars_name
    on cars (name);

create index if not exists idx_cars_type
    on cars (type);

create index if not exists idx_cars_color
    on cars (color);

create table if not exists cars_features
(
    car_id     int not null references cars (id),
    feature_id int not null references car_features (id),
    primary key (car_id, feature_id)
);


-- view to select car with details
create or replace view car_previews
            (id, created_at, updated_at, deleted_at, name, type_id, type, color_id, color, speed_range_km,
             car_features_ids, car_features)
as
SELECT c.id,
       c.created_at,
       c.updated_at,
       c.deleted_at,
       c.name,
       c.type                       AS type_id,
       ct.name                      AS type,
       c.color                      AS color_id,
       cc.name                      AS color,
       c.speed_range_km,
       ARRAY(SELECT f.id
             FROM car_features f
                      JOIN cars_features j ON f.id = j.feature_id
                      JOIN cars cc_1 ON cc_1.id = j.car_id
             WHERE j.car_id = c.id) AS car_features_ids,
       ARRAY(SELECT f.name
             FROM car_features f
                      JOIN cars_features j ON f.id = j.feature_id
                      JOIN cars cc_1 ON cc_1.id = j.car_id
             WHERE j.car_id = c.id) AS car_features
FROM cars c
         JOIN car_colors cc ON c.color = cc.id
         JOIN car_types ct ON ct.id = c.type;





