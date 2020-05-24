package domain_test

import(
	"bettertomorrow/context/customer/domain"
	"testing"
)

func TestAggregatedWallet(t *testing.T) {
	
	aggregatedWallet1 := domain.AggregatedWallet{1, []string{"usd", "pln", "nzd"}, 200 }
	aggregatedWallet2 := domain.AggregatedWallet{1, []string{"pln", "nzd", "usd"}, 200 }
	
	if aggregatedWallet1.Equal(aggregatedWallet2) {
		t.Logf("equals ok, ag1 : %v, ag2 : %v \n", aggregatedWallet1, aggregatedWallet2)
	} else {
		t.Errorf("equals failed, ag1 : %v, ag2 : %v \n", aggregatedWallet1, aggregatedWallet2)
	}

}