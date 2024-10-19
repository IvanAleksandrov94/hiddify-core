package hcore

import (
	"fmt"
	"time"

	common "github.com/hiddify/hiddify-core/v2/common"
	"github.com/sagernet/sing/common/observable"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewObserver[T any](listenerBufferSize int) *observable.Observer[T] {
	return observable.NewObserver(observable.NewSubscriber[T](listenerBufferSize), listenerBufferSize)
}

var logObserver = NewObserver[*LogMessage](1)

func Log(level LogLevel, typ LogType, message ...any) {
	if level != LogLevel_DEBUG {
		fmt.Printf("%s %s %s\n", level, typ, message)
	}
	logObserver.Emit(&LogMessage{
		Level:   level,
		Type:    typ,
		Time:    timestamppb.New(time.Now()),
		Message: fmt.Sprint(message...),
	})
}

func (s *CoreService) LogListener(req *common.Empty, stream grpc.ServerStreamingServer[LogMessage]) error {
	logSub, stopch, _ := logObserver.Subscribe()
	defer logObserver.UnSubscribe(logSub)

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-stopch:
			return nil
		case info := <-logSub:
			stream.Send(info)
			// case <-time.After(500 * time.Millisecond):
		}
	}
}
