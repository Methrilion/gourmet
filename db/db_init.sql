SET client_encoding = 'UTF-8';
-- This controls whether ordinary string literals ('...') treat backslashes literally, as specified in the SQL standard.
-- Beginning in PostgreSQL 9.1, the default is on (prior releases defaulted to off).
-- Applications can check this parameter to determine how string literals will be processed.
-- The presence of this parameter can also be taken as an indication that the escape string syntax (E'...') is supported.
-- Escape string syntax should be used if an application desires backslashes to be treated as escape characters.
-- SET standart_conforming_string = on;
-- This parameter is normally on. When set to off, it disables validation of the function body string during CREATE FUNCTION.
-- Disabling validation avoids side effects of the validation process and avoids false positives due to problems such as forward references.
-- Set this parameter to off before loading functions on behalf of other users; pg_dump does so automatically.
SET check_function_bodies = false;
-- Controls which message levels are sent to the client. Valid values are DEBUG5, DEBUG4, DEBUG3, DEBUG2, DEBUG1, LOG, NOTICE, WARNING, ERROR, FATAL, and PANIC.
-- Each level includes all the levels that follow it. The later the level, the fewer messages are sent. The default is NOTICE.
-- Note that LOG has a different rank here than in log_min_messages.
SET client_min_messages = WARNING;

-- Full parameter information  https://www.postgresql.org/docs/9.4/static/manage-ag-tablespaces.html
SET default_tablespace = '';

-- This controls whether CREATE TABLE and CREATE TABLE AS include an OID column in newly-created tables,
-- if neither WITH OIDS nor WITHOUT OIDS is specified. It also determines whether OIDs will be included in tables created by SELECT INTO.
-- The parameter is off by default; in PostgreSQL 8.0 and earlier, it was on by default.
--
-- The use of OIDs in user tables is considered deprecated, so most installations should leave this variable disabled.
-- Applications that require OIDs for a particular table should specify WITH OIDS when creating the table.
-- This variable can be enabled for compatibility with old applications that do not follow this behavior.
SET default_with_oids = false;

-- this table acts as a lock, preventing another DB client from attempting
-- to convert a database that is in the process of initialization (must
-- be the first thing we create, and we delete it as the last action of
-- this script)
CREATE TABLE db_conversion_lock (c INTEGER);





CREATE TABLE currency (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    code VARCHAR(16) NOT NULL
);

CREATE TABLE rates_of_exchange (
    id SERIAL PRIMARY KEY,
    from_id INTEGER NOT NULL,
    to_id INTEGER NOT NULL,
    price DECIMAL NOT NULL,
    FOREIGN KEY (from_id) REFERENCES currency(id),
    FOREIGN KEY (to_id) REFERENCES currency(id)
);

CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    description TEXT NOT NULL,
    currency_id INTEGER NOT NULL,
    status BOOLEAN NOT NULL,
    FOREIGN KEY (currency_id) REFERENCES currency(id)
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE prices (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    location_id INTEGER NOT NULL,
    price DECIMAL NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (location_id) REFERENCES locations(id)
);

CREATE TABLE positions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(256) NOT NULL,
    last_name VARCHAR(256) NOT NULL,
    position_id INTEGER NOT NULL,
    location_id INTEGER NOT NULL,
    FOREIGN KEY (position_id) REFERENCES positions(id),
    FOREIGN KEY (location_id) REFERENCES locations(id)
);

CREATE TABLE methods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL
);

CREATE TABLE receipts (
    id SERIAL PRIMARY KEY,
    employee_id INTEGER NOT NULL,
    location_id INTEGER NOT NULL,
    method_id INTEGER NOT NULL,
    datetime TIMESTAMP NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    FOREIGN KEY (location_id) REFERENCES locations(id),
    FOREIGN KEY (method_id) REFERENCES methods(id)
);

CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,
    receipt_id INTEGER NOT NULL,
    price_id INTEGER NOT NULL,
    amount FLOAT NOT NULL,
    price DECIMAL NOT NULL,
    result DECIMAL NOT NULL,
    FOREIGN KEY (receipt_id) REFERENCES receipts(id)
);


INSERT INTO products(name, description) VALUES
    ('Яблоки', 'Без ГМО'),
    ('Апельсины', 'Оранжевые'),
    ('Бананы', 'Кушать аккуратно!'),
    ('Молоко Советское', 'То самое молоко'),
    ('Хлеб Заводский', 'Самый вкусный'),
    ('Тетрадь', 'Просто тетрадь');

INSERT INTO currency(name, code) VALUES
    ('Белорусский рубль', 'BYN'),
    ('Российский рубль', 'RUR'),
    ('Доллар США', 'USD');

INSERT INTO positions(name, description) VALUES
    ('Продавец', 'Хорошая позиция'),
    ('Менеджер', 'Менеджер...'),
    ('Директор', 'Главный');

INSERT INTO locations(name, description, currency_id, status) VALUES
    ('Точка1', 'Точка продажи номер 1', 1, true),
    ('Точка2', 'Точка продажи номер 2', 1, true),
    ('Точка3', 'EVIL POINT', 2, true);

INSERT INTO prices(product_id, location_id, price) VALUES
    (1, 1, 200),
    (2, 1, 250),
    (3, 1, 500),
    (4, 1, 100),
    (5, 1, 100),
    (1, 2, 210),
    (2, 2, 230),
    (5, 2, 100),
    (6, 2, 200);

INSERT INTO employees(first_name, last_name, position_id, location_id) VALUES
    ('Витько', 'сын Олегов', 1, 1),
    ('Мария', 'Печеньки', 1, 2),
    ('Микитко', 'сын Иванов', 3, 1);

INSERT INTO methods(name) VALUES
    ('Налом'),
    ('Карточкой'),
    ('Мобилой');

INSERT INTO receipts(employee_id, location_id, method_id, datetime) VALUES
    (1, 1, 1, '2016-06-22 19:10:25-07'),
    (2, 1, 1, '2016-06-23 19:12:25-07'),
    (1, 1, 1, '2016-07-24 11:11:25-07');

INSERT INTO purchases(receipt_id, price_id, amount, price, result) VALUES
    (1, 1, 1.1, 200, 220),
    (1, 2, 2, 250, 500),
    (1, 3, 1, 500, 500),
    (2, 1, 1, 200, 200),
    (3, 3, 1, 500, 500),
    (3, 3, 1, 500, 500),
    (3, 3, 1, 500, 500),
    (1, 3, 1, 500, 500),
    (3, 3, 2, 500, 1000);


-- remove lock that prevents concurrent db conversion; must be the last thing we do here
DROP TABLE db_conversion_lock;

SELECT * FROM products;
SELECT * FROM currency;
SELECT * FROM positions;
SELECT * FROM locations;
SELECT * FROM employees;
SELECT * FROM receipts;
SELECT * FROM purchases;

SELECT products.name, locations.name, price FROM prices, products, locations WHERE prices.product_id = products.id AND prices.location_id = locations.id;
SELECT COUNT(*) FROM prices, products, locations WHERE prices.product_id = products.id AND prices.location_id = locations.id;
SELECT * FROM purchases, receipts WHERE purchases.receipt_id = receipts.id AND receipts.datetime < '2016-07-24 11:11:25-07' AND receipts.datetime > '2016-05-24 11:11:25-07';
SELECT SUM(amount), SUM(result) FROM purchases, receipts WHERE purchases.price_id = 3 AND purchases.receipt_id = receipts.id AND receipts.datetime < '2016-07-24 11:11:25-07' AND receipts.datetime > '2016-05-24 11:11:25-07';

-- Get 

