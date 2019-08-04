package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func ListBuckets() {
	// https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/lambda/aws-go-sdk-lambda-example-configure-function-for-notification.go
  // create session
  sess := session.Must(session.NewSession())
	
  // create session from config object 
  sess2 := session.Must(session.NewSession(&aws.Config{
		  Region: aws.String("us-east-1"),
  }))

  // more complex config object

  svc := s3.New(sess)

  result, err := svc.ListBuckets(nil)

  if err != nil {
	  // exitErrorf("Unable to list buckets, %v", err)
	  panic( err)
  }

  // fmt.Println("Buckets:")

  for _, b := range result.Buckets {
	  buf.WriteString(
		  aws.StringValue(b.Name), 
	  )
	  // fmt.Printf("* %s created on %s\n",
	  // 	aws.StringValue(b.Name), 
	  // 	aws.TimeValue(b.CreationDate),
	  // )
  }
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest, ctx context.Context) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)
	Ping(request, ctx)

// func Ping(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {

	// resp := Response{
	// 	StatusCode:      200,
	// 	IsBase64Encoded: false,
	// 	Body:            buf.String(),
	// 	Headers: map[string]string{
	// 		"Content-Type":           "application/json",
	// 		"X-MyCompany-Func-Reply": "hello-handler",
	// 	},
	// }

	// return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
	resp := events.APIGatewayProxyResponse{
		Body: bug.String(), 
		StatusCode: 200
	}


	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
