package csv

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/obrunoapolinario/kNN/internal/domain"
)

func ReadCustomersFromCSV(filePath string) ([]domain.Customer, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var customers []domain.Customer

	_, err = reader.Read() // Skip the header row
	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		customerId, _ := strconv.Atoi(record[0])
		country := record[1]
		age, _ := strconv.Atoi(record[2])
		gender := record[3]
		membershipTier := record[4]
		totalOrders, _ := strconv.Atoi(record[5])
		totalSpendUSD, _ := strconv.ParseFloat(record[6], 64)
		avgOrderValueUSD, _ := strconv.ParseFloat(record[7], 64)
		daysSinceLastPurchase, _ := strconv.Atoi(record[8])
		preferredCategory := record[9]
		preferredDevice := record[10]
		preferredPaymentMethod := record[11]
		acquisitionChannel := record[12]
		reviewsGiven, _ := strconv.Atoi(record[13])
		avgReviewScore, _ := strconv.ParseFloat(record[14], 64)
		returnsMade, _ := strconv.Atoi(record[15])
		wishlistItems, _ := strconv.Atoi(record[16])
		newsletterSubscribed, _ := strconv.ParseBool(record[17])
		churned, _ := strconv.Atoi(record[18])

		customer := domain.Customer{
			CustomerID:             strconv.Itoa(customerId),
			Country:                country,
			Age:                    age,
			Gender:                 gender,
			MembershipTier:         membershipTier,
			TotalOrders:            totalOrders,
			TotalSpendUSD:          totalSpendUSD,
			AvgOrderValueUSD:       avgOrderValueUSD,
			DaysSinceLastPurchase:  daysSinceLastPurchase,
			PreferredCategory:      preferredCategory,
			PreferredDevice:        preferredDevice,
			PreferredPaymentMethod: preferredPaymentMethod,
			AcquisitionChannel:     acquisitionChannel,
			ReviewsGiven:           reviewsGiven,
			AvgReviewScore:         avgReviewScore,
			ReturnsMade:            returnsMade,
			WishlistItems:          wishlistItems,
			NewsletterSubscribed:   newsletterSubscribed,
			Churned:                churned,
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
