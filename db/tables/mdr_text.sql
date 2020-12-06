create table mdr_text
(
    id             integer not null
        constraint mdr_text_pkey
            primary key,
    text_type_code text,
    text           text
);