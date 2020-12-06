create table maude_report
(
    id                      serial  not null
        constraint maude_report_pkey
            primary key,
    event_location          text,
    report_to_fda           boolean,
    event_type              text,
    report_number           text,
    type_of_report          text,
    product_problem_flag    text,
    date_received           date,
    date_of_event           date,
    report_date             date,
    date_facility_aware     date,
    number_devices_in_event integer,
    manufacturer_name       text,
    device_id_fk            integer not null
        constraint device_id_fk
            references device,
    mdr_text_id_fk          integer not null
        constraint mdr_text_id_fk
            references mdr_text
);

create unique index maude_report_device_id_fk_uindex
    on maude_report (device_id_fk);

create unique index maude_report_mdr_text_id_fk_uindex
    on maude_report (mdr_text_id_fk);