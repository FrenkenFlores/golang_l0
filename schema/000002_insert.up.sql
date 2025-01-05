-- main orders table
INSERT INTO orders (
    order_uid,
    track_number,
    entry,
    locale,
    internal_signature,
    customer_id,
    delivery_service,
    shardkey,
    sm_id,
    date_created,
    oof_shard
)
VALUES
(
    'b563feb7b2b84b6test',   -- order_uid
    'WBILMTESTTRACK',        -- track_number
    'WBIL',                  -- entry
    'en',                    -- locale
    NULL,                      -- internal_signature
    'test',                  -- customer_id
    'meest',                 -- delivery_service
    '9',                     -- shardkey
    99,                      -- sm_id
    '2021-11-26T06:22:19Z',  -- date_created
    '1'                      -- oof_shard
);

-- one-to-one relationship with the orders table
INSERT INTO delivery (
    order_id,
    name,
    phone,
    zip,
    city,
    address,
    region,
    email
)
VALUES
(
    1,                    -- Replace with the actual `id` from the `orders` table
    'Test Testov',        -- name
    '+9720000000',        -- phone
    '2639809',            -- zip
    'Kiryat Mozkin',      -- city
    'Ploshad Mira 15',    -- address
    'Kraiot',             -- region
    'test@gmail.com'      -- email
);

-- one-to-one relationship with the orders table
INSERT INTO payment (
    order_id,
    transaction,
    request_id,
    currency,
    provider,
    amount,
    payment_dt,
    bank,
    delivery_cost,
    goods_total,
    custom_fee
) VALUES (
    1, -- Replace with the actual order_id from the orders table
    'b563feb7b2b84b6test',
    NULL, -- Empty string for request_id
    'USD',
    'wbpay',
    1817,
    1637907727,
    'alpha',
    1500,
    317,
    0
);


-- one-to-many relationship with the orders table
INSERT INTO items (
    order_id,
    chrt_id,
    track_number,
    price,
    rid,
    name,
    sale,
    size,
    total_price,
    nm_id,
    brand,
    status
) VALUES (
    1, -- Replace with the actual order_id from the orders table
    9934930,
    'WBILMTESTTRACK',
    453,
    'ab4219087a764ae0btest',
    'Mascaras',
    30,
    '0',
    317,
    2389212,
    'Vivienne Sabo',
    202
);
