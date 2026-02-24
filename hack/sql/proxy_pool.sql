-- 代理 IP 池表
CREATE TABLE IF NOT EXISTS proxy_pool (
    id          BIGSERIAL       PRIMARY KEY,
    host        VARCHAR(255)    NOT NULL DEFAULT '',           -- 代理主机地址
    port        INT             NOT NULL DEFAULT 0,            -- 代理端口
    username    VARCHAR(255)    NOT NULL DEFAULT '',           -- 认证用户名
    password    VARCHAR(255)    NOT NULL DEFAULT '',           -- 认证密码
    status      SMALLINT        NOT NULL DEFAULT 1,            -- 状态: 1=启用 0=禁用
    remark      VARCHAR(255)    NOT NULL DEFAULT '',           -- 备注
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE  proxy_pool IS '代理IP池';
COMMENT ON COLUMN proxy_pool.host IS '代理主机地址';
COMMENT ON COLUMN proxy_pool.port IS '代理端口';
COMMENT ON COLUMN proxy_pool.username IS '认证用户名';
COMMENT ON COLUMN proxy_pool.password IS '认证密码';
COMMENT ON COLUMN proxy_pool.status IS '状态: 1=启用 0=禁用';
COMMENT ON COLUMN proxy_pool.remark IS '备注';

-- 插入初始代理
INSERT INTO proxy_pool (host, port, username, password, status) VALUES
('81.22.133.221', 12323, '14a695a9e63f2', 'b951f67305', 1);
