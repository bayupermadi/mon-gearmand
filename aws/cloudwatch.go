package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/spf13/viper"
)

func CW(vmetric float64, logpath string) {

	awsKeyID := viper.Get("app.smtp.ses.aws-key-id").(string)
	awsSecretKey := viper.Get("app.smtp.ses.aws-secret-key").(string)
	awsRegion := viper.Get("app.smtp.ses.aws-region").(string)

	setConfiguration(awsKeyID, awsSecretKey, awsRegion)

	sess := startNewSession()

	// Create new cloudwatch client.
	svc := cloudwatch.New(sess)

	_, err := svc.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Germand"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("LogSize"),
				Unit:       aws.String("Megabytes"),
				Value:      aws.Float64(vmetric),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("Files"),
						Value: aws.String(logpath),
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println("Error adding metrics:", err.Error())
		return
	}

}
