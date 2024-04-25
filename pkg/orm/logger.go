package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/xian1367/layout-go-zero/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// Logger 操作对象，实现 gormLogger.Interface
type Logger struct {
	Logger        logx.Logger
	SlowThreshold time.Duration
}

// NewGormLogger 外部调用。实例化一个 GormLogger 对象，示例：
//
//	DB, err := orm.Open(dbConfig, &orm.Config{
//	    Logger: logger.NewGormLogger(),
//	})
func NewGormLogger() Logger {
	return Logger{
		Logger:        logx.WithContext(context.Background()), // 使用全局的 logger.Logger 对象
		SlowThreshold: 200 * time.Millisecond,                 // 慢查询阈值，单位为千分之一秒
	}
}

// LogMode 实现 gormLogger.Interface 的 LogMode 方法
func (l Logger) LogMode(level logger.LogLevel) logger.Interface {
	return Logger{
		Logger:        l.Logger,
		SlowThreshold: l.SlowThreshold,
	}
}

// Info 实现 gormLogger.Interface 的 Info 方法
func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	l.logger().WithContext(ctx).Debugf(str, args...)
}

// Warn 实现 gormLogger.Interface 的 Warn 方法
func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	l.logger().WithContext(ctx).Errorf(str, args...)
}

// Error 实现 gormLogger.Interface 的 Error 方法
func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	l.logger().WithContext(ctx).Errorf(str, args...)
}

// Trace 实现 gormLogger.Interface 的 Trace 方法
func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// 获取运行时间
	elapsed := time.Since(begin)
	// 获取 SQL 请求和返回条数
	sql, rows := fc()

	// 通用字段
	logFields := []logx.LogField{
		logx.Field("sql", sql),
		logx.Field("time", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)),
		logx.Field("rows", rows),
	}
	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.logger().WithContext(ctx).Errorw("Database ErrRecordNotFound", logFields...)
		} else {
			// 其他错误使用 error 等级
			logFields = append(logFields, logx.Field("error", err.Error()))
			l.logger().WithContext(ctx).Errorw("Database Error", logFields...)
		}
	}

	// 慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.Logger.WithContext(ctx).Sloww("Database Slow Log", logFields...)
	}

	// 记录所有 SQL 请求
	if config.Get().Debug {
		l.logger().WithContext(ctx).Debugw("Database Query", logFields...)
	}
}

// logger 内用的辅助方法，确保 Logger 内置信息 Caller 的准确性
func (l Logger) logger() logx.Logger {
	return l.Logger
}
