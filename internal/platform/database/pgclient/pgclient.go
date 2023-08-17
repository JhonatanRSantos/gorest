package pgclient

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	connTimeout           = time.Second * 30
	connMaxLifeTime       = time.Second * 30
	connMaxLifeTimeJitter = time.Second * 2
	connMaxIdleTime       = time.Second * 10
	maxConn               = 10
	minConn               = 2
	healthCheckPeriod     = time.Second * 5
)

type PGClient struct {
	connPool *pgxpool.Pool
}

type PGClientConfig struct {
	Host   string
	Port   uint16
	User   string
	Pass   string
	DBName string
}

func NewClient(ctx context.Context, config PGClientConfig) (*PGClient, error) {
	pgPoolConfig, _ := pgxpool.ParseConfig("")

	pgPoolConfig.ConnConfig.Host = config.Host
	pgPoolConfig.ConnConfig.Port = config.Port
	pgPoolConfig.ConnConfig.User = config.User
	pgPoolConfig.ConnConfig.Password = config.Pass
	pgPoolConfig.ConnConfig.Database = config.DBName
	pgPoolConfig.ConnConfig.ConnectTimeout = connTimeout
	pgPoolConfig.ConnConfig.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	// pgPoolConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
	// 	pgxUUID.Register(conn.TypeMap())
	// 	return nil
	// }
	pgPoolConfig.MaxConnLifetime = connMaxLifeTime
	pgPoolConfig.MaxConnLifetimeJitter = connMaxLifeTimeJitter
	pgPoolConfig.MaxConnIdleTime = connMaxIdleTime
	pgPoolConfig.MaxConns = maxConn
	pgPoolConfig.MinConns = minConn
	pgPoolConfig.HealthCheckPeriod = healthCheckPeriod
	connPool, err := pgxpool.NewWithConfig(ctx, pgPoolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new connection pool. %s", err)
	}

	if err = connPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database. %s", err)
	}

	return &PGClient{
		connPool: connPool,
	}, nil
}

func (pgc *PGClient) Close() {
	pgc.connPool.Close()
}

func (pgc *PGClient) GetConnection(ctx context.Context) (*pgxpool.Conn, error) {
	return pgc.connPool.Acquire(ctx)
}
