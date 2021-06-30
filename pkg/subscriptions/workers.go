package subscriptions

import (
	"log"

	"github.com/robfig/cron"
)

const (
	everydayWokersCron = "0 6 * * *"
)

// WorkerPool ...
type WorkerPool interface {
	Start() error
	Stop()
}

// workerPool ...
type workerPool struct {
	workersIDS        []cron.EntryID
	SubscriptionStore SubscriptionStore
	cron              *cron.Cron
}

// NewWorkerPool ...
func NewWorkerPool(ss SubscriptionStore) WorkerPool {
	return &workerPool{
		workersIDS:        make([]cron.EntryID, 0),
		SubscriptionStore: ss,
	}
}

// Start ...
func (wp *workerPool) Start() error {
	log.Println("[Workers][Subscriptions][Start]")

	wp.cron = cron.New()

	// work: check with cron if any subscription is going to expire
	id, err := wp.cron.AddFunc(everydayWokersCron, wp.checkSubscriptionsToExpire)
	if err != nil {
		return err
	}
	wp.workersIDS = append(wp.workersIDS, id)

	// work: check with cron if any subscription has to expired
	id, err = wp.cron.AddFunc(everydayWokersCron, wp.checkExpiredSubscriptions)
	if err != nil {
		return err
	}
	wp.workersIDS = append(wp.workersIDS, id)

	wp.cron.Start()

	return nil
}

// Stop ...
func (wp *workerPool) Stop() {
	log.Println("[Workers][Subscriptions][Stop]")

	wp.cron.Stop()

	for _, v := range wp.workersIDS {
		wp.cron.Remove(v)
	}
}

func (wp *workerPool) checkSubscriptionsToExpire() {
	log.Println("[Workers][Subscriptions][checkSubscriptionsToExpire][StartCron]")

	transactions, err := wp.SubscriptionStore.ListTransactionsToExpire()
	if err != nil {
		log.Println("[Workers][Subscriptions][checkSubscriptionsToExpire][Error]", err)
		return
	}

	// create new subcription transaction for each user
	for i := 0; i < len(transactions); i++ {
		lastTransaction := transactions[i]

		_, err := wp.SubscriptionStore.Subscribe(QueryTransaction{
			UserID:         lastTransaction.UserID,
			SubscriptionID: lastTransaction.SubscriptionID,
			ProviderID:     lastTransaction.ProviderID,
		})
		if err != nil {
			log.Println("[Workers][Subscriptions][checkSubscriptionsToExpire][Error] index=", i, err)
			return
		}

		// TODO(ca): send notification by user_id
	}

	log.Println("[Workers][Subscriptions][checkSubscriptionsToExpire][EndCron]")
}

func (wp *workerPool) checkExpiredSubscriptions() {
	log.Println("[Workers][Subscriptions][checkExpiredSubscriptions][StartCron]")

	transactions, err := wp.SubscriptionStore.ListExpiredTransactions()
	if err != nil {
		log.Println("[Workers][Subscriptions][checkExpiredSubscriptions][Error]", err)
		return
	}

	subscription, err := wp.SubscriptionStore.GetSubscription(QuerySubscription{
		Price: float64(0),
	})
	if err != nil {
		log.Println("[Workers][Subscriptions][checkExpiredSubscriptions][Error]", err)
		return
	}

	// create new subcription transaction for each user
	for i := 0; i < len(transactions); i++ {
		lastTransaction := transactions[i]

		_, err := wp.SubscriptionStore.Unsubscribe(QueryTransaction{
			UserID:         lastTransaction.UserID,
			SubscriptionID: lastTransaction.SubscriptionID,
		})
		if err != nil {
			log.Println("[Workers][Subscriptions][checkExpiredSubscriptions][Error] index=", i, err)
			return
		}

		_, err = wp.SubscriptionStore.Subscribe(QueryTransaction{
			UserID:         lastTransaction.UserID,
			SubscriptionID: subscription.ID,
			ProviderID:     ProviderFree,
		})
		if err != nil {
			log.Println("[Workers][Subscriptions][checkExpiredSubscriptions][Error] index=", i, err)
			return
		}

		// TODO(ca): send notification by user_id
	}

	log.Println("[Workers][Subscriptions][checkExpiredSubscriptions][EndCron]")
}
