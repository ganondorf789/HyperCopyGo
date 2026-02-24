-- 持仓表
CREATE TABLE IF NOT EXISTS position (
    id                  BIGSERIAL       PRIMARY KEY,
    "user"              VARCHAR(255)    NOT NULL DEFAULT '',           -- 用户钱包地址
    symbol              VARCHAR(64)     NOT NULL DEFAULT '',           -- 交易对符号
    position_size       NUMERIC(30,8)   NOT NULL DEFAULT 0,           -- 持仓数量（负数为空头）
    entry_price         NUMERIC(30,8)   NOT NULL DEFAULT 0,           -- 开仓均价
    mark_price          NUMERIC(30,8)   NOT NULL DEFAULT 0,           -- 标记价格
    liq_price           NUMERIC(30,8)   NOT NULL DEFAULT 0,           -- 强平价格
    leverage            INT             NOT NULL DEFAULT 1,            -- 杠杆倍数
    margin_balance      NUMERIC(30,8)   NOT NULL DEFAULT 0,           -- 保证金余额
    position_value_usd  NUMERIC(30,8)   NOT NULL DEFAULT 0,           -- 持仓价值(USD)
    unrealized_pnl      NUMERIC(30,8)   NOT NULL DEFAULT 0,           -- 未实现盈亏
    funding_fee         NUMERIC(30,8)   NOT NULL DEFAULT 0,           -- 资金费用
    margin_mode         VARCHAR(16)     NOT NULL DEFAULT 'cross',     -- 保证金模式 cross/isolated
    labels              TEXT            NOT NULL DEFAULT '',           -- 标签，逗号分隔
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  position IS '持仓表';
COMMENT ON COLUMN position."user" IS '用户钱包地址';
COMMENT ON COLUMN position.symbol IS '交易对符号';
COMMENT ON COLUMN position.position_size IS '持仓数量（负数为空头）';
COMMENT ON COLUMN position.entry_price IS '开仓均价';
COMMENT ON COLUMN position.mark_price IS '标记价格';
COMMENT ON COLUMN position.liq_price IS '强平价格';
COMMENT ON COLUMN position.leverage IS '杠杆倍数';
COMMENT ON COLUMN position.margin_balance IS '保证金余额';
COMMENT ON COLUMN position.position_value_usd IS '持仓价值(USD)';
COMMENT ON COLUMN position.unrealized_pnl IS '未实现盈亏';
COMMENT ON COLUMN position.funding_fee IS '资金费用';
COMMENT ON COLUMN position.margin_mode IS '保证金模式 cross/isolated';
COMMENT ON COLUMN position.labels IS '标签，逗号分隔';

CREATE INDEX idx_position_user ON position("user");
CREATE INDEX idx_position_symbol ON position(symbol);
