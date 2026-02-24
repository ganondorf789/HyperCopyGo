-- 跟踪钱包表
CREATE TABLE IF NOT EXISTS my_track_wallet (
    id              BIGSERIAL    PRIMARY KEY,
    user_id         BIGINT       NOT NULL DEFAULT 0,          -- 所属用户ID
    wallet          VARCHAR(255) NOT NULL DEFAULT '',         -- 跟踪的钱包地址
    remark          VARCHAR(255) NOT NULL DEFAULT '',         -- 备注
    enable_notify   SMALLINT     NOT NULL DEFAULT 0,          -- 是否开启通知 0:关 1:开
    notify_action   VARCHAR(64)  NOT NULL DEFAULT '',         -- 通知动作 1:开仓 2:平仓 3:加仓 4:减仓，逗号分隔
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
