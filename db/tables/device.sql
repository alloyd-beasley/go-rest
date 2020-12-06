create table device
(
    id                         integer not null
        constraint device_pkey
            primary key,
    manufacturer_d_address_1   text,
    manufacturer_d_state       text,
    manufacturer_d_postal_code text,
    manufacturer_d_city        text,
    manufacturer_d_country     text,
    manufacturer_d_name        text,
    lot_number                 text,
    model_number               text,
    generic_name               text,
    brand_name                 text
);