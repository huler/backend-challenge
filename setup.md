# Setup instructions

There are three parts to this codebase,

1. The cloudformation templates that define the infrastructure.
2. The golang code that will serve any api requests
3. The React client application.

We would suggest that you set the infrastructure up first, and then configure the React application.

## AWS installation instructions

In order to set-up the application on AWS.

- Login to the console.

- Create an s3 bucket which you will use to store the code and cloudFormation templates.

- In the s3 bucket create a src directory. Execute the build.sh script (in the api directory, it will create a dist directory) to build the lambda functions and upload them to the src directory in the s3 bucket.

- In the s3 bucket create a folder called cfn. Upload the CFN file from the /cfn directory.

- Click on the main.yaml file and copy its Object URL to your clipboard.

- From the console load the cloudFormation console. Select Create Stack - with new resources (standard). This will load the Create Stack menu.

- Paste the location of the main.yaml file into the Amazon S3 URL field. Click on next.

- Give the stack a name (huler)

- Enter the name of the bucket that you created in the second step.

- Enter the name of the prefix you wish to use for all of the resources in this stack (huler)

- Click next.

- Scroll to the bottom of the page and click on the two tick boxes. Then click on create stack. Don't worry the stack does not spin up any expensive resources. It will take about five minutes for the script to run.

- Once the infrastructure has been created you will need to find the endpoint for the api gateway. Go to the resources tab of your stack, and look for the resource with type `AWS::ApiGateway::RestApi`.

- In the Physical ID column will be a 10 character id (something like 9qa7ws9f9l). You will also need the region you deployed your stack to. For example, if your Physical ID is 9qa7ws9f9l, and this was deployed to the eu-west-1 region (Ireland), the base location for your api endpoints will be `https://9qa7ws9f9l.execute-api.eu-west-1.amazonaws.com`

- Check by entering `https://9qa7ws9f9l.execute-api.eu-west-1.amazonaws.com/Prod/getresults` in to your browser, you will get some results back.

## Client setup instructions

- Now we need to get the app working. cd into the client directory.

- type

```
yarn install
```

- In the client directory there is a file called `.env.example`. Copy this file and rename it to `.env`, modify it's content to the base location for your api endpoints from the section above.

- type

```
yarn start
```

- The app should load in your browser.

- Run through the survey. If everything goes ok you should get a message thanking you for your input, and you should see a post request to the back end.

- You can now check the back-end to see if everything has worked ok. From the AWS console load dynamo and check out the data in the dynamo bucket in this stack.

- If you refresh the page on the survey application, you will now be able to go to the results page.
