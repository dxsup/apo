package serviceoverview

import (
	"time"

	"github.com/CloudDetail/apo/backend/pkg/repository/clickhouse"

	"github.com/CloudDetail/apo/backend/pkg/model/response"
	"github.com/CloudDetail/apo/backend/pkg/repository/database"
	"github.com/CloudDetail/apo/backend/pkg/repository/prometheus"
)

var _ Service = (*service)(nil)

type Service interface {
	GetServiceMoreUrl(startTime time.Time, endTime time.Time, step time.Duration, serviceNames string, sortRule SortType) (res []response.ServiceDetail, err error)
	GetInstances(startTime time.Time, endTime time.Time, step time.Duration, serviceName string, endPoint string) (res response.InstancesRes, err error)
	GetThreshold(level string, serviceName string, endPoint string) (res response.GetThresholdResponse, err error)
	SetThreshold(level string, serviceName string, endPoint string, latency float64, errorRate float64, tps float64, log float64) (res response.SetThresholdResponse, err error)
	GetServicesAlert(startTime time.Time, endTime time.Time, step time.Duration, serviceNames []string, returnData []string) (res []response.ServiceAlertRes, err error)
	GetServicesEndPointData(startTime time.Time, endTime time.Time, step time.Duration, serviceNames string, sortRule SortType) (res []response.ServiceEndPointsRes, err error)
}
type service struct {
	dbRepo   database.Repo
	promRepo prometheus.Repo
	chRepo   clickhouse.Repo
}

func New(chRepo clickhouse.Repo, dbRepo database.Repo, promRepo prometheus.Repo) Service {
	return &service{
		dbRepo:   dbRepo,
		promRepo: promRepo,
		chRepo:   chRepo,
	}
}
