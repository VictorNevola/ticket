-- +migrate Up
CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    tax_id TEXT NOT NULL UNIQUE,
    secret_key TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE promotions (
    id UUID PRIMARY KEY,
    company_id UUID REFERENCES companies(id),
    name TEXT NOT NULL,
    text_message_in_progress TEXT NOT NULL,
    text_message_success TEXT NOT NULL,
    start_date TIMESTAMPTZ NOT NULL,
    end_date TIMESTAMPTZ NOT NULL,
    qty_max_users INT NOT NULL,
    vouchers_per_user INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_promotions_company_id ON promotions(company_id);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE users_in_promotions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    promotion_id UUID NOT NULL REFERENCES promotions(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_in_promotions_user_id ON users_in_promotions(user_id);
CREATE INDEX idx_users_in_promotions_promotion_id ON users_in_promotions(promotion_id);

CREATE TABLE vouchers (
    id UUID PRIMARY KEY,
    voucher_hash TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    confirmed_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE voucher_usages (
    id UUID PRIMARY KEY,
    promotion_id UUID NOT NULL REFERENCES promotions(id),
    voucher_id UUID NOT NULL REFERENCES vouchers(id),
    user_id UUID NOT NULL REFERENCES users(id)
);

CREATE INDEX idx_voucher_usages_promotion_id ON voucher_usages(promotion_id);
CREATE INDEX idx_voucher_usages_voucher_id ON voucher_usages(voucher_id);
CREATE INDEX idx_voucher_usages_user_id ON voucher_usages(user_id);

-- +migrate Down
DROP TABLE companies;
DROP TABLE promotions;
DROP TABLE users_in_promotions;
DROP TABLE users;
DROP TABLE voucher_usages;
DROP TABLE vouchers;