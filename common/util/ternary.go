package util

import (
	"database/sql"
	"time"
)

/**
 * Perfect place for generics use, but GO does not have them so c&p of methods with different types must be done
 * 
 * In functions below curring is used.
 */

func IfNullString(cond bool) func(sql.NullString, sql.NullString) sql.NullString {
	return func(a sql.NullString, b sql.NullString) sql.NullString {
		if cond {
			return a
		} else {
			return b
		}
	}
}


func IfNullTime(cond bool) func(sql.NullTime, sql.NullTime) sql.NullTime {
	return func(a sql.NullTime, b sql.NullTime) sql.NullTime {
		if cond {
			return a
		} else {
			return b
		}
	}
	
}

/************/

func IfString(cond bool) func(string, string) string {
	return func(a string, b string) string {
		if cond {
			return a
		} else {
			return b
		}
	}
	
}

func IfTime(cond bool) func(time.Time, time.Time) time.Time {
	return func(a time.Time, b time.Time) time.Time {
		if cond {
			return a
		} else {
			return b
		}
	}
	
}