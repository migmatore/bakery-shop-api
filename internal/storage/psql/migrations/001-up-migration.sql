create table public.sellers
(
    id   bigint primary key,
    name varchar(50)
);

create table public.products
(
    id        bigint primary key,
    name      varchar(50),
    seller_id bigint references public.sellers (id) on delete cascade
);

create table public.customers
(
    id   bigint primary key,
    name varchar(50)
);

create table public.orders
(
    id          bigint primary key,
    product_id  bigint references public.products (id) on delete cascade,
    customer_id bigint references public.customers (id) on delete cascade
);

insert into public.sellers (id, name)
values (1, 'seller1'),
       (2, 'seller2'),
       (3, 'seller3');

insert into public.products (id, name, seller_id)
values (1, 'product1', 1),
       (2, 'product2', 1),
       (3, 'product3', 2),
       (4, 'product4', 3);

insert into public.customers (id, name)
values (1, 'customer1'),
       (2, 'customer2'),
       (3, 'customer3'),
       (4, 'customer4');

insert into public.orders (id, product_id, customer_id)
values (1, 1, 1),
       (2, 1, 2),
       (3, 2, 4),
       (4, 1, 3),
       (5, 3, 1),
       (6, 4, 2);