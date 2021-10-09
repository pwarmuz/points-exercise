package points

import (
	"testing"
	"time"
)

// This is a convoluted illustration of a test suite pattern for the statements struct
type setup struct {
	Payer  string
	Points int
	Date   string
}

// tcSetup creates the default layout that is expected for these tests
func tcSetup() []setup {
	return []setup{
		{Payer: "Payer01", Points: 100, Date: "1/3 10AM"},
		{"Payer02", 400, "1/7 7PM"},
		{"Payer01", 600, "1/9 7PM"},
		{"Payer01", 900, "1/9 10PM"},
	}
}

// TestStatementsCredit is a test suite to run
// creditPoints which verifies the points
// creditPriority which verifies the credit priority
// creditListBalance which verifies the balance
func TestStatementsCredit(t *testing.T) {
	stmt := NewStatement()

	testCase := struct {
		Setup []setup
		Want  map[string]Credit
	}{
		Setup: tcSetup(),
		Want: map[string]Credit{
			"Payer01": {Points: 1600, Priority: 0},
			"Payer02": {Points: 400, Priority: 1},
		},
	}

	tests := []struct {
		title string
		run   func(t *testing.T, stmt *Statement, tc map[string]Credit)
	}{
		{"verifyCreditPoints", creditPoints},
		{"verifyCreditPriority", creditPriority},
		{"verifyCreditListBalance", creditListBalance},
	}
	// statement's Update method is used to verify that it's results are accurate
	for _, v := range testCase.Setup {
		stmt.Update(v.Payer, v.Points, v.Date)
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			test.run(t, stmt, testCase.Want)
		})
	}
}

func creditPoints(t *testing.T, stmt *Statement, tc map[string]Credit) {
	for payer, credit := range stmt.Balance {
		if credit.Points != tc[payer].Points {
			t.Error("failed equivalent points")
		}
	}
}
func creditPriority(t *testing.T, stmt *Statement, tc map[string]Credit) {
	for payer, credit := range stmt.Balance {
		if credit.Priority != tc[payer].Priority {
			t.Error("failed equivalent Priority")
		}
	}
}
func creditListBalance(t *testing.T, stmt *Statement, tc map[string]Credit) {
	list := stmt.ListBalance()
	for i, transaction := range list {
		t.Log(transaction.Points)
		if transaction.Points != tc[transaction.Payer].Points {
			t.Error("failed equivalent Transaction Points")
		}
		if i != tc[list[i].Payer].Priority {
			t.Error("failed equivalent Payer to Priority verification")
		}
	}
}

// TestsStatementsTransaction is a test suite to run
// transactionPayer which verifies the consistency of payers
// transactionPoints which verifies the points are consistent and accurate
func TestStatementsTransactions(t *testing.T) {
	stmt := NewStatement()

	testCase := struct {
		Setup []setup
		Want  map[time.Time]Transaction
	}{
		Setup: tcSetup(),
		Want: map[time.Time]Transaction{
			ParseLayout("1/3 10AM"): {Payer: "Payer01", Points: 100},
			ParseLayout("1/7 7PM"):  {Payer: "Payer02", Points: 400},
			ParseLayout("1/9 7PM"):  {Payer: "Payer01", Points: 600},
			ParseLayout("1/9 10PM"): {Payer: "Payer01", Points: 900},
		},
	}

	tests := []struct {
		title string
		run   func(t *testing.T, stmt *Statement, tc map[time.Time]Transaction)
	}{
		{"verifyTransactionPayer", transactionPayer},
		{"verifyTransactionPoints", transactionPoints},
	}
	// statement's Update method is used to verify that it's results are accurate
	for _, v := range testCase.Setup {
		stmt.Update(v.Payer, v.Points, v.Date)
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			test.run(t, stmt, testCase.Want)
		})
	}
}

func transactionPayer(t *testing.T, stmt *Statement, tc map[time.Time]Transaction) {
	for ts, transaction := range stmt.Transactions {
		if transaction.Payer != tc[ts].Payer {
			t.Error("failed equivalent points")
		}
	}
}
func transactionPoints(t *testing.T, stmt *Statement, tc map[time.Time]Transaction) {
	for ts, transaction := range stmt.Transactions {
		if transaction.Points != tc[ts].Points {
			t.Error("failed equivalent Priority")
		}
	}
}
