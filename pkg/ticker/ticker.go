package ticker

import (
	"go.uber.org/zap"
)

func Init() {
	errCh := make(chan error)

	//错误处理
	go func() {
		for err := range errCh {
			zap.L().Error("ticker in pkg happen err:%v", zap.Error(err))
		}
	}()
}
