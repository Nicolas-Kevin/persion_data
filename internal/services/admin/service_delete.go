package admin

import (
	"go-gin-api/configs"
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/pkg/password"
	"go-gin-api/internal/repository/mysql"
	"go-gin-api/internal/repository/mysql/admin"
	"go-gin-api/internal/repository/redis"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := admin.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), redis.WithTrace(ctx.Trace()))
	return
}
