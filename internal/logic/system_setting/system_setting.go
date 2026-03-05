package system_setting

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	v1 "demo/api/system_setting/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterSystemSetting(&sSystemSetting{})
}

type sSystemSetting struct{}

func (s *sSystemSetting) Get(ctx context.Context) (res *v1.SystemSettingGetRes, err error) {
	var setting entity.SystemSetting
	err = dao.SystemSetting.Ctx(ctx).Scan(&setting)
	if err != nil {
		return nil, err
	}
	if setting.Id == 0 {
		return nil, fmt.Errorf("系统设置未初始化")
	}
	return &v1.SystemSettingGetRes{
		SystemSettingItem: model.SystemSettingItem{
			Id:                     setting.Id,
			MarketMinutes:          setting.MarketMinutes,
			MarketNewPositionCount: setting.MarketNewPositionCount,
			CreatedAt:              setting.CreatedAt,
			UpdatedAt:              setting.UpdatedAt,
		},
	}, nil
}

func (s *sSystemSetting) Update(ctx context.Context, in v1.SystemSettingUpdateReq) error {
	count, err := dao.SystemSetting.Ctx(ctx).Count()
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = dao.SystemSetting.Ctx(ctx).Data(g.Map{
			"market_minutes":            in.MarketMinutes,
			"market_new_position_count": in.MarketNewPositionCount,
		}).Insert()
		return err
	}

	_, err = dao.SystemSetting.Ctx(ctx).Data(g.Map{
		"market_minutes":            in.MarketMinutes,
		"market_new_position_count": in.MarketNewPositionCount,
	}).Update()
	return err
}
