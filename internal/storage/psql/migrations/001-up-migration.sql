CREATE TABLE IF NOT EXISTS categories
(
    category_id INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(50)  NOT NULL UNIQUE,
    description VARCHAR(250) NULL
);

CREATE TABLE IF NOT EXISTS recipes
(
    recipe_id   INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(150) NOT NULL UNIQUE,
    description VARCHAR(250) NOT NULL,
    notes       TEXT         NOT NULL
);

CREATE TABLE IF NOT EXISTS company_addresses
(
    company_address_id INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    region             VARCHAR(200) NOT NULL,
    city               VARCHAR(200) NOT NULL,
    street             VARCHAR(200) NOT NULL,
    house_number       VARCHAR(10)  NULL,
    building_number    VARCHAR(5)   NULL
);

CREATE TABLE IF NOT EXISTS suppliers
(
    supplier_id        INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name               VARCHAR(100) NOT NULL UNIQUE,
    company_address_id INTEGER      NULL REFERENCES company_addresses (company_address_id) ON DELETE SET NULL,
    phone_number       VARCHAR(17)  NOT NULL
);

CREATE TABLE IF NOT EXISTS manufacturers
(
    manufacturer_id    INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name               VARCHAR(150) NOT NULL UNIQUE,
    company_address_id INTEGER      NULL REFERENCES company_addresses (company_address_id) ON DELETE SET NULL,
    supplier_id        INTEGER      NULL REFERENCES suppliers (supplier_id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS positions
(
    position_id INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(150) NOT NULL UNIQUE,
    description VARCHAR(250) NULL
);

CREATE TABLE IF NOT EXISTS employees
(
    employee_id   INTEGER     NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name    VARCHAR(50) NOT NULL,
    last_name     VARCHAR(50) NOT NULL,
    patronymic    VARCHAR(50) NULL,
    phone_number  VARCHAR(17) NOT NULL,
    email         VARCHAR(50) NULL UNIQUE,
    password_hash VARCHAR(64) NULL UNIQUE,
    position_id   INTEGER     NOT NULL REFERENCES positions (position_id) ON DELETE CASCADE,
    company_id    INTEGER     NOT NULL REFERENCES manufacturers (manufacturer_id) ON DELETE CASCADE,
    admin         BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS products
(
    product_id         INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name               VARCHAR(50)  NOT NULL,
    image_path         VARCHAR(50)  NOT NULL,
    description        VARCHAR(250) NULL,
    price              NUMERIC      NOT NULL CHECK (price > 0),
    manufacturing_date DATE         NOT NULL DEFAULT CURRENT_DATE CHECK (manufacturing_date <= expiration_date),
    expiration_date    DATE         NOT NULL,
    category_id        INTEGER      NULL REFERENCES categories (category_id) ON DELETE SET NULL,
    recipe_id          INTEGER      NULL REFERENCES recipes (recipe_id) ON DELETE SET NULL,
    manufacturer_id    INTEGER      NOT NULL REFERENCES manufacturers (manufacturer_id) ON DELETE CASCADE,
    unit_stock         INTEGER      NOT NULL DEFAULT 0,
    created_at         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP    NULL
);

CREATE TABLE IF NOT EXISTS weight_units
(
    weight_unit_id INTEGER    NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name           VARCHAR(5) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS ingredients
(
    ingredient_id      INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name               VARCHAR(100) NOT NULL,
    description        VARCHAR(250) NULL,
    remaining_quantity INTEGER      NOT NULL,
    weight_unit_id     INTEGER      NOT NULL REFERENCES weight_units (weight_unit_id) ON DELETE RESTRICT,
    supplier_id        INTEGER      NOT NULL REFERENCES suppliers (supplier_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS recipe_ingredients
(
    recipe_ingredient_id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    recipe_id            INTEGER NOT NULL REFERENCES recipes (recipe_id) ON DELETE CASCADE,
    ingredient_id        INTEGER NOT NULL REFERENCES ingredients (ingredient_id) ON DELETE CASCADE,
    quantity             NUMERIC NOT NULL CHECK ( quantity > 0 ),
    weight_unit_id       INTEGER NOT NULL REFERENCES weight_units (weight_unit_id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS payment_methods
(
    payment_method_id INTEGER     NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name              VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS receipts
(
    receipt_id        INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    payment_details   TEXT    NOT NULL,
    payment_method_id INTEGER NOT NULL REFERENCES payment_methods (payment_method_id) ON DELETE CASCADE,
    amount            NUMERIC NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS order_statuses
(
    order_status_id INTEGER     NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name            VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS delivery_methods
(
    delivery_method_id INTEGER     NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name               VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS delivery_addresses
(
    delivery_address_id INTEGER      NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    region              VARCHAR(200) NOT NULL,
    city                VARCHAR(200) NOT NULL,
    street              VARCHAR(200) NOT NULL,
    house_number        VARCHAR(10)  NOT NULL,
    building_number     VARCHAR(5)   NULL,
    apartment_number    VARCHAR(5)   NULL
);

CREATE TABLE IF NOT EXISTS wish_lists
(
    wish_list_id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
);

CREATE TABLE IF NOT EXISTS wish_list_item
(
    wish_list_item_id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    wish_list_id      INTEGER NOT NULL REFERENCES wish_lists (wish_list_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS carts
(
    cart_id        INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    products_count INTEGER NOT NULL DEFAULT 0,
    total_price    NUMERIC NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS cart_items
(
    cart_item_id INTEGER NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    product_id   INTEGER NOT NULL REFERENCES products (product_id) ON DELETE CASCADE,
    quantity     INTEGER NOT NULL DEFAULT 1,
    price        NUMERIC NOT NULL CHECK (price > 0),
    cart_id      INTEGER NOT NULL REFERENCES carts (cart_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS customers
(
    customer_id         INTEGER     NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name          VARCHAR(50) NOT NULL,
    last_name           VARCHAR(50) NOT NULL,
    patronymic          VARCHAR(50) NULL,
    image_path          VARCHAR(50) NULL,
    phone_number        VARCHAR(17) NOT NULL,
    email               VARCHAR(50) NULL UNIQUE,
    password_hash       VARCHAR(64) NULL UNIQUE,
    delivery_address_id INTEGER     NULL REFERENCES delivery_addresses (delivery_address_id) ON DELETE SET NULL,
    cart_id             INTEGER     NOT NULL REFERENCES carts (cart_id) ON DELETE CASCADE,
    wish_list_id        INTEGER     NOT NULL REFERENCES wish_lists (wish_list_id) ON DELETE RESTRICT,
    created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS orders
(
    order_id            INTEGER   NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    customer_id         INTEGER   NOT NULL REFERENCES customers (customer_id) ON DELETE CASCADE,
    cart_id             INTEGER   NOT NULL REFERENCES carts (cart_id) ON DELETE CASCADE,
    receipt_id          INTEGER   NOT NULL REFERENCES receipts (receipt_id) ON DELETE CASCADE,
    order_status_id     INTEGER   NOT NULL REFERENCES order_statuses (order_status_id) ON DELETE CASCADE,
    delivery_address_id INTEGER   NOT NULL REFERENCES delivery_addresses (delivery_address_id) ON DELETE CASCADE,
    delivery_method_id  INTEGER   NOT NULL REFERENCES delivery_methods (delivery_method_id) ON DELETE CASCADE,
    order_date          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO payment_methods(name)
VALUES ('Оплата онлайн'),
       ('Оплата наличными после доставки'),
       ('Оплата картой после доставки'),
       ('Смешанная оплата после доставки(наличные и карта)');
INSERT INTO order_statuses(name)
VALUES ('В обработке'),
       ('Доставляется'),
       ('Доставлен');
INSERT INTO delivery_methods(name)
VALUES ('Самомывоз'),
       ('Курьер');
INSERT INTO weight_units(name)
VALUES ('кг'),
       ('г'),
       ('л'),
       ('мл');