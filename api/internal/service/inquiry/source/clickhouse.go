package source

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gotomicro/ego/core/elog"

	"github.com/clickvisual/clickvisual/api/pkg/model/view"
)

type ClickHouse struct {
	s *Source
}

func NewClickHouse(s *Source) *ClickHouse {
	return &ClickHouse{s}
}

func (c *ClickHouse) Databases() (res []string, err error) {
	return c.queryStringArr("SHOW DATABASES")
}

func (c *ClickHouse) Tables(database string) (res []string, err error) {
	return c.queryStringArr(fmt.Sprintf("SHOW TABLES FROM %s", database))
}

func (c *ClickHouse) Columns(database, table string) (res []view.Column, err error) {
	conn, err := sql.Open("clickhouse", c.s.GetDSN())
	if err != nil {
		elog.Error("ClickHouse", elog.Any("step", "sql.error"), elog.String("error", err.Error()))
		return
	}
	conn.SetConnMaxIdleTime(time.Minute * 3)
	defer func() { _ = conn.Close() }()
	query := fmt.Sprintf("select name, type from system.columns where database = '%s' and table = '%s'", database, table)
	list, err := c.doQuery(conn, query)
	if err != nil {
		return
	}
	for _, row := range list {
		res = append(res, view.Column{
			Field: row["name"].(string),
			Type:  row["type"].(string),
		})
	}
	return
}

func (c *ClickHouse) Exec(s string) (err error) {
	obj, err := sql.Open("clickhouse", c.s.GetDSN())
	if err != nil {
		elog.Error("ClickHouse", elog.Any("step", "open"), elog.String("error", err.Error()))
		return
	}
	defer func() { _ = obj.Close() }()
	_, err = obj.Exec(s)
	return
}

func (c *ClickHouse) Query(s string) (res []map[string]interface{}, err error) {
	elog.Info("ClickHouse", elog.FieldComponent("Query"), elog.String("s", s))
	return
}

func (c *ClickHouse) queryStringArr(sq string) (res []string, err error) {
	obj, err := sql.Open("clickhouse", c.s.GetDSN())
	if err != nil {
		elog.Error("ClickHouse", elog.Any("step", "open"), elog.String("error", err.Error()))
		return
	}
	defer func() { _ = obj.Close() }()
	// query databases
	rows, err := obj.Query(sq)
	if err != nil {
		elog.Error("ClickHouse", elog.Any("step", "query"), elog.String("error", err.Error()))
		return
	}
	for rows.Next() {
		var tmp string
		errScan := rows.Scan(&tmp)
		if errScan != nil {
			elog.Error("source", elog.String("err", errScan.Error()))
			continue
		}
		res = append(res, tmp)
	}
	return
}

func (c *ClickHouse) doQuery(ins *sql.DB, sql string) (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, 0)
	rows, err := ins.Query(sql)
	if err != nil {
		return
	}
	defer func() { _ = rows.Close() }()
	cts, _ := rows.ColumnTypes()
	var (
		fields = make([]string, len(cts))
		values = make([]interface{}, len(cts))
	)
	for idx, field := range cts {
		fields[idx] = field.Name()
	}
	for rows.Next() {
		line := make(map[string]interface{}, 0)
		for idx := range values {
			fieldValue := reflect.ValueOf(&values[idx]).Elem()
			values[idx] = fieldValue.Addr().Interface()
		}
		if err = rows.Scan(values...); err != nil {
			elog.Error("ClickHouse", elog.Any("step", "doQueryNext"), elog.Any("error", err.Error()))
			return
		}
		elog.Debug("ClickHouse", elog.Any("fields", fields), elog.Any("values", values))
		for k := range fields {
			elog.Debug("ClickHouse", elog.Any("fields", fields[k]), elog.Any("values", values[k]))
			line[fields[k]] = values[k]
		}
		res = append(res, line)
	}
	if err = rows.Err(); err != nil {
		elog.Error("ClickHouse", elog.Any("step", "doQuery"), elog.Any("error", err.Error()))
		return
	}
	return
}

func (c *ClickHouse) ClusterInfo() (isCluster, isShard, isReplica int, clusters []string, err error) {
	cs, err := c.clusters()
	clusterMap := make(map[string]interface{})
	clusters = make([]string, 0)
	if err != nil {
		return
	}
	for _, clu := range cs {
		clusterMap[clu.Cluster] = struct{}{}
		if clu.ReplicaNum > 1 {
			isReplica = 1
			isCluster = 1
		}
		if clu.ShardNum > 1 {
			isShard = 1
			isCluster = 1
		}
	}
	for clu := range clusterMap {
		clusters = append(clusters, clu)
	}
	return
}

func (c *ClickHouse) clusters() (res []view.Cluster, err error) {
	obj, err := sql.Open("clickhouse", c.s.GetDSN())
	if err != nil {
		elog.Error("ClickHouse", elog.Any("step", "open"), elog.FieldErr(err))
		return
	}
	defer func() { _ = obj.Close() }()
	// query databases
	rows, err := obj.Query("SELECT cluster,shard_num,shard_weight,replica_num,host_name,host_address,port from `system`.`clusters`")
	if err != nil {
		elog.Error("ClickHouse", elog.Any("step", "query"), elog.FieldErr(err))
		return
	}
	for rows.Next() {
		var cluster string
		var shard_num int
		var shard_weight int
		var replica_num int
		var host_name string
		var host_address string
		var port int
		// var is_local int
		// var user string
		// var default_database string
		// var errors_count int
		// var estimated_recovery_time int
		errScan := rows.Scan(&cluster, &shard_num, &shard_weight, &replica_num, &host_name, &host_address, &port)
		if errScan != nil {
			elog.Error("source", elog.FieldErr(err))
			continue
		}
		if strings.HasPrefix(cluster, "test_") {
			continue
		}
		res = append(res, view.Cluster{
			Cluster:     cluster,
			ShardNum:    shard_num,
			ReplicaNum:  replica_num,
			HostName:    host_name,
			HostAddress: host_address,
			Port:        port,
		})
	}
	return
}
