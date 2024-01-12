package services

import "github.com/simpleforce/simpleforce"

func queryGetAll(query string, client *simpleforce.Client, result chan<- any) error {
	for {
		queryResults, err := client.Query(query)
		if err != nil {
			return err
		}

		for _, record := range queryResults.Records {
			result <- record
		}

		if queryResults.Done {
			break
		}

		query = queryResults.NextRecordsURL

	}
	return nil
}
