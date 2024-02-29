package DB

import (
	"context"
)

////////////////////////////////////////////////
//Jobs

const (
	all       = "SELECT * FROM jobs WHERE start_time > CURRENT_DATE"
	pending   = "SELECT * FROM jobs WHERE start_time > CURRENT_DATE AND finalized = false"
	finalized = "SELECT * FROM jobs WHERE start_time > CURRENT_DATE AND finalized = true"
)

// TODO: Figure out error handling for address errors
func (pg *postgres) GetJobsByStatus(ctx context.Context, status string) ([]Job, error) {
	var jobs []Job
	var query string
	switch status {
	case "all":
		query = all
	case "pending":
		query = pending
	case "finalized":
		query = finalized
	}

	rows, err := pg.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		LoadAddrID   int
		UnloadAddrID int
	)
	for rows.Next() {
		var j Job
		if err := rows.Scan(
			&j.ID,
			&j.CustomerID,
			&LoadAddrID,
			&UnloadAddrID,
			&j.StartTime,
			&j.HoursLabor,
			&j.Finalized,
			&j.Rooms,
			&j.Pack,
			&j.Unpack,
			&j.Load,
			&j.Unload,
			&j.Clean,
			&j.Milage,
			&j.Cost,
		); err != nil {
			return nil, err
		}
		//need to figure out error handling here
		j.LoadAddr, _ = getAddr(ctx, LoadAddrID)
		j.UnloadAddr, _ = getAddr(ctx, UnloadAddrID)
		jobs = append(jobs, j)
	}
	return jobs, nil
}

const addrQuery = "SELECT * FROM addresses WHERE address_id = $1"

func getAddr(ctx context.Context, addrID int) (Address, error) {
	var a Address
	row := PgInstance.db.QueryRow(ctx, addrQuery, addrID)
	err := row.Scan(
		&a.AddressID,
		&a.Street,
		&a.City,
		&a.State,
		&a.Zip,
		&a.ResType,
		&a.Flights,
		&a.AptNum,
	)
	if err != nil {
		return a, err
	}
	return a, nil
}