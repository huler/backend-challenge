# Introduction

When we launched Huler in April 2021, it was a busy time getting all the required functionality working as expected and getting the product ready to launch. As a business, we are always looking to improve the ways we work and improve the user experience for our customers.

This repository is somewhat reflective of our previous codebase (we've exaggerated some parts of it). We've done this for two reasons.

1. It gives you an accurate idea of where our codebases have evolved from, and the challenges faced.

2. It tells us about your skills in working with the lineage of our codebase, your ability to spot problems, prioritise issues, suggest solutions, as well as your understanding of the technologies we used.

**We are not looking for you to fix every problem in this codebase, we are looking for you to demonstrate your skills in identifying and solving problems.**

## The premise

For this code review we would like you to imagine that a member of the Huler team has built an initial implementation of an employee satisfaction survey.
This survey is backed by two serverless lambda functions: One to POST results and one to GET results. API Gateway is used to route HTTP requests to these lambdas. 
Some of the functional elements of this survey tool don't make sense, but for the sake of argument let's assume it was built accurately to spec.

This is an initial implementation; you should assume it is likely that we will want to expand and build on this first implementation.

### Part 1

We would like you to imagine that you are reviewing this codebase from a technical point of view,
- Are there any bugs?
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
We would like you to take some time to address one or two of the problems you found in the first part of this challenge.

We would encourage you to tackle one of the harder problems you found in the first part. Refactoring the api to use the most appropriate status code might be valuable, but it's not going to show off your skills.

Please fork this repository, commit your solution to a branch, setup a PR into the master branch and give us access.

### Part 3 (optional)
Should you wish to further demonstrate your skills, we would like to see how you would approach the same premise from scratch.

We would like you to build a simple employee satisfaction survey API with two endpoints, using Go and MySQL, following best practices and industry standards.

You should use the same premise as above, but you should not use any of the code in this repository. You can use any libraries or frameworks you like, but we would like to see how you would approach this problem from scratch.

Please create a new repository for this part of the challenge, and give us access to it.

## AI Tool Usage Policy
- **For Part 1 and Part 2:** We kindly request that you **do not use AI assistance tools** (e.g., GitHub Copilot, ChatGPT). The primary goal of these parts is to understand your individual approach to code analysis, problem identification, and solution implementation.
- **For Part 3 (Optional):** You **are permitted** to use AI tools for this part. If you do, please document which tools you used and how you used them in an `AI.md` file in the root of your repository.

## Thank you
Thank you for taking the time to complete this challenge. Good luck!