-- 交易员 PnL 历史表
-- 每条记录: 一个交易员 + 一个时间框架 + 完整 PnL 数据点列表(JSONB)
CREATE TABLE IF NOT EXISTS pnl_history (
    id          BIGSERIAL       PRIMARY KEY,
    "user"      VARCHAR(255)    NOT NULL DEFAULT '',           -- 交易员钱包地址
    timeframe   VARCHAR(16)     NOT NULL DEFAULT '',           -- 时间框架: 1D, 7D, 30D, All
    pnl_list    JSONB           NOT NULL DEFAULT '[]',         -- PnL 数据点列表 [{"ts":1771322655316,"v":"0"}, ...]
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  pnl_history IS '交易员PnL历史';
COMMENT ON COLUMN pnl_history."user" IS '交易员钱包地址';
COMMENT ON COLUMN pnl_history.timeframe IS '时间框架: 1D, 7D, 30D, All';
COMMENT ON COLUMN pnl_history.pnl_list IS 'PnL数据点列表';

CREATE UNIQUE INDEX idx_pnl_history_user_timeframe ON pnl_history("user", timeframe);
