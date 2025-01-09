# Introduction

When we launched Huler in April 2021, it was a busy time getting all the required functionality working as expected and getting the product ready to launch. As a business, we are always looking to improve the ways we work and improve the user experience for our customers.

This repository is somewhat reflective of our codebase (we've exaggerated some parts of it). We've done this for two reasons,

1. It gives you an accurate idea of what it is like working with our codebase. We are currently in a phase of development where we are adding new features, but we also have some time available to improve what we already have in place.

2. It tells us about your skills in working with our codebase, your ability to spot problems, prioritise issues, suggest solutions, as well as your understanding of the technologies we use.

## The premise

For this code review we would like you to imagine that a member of the Huler team has built an initial implementation of an employee satisfaction survey. This survey is backed by two serverless lambda functions: One to POST results and one to GET results. API Gateway is used to route HTTP requests to these lambdas. 
Some of the functional elements of this survey tool don't make sense, but for the sake of argument let's assume it was built accurately to spec.

This is an initial implementation; you should assume it is likely that we will want to expand and build on this first implementation.

### Part 1

We would like you to imagine that you are reviewing this codebase from a technical point of view,

- Are there are any bugs?
- Have we made any decisions that will make future development harder?
- Can we make the codebase easier to work with?
- Have we made any decisions that will cause problems as we scale to a large number of users?

For each problem you find, we would like you to add three sentences to the problems.md file,

1. What is the problem?
2. Why is it a problem?
3. How would you fix this problem?

As an example,
"The api always returns status code 5xx if there are any problems. This means that the React application can't differentiate between a server-side issue and a client-side issue. I would return the appropriate status code, 4xx in this case if there are any problems with the data provided by the user."

There are many problems with this codebase, we don’t expect you to find and document all of them. Take some time to understand the codebase, and concentrate on what you feel are the biggest problems. As a rough guide, aim to document the 5 problems you would prioritise fixing in this codebase. Feel free to write a brief list summarising problems outside of your top 5.

Some of the problems you find may not have a clear solution based on the information we have given you here. These points will form discussion points for the next stage of our interview process, that will be an opportunity for us to discuss the possible solutions and the tradeoffs involved.

You should assume that there is some commercial pressure to deliver this tool so refactoring is possible but migrating to completely different technologies is not.

### Part 2

We would like you to take some time to address one or two of the problems you found in the first part of this code review.

We would encourage you to tackle one of the harder problems you found in the first part. Refactoring the api to use the most appropriate status code might be valuable, but it's not going to show off your skills.

Please fork this repository, commit your solution to a feature branch, setup a PR into the master branch and give us access. 

## Thank you

Thank you for taking the time to complete this code review. Good luck!

