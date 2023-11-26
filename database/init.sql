CREATE TABLE Delivery (
    delivery_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    phone VARCHAR(30),
    zip VARCHAR(30),
    city VARCHAR(255),
    address VARCHAR(255),
    region VARCHAR(255),
    email VARCHAR(255)
);

CREATE TABLE Payment (
    payment_id SERIAL PRIMARY KEY,
    transaction VARCHAR(255),
    request_id VARCHAR(255),
    currency VARCHAR(3),
    provider VARCHAR(255),
    amount INT,
    payment_dt INT,
    bank VARCHAR(255),
    delivery_cost INT,
    goods_total INT,
    custom_fee INT
);

CREATE TABLE "Order" (
    order_id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255),
    track_number VARCHAR(255),
    entry VARCHAR(255),
    delivery_id INT REFERENCES Delivery(delivery_id),
    payment_id INT REFERENCES Payment(payment_id),
    locale VARCHAR(5),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey VARCHAR(10),
    sm_id INT,
    date_created TIMESTAMP,
    oof_shard VARCHAR(10)
);

CREATE TABLE Item (
    item_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES "Order"(order_id),
    chrt_id INT,
    track_number VARCHAR(255),
    price INT,
    rid VARCHAR(255),
    name VARCHAR(255),
    sale INT,
    size VARCHAR(50),
    total_price INT,
    nm_id INT,
    brand VARCHAR(255),
    status INT
);