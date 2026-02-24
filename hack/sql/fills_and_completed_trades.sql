-- 成交记录表 (Fills)
-- 每条记录对应 Hyperliquid API 返回的一条 Fill
CREATE TABLE IF NOT EXISTS fills (
    id              BIGSERIAL       PRIMARY KEY,
    "user"          VARCHAR(255)    NOT NULL DEFAULT '',           -- 交易员钱包地址
    coin            VARCHAR(64)     NOT NULL DEFAULT '',           -- 币种
    dir             VARCHAR(32)     NOT NULL DEFAULT '',           -- 方向描述
    side            VARCHAR(16)     NOT NULL DEFAULT '',           -- 买卖方向 B/A
    px              NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- 成交价格
    sz              NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- 成交数量
    closed_pnl      NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- 已实现盈亏
    fee             NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- 手续费
    fee_token       VARCHAR(32)     NOT NULL DEFAULT '',           -- 手续费币种
    builder_fee     NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- builder手续费
    hash            VARCHAR(255)    NOT NULL DEFAULT '',           -- 交易哈希
    oid             BIGINT          NOT NULL DEFAULT 0,            -- 订单ID
    tid             BIGINT          NOT NULL DEFAULT 0,            -- 成交ID
    crossed         BOOLEAN         NOT NULL DEFAULT FALSE,        -- 是否crossed
    start_position  NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- 成交前持仓
    fill_time       BIGINT          NOT NULL DEFAULT 0,            -- 成交时间戳(ms)
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  fills IS '成交记录';
COMMENT ON COLUMN fills."user" IS '交易员钱包地址';
COMMENT ON COLUMN fills.coin IS '币种';
COMMENT ON COLUMN fills.dir IS '方向描述';
COMMENT ON COLUMN fills.side IS '买卖方向 B/A';
COMMENT ON COLUMN fills.px IS '成交价格';
COMMENT ON COLUMN fills.sz IS '成交数量';
COMMENT ON COLUMN fills.closed_pnl IS '已实现盈亏';
COMMENT ON COLUMN fills.fee IS '手续费';
COMMENT ON COLUMN fills.fee_token IS '手续费币种';
COMMENT ON COLUMN fills.builder_fee IS 'builder手续费';
COMMENT ON COLUMN fills.hash IS '交易哈希';
COMMENT ON COLUMN fills.oid IS '订单ID';
COMMENT ON COLUMN fills.tid IS '成交ID';
COMMENT ON COLUMN fills.crossed IS '是否crossed';
COMMENT ON COLUMN fills.start_position IS '成交前持仓';
COMMENT ON COLUMN fills.fill_time IS '成交时间戳(ms)';

CREATE INDEX idx_fills_user ON fills("user");
CREATE INDEX idx_fills_user_coin ON fills("user", coin);
CREATE INDEX idx_fills_tid ON fills(tid);

-- 已完成交易表 (Completed Trades)
-- 由 Fills 计算聚合，每条记录对应一笔完整的开平仓交易
CREATE TABLE IF NOT EXISTS completed_trades (
    id              BIGSERIAL       PRIMARY KEY,
    "user"          VARCHAR(255)    NOT NULL DEFAULT '',           -- 交易员钱包地址
    coin            VARCHAR(64)     NOT NULL DEFAULT '',           -- 币种
    side            VARCHAR(16)     NOT NULL DEFAULT '',           -- 方向 long/short
    entry_px        NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- 开仓均价
    close_px        NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- 平仓均价
    sz              NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- 交易数量
    closed_pnl      NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- 已实现盈亏
    total_fee       NUMERIC(30,10)  NOT NULL DEFAULT 0,            -- 总手续费
    open_time       BIGINT          NOT NULL DEFAULT 0,            -- 开仓时间戳(ms)
    close_time      BIGINT          NOT NULL DEFAULT 0,            -- 平仓时间戳(ms)
    duration_ms     BIGINT          NOT NULL DEFAULT 0,            -- 持仓时长(ms)
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  completed_trades IS '已完成交易';
COMMENT ON COLUMN completed_trades."user" IS '交易员钱包地址';
COMMENT ON COLUMN completed_trades.coin IS '币种';
COMMENT ON COLUMN completed_trades.side IS '方向 long/short';
COMMENT ON COLUMN completed_trades.entry_px IS '开仓均价';
COMMENT ON COLUMN completed_trades.close_px IS '平仓均价';
COMMENT ON COLUMN completed_trades.sz IS '交易数量';
COMMENT ON COLUMN completed_trades.closed_pnl IS '已实现盈亏';
COMMENT ON COLUMN completed_trades.total_fee IS '总手续费';
COMMENT ON COLUMN completed_trades.open_time IS '开仓时间戳(ms)';
COMMENT ON COLUMN completed_trades.close_time IS '平仓时间戳(ms)';
COMMENT ON COLUMN completed_trades.duration_ms IS '持仓时长(ms)';

CREATE INDEX idx_completed_trades_user ON completed_trades("user");
CREATE INDEX idx_completed_trades_user_coin ON completed_trades("user", coin);
