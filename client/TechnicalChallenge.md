# Introduction and thank you

Recruitment is a two way thing. You've got to like us and what we're all about, and we have to make sure you have the qualities we are looking for in spades. All too often recruitment processes are pretty abstract and don't relate well to the day-to-day roles that are being recruited for. Therefore we decided to make this technical challenge super realistic.

We launched Huler in April 2021, it was a busy time working to get all the required functonality working as expected and get everything ready in time. We're not going to lie, this meant that we didn't produce the best code we've ever seen. Our codebase works, but there are definitely areas for improvement. To be fair, we've seen worse code too.

This technical challenge is typical of our codebase (we've exaggerated some parts of it). We've done this for two reasons,

1. It gives you an accurate idea of what it is like working with our codebase. We are currently in a phase of development where we are adding new features, but we also have some time available to improve what we already have in place. If working with our codebase and technologies is not for you, then now is the best time to find out!

2. It tells us about your skills in working with our codebase, your ability to spot problems, prioritise issues, suggest solutions, as well as your understanding of the technologies we use.

## The premise

For this technical challenge we would like you to imagine that a member of the Huler team has built an initial implementation of an employee satisfaction survey. Some of the functional elements of this survey tool don't make sense, but for the sake of argument let's assume it was built accurately to spec.

This is an initial implementation, you should assume it is likely that we will want to expand and build on this first implementation.

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

There are many problems with this codebase, don't try to find and document all of them. Take some time to understand the codebase, and concentrate on what you feel are the biggest problems. As a rough guide, aim to document what you feel are the biggest 5-10 problems in this codebase.

Some of the problems you find may not have a clear solution based on the information we have given you here. You should assume that there is some commercial pressure to deliver this tool, but there is enough time to complete any required refactoring, i.e. migrating to a completely different technology will not be possible.

### Part 2

For the second part, we would like you to take some time to address one of the problems you found in the first part of this technical challenge.

You will need to build and deploy the application to the AWS infrastructure. The infrastructure is provisioned on a per-request/on-demand basis which means the cost will be negligible (pennies at the most).

### A quick word on cloudformation

We have provisioned the infrastructure using raw cloudformation. This is also the case for our codebase, although it is better organised than what we have put together here.

We are planning to move the majority of our cloudformation infrastructure to CDK over the next 6 months or so.

You are welcome to look at the cloudformation but we would recommend that you don't spend much time with it - concentrate on the javascript and Go parts of this codebase.

## Thank you

Once you are done we would appreciate you checking your response into a private github repo and inviting us to that repo so that we can see what you've been up to!

We hope that you enjoy this challenge, and look forward to receiving your response. Rest assured, we understand the commitment that you are undertaking. We have spent a lot of time preparing this challenge, and will put a lot of effort into reviewing what we get back.

Good Luck!!
