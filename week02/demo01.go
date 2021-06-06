package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func daosql() error {
	return errors.Wrap(sql.ErrNoRows, "dao failed")
}

func servicesql() error {
	return errors.WithMessage(daosql(), "service failed")
}

func main() {
	err := servicesql()
	if errors.Cause(err) == sql.ErrNoRows {
		// fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		fmt.Printf("got err, %+v\n", err)
	}
}
