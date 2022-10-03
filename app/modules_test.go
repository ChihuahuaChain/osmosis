package app

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	"gotest.tools/assert"
	"pgregory.net/rapid"
)

func TestOrderEndBlockers_Determinism(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewOsmosisApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, DefaultNodeHome, 5, simapp.EmptyAppOptions{}, GetWasmEnabledProposals(), EmptyWasmOpts)

	rapid.Check(t, func(t *rapid.T) {
		a := OrderEndBlockers(app.mm.ModuleNames())
		b := OrderEndBlockers(app.mm.ModuleNames())

		fmt.Println("=================")
		fmt.Println("A EndBlockers:", a)
		fmt.Println("B EndBlockers:", b)
		fmt.Println("EQUAL:", reflect.DeepEqual(a, b))
		fmt.Println("=================")

		assert.DeepEqual(t, a, b)
	})
}