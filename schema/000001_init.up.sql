create table order_main
(
    order_id      bigint            not null
        primary key,
    buyer_address varchar(255),
    buyer_email   varchar(255),
    buyer_name    varchar(255),
    buyer_phone   varchar(255),
    create_time   timestamp,
    order_amount  numeric(19, 2)    not null,
    order_status  integer default 0 not null,
    update_time   timestamp
);

alter table order_main
    owner to postgres;

create table product_category
(
    category_id   integer not null
        primary key,
    category_name varchar(255),
    category_type integer
        constraint uk_6kq6iveuim6wd90cxo5bksumw
            unique,
    create_time   timestamp,
    update_time   timestamp
);

alter table product_category
    owner to postgres;

create table product_info
(
    product_id          varchar(255)   not null
        primary key,
    category_type       integer default 0,
    create_time         timestamp,
    product_description varchar(255),
    product_icon        varchar(255),
    product_name        varchar(255)   not null,
    product_price       numeric(19, 2) not null,
    product_status      integer default 0,
    product_stock       integer        not null
        constraint product_info_product_stock_check
            check (product_stock >= 0),
    update_time         timestamp
);

alter table product_info
    owner to postgres;

create table users
(
    id       bigint  not null
        primary key,
    active   boolean not null,
    address  varchar(255),
    email    varchar(255)
        constraint uk_sx468g52bpetvlad2j9y0lptc
            unique,
    name     varchar(255),
    password varchar(255),
    phone    varchar(255),
    role     varchar(255)
);

alter table users
    owner to postgres;

create table cart
(
    user_id bigint not null
        primary key
        constraint fkg5uhi8vpsuy0lgloxk2h4w5o6
            references users
);

alter table cart
    owner to postgres;

create table product_in_order
(
    id                  bigint         not null
        primary key,
    category_type       integer        not null,
    count               integer
        constraint product_in_order_count_check
            check (count >= 1),
    product_description varchar(255)   not null,
    product_icon        varchar(255),
    product_id          varchar(255),
    product_name        varchar(255),
    product_price       numeric(19, 2) not null,
    product_stock       integer
        constraint product_in_order_product_stock_check
            check (product_stock >= 0),
    cart_user_id        bigint
        constraint fkhnivo3fl2qtco3ulm4mq0mbr5
            references cart,
    order_id            bigint
        constraint fkt0sfj3ffasrift1c4lv3ra85e
            references order_main
);

alter table product_in_order
    owner to postgres;

