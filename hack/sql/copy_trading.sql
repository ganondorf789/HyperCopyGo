CREATE TABLE copy_trading (
    id                              BIGSERIAL PRIMARY KEY,
    user_id                         BIGINT      NOT NULL DEFAULT 0,
    target_wallet                   VARCHAR(255) NOT NULL DEFAULT '',
    target_wallet_platform          VARCHAR(64)  NOT NULL DEFAULT '',
    remark                          VARCHAR(255) NOT NULL DEFAULT '',
    leverage                        INT          NOT NULL DEFAULT 1,
    margin_mode                     INT          NOT NULL DEFAULT 1,
    follow_model                    INT          NOT NULL DEFAULT 1,
    follow_model_value              NUMERIC(20,8) NOT NULL DEFAULT 0,
    min_value                       NUMERIC(20,8) NOT NULL DEFAULT 0,
    max_value                       NUMERIC(20,8) NOT NULL DEFAULT 0,
    max_margin_usage                NUMERIC(10,4) NOT NULL DEFAULT 0,
    tp_value                        NUMERIC(10,4) NOT NULL DEFAULT 0,
    sl_value                        NUMERIC(10,4) NOT NULL DEFAULT 0,
    opt_reverse_follow_order        INT          NOT NULL DEFAULT 0,
    opt_followup_decrease           INT          NOT NULL DEFAULT 0,
    opt_followup_increase           INT          NOT NULL DEFAULT 0,
    opt_forced_liquidation_protection INT        NOT NULL DEFAULT 0,
    opt_position_increase_opening   INT          NOT NULL DEFAULT 0,
    opt_slippage_protection         INT          NOT NULL DEFAULT 0,
    symbol_list_type                VARCHAR(16)  NOT NULL DEFAULT 'WHITE',
    symbol_list                     TEXT         NOT NULL DEFAULT '',
    main_wallet                     VARCHAR(255) NOT NULL DEFAULT '',
    main_wallet_platform            VARCHAR(64)  NOT NULL DEFAULT '',
    status                          INT          NOT NULL DEFAULT 1,
    created_at                      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at                      TIMESTAMPTZ  NOT NULL DEFAULT NOW()
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
