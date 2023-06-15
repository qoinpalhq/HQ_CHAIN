package main

import (
	"fmt"
	"sync"

	bl "github.com/qoinpalhq/HQ_CHAIN/blockchain"
	// "context"
	// "time"

	ev "github.com/qoinpalhq/HQ_CHAIN/events"
)

func main() {

		var wg sync.WaitGroup
		var mp *bl.Mempool
		e := ev.NewEventStream()
	
		wg.Add(1)
		go func() {
			defer wg.Done()
			chain := bl.InitializeChain(e)
			mp = bl.GetSingleMempoolInstance(e)
			fmt.Println(chain, mp)
		}()
	
		wg.Wait() // Wait for chain initialization to complete
	
		trx1 := bl.CreateTransaction([]byte("Jason"), []byte("Qoinpal"), 100)
		trx2 := bl.CreateTransaction([]byte("Kendrick"), []byte("Dayo"), 200)
		trx3 := bl.CreateTransaction([]byte("Kendrick"), []byte("Dayo"), 200)
	
		bl.AddTransactionToMempool(trx1, mp)
		bl.AddTransactionToMempool(trx2, mp)
		bl.AddTransactionToMempool(trx3, mp)
	
		// Continue with the rest of your code...
	

	
}
