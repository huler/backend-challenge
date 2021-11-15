# Introduction

When we launched Huler in April 2021, it was a busy time working to get all the required functionality working as expected and getting the product ready to launch. As a business, we are always looking to improve the ways we work and improve the user experience for our customers.

This technical challenge is typical of our codebase (we've exaggerated some parts of it). We've done this for two reasons,

1. It gives you an accurate idea of what it is like working with our codebase. We are currently in a phase of development where we are adding new features, but we also have some time available to improve what we already have in place.

2. It tells us about your skills in working with our codebase, your ability to spot problems, prioritise issues, suggest solutions, as well as your understanding of the technologies we use.

## The premise

For this technical challenge we would like you to imagine that a member of the Huler team has built an initial implementation of an employee satisfaction survey. Some of the functional elements of this survey tool don't make sense, but for the sake of argument let's assume it was built accurately to spec.

This is an initial implementation; you should assume it is likely that we will want to expand and build on this first implementation.

### Part 1

We would like you to imagine that you are reviewing this codebase from a technical point of view,

- Are there are any bugs?
- Have we made any decisions that will make future development harder?
- Can we make the codebase easier to work with?
- Have we made any decisions that will cause problems as we scale to a large number of users?

For each problem you find, we would like you to write three sentences,

1. What is the problem?
2. Why is it a problem?
3. How would you fix this problem?

As an example,
"The api always returns status code 5xx if there are any problems. This means that the React application can't differentiate between a server-side issue and a client-side issue. I would return the appropriate status code, 4xx in this case if there are any problems with the data provided by the user."

There are many problems with this codebase, we don’t expect you to find and document all of them. Take some time to understand the codebase, and concentrate on what you feel are the biggest problems. As a rough guide, aim to document the 5 problems you would prioritise fixing in this codebase.

Some of the problems you find may not have a clear solution based on the information we have given you here. These points will form discussion points for the next stage of our interview process, that will be an opportunity for us to discuss the possible solutions and the tradeoffs involved.

You should assume that there is some commercial pressure to deliver this tool, but there is enough time to complete any required refactoring, i.e. migrating to a completely different technology will not be possible.

### Part 2

For the second part, we would like you to take some time to address one or two of the problems you found in the first part of this technical challenge.

We would encourage you to tackle one of the harder problems you found in the first part. Refactoring the api to use the most appropriate status code might be valuable, but it's not the hardest problem to solve in this codebase.

As a rough guide, if we were to look at the commit history then we would likely see between 100 and 200 lines of code in total that have been added/removed/modified.

You will need to build and deploy the application to the AWS infrastructure. The infrastructure is provisioned on a per-request/on-demand basis and should fall in to the "Always Free" tier. We would encourage you to pull down the infrastructure when you are finished - the dynamo tables are configured to be deleted even if there is data in them.

### A quick word on cloudformation

We have provisioned the infrastructure using raw cloudformation. This is also the case for our codebase, although it is better organised than what we have put together here.

We are planning to move the majority of our cloudformation infrastructure to CDK over the next 6 months or so.

You might feel that some of the problems you find in the first part are best solved with changes to our infrastructure. We can discuss these suggested changes, but we would encourage you to leave the infrastructure as it is - diagnosing cloudformation issues can be time consuming and don't provide a great deal of value to this process.

## Thank you

Once you are done we would appreciate you checking your response into a private github repo and invite 'DewiWilliamsMCG' (dewi.williams@myclevergroup.com).
