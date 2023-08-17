# Backend framework comparison

Table of contents:

- [Backend framework comparison](#backend-framework-comparison)
  - [Introduction](#introduction)
  - [Framework selection](#framework-selection)
  - [Comparison metrics](#comparison-metrics)
    - [Community support](#community-support)
      - [Metric weight](#metric-weight)
    - [Documentation Quality](#documentation-quality)
      - [Metric weight](#metric-weight-1)
      - [Results](#results)
    - [Community contributions](#community-contributions)
      - [Open-Closed issue ratio](#open-closed-issue-ratio)
        - [Metric weight](#metric-weight-2)
      - [Pull requests](#pull-requests)
        - [Metric weight](#metric-weight-3)
      - [Stars](#stars)
        - [Metric weight](#metric-weight-4)
    - [Database support](#database-support)
      - [Metric weight](#metric-weight-5)
      - [Results](#results-1)
    - [Security](#security)
      - [Metric weight](#metric-weight-6)
      - [Results](#results-2)
    - [Is compiled](#is-compiled)
    - [Testing](#testing)
      - [Metric weight](#metric-weight-7)
      - [Results](#results-3)
    - [Release recency](#release-recency)
      - [Metric weight](#metric-weight-8)
  - [Framework comparison chart](#framework-comparison-chart)
  - [Comparison result](#comparison-result)
  - [Prototyping conclusion](#prototyping-conclusion)
    - [Gin](#gin)
      - [Summary](#summary)
    - [Django](#django)
      - [Summary](#summary-1)
    - [Conclusion](#conclusion)
  - [References](#references)

## Introduction

This document describes the research for selecting a backend framework for a RESTful API server. This document aims to compare the most popular backend frameworks and provide a substantiated choice for the framework to be used in future development.

## Framework selection

The chart below sums up the most common RESTful HTTP API frameworks. The list of languages to pick was based on online literature research. Python was selected because it is used by many large companies, such as Instagram and Spotify, for their backend services [8], and it is almost always on the top 3 picks for an API [10] [11]. Java is also included because many professionals use it due to its maturity [8] [9]. Go and Rust was also picked because they are new and trending languages for backend development, and since they are new, they have a lot of built-in security and improvements. Many online top-list sources also list Rust and Go as great languages for backend development.

Other languages, such as Ruby and PHP, were not included because they are less popular than the ones listed above, and the aim was to collect at least languages, where at least one is compiled, and one is not.

- `Java`: 2 years experience from University. Java compiles to bytecode, and its JVM can be installed on any machine, so running the backend server on any system is possible. Java is an old, well-mature system. [2]
- `Go`: 5 months of internship experience from one member, new to the other member. Go is a relatively new framework (made in 2007) developed by Google aimed to solve their networking problems and write DevOps solutions [1]. It is used chiefly for compiled server-side applications. A good solution for writing backend API servers.
- `Python`: 5 months of internship experience from one member, new to the other member. Python is an interpreted language that uses JIT (Just in Time) execution to compile and run code when it is executed. Python is a prevalent language both in backend development and in general programming. [3] [5]
- `Rust`: New to both members, chosen for its recent popularity and choice by many big companies to rewrite their C or C++ applications to Rust. Its primary design principle is to be a language like C++, but with way better security, so it is mainly used for low-level programming, offering high performance [4]. Many companies, like AWS [6], also plan to rewrite their services in Rust. Thus, Rust can be great for writing efficient, threaded CLI applications.

## Comparison metrics

This chapter describes the comparison metrics used in the summing-up and comparison table.

### Community support

The "Community support" metric was assessed based on the number of answered questions on the framework's StackOverflow tag(s). The numbers in the table above are deviations from an average calculated across all frameworks.

The average is: `105128`

- Spring: `165822`
- Gin: `592`
- Django: `253863`
- Rocket: `238`

Based on these scores, a framework is graded on a scale of 1 (very bad) to 5 (outstanding). The above-calculated baseline average is a 3.

| Grade | Framework / average |
|-|-|
| 5 | > 2 |
| 4 | 1 - 2 |
| 3 | 0.5 - 1 |
| 2 | 0.1 - 0.5 |
| 1 | < 0.1 |

The data was collected on `15 February 2023`.

#### Metric weight

The weight for the community support metric is a 5 (very important) because this metric describes how active the community is. If it is more activity, there are more questions asked. Thus it is easier to solve problems and find solutions to questions, which increases productivity.

### Documentation Quality

The "Documentation quality" metric was assessed based on the following criteria:

| Mark | Answers question |
|--|--|
| F | Is a documentation index available, and is that index or the documentation itself searchable? |
| E | Does the documentation provide real-life examples? |
| S | Is a project quick to set up? (can be set up within 15 minutes) |

The more "yes" answers, the better the documentation quality. These criteria were chosen based on developers' most common problems when using a framework. Based on this, we provide a 1 to 5 rating.

| Grade | Combination (questions answered) |
|-|-|
| 5 | [F] [E] [S] |
| 4 | [F] [E] |
| 3 | [E] [F] |
| 2 | [E] [S] |
| 1 | Other combination(s) or none |

For us, `findability` is the most important documentation quality. `Findability` means whenever we want to find something specific (how to add CORS to the router), it is easy to find (via search or index). Therefore, it ensures that we can find solutions quickly.

The second most important is `real-life examples`. Documentation can often be verbose and overwhelming with many parameters. Examples ensure that we can quickly write the feature.

The least essential documentation quality is how easy it is to set up a project since this quality does not matter anymore once it is set up.

The grades are then defined based on the combinations of these questions and their importance to us.

#### Metric weight

The weight for the documentation quality metric is a 5 (very important) because this metric describes how easy it is to find solutions to problems. If it is easy to find solutions, it is easier to solve problems, which increases productivity.

#### Results

- Spring answers questions: [E] [F] [S]
- Gin answers questions: [E] [F] [S]
- Django answers questions: [E] [F] [S]
- Rocket answers questions: [E] [F] [S]

### Community contributions

The "Community contributions" metric was assessed based on the number of stars, pull requests, and the ratio of open-to-closed issue tickets. The community contributions metric is rated from 1 (very bad) to 5 (outstanding). The lowest score for 1 is calculated by taking the lowest score from the three metrics, and the highest score for 5 is calculated by taking the highest score from the three metrics. The scales for the scores in between are divided equally as much as possible.

The data was collected on `15 February 2023`.

#### Open-Closed issue ratio

Average of Open-Closed issue ratio: `0.122`

- Rocket: `0.079`
- Django: `0.028`
- Gin: `0.327`
- Spring: `0.056`

| Grade | Framework / average |
|-|-|
| 5 | < 0.25 |
| 4 | 0.25 - 1 |
| 3 | 1 - 2 |
| 2 | 2 - 2.5 |
| 1 | > 2.5 |

##### Metric weight

The weight for this metric is 4/5. This is an important community metric because the more issues are closed, the more active and involved the community is with the framework. The community is more likely to help with problems, and the framework is more likely to be updated with new features. A more active community means the framework is more future-proof.

#### Pull requests

Average of pull requests: `10566`

- Rocket: `698`
- Django: `16505`
- Gin: `1984`
- Spring: `23077`

| Grade | Framework / average |
|-|-|
| 5 | > 2 |
| 4 | 1.5 - 2 |
| 3 | 1 - 1.5 |
| 2 | 0.1 - 1 |
| 1 | < 0.1 |

##### Metric weight

The weight for this metric is 3/5. This is a moderately important community metric because the more pull requests are made, the more active the community is in further developing the framework. This activity means many bugs are fixed, new features are added, and the framework is more mature.

#### Stars

Average of stars: `51446`

- Spring: `50760`
- Gin: `66466`
- Django: `68776`
- Rocket: `19785`

| Grade | Framework / average |
|-|-|
| 5 | > 1.3 |
| 4 | 1 - 1.3 |
| 3 | 0.7 - 1 |
| 2 | 0.4 - 0.7 |
| 1 | < 0.4 |

##### Metric weight

The weight for this metric is 2/5. It is a less important community metric because the more stars a framework has, the more popular it is. This popularity is just a user preference, so this metric is unimportant.

### Database support

The "Database support" metric was assessed based on driver support for the following databases:

| Mark | Meaning |
|--|--|
| Fi | Firebase |
| My | MySQL |
| Mo | MongoDB |
| Po | PostgreSQL |

The following databases were chosen because MongoDB is the most popular NoSQL database (document database), MySQL is the most popular free SQL database, and PostgreSQL is the world's most advanced open-source object-relational database management system [7]. The "Firebase" database was chosen because it is a free, popular cloud database. The more of the above databases are supported, the better the database support.

| Grade | Combination (database is supported) |
|-|-|
| 5 | [Po] [My] [Mo] [Fi] |
| 4 | [Po] [My] [Mo] or [Po] [Mo] [Fi] |
| 3 | [Po] [My] or [Po] [Mo] Or [Po] [Fi] |
| 2 | [Mo] [My] or [My] [Fi] |
| 1 | Other combination(s) or none |

For us, `PostgreSQL` is the most crucial database to support. It is because PostgreSQL is free, open-source, and we have experimented with it. Then, `MySql` is the second most important database to support because it is another relational database that is free and widely used. PostgreSQL and MySQL are the most used databases, according to a survey made by Stackoverflow [12]. `MongoDB` is the third most important database to support because it is a document database that is free and the most used document database. We also wanted to include a no-SQL database to investigate if the framework supports various databases. For this reason, we chose `Firebase` as the fourth most important database because it is a free cloud database.

#### Metric weight

The weight for this metric is 2/5. It is because almost all frameworks support some database, and for a RESTful API, it is less critical which database is supported as that it does support a database.

#### Results

The results are as follows:

- Spring supports: `PostgreSQL`, `MySQL`, `MongoDB`
- Gin supports: `PostgreSQL`, `MySQL`, `MongoDB`, `Firebase`
- Django supports: `PostgreSQL`, `MySQL`, `MongoDB`, `Firebase`
- Rocket supports: `PostgreSQL`, `MySQL`, `MongoDB`

### Security

When choosing security features, TLS encryption, bearer authentication, and SSO (OAuth 2) were selected because they are the most common security features to authenticate in RESTful HTTP APIs.

| Mark | Meaning |
|--|--|
| T | TLS encryption |
| B | Bearer authentication |
| S | SSO |

The more of the above security features are supported, the better the security support.

| Grade | Combination (supported security) |
|-|-|
| 5 | [T] [B] [S] |
| 4 | [T] [B] |
| 3 | [T] [S] |
| 2 | [B] or [S] |
| 1 | Other combination(s) or none |

When considering which security feature is the most important, we chose first TLS because of encryption through the public network, then a basic bearer authentication because it is the most common and simple authentication method, and finally SSO because it is a more advanced authentication method, but not as crucial as a basic bearer authentication.

#### Metric weight

We gave a 5/5 weight for the security metric because securing API communication to the server from the public network is crucial.

#### Results

The results are as follows:

- Spring supports: `TLS`, `Bearer authentication`, `SSO`
- Gin supports: `TLS`, `Bearer authentication`, `SSO`
- Django supports: `TLS`, `Bearer authentication`, `SSO`
- Rocket supports: `None (cookies only)`

### Is compiled

We prefer, in general, compiled languages over interpreted ones because interpreted languages need either a virtual machine software installed separately or an interpreter app. This requirement means there is an overhead in performance (speed, RAM usage, etc.), and it may be less convenient to deploy in a Docker container or on a VPS server.

When comparing frameworks, we want to look at both compiled and interpreted languages to see their differences, so we chose to give this metric no weight. In this comparison, we will select the best framework that is compiled and one that is not compiled.

### Testing

The "Testing" metric was assessed based on support for the following testing methodologies directly in the framework's language:

| Mark | Meaning |
|--|--|
| U | Unit testing |
| E | End-to-end testing |
| I | Integration testing |

The more of the above testing methodologies are supported, the better the testing support.

| Grade | Combination (supported testing methodology) |
|-|-|
| 5 | [U] [E] [i] |
| 4 | [U] [E] |
| 3 | [U] [I] |
| 2 | [U] or [E] |
| 1 | Other combination(s) or none |

Unit testing is the most crucial testing methodology to support because we can test the proper functionality of individual functions, components, and elements that comprise the entire software. End-to-end testing is the second most important testing methodology to support because we can ensure that the software behaves on the user interaction level as expected. Finally, integration testing is the third most important testing methodology to support. With integration testing, we can ensure that our servers can correctly communicate with the database and that their API endpoints work as intended.

#### Metric weight

The weight for this metric is 3/5. Testing is essential. However, testing is only vital if the framework is easy to use and the documentation is good. Testing comes after usability and documentation.

#### Results

The results are as follows:

- Spring supports: `Unit testing`, `End-to-end testing`, `Integration testing
- Gin supports: `Unit testing
- Django supports: `Unit testing
- Rocket supports: `Unit testing

### Release recency

The "Release recency" metric was assessed based on the recency of the last public release (including non-final releases):

| Grade | Meaning |
|--|--|
| 5 | < 2 weeks ago |
| 4 | 2 weeks - 1 month ago |
| 3 | 1 - 3 months ago |
| 2 | 3 - 6 months ago |
| 1 | > 6 months ago |

#### Metric weight

The more recent the last public release, the better the release recency. The weight for this metric is 5/5. It is the most important metric for us. We want to use a framework that is actively developed and maintained. This activity means the framework is not abandoned and will be updated with new features and bug fixes. Furthermore, the more recent the last public release, the more likely to have the most recent security patches and technologies.

## Framework comparison chart

The metrics' weights are defined with a range from 1 to 5, where 1 is the least, and 5 is the most important to consider.

| Comparison metric | Metric weight | Spring (Java) | Gin (Go) | Django (Python) | Rocket (Rust) |
|-------------------|--|-----|-----|--------|--------|
| Documentation quality | 5/5 | 5/5 | 5/5 | 5/5 | 5/5 |
| Community support | 5/5 | 4/5 | 1/5 | 5/5 | 1/5 |
| Community contributions - Stars | 2/5 | 3/5 | 4/5 | 5/5 | 1/5 |
| Community contributions - Pull requests | 3/5 | 5/5 | 2/5 | 4/5 | 1/5 |
| Community contributions - Open-Closed ratio | 4/5 | 4/5 | 1/5 | 5/5 | 4/5 |
| Database support | 2/5 | 4/5 | 5/5 | 5/5 | 4/5 |
| Security | 5/5 | 5/5 | 5/5 | 5/5 | 1/5 |
| Testing | 3/5 | 5/5 | 2/5 | 2/5 | 2/5 |
| Release recency | 5/5 | 3/5 | 3/5 | 5/5 | 1/5 |
| Is compiled | N/A | NO | YES | NO | YES |

## Comparison result

Based on the data collected for various backend frameworks, the two selected frameworks that will be further investigated and compared are Django and Gin.

Django was selected because it displayed the best results overall across all but two of our comparison metrics. It is not run in a compiled language and does not offer end-to-end and integration testing. However, its extensive community support weighs its performance, and only Spring offers all our testing criteria.

While Spring showed the second-best results across our criteria, we chose not to use it for this project. Our first reason for doing so was to compare the performance of compiled and non-compiled frameworks. As we had already settled on researching Django, an interpreter-powered framework, we opted for our second choice to be one of the compiled frameworks, of which Spring is not. Our second reason was our interest in working with a new framework in terms of experience. From our comparison list, all frameworks except Spring have fewer than half a year of experience among us.

Lastly, we selected Gin as the second framework to be prototyped since it has better database support over Spring. Also, Gin scores higher in Community preference by the number of stars. In general, Gin (and also Spring and Django) significantly outperforms Rust's Rocket. Rocket's community support is small, its security is quite basic and limited, and it only sometimes gets updated.

## Prototyping conclusion

In this chapter, we sum up our experiences while prototyping the two frameworks. We collected the pros and cons for each framework. Based on the collected data, we will decide which framework is the best for a RESTful HTTP API.

### Gin

- Pro:
  - Easy to install and set up in a Go project
  - Has built-in functions to handle HTTP responses
  - Has a built-in way to handle path parameters and queries
  - Easy to group endpoints
  - Easy to plug in middleware
  - Provides complete control and customization over the HTTP server
  - Easy to work with router contexts
  - Straightforward to unit test with Go's testing framework
- Con:
  - Because of Go, error handling is very verbose and requires a lot of code, making the code long and hard to read
  - Because of Go, the code is not very compact because of the lack of turnery operators, try-catch blocks, extending structs, etc.
  - Does not manage the database by itself, requiring an external library for database object management
  - Does not have any built-in authentication or security, requires external library or custom code for that
  - Error handling for the database requires a lot of extra code
  - No data seeding

#### Summary

One of the things that I found really simple and convenient about Gin is that is uses Go, and thus it is super simple to install and manage dependencies. Go has a built-in package manager called "go modules" that makes it easy to manage packages and dependencies. Additionally, running and compiling the program is incredibly fast, which saves a lot of time in development.

However, one downside of Go is that its syntax is really verbose. Due to its C-like syntax, I had to write a lot of code to make simple things work. For instance, handling errors in Go is easy but it has a very weird way of handling errors. Errors are always returned instead of being thrown, which makes the code very long as I had to if-check all functions that return errors if their returned error is nil or not. This is a bit annoying, but simple enough to use and understand

Moreover, the previously conducted research showed a less extensive community support for Gin than other popular languages, such as Django or Spring. Nevertheless, I was pleasantly surprised to find a wide range of open-source packages and libraries available for use in my project. Additionally, I found a lot of helpful articles and answers on StackOverflow on how to use Gin and solve common problems with it.

What was also great about Gin, is that there were no magic involved. Even if I had to write everything myself, I understood what does what, how and where it is done. Thus extending, customizing and tweaking the API was very easy and straightforward.

Overall, my experience with Gin has been a bit challenging due to its verbose syntax thanks to Go and error handling approach. However, I appreciate its fast performance, easy dependency management, and the availability of open-source packages. Once the initial setup is done to layout the structure, the rest of the development is very fast and easy.

### Django

- Pro:
  - Easy, automated setup
  - Tightly integrates with the database through built-in ORM
  - Has pre-built classes for defining endpoints that require effectively no code to configure
  - Infers many API properties and settings automatically from Model and Serialization structures
  - Authentication is easy to set up and integrate into built-in security
  - Automatic security setup (hashing, validation, permissions)
  - Automatic API documentation and browsable API page
  - High-quality documentation
- Con:
  - Working with data from multiple Models requires writing custom logic
  - The high-level nature of the framework obfuscates certain basic HTTP server features (e.g., error code configuration)
  - Customizing specific API properties requires overriding built-in classes
  - Minimal support for integration testing (only basic requests via Python library)

#### Summary

In short, the experience with Django, from setup to "production", has been a breeze. Creating the project and database models barely require any guidance as the automated setup script and intuitive ORM created a great jumping-off point. Creating and configuring endpoints also required practically no scripting, in some cases being finished in just 5 lines of code. Django is able to infer and auto-configure most aspects of an API based off the code that makes up the models and their serializers. However, this simplicity also proved to be a double-edged sword when it came to extending the built-in functionality, like in the cases of providing custom error message, API action behaviour, or lookup logic. Engaging in this kind of customization can quickly turn into a mess of function overrides and raw Django HTTP object manupulation, which we luckily did not have to do too much for this assignment.

In the rare cases that Django's automation and logic structures failed to guide us, the excellent official documentation and community support got things back on track very quickly. In fact, community contributions also saved us time when it came to expanding Django's authentication functionality with JWT, but having multiple high-quality framework plugins to choose from. Django's documentation quality also extends to the documentation we can produce for the project, with a browsable API-schema being automatically generated as code is written and OpenAPI specification schema generation available with a single command.

Unfortunately, one area in which Django did let us down is testing. While it does offer a custom test suite for unit tests, API integration tests left a lot more to be desired as we had to use Python's requests library, with little extension on Django's end. Even though this did not make testing impossible, it did make the process more verbose and time-consuming with lengthy request parameter syntax and manual workarounds for test sequence ordering being necessary.

Even with its lower-level flaws, Django offers a refreshing take on building APIs that prioritizes simplicity and speed. This approach made it a fantastic framework to use for this rapid prototyping assignment and will undoubtedly save time when used in a larger-scale project.

### Conclusion

For the final framework, we chose Python's Django. This choice was primarily driven by Django's surprising ease of development. With much of the API requiring no code beyond built-in class instantiation and many API properties being inferred or auto-configured by the framework, using Django will save significant development time. Another major factor in our choice was the wealth of public documentation and tutorials, which undoubtedly aid in customizing aspects of the framework when needed.

We chose something other than Go's Gin because writing most of the functionality from scratch takes a lot of time. It is not difficult to work with Go and Gin at all. However, it takes a lot of code to cover all functionality and error handling. In addition, the fact that Gin does not manage the database by itself makes it less appealing. Gin only handles the HTTP requests: structures destructures and provides endpoints. The developer must implement other functionality, such as authentication, database communication, response objects, etc. While not a problem, it requires a lot of extra code to be written.

---

## References

[1] K. Chris, "What is go? Golang programming language meaning explained," freeCodeCamp.org, 07-Oct-2021. [Online]. Available: <https://www.freecodecamp.org/news/what-is-go-programming-language/>. [Accessed: 14-Feb-2023].

[2] A. Binstock, "Oracle brandvoice: Java's 20 years of innovation," Forbes, 20-May-2015. [Online]. Available: <https://www.forbes.com/sites/oracle/2015/05/20/javas-20-years-of-innovation/>. [Accessed: 14-Feb-2023].

[3] JetBrains, "Python programming - the state of developer ecosystem in 2021 infographic," JetBrains, 19-Jul-2021. [Online]. Available: <https://www.jetbrains.com/lp/devecosystem-2021/python/>. [Accessed: 14-Feb-2023].

[4] K. Wróbel, "Why is Rust programming language so popular?," Codilime.com, 25-Mar-2022. [Online]. Available: <https://codilime.com/blog/why-is-rust-programming-language-so-popular/>. [Accessed: 14-Feb-2023].

[5] OfferZen, "State of the netherlands' software developer nation 2022," OfferZen, 15-Jun-2022. [Online]. Available: <https://www.offerzen.com/reports/software-developer-netherlands/>. [Accessed: 14-Feb-2023].

[6] M. Asay, "Why AWS loves Rust, and how we'd like to help," Amazon, 24-Nov-2020. [Online]. Available: <https://aws.amazon.com/blogs/opensource/why-aws-loves-rust-and-how-wed-like-to-help/>. [Accessed: 14-Feb-2023].

[7] F. Andrade, "Top databases to use in 2022: What is the right database for your use case?," Medium, 30-Sep-2022. [Online]. Available: <https://towardsdatascience.com/top-databases-to-use-in-2022-what-is-the-right-database-for-your-use-case-bb8d3f183b21/>. [Accessed: 14-Feb-2023].

[8] M. Pruciak, "The best programming languages to use in your backend in 2022," Agile Software Development Agency in Europe, 31-Jan-2022. [Online]. Available: <https://www.ideamotive.co/blog/best-programming-languages-to-use-in-your-backend/>. [Accessed: 21-Feb-2023].

[9] Codetheorem, "Top 10 backend languages for web developers in 2022," Codetheorem, 02-Aug-2022. [Online]. Available: <https://codetheorem.co/blogs/top-backend-languages/>. [Accessed: 21-Feb-2023].

[10] Javinpaul, "10 best backend frameworks for web development in 2023," Medium, 09-Mar-2022. [Online]. Available: <https://medium.com/javarevisited/10-best-backend-frameworks-for-web-development-8d19e337f774/>. [Accessed: 21-Feb-2023].

[11] APINinjas, "Top 5 programming languages in 2022," API Ninjas Blog, 12-Jul-2022. [Online]. Available: <https://api-ninjas.com/blog/top-5-programming-languages-in-2022/>. [Accessed: 21-Feb-2023].

[12] A. Mehta, "Top 15 databases to use in 2022 and beyond," Appinventiv, 08-Feb-2023. [Online]. Available: <https://appinventiv.com/blog/top-web-app-database-list/>. [Accessed: 23-Feb-2023].
