package logging

import (
	"fmt"
	"go.uber.org/zap"
	"helloWorld/pkg/config/config"
	"net/url"
	"os"
)

func NewLog(c config.LogConfig) *zap.Logger {
	//log := initKafkaLogs(c.KafkaLogConfig)
	log := initFileLogs(c.FileLogConfig)
	defer func(log *zap.Logger) {
		if err := log.Sync(); err != nil {
			_ = fmt.Errorf("log defer error %v", err)
		}
	}(log)
	return log
}

func initFileLogs(c config.FileLogConfig) *zap.Logger {
	os.OpenFile(c.LogPath, os.O_RDONLY|os.O_CREATE, 0666)
	conf := zap.NewProductionConfig()
	conf.OutputPaths = []string{"stdout", c.LogPath}

	var err error
	logger, err := conf.Build()
	if err != nil {
		fmt.Printf("Log configuration failed! %e\n", err)
	}
	return logger
}

func initKafkaLogs(cc config.KafkaLogConfig) *zap.Logger {
	kafkaUrl := url.URL{Scheme: "kafka", Host: cc.KafkaCluster, RawQuery: "topic=" + cc.Topic}
	fmt.Printf("Logging configuration started with Kafka Sink: %v\n", kafkaUrl.String())
	if err := zap.RegisterSink("kafka", InitKafkaSink); err != nil {
		fmt.Errorf("Log Kafka registeration failed; %v \n", err)
		panic(err)
	}
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", kafkaUrl.String()}
	if logger, err := c.Build(); err != nil {
		fmt.Errorf("Log build failed; %v \n", err)
		panic(err)
	} else {
		return logger
	}
}
