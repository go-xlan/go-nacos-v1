package nacosv1

import (
	"os"
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func MustNewNacosClientParam(config *Config, zapLog *zaplog.Zap, nacosOptions []constant.ClientOption) vo.NacosClientParam {
	endpoint := config.Endpoint
	if endpoint == "" {
		endpoint = must.Nice(os.Getenv("NACOS_ADDR"))
	}
	zapLog.LOG.Debug("nacos", zap.String("endpoint", endpoint))

	endpoint2s := strings.Split(endpoint, ":")
	nacosIp := endpoint2s[0]
	nacosPortNum := rese.C1(strconv.Atoi(endpoint2s[1]))

	var opts = []constant.ClientOption{
		constant.WithEndpoint(endpoint),
		withAppName(config.AppName),
		constant.WithNamespaceId(config.Namespace),
	}

	opts = append(opts, nacosOptions...)

	clientConfig := constant.NewClientConfig(opts...)

	serverConfig := constant.NewServerConfig(nacosIp, uint64(nacosPortNum))

	param := vo.NacosClientParam{
		ClientConfig:  clientConfig,
		ServerConfigs: []constant.ServerConfig{*serverConfig},
	}
	return param
}

func withAppName(appName string) constant.ClientOption {
	return func(config *constant.ClientConfig) {
		config.AppName = appName
	}
}
