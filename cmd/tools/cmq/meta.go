package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

type opts struct {
	Hosts       string
	Port        int
	Keyspace    string
	Consistency string
	Username    string
	Password    string
	Timeout     time.Duration
	Retries     int
	PageSize    int
}

type metadataClient struct {
	session     *gocql.Session
	consistency gocql.Consistency
	retries     int
	pageSize    int
}

const (
	protocolVersion = 4
	numConns        = 1
)

func newMetadataClient(opts *opts) (*metadataClient, error) {

	cluster := gocql.NewCluster(opts.Hosts) // TODO: add support for multiple hosts

	cluster.Port = opts.Port
	cluster.Keyspace = opts.Keyspace

	if len(opts.Username) > 0 || len(opts.Password) > 0 {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: opts.Username,
			Password: opts.Password,
		}
	}

	cluster.Consistency = gocql.ParseConsistency(opts.Consistency)

	cluster.NumConns = numConns
	cluster.ProtoVersion = protocolVersion
	cluster.Timeout = opts.Timeout
	cluster.PageSize = opts.PageSize

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("CreateSession: %v", err)
	}

	return &metadataClient{
		session:     session,
		consistency: cluster.Consistency,
		retries:     opts.Retries,
		pageSize:    opts.PageSize,
	}, nil
}

func (m *metadataClient) QueryRow(cql string) (row map[string]interface{}, err error) {

	row = make(map[string]interface{})
	err = m.session.Query(cql).MapScan(row)
	return
}

func (m *metadataClient) Close() {
	m.session.Close()
}

type pagedScanIter struct {
	query *gocql.Query
	iter  *gocql.Iter
}

func (m *metadataClient) pagedQueryIter(cql string) *pagedScanIter {

	query := m.session.Query(cql).Consistency(m.consistency).PageSize(m.pageSize).RetryPolicy(&gocql.SimpleRetryPolicy{NumRetries: m.retries})
	iter := query.Iter()

	if iter == nil {
		return nil
	}

	return &pagedScanIter{query: query, iter: iter}
}

func (it *pagedScanIter) scan(dest ...interface{}) bool {

	if !it.iter.Scan(dest...) {

		pageState := it.iter.PageState()

		if len(pageState) == 0 {
			return false
		}

		it.iter.Close()

		it.iter = it.query.PageState(pageState).Iter()
	}

	return true
}

func (it *pagedScanIter) close() error {

	err := it.iter.Close()

	if err != nil {
		fmt.Printf("pagedScanIter close error: %v\n", err)
	}

	return err
}

type filter struct {
	dest  string
	cg    string
	ext   string
	where map[string]string
}

func (m *metadataClient) Query(table string, f filter) (row map[string]interface{}, err error) {

	/*
		Query(
			"consumer_group_extents",
			select{"destination.path", "destination.uuid"},
			where{cg: cgUUID, ext: extUUID},
			filter{"destination.status": 5},
		)
	*/

	/*
		row = make(map[string]interface{})

		var where []string

		// err = m.session.Query(cql).MapScan(row)
	*/

	return
}
