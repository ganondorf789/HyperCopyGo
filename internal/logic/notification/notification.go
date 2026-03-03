package notification

import (
	"context"

	v1 "demo/api/notification/v1"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func init() {
	service.RegisterNotification(&sNotification{})
}

type sNotification struct{}

// allCategories 所有通知分类，保证 1 级页面固定顺序
var allCategories = []string{"public", "copy_trading", "whale", "track", "market"}

// summaryRow 原生 SQL 聚合行
type summaryRow struct {
	Category    string `json:"category"`
	UnreadCount int    `json:"unread_count"`
}

func (s *sNotification) Summary(ctx context.Context, userId int64) (res *v1.NotificationSummaryRes, err error) {
	// 1. 各分类未读数
	sql := `
		SELECT n.category, COUNT(*) AS unread_count
		FROM notification n
		LEFT JOIN notification_read nr ON nr.notification_id = n.id AND nr.user_id = $1
		WHERE (n.user_id = $1 OR n.user_id = 0)
		  AND n.status = 1
		  AND nr.id IS NULL
		GROUP BY n.category
	`
	var rows []summaryRow
	err = dao.Notification.DB().GetScan(ctx, &rows, sql, userId)
	if err != nil {
		return nil, err
	}
	unreadMap := make(map[string]int, len(rows))
	for _, r := range rows {
		unreadMap[r.Category] = r.UnreadCount
	}

	// 2. 各分类最新一条消息
	latestSQL := `
		SELECT DISTINCT ON (category) id, category, title, content, level, created_at
		FROM notification
		WHERE (user_id = $1 OR user_id = 0) AND status = 1
		ORDER BY category, created_at DESC
	`
	type latestRow struct {
		entity.Notification
	}
	var latestRows []latestRow
	err = dao.Notification.DB().GetScan(ctx, &latestRows, latestSQL, userId)
	if err != nil {
		return nil, err
	}
	latestMap := make(map[string]*model.NotificationPreview, len(latestRows))
	for _, r := range latestRows {
		latestMap[r.Category] = &model.NotificationPreview{
			Id:        r.Id,
			Title:     r.Title,
			Content:   r.Content,
			Level:     r.Level,
			CreatedAt: r.CreatedAt,
		}
	}

	// 3. 组装固定顺序的分类列表
	categories := make([]model.NotificationCategorySummary, 0, len(allCategories))
	for _, cat := range allCategories {
		categories = append(categories, model.NotificationCategorySummary{
			Category:    cat,
			UnreadCount: unreadMap[cat],
			Latest:      latestMap[cat],
		})
	}

	return &v1.NotificationSummaryRes{Categories: categories}, nil
}

func (s *sNotification) List(ctx context.Context, userId int64, in v1.NotificationListReq) (res *v1.NotificationListRes, err error) {
	m := dao.Notification.Ctx(ctx).
		Where("(user_id = ? OR user_id = 0)", userId).
		Where(do.Notification{Category: in.Category, Status: 1})

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var items []entity.Notification
	err = m.Page(in.Page, in.PageSize).
		OrderDesc(dao.Notification.Columns().Id).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	// 批量查询已读状态
	readSet := make(map[int64]bool)
	if len(items) > 0 {
		ids := make([]int64, 0, len(items))
		for _, n := range items {
			ids = append(ids, n.Id)
		}
		var reads []entity.NotificationRead
		err = dao.NotificationRead.Ctx(ctx).
			Where(do.NotificationRead{UserId: userId}).
			Where("notification_id IN(?)", ids).
			Scan(&reads)
		if err != nil {
			return nil, err
		}
		for _, r := range reads {
			readSet[r.NotificationId] = true
		}
	}

	list := make([]model.NotificationItem, 0, len(items))
	for _, e := range items {
		list = append(list, model.NotificationItem{
			Id:        e.Id,
			Category:  e.Category,
			Title:     e.Title,
			Content:   e.Content,
			RefId:     e.RefId,
			RefType:   e.RefType,
			Level:     e.Level,
			IsRead:    readSet[e.Id],
			CreatedAt: e.CreatedAt,
		})
	}

	return &v1.NotificationListRes{
		List:  list,
		Total: total,
		Page:  in.Page,
	}, nil
}

func (s *sNotification) Read(ctx context.Context, userId int64, ids []int64) error {
	for _, nid := range ids {
		count, err := dao.NotificationRead.Ctx(ctx).
			Where(do.NotificationRead{UserId: userId, NotificationId: nid}).
			Count()
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		_, err = dao.NotificationRead.Ctx(ctx).Data(do.NotificationRead{
			UserId:         userId,
			NotificationId: nid,
		}).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sNotification) ReadAll(ctx context.Context, userId int64, category string) error {
	sql := `
		INSERT INTO notification_read (user_id, notification_id)
		SELECT $1, n.id
		FROM notification n
		LEFT JOIN notification_read nr ON nr.notification_id = n.id AND nr.user_id = $1
		WHERE (n.user_id = $1 OR n.user_id = 0)
		  AND n.category = $2
		  AND n.status = 1
		  AND nr.id IS NULL
	`
	_, err := dao.Notification.DB().Exec(ctx, sql, userId, category)
	return err
}
