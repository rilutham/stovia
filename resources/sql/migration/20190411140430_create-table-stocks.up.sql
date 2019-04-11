CREATE TABLE IF NOT EXISTS stocks (
    id BIGSERIAL PRIMARY KEY,
    code text NOT NULL UNIQUE,
    company_name text NOT NULL,
    year text NOT NULL,
    quarter integer NOT NULL,
    current_price numeric DEFAULT 0.0,
    total_equity numeric NOT NULL,
    net_profit numeric NOT NULL,
    number_of_shares bigint NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
