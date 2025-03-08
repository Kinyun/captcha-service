package models

import (
	"captcha-service/app/config/constant"
	"captcha-service/app/logger/singleton"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

func IsJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

func ServiceLog(ctx context.Context, begin time.Time, param, request, response interface{}, err error) {
	var (
		requestID = FromContext(ctx)
		fields    = []zap.Field{zap.Duration("took", time.Since(begin))}
	)
	if param != nil {
		parB, _ := json.Marshal(param)
		fields = append(fields, zap.Any("param", json.RawMessage(parB)))
	}
	if request != nil {
		reqB, err := json.Marshal(request)
		if err != nil {
			fields = append(fields, zap.Any("request", json.RawMessage(request.([]byte))))
		} else {
			fields = append(fields, zap.Any("request", json.RawMessage(reqB)))
		}
	}
	if response != nil {
		resB, _ := json.Marshal(response)
		fields = append(fields, zap.Any("response", json.RawMessage(resB)))
	}

	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	singleton.WithRequestID(requestID).Info(constant.LLvlService, fields...)
}
