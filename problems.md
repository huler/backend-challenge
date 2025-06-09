# Problems

**Problem 1**

_What is the problem?_

The survey submissions rely on email IDs as the main identifier. This means that if someone submits multiple times,
their latest entry will replace any previous ones, which can lead to data loss.

_Why is it a problem?_

This violates data integrity and prevents users from submitting multiple surveys or tracking historical feedback.
The efforts to analyze data trends or perform audits can be challenging.

_How would you fix it?_

Primary keys implemented as UUIDs enable the storage system to assign a distinct identifier to every individual response. The email address can be used 
as a secondary index to help achieve better query performance.

**Problem 2**

_What is the problem?_

The fetchDepartmentData method executes a full table scan to filter results using the sk attribute (DEPARTMENT#X). Every GET endpoint call makes DynamoDB to scan every record in the table before applying result filters in memory regardless of how many items match. 
Large datasets slow down scan operations because they lack efficiency during processing.

_Why is it a problem?_

Scan operations require extensive resources because they examine all table records but discard those records which fail to meet the filter conditions.
As datasets increase in size, they result in elevated AWS expenses and extended response durations. Regular fetch requests to DynamoDB endpoints often exceed usage limits which negatively impacts user experience.

_How would you fix this problem?_

We can optimize the schema for Query operations which will help in enhanced performance and scalability. To achieve better performance, we must modify the table configuration so that department becomes the GSI partition key. 
With targeted Queries, DynamoDB reaches specific items which delivers quick responses and reduces costs.


**Problem 3**

_What is the problem?_

There are no validations for the POST survey Lambda function for the incoming client data. The Lambda function works with the assumption that incoming request JSON contains properly formatted and structured relevant fields.
If a required field is missing or is of the incorrect type, or if harmful data is provided, the function goes on without any checks or breaks.

_Why is it a problem?_

If data isn’t validated properly, it can cause problems when the data is integrated into a new system. Additionally, invalid entries in the DB can lead to crashes when code tries to 
look for a non-existent field. It also makes subsequent data analysis more complicated if malformed input is stored. Furthermore, omission of input validation can also pose security threats. 

_How would you fix this problem?_

We must implement strict input validation and return specific errors. It's crucial to verify that all necessary fields have the correct type before we use them. One way to check if types are correct is by using Go structs with JSON tags and then decoding into these structs. 
If the decoding fails, it indicates that validation hasn't been successful. We can also add simple checks like if survey.Rating < 1 || survey.Rating > 5 { return clientError("Rating out of range") }. 

**Problem 4**

_What is the problem?_

The Lambda function initializes the AWS clients such as DynamoDB sessions inside the handler for each request using session.New() instead of doing it once for both lambdas. The Database connections are expensive and shouldn't be done repeatedly.

_Why is it a problem?_

Recreating clients every time we invoke a function can really slow things down, whether it's a cold start or not. AWS Lambda tries to reuse the execution environment for subsequent calls whenever it can, so if we don’t take advantage of that, 
we’re just wasting time and CPU on setting things up for each request. This can cause delays and drive up costs. According to AWS best practices, it's a good idea to set up our SDK clients and database connections outside the handler. 
This way, they can be reused for multiple invocations. If we skip this step, each Lambda call will waste time creating new connections to DynamoDB, which is far from efficient.

_How would you fix this problem?_

To optimize performance, we should make sure to initialize and cache any heavy resources outside the handler function. This approach ensures that during a cold start, the client is created just once, and for warm invocations, it gets reused. 
This takes advantage of the Lambda execution environment's ability to reuse resources, ultimately boosting performance.

**Problem 5**

_What is the problem?_

In the bumpResponseCount function, it uses UpdateItem with predefined responses.
`UpdateExpression: aws.String("SET responses = if_not_exists(responses, :start) + :inc"),
`
We can end up with race conditions when there’s a spike in concurrency. Operations done with SET don’t guarantee atomicity across concurrent writes the way ADD does in DynamoDB. If multiple updates hit at once, SET can cause increments to get dropped, 
so it’s not reliable for counters in high-traffic scenarios.

_Why is it a problem?_

In times of high-traffic, if several surveys are submitted at once, they can end up overwriting each other’s counts. This can result in fewer responses being reported than what actually came in, which messes up our analytics and can lead to a lack of trust in the data.
For example, we might think we have 100 responses, but it could actually only be 112. It also makes auditing difficult and once that data is lost, there’s no going back to fix it.

_How would you fix this problem?_

We should instead use ADD operation here.
`UpdateExpression: aws.String("ADD responses = if_not_exists(responses, :start) + :inc"),
`
Using ADD operation helps us safely submit surveys and avoid any over-writes. 



