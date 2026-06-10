type Customer struct {
	CustomerID string

	Country        string
	Age            int
	Gender         string
	MembershipTier string

	TotalOrders      int
	TotalSpendUSD    float64
	AvgOrderValueUSD float64

	DaysSinceLastPurchase int

	PreferredCategory      string
	PreferredDevice        string
	PreferredPaymentMethod string
	AcquisitionChannel     string

	ReviewsGiven   int
	AvgReviewScore float64

	ReturnsMade   int
	WishlistItems int

	NewsletterSubscribed bool

	Churned int
}