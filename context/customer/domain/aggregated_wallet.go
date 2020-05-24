package domain

import (
	"sort"
	"github.com/google/go-cmp/cmp"
)

type AggregatedWallet struct {
	CustomerId uint
	Currencies []string
	Total float64
}

/**
 * Equals methof to comapre AggrefatedWallet objects  
 *
 * Temporary arrays has to be created for comapring slices (Slices must to sroted, without changing the base object)
 */
func (aw AggregatedWallet) Equal(other AggregatedWallet) bool {
	otherSortedCurrencies := make([]string, len(other.Currencies))
	otherSortedCurrencies = append(other.Currencies[:0:0], other.Currencies...)
	sort.Sort(sort.StringSlice(otherSortedCurrencies))
	thisSortedCurrencies := make([]string, len(aw.Currencies))
	thisSortedCurrencies = append(aw.Currencies[:0:0], aw.Currencies...)
	sort.Sort(sort.StringSlice(thisSortedCurrencies))
	return aw.CustomerId == other.CustomerId && aw.Total == other.Total && cmp.Equal(thisSortedCurrencies, otherSortedCurrencies)
}

