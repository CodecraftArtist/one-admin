package elastic

import (
	"context"
	"github.com/olivere/elastic/v7"
	esConfig "github.com/olivere/elastic/v7/config"
	"starter/pkg/app"
)

type config struct {
	URL         string `toml:"url"`
	Index       string `toml:"index"`
	Username    string `toml:"username"`
	Password    string `toml:"password"`
	Shards      int    `toml:"shards"`
	Replicas    int    `toml:"replicas"`
	Sniff       bool   `toml:"sniff"`
	HealthCheck bool   `toml:"health"`
	InfoLog     string `toml:"info_log"`
	ErrorLog    string `toml:"error_log"`
	TraceLog    string `toml:"trace_log"`
}

var (
	// ES elastic search 的连接资源
	ES   *elastic.Client
	conf config
	err  error
)

func (config config) ElasticSearchConfig() *esConfig.Config {
	err = app.Config().Bind("application", "elasticsearch", &conf)
	if err == app.ErrNodeNotExists {
		return nil
	}
	return &esConfig.Config{
		URL:         conf.URL,
		Index:       conf.Index,
		Username:    conf.Username,
		Password:    conf.Password,
		Shards:      conf.Shards,
		Replicas:    conf.Replicas,
		Sniff:       &conf.Sniff,
		Healthcheck: &conf.HealthCheck,
		Infolog:     conf.InfoLog,
		Errorlog:    conf.ErrorLog,
		Tracelog:    conf.TraceLog,
	}
}

// Start 连接到 es
func Start() {
	option := conf.ElasticSearchConfig()
	if option == nil {
		return
	}
	ES, err = elastic.NewClientFromConfig(option)
	if err != nil {
		app.Logger().WithField("log_type", "pkg.elastic.es").Error(err)
	}
}

// Insert 向es写入数据
func Insert(index string, body interface{}) *elastic.IndexResponse {
	rs, err := ES.Index().Index(index).BodyJson(body).Do(context.Background())
	if err != nil {
		app.Logger().WithField("log_type", "pkg.elastic.es").Error("es write error: ", err)
	}

	return rs
}

// InsertString 写入字符串
func InsertString(index string, body string) *elastic.IndexResponse {
	rs, err := ES.Index().Index(index).BodyString(body).Do(context.Background())
	if err != nil {
		app.Logger().WithField("log_type", "pkg.elastic.es").Error("es write error: ", err)
	}
	return rs
}
