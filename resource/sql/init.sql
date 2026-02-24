-- 用户表
CREATE TABLE IF NOT EXISTS "user" (
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(64)  NOT NULL UNIQUE,
    password   VARCHAR(128) NOT NULL,
    nickname   VARCHAR(64)  NOT NULL DEFAULT '',
    avatar     VARCHAR(255) NOT NULL DEFAULT '',
    email      VARCHAR(128) NOT NULL DEFAULT '',
    phone      VARCHAR(32)  NOT NULL DEFAULT '',
    status     SMALLINT     NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  "user" IS '用户表';
COMMENT ON COLUMN "user".username IS '用户名';
COMMENT ON COLUMN "user".password IS '密码';
COMMENT ON COLUMN "user".nickname IS '昵称';
COMMENT ON COLUMN "user".avatar IS '头像';
COMMENT ON COLUMN "user".email IS '邮箱';
COMMENT ON COLUMN "user".phone IS '手机号';
COMMENT ON COLUMN "user".status IS '状态 1:正常 0:禁用';

-- 后台管理员表
CREATE TABLE IF NOT EXISTS "admin" (
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(64)  NOT NULL UNIQUE,
    password   VARCHAR(128) NOT NULL,
    realname   VARCHAR(64)  NOT NULL DEFAULT '',
    role       VARCHAR(32)  NOT NULL DEFAULT 'admin',
    status     SMALLINT     NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  "admin" IS '后台管理员表';
COMMENT ON COLUMN "admin".username IS '用户名';
COMMENT ON COLUMN "admin".password IS '密码';
COMMENT ON COLUMN "admin".realname IS '真实姓名';
COMMENT ON COLUMN "admin".role IS '角色 admin/super_admin';
COMMENT ON COLUMN "admin".status IS '状态 1:正常 0:禁用';

-- 跟单交易配置表
CREATE TABLE IF NOT EXISTS copy_trading (
    id                                BIGSERIAL    PRIMARY KEY,
    user_id                           BIGINT       NOT NULL DEFAULT 0,          -- 所属用户ID
    target_wallet                     VARCHAR(255) NOT NULL DEFAULT '',         -- 目标钱包地址
    target_wallet_platform            VARCHAR(64)  NOT NULL DEFAULT '',         -- 目标钱包平台
    remark                            VARCHAR(255) NOT NULL DEFAULT '',         -- 备注
    leverage                          INT          NOT NULL DEFAULT 1,          -- 杠杆倍数
    margin_mode                       INT          NOT NULL DEFAULT 1,          -- 保证金模式 1:逐仓 2:全仓
    follow_model                      INT          NOT NULL DEFAULT 1,          -- 跟单模式 1:固定金额 2:固定比例
    follow_model_value                NUMERIC(20,8) NOT NULL DEFAULT 0,         -- 跟单模式值
    min_value                         NUMERIC(20,8) NOT NULL DEFAULT 0,         -- 最小下单金额
    max_value                         NUMERIC(20,8) NOT NULL DEFAULT 0,         -- 最大下单金额
    max_margin_usage                  NUMERIC(10,4) NOT NULL DEFAULT 0,         -- 最大保证金使用率
    tp_value                          NUMERIC(10,4) NOT NULL DEFAULT 0,         -- 止盈比例
    sl_value                          NUMERIC(10,4) NOT NULL DEFAULT 0,         -- 止损比例
    opt_reverse_follow_order          INT          NOT NULL DEFAULT 0,          -- 反向跟单 0:关 1:开
    opt_followup_decrease             INT          NOT NULL DEFAULT 0,          -- 跟随减仓 0:关 1:开
    opt_followup_increase             INT          NOT NULL DEFAULT 0,          -- 跟随加仓 0:关 1:开
    opt_forced_liquidation_protection INT          NOT NULL DEFAULT 0,          -- 强平保护 0:关 1:开
    opt_position_increase_opening     INT          NOT NULL DEFAULT 0,          -- 加仓开仓 0:关 1:开
    opt_slippage_protection           INT          NOT NULL DEFAULT 0,          -- 滑点保护 0:关 1:开
    symbol_list_type                  VARCHAR(16)  NOT NULL DEFAULT 'WHITE',    -- 交易对列表类型 WHITE:白名单 BLACK:黑名单
    symbol_list                       TEXT         NOT NULL DEFAULT '',         -- 交易对列表,逗号分隔
    main_wallet                       VARCHAR(255) NOT NULL DEFAULT '',         -- 主钱包地址
    main_wallet_platform              VARCHAR(64)  NOT NULL DEFAULT '',         -- 主钱包平台
    status                            INT          NOT NULL DEFAULT 1,          -- 状态 0:停用 1:启用
    created_at                        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at                        TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  copy_trading IS '跟单交易配置表';
COMMENT ON COLUMN copy_trading.user_id IS '所属用户ID';
COMMENT ON COLUMN copy_trading.target_wallet IS '目标钱包地址';
COMMENT ON COLUMN copy_trading.target_wallet_platform IS '目标钱包平台';
COMMENT ON COLUMN copy_trading.remark IS '备注';
COMMENT ON COLUMN copy_trading.leverage IS '杠杆倍数';
COMMENT ON COLUMN copy_trading.margin_mode IS '保证金模式 1:逐仓 2:全仓';
COMMENT ON COLUMN copy_trading.follow_model IS '跟单模式 1:固定金额 2:固定比例';
COMMENT ON COLUMN copy_trading.follow_model_value IS '跟单模式值';
COMMENT ON COLUMN copy_trading.min_value IS '最小下单金额';
COMMENT ON COLUMN copy_trading.max_value IS '最大下单金额';
COMMENT ON COLUMN copy_trading.max_margin_usage IS '最大保证金使用率';
COMMENT ON COLUMN copy_trading.tp_value IS '止盈比例';
COMMENT ON COLUMN copy_trading.sl_value IS '止损比例';
COMMENT ON COLUMN copy_trading.opt_reverse_follow_order IS '反向跟单 0:关 1:开';
COMMENT ON COLUMN copy_trading.opt_followup_decrease IS '跟随减仓 0:关 1:开';
COMMENT ON COLUMN copy_trading.opt_followup_increase IS '跟随加仓 0:关 1:开';
COMMENT ON COLUMN copy_trading.opt_forced_liquidation_protection IS '强平保护 0:关 1:开';
COMMENT ON COLUMN copy_trading.opt_position_increase_opening IS '加仓开仓 0:关 1:开';
COMMENT ON COLUMN copy_trading.opt_slippage_protection IS '滑点保护 0:关 1:开';
COMMENT ON COLUMN copy_trading.symbol_list_type IS '交易对列表类型 WHITE:白名单 BLACK:黑名单';
COMMENT ON COLUMN copy_trading.symbol_list IS '交易对列表,逗号分隔';
COMMENT ON COLUMN copy_trading.main_wallet IS '主钱包地址';
COMMENT ON COLUMN copy_trading.main_wallet_platform IS '主钱包平台';
COMMENT ON COLUMN copy_trading.status IS '状态 0:停用 1:启用';

-- 钱包表
CREATE TABLE IF NOT EXISTS wallet (
    id                  BIGSERIAL    PRIMARY KEY,
    user_id             BIGINT       NOT NULL DEFAULT 0,          -- 所属用户ID
    address             VARCHAR(255) NOT NULL DEFAULT '',         -- 钱包地址
    api_wallet_address  VARCHAR(255) NOT NULL DEFAULT '',         -- API Wallet Address
    api_secret_key      VARCHAR(255) NOT NULL DEFAULT '',         -- API Secret Key
    remark              VARCHAR(255) NOT NULL DEFAULT '',         -- 备注
    status              SMALLINT     NOT NULL DEFAULT 1,          -- 状态 1:正常 0:禁用
    created_at          TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  wallet IS '钱包表';
COMMENT ON COLUMN wallet.user_id IS '所属用户ID';
COMMENT ON COLUMN wallet.address IS '钱包地址';
COMMENT ON COLUMN wallet.api_wallet_address IS 'API Wallet Address';
COMMENT ON COLUMN wallet.api_secret_key IS 'API Secret Key';
COMMENT ON COLUMN wallet.remark IS '备注';
COMMENT ON COLUMN wallet.status IS '状态 1:正常 0:禁用';

-- 索引
CREATE INDEX idx_user_status ON "user"(status);
CREATE INDEX idx_admin_status ON "admin"(status);
CREATE INDEX idx_copy_trading_user_id ON copy_trading(user_id);
CREATE INDEX idx_copy_trading_status ON copy_trading(status);
CREATE INDEX idx_wallet_user_id ON wallet(user_id);
CREATE INDEX idx_wallet_status ON wallet(status);

-- 跟踪钱包表
CREATE TABLE IF NOT EXISTS my_track_wallet (
    id              BIGSERIAL    PRIMARY KEY,
    user_id         BIGINT       NOT NULL DEFAULT 0,          -- 所属用户ID
    wallet          VARCHAR(255) NOT NULL DEFAULT '',         -- 跟踪的钱包地址
    remark          VARCHAR(255) NOT NULL DEFAULT '',         -- 备注
    enable_notify   SMALLINT     NOT NULL DEFAULT 0,          -- 是否开启通知 0:关 1:开
    notify_action   VARCHAR(64)  NOT NULL DEFAULT '',         -- 通知动作 1:开仓 2:平仓 3:加仓 4:减仓
    lang            VARCHAR(16)  NOT NULL DEFAULT 'zh',       -- 语言
    status          SMALLINT     NOT NULL DEFAULT 1,          -- 状态 1:正常 0:禁用
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  my_track_wallet IS '跟踪钱包表';
COMMENT ON COLUMN my_track_wallet.user_id IS '所属用户ID';
COMMENT ON COLUMN my_track_wallet.wallet IS '跟踪的钱包地址';
COMMENT ON COLUMN my_track_wallet.remark IS '备注';
COMMENT ON COLUMN my_track_wallet.enable_notify IS '是否开启通知 0:关 1:开';
COMMENT ON COLUMN my_track_wallet.notify_action IS '通知动作 1:开仓 2:平仓 3:加仓 4:减仓';
COMMENT ON COLUMN my_track_wallet.lang IS '语言';
COMMENT ON COLUMN my_track_wallet.status IS '状态 1:正常 0:禁用';

CREATE INDEX idx_my_track_wallet_user_id ON my_track_wallet(user_id);
CREATE INDEX idx_my_track_wallet_status ON my_track_wallet(status);

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
