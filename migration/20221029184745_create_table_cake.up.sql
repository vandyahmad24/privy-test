create table cake
(
    id          int auto_increment primary key,
    title       varchar(255),
    description varchar(255),
    rating      float,
    image       varchar(255),
    created_at  datetime(3)  null,
    updated_at  datetime(3)  null
);

