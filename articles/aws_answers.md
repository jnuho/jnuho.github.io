
- Exam Objectives
  - According to the AWS Certified Cloud Practitioner Exam Guide (version 1.4), the AWS Certified Cloud Practitioner (CLF-C01) examination is “intended for individuals who have the knowledge and skills necessary to effectively demonstrate an overall understanding of the AWS Cloud, independent of specific technical roles addressed by other AWS certifica- tions” (for example, solution architects or SysOps administrators).
  - The AWS Cloud and its basic global infrastructure
  - AWS Cloud architectural principles
  - The AWS Cloud value proposition
  - Key AWS services along with their common use cases (for example, highly available web applications or data analysis)
  - The basic security and compliance practices relating to the AWS platform and the shared security model
  - AWS billing, account management, and pricing models
  - Documentation and technical assistance resources
  - Basic characteristics for deploying and operating in the AWS Cloud AWS recommends that “candidates have at least six months of experience with the AWS Cloud in any role, including technical, managerial, sales, purchasing, or financial.” They should also possess general knowledge of information technology and application servers and their uses in the AWS Cloud.

- Assessment Test

1. Which of the following describes the cloud design principle of scalability?
A. The ability to automatically increase available compute resources to meet growing user
demand
B. The ability to route incoming client requests between multiple application servers
C. The ability to segment physical resources into multiple virtual partitions
D. The ability to reduce production costs by spreading capital expenses across many
accounts
1. A. A scalable deployment will automatically “scale up” its capacity to meet growing user
demand without the need for manual interference. See Chapter 1.

2. Which of the following best describes the cloud service model known as infrastructure as a
service (IaaS)?
A. End user access to software applications delivered over the internet
B. Access to a simplified interface through which customers can directly deploy
their application code without having to worry about managing the underlying
infrastructure
C. Customer rental of the use of measured units of a provider’s physical compute, storage,
and networking resources
D. Abstracted interfaces built to manage clusters of containerized workloads
2. C. IaaS is a model that gives customers access to virtualized units of a provider’s physical
resources. IaaS customers manage their infrastructure much the way they would local,
physical servers. See Chapter 1.

3. How does AWS ensure that no single customer consumes an unsustainable proportion of available resources?
A. AWS allows customers to consume as much as they’re willing to pay for, regardless of
general availability.
B. AWS imposes default limits on the use of its service resources but allows customers to
request higher limits.
C. AWS imposes hard default limits on the use of its service resources.
D. AWS imposes default limits on the use of its services by Basic account holders;
Premium account holders face no limits.
3. B. AWS applies usage limits on most features of its services. However, in many cases, you
can apply for a limit to be lifted. See Chapter 2.

4. The AWS Free Tier is designed to give new account holders the opportunity to get to know
how their services work without necessarily costing any money. How does it work?
A. You get service credits that can be used to provision and launch a few typical
workloads.
B. You get full free access to a few core AWS services for one month.
C. You get low-cost access to many core AWS services for three months.
D. You get free lightweight access to many core AWS services for a full 12 months.
4. D. The Free Tier offers you free lightweight access to many core AWS services for a full
12 months. See Chapter 2.

5. AWS customers receive “production system down” support within one hour when they
subscribe to which support plan(s)?
A. Enterprise.
B. Business and Enterprise.
C. Developer and Basic.
D. All plans get this level of support.
5. B. “Production system down” support within one hour is available only to subscribers to
the Business or Enterprise support plans. See Chapter 3.

6. AWS customers get full access to the AWS Trusted Advisor best practice checks when they
subscribe to which support plan(s)?
A. All plans get this level of support.
B. Basic and Business.
C. Business and Enterprise.
D. Developer, Business, and Enterprise.
6. D. All support plans come with full access to Trusted Advisor except for the (free) Basic
plan. See Chapter 3.

7. The AWS Shared Responsibility Model illustrates how AWS itself (as opposed to its
customers) is responsible for which aspects of the cloud environment?
A. The redundancy and integrity of customer-added data
B. The underlying integrity and security of AWS physical resources
C. Data and configurations added by customers
D. The operating systems run on EC2 instances
7. B. According to the Shared Responsibility Model, AWS is responsible for the underlying
integrity and security of AWS physical resources, but not the integrity of the data and
configurations added by customers. See Chapter 4.

8. Which of these is a designation for two or more AWS data centers within a single
geographic area?
A. Availability Zone
B. Region
C. Network subnet
D. Geo-unit
8. A. An Availability Zone is one of two or more physical data centers located within a single
AWS Region. See Chapter 4.

9. How, using security best practices, should your organization’s team members access your
AWS account resources?
A. Only a single team member should be given any account access.
B. Through a jointly shared single account user who’s been given full account-wide
permissions.
C. Through the use of specially created users, groups, and roles, each given the fewest
permissions necessary.
D. Ideally, resource access should occur only through the use of access keys.
9. C. Team members should each be given identities (as users, groups, and/or roles) configured
with exactly the permissions necessary to do their jobs and no more. See Chapter 5.

10. Which of the following describes a methodology that protects your organization’s data
when it’s on-site locally, in transit to AWS, and stored on AWS?
A. Client-side encryption
B. Server-side encryption
C. Cryptographic transformation
D. Encryption at rest
10. A. End-to-end encryption that protects data at every step of its life cycle is called client-side
encryption. See Chapter 5.

11. What authentication method will you use to access your AWS resources remotely through
the AWS Command Line Interface (CLI)?
A. Strong password
B. Multifactor authentication
C. SSH key pairs
D. Access keys
11. D. AWS CLI requests are authenticated through access keys. See Chapter 6.

12. Which of these is the primary benefit from using resource tags with your AWS assets?
A. Tags enable the use of remote administration operations via the AWS CLI.
B. Tags make it easier to identify and administrate running resources in a busy AWS
account.
C. Tags enhance data security throughout your account.
D. Some AWS services won’t work without the use of resource tags.
12. B. Resource tags—especially when applied with consistent naming patterns—can make it
easier to visualize and administrate resources on busy accounts. See Chapter 6.

13. What defines the base operating system and software stack that will be available for a new
Elastic Compute Cloud (EC2) instance when it launches?
A. The Virtual Private Cloud (VPC) into which you choose to launch your instance.
B. The instance type you select.
C. The Amazon Machine Image (AMI) you select.
D. You don’t need to define the base OS—you can install that once the instance launches.
13. C. The AMI you select while configuring your new instance defines the base OS. See
Chapter 7.

14. Which of the following AWS compute services offers an administration experience that
most closely resembles the way you would run physical servers in your own local data
center?
A. Simple Storage Service (S3)
B. Elastic Container Service (ECS)
C. Elastic Compute Cloud (EC2)
D. Lambda
14. C. You can administrate EC2 instances using techniques that are similar to the way you’d
work with physical servers. See Chapter 7.

15. Which of the following AWS object storage services offers the lowest ongoing charges, but
at the cost of some convenience?
A. Glacier
B. Storage Gateway
C. Simple Storage Service (S3)
D. Elastic Block Store (EBS)
15. A. Amazon Glacier can reliably store large amounts of data for a very low price but
requires CLI or SDK administration access, and retrieving your data can take hours. See

16. Which of the following AWS storage services can make the most practical sense for
petabyte-sized archives that currently exist in your local data center?
A. Saving to a Glacier Vault
B. Saving to a Simple Storage Service (S3) bucket
C. Saving to an Elastic Block Store (EBS) volume
D. Saving to an AWS Snowball device
16. D. You can transfer large data stores to the AWS cloud (to S3 buckets) by having Amazon
send you a Snowball device to which you copy your data and which you then ship back to
Amazon. See Chapter 8.

17. Which of the following will provide the most reliable and scalable relational database
experience on AWS?
A. Relational Database Service (RDS)
B. Running a database on an EC2 instance
C. DynamoDB
D. Redshift
17. A. RDS offers a managed and highly scalable database environment for most popular
relational database engines (including MySQL, MariaDB, and Oracle). See Chapter 9.
18. What’s the best and simplest way to increase reliability of an RDS database instance?
A. Increase the available IOPS.
B. Choose the Aurora database engine when you configure your instance.
C. Enable Multi-AZ.
D. Duplicate the database in a second AWS Region.
18. C. Multi-AZ will automatically replicate your database in a second Availability Zone for
greater reliability. It will, of course, also double your costs. See Chapter 9.
19. How does AWS describe an isolated networking environment into which you can launch
compute resources while closely controlling network access?
A. Security group
B. Virtual private cloud (VPC)
C. Availability Zone
D. Internet gateway
19. B. A VPC is an isolated networking environment into which you can launch compute
resources while closely controlling network access. See Chapter 10.
20. What service does AWS use to provide a content delivery network (CDN) for its customers?
A. VPC peering
B. Internet gateway
C. Route 53
D. CloudFront
20. D. CloudFront is a content delivery network (CDN) that distributes content through its
global network of edge locations. See Chapter 10.
21. What is Amazon’s Git-compliant version control service for integrating your source code
with AWS resources?
A. CodeCommit
B. CodeBuild
C. CodeDeploy
D. Cloud9
21. A. CodeCommit is a Git-compliant version control service for integrating your source code
with AWS resources. See Chapter 11.
22. Which AWS service allows you to build a script-like template representing complex resource
stacks that can be used to launch precisely defined environments involving the full range of
AWS resources?
A. LightSail
B. EC2
C. CodeDeploy
D. CloudFormation
22. D. CloudFormation templates can represent complex resource stacks that can be used
to launch precisely defined environments involving the full range of AWS resources. See
Chapter 11.
23. What is Amazon Athena?
A. A service that permits queries against data stored in Amazon S3
B. A service that permits processing and analyzing of real-time video and data streams
C. A NoSQL database engine
D. A Greece-based Amazon Direct Connect service partner
23. A. Amazon Athena is a managed service that permits queries against S3-stored data. See Chapter 13.
24. What is Amazon Kinesis?
A. A service that permits queries against data stored in Amazon S3
B. A service that permits processing and analyzing of real-time video and data streams
C. A NoSQL database engine
D. A Greece-based Amazon Direct Connect service partner
24. B. Amazon Kinesis allows processing and analyzing of real time video and data streams.
See Chapter 13.
25. What is Amazon Cognito?
A. A service that can manage authentication and authorization for your public-facing
applications
B. A service that automates the administration of authentication secrets used by your
AWS resources
C. A service that permits processing and analyzing of real-time video and data streams
D. A relational database engine

25. A. Amazon Cognito can manage authentication and authorization for your public-facing
applications. See Chapter 13.


- Chapter1. Cloud: Exam Essentials
  - Understand how a large and geographically dispersed infrastructure improves service quality.
The sheer scale and geographic redundancy of the physical compute and networking
resources owned by AWS mean that the company is able to guarantee a level of reliability
and availability that would be hard to reproduce in any other environment.
  - Understand how metered, pay-per-use pricing makes for flexible compute options. Access
to cloud infrastructure—sometimes for pennies per hour—makes it possible to experiment,
sandbox, and regularly reassess and update application stacks.
  - Understand that cloud services come on a wide range of forms. IaaS gives you near-full
control over virtualized hardware resources, closely emulating the way you would admin-
istrate actual physical servers. PaaS products abstract the underlying infrastructure, pro-
viding a simplified interface for you to add your application code. SaaS products provide
services over a public network directly to end users.
  - Understand how serverless computing can be both cheap and efficient. Serverless services
like AWS Lambda allow you access to AWS compute power for up to 15 minutes for a
single function. This lets you operate code in response to real-time event triggers.
  - Understand how scalability allows applications to grow to meet need. A cloud-optimized
application allows for automated provisioning of server instances that are designed from
scratch to perform a needed compute function within an appropriate network environment.
  - Understand how elasticity matches compute power to both rising and falling demand.
The scaling services of a cloud provider—like AWS Auto Scaling—should be configured to
force compliance with your budget and application needs. You set the upper and lower lim-
its, and the scaler handles the startups and shutdowns to optimize operations in between
those limits.



Chapter1. Review Questions

1.
Which of the following does not contribute significantly to the operational value of a large
cloud provider like AWS?
A. Multiregional presence
B. Highly experienced teams of security engineers
C. Deep experience in the retail sphere
D. Metered, pay-per-use pricing

1. C. Having globally distributed infrastructure and experienced security engineers makes a
provider’s infrastructure more reliable. Metered pricing makes a wider range of workloads
possible.

2.
Which of the following are signs of a highly available application? (Select TWO.)
A. A failure in one geographic region will trigger an automatic failover to resources in a
different region.
B. Applications are protected behind multiple layers of security.
C. Virtualized hypervisor-driven systems are deployed as mandated by company policy.
D. Spikes in user demand are met through automatically increasing resources.

2. A, D. Security and virtualization are both important characteristics of successful cloud
workloads, but neither will directly impact availability.

3.
How does the metered payment model make many benefits of cloud computing possible?
(Select TWO.)
A. Greater application security is now possible.
B. Experiments with multiple configuration options are now cost-effective.
C. Applications are now highly scalable.
D. Full-stack applications are possible without the need to invest in capital expenses.

3. B, D. Security and scalability are important cloud elements but are not related to metered
pricing.

4.
Which of the following are direct benefits of server virtualization? (Select TWO.)
A. Fast resource provisioning and launching
B. Efficient (high-density) use of resources
C. Greater application security
D. Elastic application designs

4. A, B. Security and elasticity are important but are not directly related to server
virtualization.

5.
What is a hypervisor?
A. Hardware device used to provide an interface between storage and compute modules
B. Hardware device used to provide an interface between networking and compute
modules
C. Software used to log and monitor virtualized operations
D. Software used to administrate virtualized resources run on physical infrastructureReview Questions

5. D. A hypervisor is software (not hardware) that administrates virtualized operations.

6.
Which of the following best describes server virtualization?
A. “Sharding” data from multiple sources into a single virtual data store
B. Logically partitioning physical compute and storage devices into multiple smaller
virtual devices
C. Aggregating physical resources spread over multiple physical devices into a single
virtual device
D. Abstracting the complexity of physical infrastructure behind a simple web interface

6. B. Sharding, aggregating remote resources, and abstracting complex infrastructure
can all be accomplished using virtualization techniques, but they aren’t, of themselves,
virtualization.

7.
Which of the following best describes Infrastructure as a Service products?
A. Services that hide infrastructure complexity behind a simple interface
B. Services that provide a service to end users through a public network
C. Services that give you direct control over underlying compute and storage resources
D. Platforms that allow developers to run their code over short periods on cloud servers

7. C. PaaS products mask complexity, SaaS products provide end-user services, and serverless
architectures (like AWS Lambda) let developers run code on cloud servers.

8.
Which of the following best describes Platform as a Service products?
A. Services that hide infrastructure complexity behind a simple interface
B. Platforms that allow developers to run their code over short periods on cloud servers
C. Services that give you direct control over underlying compute and storage resources
D. Services that provide a service to end users through a public network

8. A. IaaS products provide full infrastructure access, SaaS products provide end-user
services, and serverless architectures (like AWS Lambda) let developers run code on cloud
servers.

9.
Which of the following best describes Software as a Service products?
A. Services that give you direct control over underlying compute and storage resources
B. Services that provide a service to end users through a public network
C. Services that hide infrastructure complexity behind a simple interface
D. Platforms that allow developers to run their code over short periods on cloud servers

9. B. IaaS products provide full infrastructure access, PaaS products mask complexity, and
serverless architectures (like AWS Lambda) let developers run code on cloud servers.

10. Which of the following best describes scalability?
A. The ability of an application to automatically add preconfigured compute resources to
meet increasing demand
B. The ability of an application to increase or decrease compute resources to match
changing demand
C. The ability to more densely pack virtualized resources onto a single physical server
D. The ability to bill resource usage using a pay-per-user modelChapter 1

10. A. Increasing or decreasing compute resources better describes elasticity. Efficient use of
virtualized resources and billing models aren’t related directly to scalability.

11. Which of the following best describes elasticity?
A. The ability to more densely pack virtualized resources onto a single physical server
B. The ability to bill resource usage using a pay-per-user model
C. The ability of an application to increase or decrease compute resources to match
changing demand
D. The ability of an application to automatically add preconfigured compute resources to
meet increasing demand

11. C. Preconfiguring compute instances before they’re used to scale up an application is an
element of scalability rather than elasticity. Efficient use of virtualized resources and billing
models aren’t related directly to elasticity.

12. Which of the following characteristics most help AWS provide such scalable services?
(Select TWO.)
A. The enormous number of servers it operates
B. The value of its capitalized assets
C. Its geographic reach
D. Its highly automated infrastructure administration systems

12. A, D. Capitalized assets and geographic reach are important but don’t have a direct impact
on operational scalability.


- Chapter2. Summary
The Free Tier offers new accounts a full year of free access to light configurations of most
AWS services. It’s meant as an opportunity to experiment and learn how your organiza-
tion’s needs can be met in the cloud. You can track how close you are to “outspending”
your Free Tier allowance from the Billing Dashboard.

You can find pricing information online by adding /pricing to the URL of an AWS
service. aws.amazon.com/s3/ , for instance, would become aws.amazon.com/s3/pricing .
It’s important to understand both the billing rates and the specific metrics used to measure
them.

The AWS Simple Monthly Calculator and AWS Total Cost of Ownership Calculator let
you anticipate real-world usage costs for AWS deployments and (in the case of the TCO cal-
culator) compare your spend with its on-premises equivalent.

Resource requests will sometimes be refused (and launches will fail) if your request
would have pushed your consumption past a service limit.

The AWS Billing Dashboard is the hub for accessing account administration, payment
and tax status management, cost monitoring, and cost control.

- Chapter2. Exam Essentials
  - Understand the value of the 12-month Free Tier. The Free Tier lets you run light services
such as the t2.micro EC2 instance type and a 30 GB SSH EBS volume. The goal is to get
you comfortable with the AWS environment so you can learn how it can be used to host
your applications.
  - Understand the value of permanent Free Tier services. Low-volume consumption includes
the retrieval of up to 10 GB of stored objects from Glacier or 62,000 outbound emails
through Amazon SES. The goal is to give you the opportunity to launch proof-of-concept
deployments.
  - Know how to access Amazon’s resource pricing online documentation. To accurately cal-
culate the true costs of an AWS deployment, you must understand the pricing for the par-
ticular level of resource you launch within a particular AWS Region. Each service resource
(like an EC2 instance) is billed by metrics unique to its characteristics.
  - Use the AWS Simple Monthly Calculator to accurately model multitiered application stack
pricing. Pricing for all variations of the core AWS services is prebuilt into the calculator,
allowing you to model pricing for multiple resource configurations.
  - Use the AWS Total Cost of Ownership Calculator to compare on-premises with AWS
deployment costs. You can conveniently compare apples to apples—capital expenses for
on-premises versus operating expenses for cloud—to know whether the AWS cloud is really
right for your workload.
  - Understand how your use of AWS services is limited by default. Access to all service
resources is restricted by default limits. In many cases, you can manually request limit
increases from AWS support.
  - Understand the value of cost management tools for avoiding costly cloud overspends. AWS
budgets can be configured to send alerts when your resource consumption approaches or
passes a preset limit. Cost Explorer provides visualizations to more easily monitor historical
and current costs. Cost and usage reports can send in-depth and ongoing CSV-formatted
data to Redshift or QuickSight for analysis. You can use cost allocation tags to more effec-
tively track and manage the source of account costs. The security and operations of multiple
AWS accounts controlled by a single company can be managed through AWS Organizations.


- Chapter2. Review Questions
1.
Which of the following EC2 services can be used without charge under the Free Tier?
A. Any single EC2 instance type as long as it runs for less than one hour per day
B. Any single EC2 instance type as long as it runs for less than 75 hours per month
C. A single t2.micro EC2 instance type instance for 750 hours per month
D. t2.micro EC2 instance type instances for a total of 750 hours per month
1.
D. Only the t2.micro instance type is Free Tier–eligible, and any combination of t2.micro
instances can be run up to a total of 750 hours per month.
2.
You want to experiment with deploying a web server on an EC2 instance. Which two of
the following resources can you include to make that work while remaining within the Free
Tier? (Select TWO.)
A. A 5 GB bucket on S3
B. A t2.micro instance type EC2 instance
C. A 30 GB solid-state Elastic Block Store (EBS) drive
D. Two 20 GB solid-state Elastic Block Store (EBS) drives

2. B, C. S3 buckets—while available in such volumes under the Free Tier—are not necessary
for an EC2 instance. Since the maximum total EBS space allowed by the Free Tier is 30 GB,
two 20 GB would not be covered.

3.
Which of the following usage will always be cost-free even after your account’s Free Tier
has expired? (Select TWO.)
A. One million API calls/month on Amazon API Gateway
B. 10 GB of data retrievals from Amazon Glacier per month
C. 500 MB/month of free storage on the Amazon Elastic Container Registry (ECR)
D. 10 custom monitoring metrics and 10 alarms on Amazon CloudWatch
3. B, D. The API calls/month and ECR free storage are available only under the Free Tier.
4.
Which of the following tools are available to ensure you won’t accidentally run past your
Free Tier limit and incur unwanted costs? (Select TWO.)
A. Automated email alerts when activity approaches the Free Tier limits
B. The Top Free Tier Services by Usage section on the Billing & Cost Management
Dashboard
C. Billing & Cost Management section on the Top Free Tier Services Dashboard
D. The Billing Preferences Dashboard
4. A, B. There is no Top Free Tier Services Dashboard or, for that matter, a Billing Preferences
Dashboard.
5.
Which of the following is likely to be an accurate source of AWS pricing information?
A. Wikipedia pages relating to a particular service
B. The AWS Command Line Interface (AWS CLI)
C. AWS online documentation relating to a particular service
D. The AWS Total Cost of Ownership Calculator
5. C. Wikipedia pages aren’t updated or detailed enough to be helpful in this respect. The
AWS CLI isn’t likely to have much (if any) pricing information. The TCO Calculator
shouldn’t be used for specific and up-to-date information about service pricing.
6.
Which of the following will probably not affect the pricing for an AWS service?
A. Requests for raising the available service limit
B. AWS Region
C. The volume of data saved to an S3 bucket
D. The volume of data egress from an Amazon Glacier vault
6. A. Pricing will normally change based on the volume of service units you consume and,
often, between AWS Regions.
7.
Which of the following is a limitation of the AWS Simple Monthly Calculator?
A. You can calculate resource use for only one service at a time.
B. Not all AWS services are included.
C. The pricing is seldom updated and doesn’t accurately reflect current pricing.
D. You’re not able to specify specific configuration parameters.
7. B. You can, in fact, calculate costs for a multiservice stack. The calculator pricing is kept
up-to-date. You can specify very detailed configuration parameters.
8.
Which of the following Simple Monthly Calculator selections will likely have an impact on
most other configuration choices on the page? (Select TWO.)
A. Calculate By Month Or Year
B. Include Multiple Organizations
C. Free Usage Tier
D. Choose Region
8. C, D. Calculate By Month Or Year is not an option, and since the calculator calculates
only cost by usage, Include Multiple Organizations wouldn’t be a useful option.
9.
Which of the following is not an included parameter in the AWS Total Cost of Ownership
Calculator?
A. The tax implications of a cloud deployment
B. Labor costs of an on-premises deployment
C. Networking costs of an on-premises deployment
D. Electricity costs of an on-premises deployment
9. A. The calculator covers all significant costs associated with an on-premises deployment
but doesn’t include local or national tax implications.
10. Which of the following AWS Total Cost of Ownership Calculator parameters is likely to
have the greatest impact on cost?
A. Currency
B. AWS Region
C. Guest OS
D. Number of servers
10. D. The currency you choose to use will have little impact on price—it’s all relative, of
course. The guest OS and region will make a difference, but it’s relatively minor.
11. Which of the following AWS documentation URLs points to the page containing an up-to-
date list of service limits?
A. https://docs.aws.amazon.com/general/latest/gr/limits.html
B. https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html
C. https://aws.amazon.com/general/latest/gr/aws_service_limits.html
D. https://docs.aws.amazon.com/latest/gr/aws_service_limits.html
11. B. The correct URL is https://docs.aws.amazon.com/general/latest/gr/
aws_service_limits.html.
12. Which of the following best describes one possible reason for AWS service limits?
A. To prevent individual customers from accidentally launching a crippling level of
resource consumption
B. To more equally distribute available resources between customers from different
regions
C. To allow customers to more gradually increase their deployments
D. Because there are logical limits to the ability of AWS resources to scale upward
12. A. Resource limits exist only within individual regions; the limits in one region don’t
impact another. There’s no logistical reason that customers can’t scale up deployments
at any rate. There are, in fact, no logical limits to the ability of AWS resources to scale
upward.
13. Is it always possible to request service limit increases from AWS?
A. Yes. All service limits can be increased.
B. No. A limit can never be increased.
C. Service limits are defaults. They can be increased or decreased on demand.
D. No. Some service limits are hard.
13. D. While most service limits are soft and can be raised on request, there are some service
limits that are absolute.
14. Which is the best place to get a quick summary of this month’s spend for your account?
A. Budgets
B. Cost Explorer
C. Cost and usage reports
D. Billing & Cost Management Dashboard
14. D. The Cost Explorer and Cost and Usage Reports pages provide more in-depth and/or
customized details. Budgets allow you to set alerts based on usage.
15. What is the main goal for creating a Usage budget type (in AWS Budgets)?
A. To correlate usage per unit cost to understand your account cost efficiency
B. To track the status of any active reserved instances on your account
C. To track particular categories of resource consumption
D. To monitor costs being incurred against your account
15. C. Reservation budgets track the status of any active reserved instances on your account.
Cost budgets monitor costs being incurred against your account. There is no budget type
that correlates usage per unit cost to understand your account cost efficiency.
16. Which of the following is not a setting you can configure in a Cost budget?
A. Period (monthly, quarterly, etc.)
B. Instance type
C. Start and stop dates
D. Owner (username of resource owner)
16. D. You can configure the period, instance type, and start/stop dates for a budget, but you
can’t filter by resource owner.
17. What is the main difference between the goals of Cost Explorer and of cost and usage
reports?
A. Cost Explorer displays visualizations of high-level historical and current account costs,
while cost and usage reports generate granular usage reports in CSV format.
B. Cost and usage reports display visualizations of high-level historical and current
account costs, while Cost Explorer generates granular usage reports in CSV format.
C. Cost Explorer lets you set alerts that are triggered by billing events, while cost and
usage reports help you visualize system events.
D. Cost and usage reports are meant to alert you to malicious intrusions, while Cost
Explorer displays visualizations of high-level historical and current account costs.
17. A. Billing events aren’t triggers for alerts. Nothing in this chapter discusses intrusion events.
18. What is the purpose of cost allocation tags?
A. To associate spend limits to automatically trigger resource shutdowns when necessary
B. To help you identify the purpose and owner of a particular running resource to better
understand and control deployments
C. To help you identify resources for the purpose of tracking your account spending
D. To visually associate account events with billing periods
18. C. Tags are passive, so they can’t automatically trigger anything. Resource tags—not cost
allocation tags—are meant to help you understand and control deployments. Tags aren’t
associated with particular billing periods.
19. Which of the following scenarios would be a good use case for AWS Organizations?
(Select TWO.)
A. A single company with multiple AWS accounts that wants a single place to
administrate everything
B. An organization that provides AWS access to large teams of its developers and admins
C. A company that’s integrated some operations with an upstream vendor
D. A company with two distinct operational units, each with its own accounting system
and AWS account
19. A, C. Companies with multiple users of resources in a single AWS account would not
benefit from AWS Organizations, nor would a company with completely separated units.
The value of AWS Organizations is in integrating the administration of related accounts.
20. Which of these tools lets you design graphs within the browser interface to track your
account spending?
A. Budgets
B. Cost Explorer
C. Reports
D. Consolidating Billing

20. B. Budgets are used to set alerts. Reports provide CSV-formatted data for offline
processing. Consolidated Billing (now migrated to AWS Organizations) is for
administrating resources across multiple AWS accounts.

- Chapter3. Summary
AWS requires you to select one of four support plans. The Basic plan is free, Developer
starts at $29, Business starts at $100, and Enterprise starts at $15,000. Nonfree plans
(beyond the minimum charge) are billed as a percentage of total monthly resource usage.
The higher the support level, the more responsive and personalized the technical trouble-
shooting and planning support you get.

The Developer Support plan is aimed at organizations still testing and planning deploy-
ments. The Business Support plan is for smaller operations running relatively light produc-
tion infrastructure. The Enterprise Support plan is ideal for large production deployments
with a global footprint that cannot tolerate downtime.

AWS provides exhaustive and regularly updated documentation using multiple styles
across a number of web platforms. Those include user guide documentation pages, the
Knowledge Center, resources specific to security, and discussion forums.

The AWS Trusted Advisor alerts users to the best-practice compliance of their running
account resources. Basic Support and Developer Support plan users get service limit and
some security information, while Business and Enterprise customers get access to all alerts.

- Chapter3.  Exam Essentials
  - Know how to choose a support plan that reflects your operational needs. The more com-
plex and expensive your AWS deployments get, the more costly a configuration mistake
becomes. You can think about more expensive AWS support levels much the same way you
already think about hiring experienced and reliable admins. Whatever it takes to design
and deploy a lean and security-hardened application is a justifiable business expense.
  - Understand the benefits of the Enterprise Support plan’s technical account manager (TAM).
The ongoing, personalized attention your account deployments receive from a TAM can
make a significant difference in the quality of support. There’s nothing like having an
expert insider involved in the planning and execution of your complex infrastructure.
  - Understand how to find AWS resource usage guidance through official AWS documentation.
The AWS user guides are available in multiple formats (including HTML, PDF, Kindle,
and, on GitHub, MarkDown) and methodically explain practical usage for AWS services at
all levels. The Knowledge Center is a large collection of FAQs covering hundreds of com-
mon problems and their solutions.
  - Understand how to use Trusted Advisor for alerts to common system misconfigurations.
The Trusted Advisor alerts are divided into five categories: Cost Optimization, Performance,
Security, Fault Tolerance, and Services Limits. You should set up an administration routine
that includes regular visits to the Trusted Advisor to see whether any important status checks
have changed.


- Chapter3. Review Questions

1.
Your company is planning a major deployment on AWS. While the design and testing stages
are still in progress, which of the following plans will provide the best blend of support and
cost savings?
A. Basic
B. Developer
C. Business
D. Enterprise
1. C. The Basic plan won’t provide any personalized support. The Developer plan is cheaper,
but there is limited access to support professionals. The Business plan does offer 24/7 email,
chat, and phone access to an engineer, so until you actually deploy, this will make the most
sense. At a $15,000 monthly minimum, the Enterprise plan won’t be cost effective.
2.
Your web development team is actively gearing up for a deployment of an ecommerce site.
During these early stages of the process, individual developers are running into frustrating
conflicts and configuration problems that are highly specific to your situation. Which of the
following plans will provide the best blend of support and cost savings?
A. Basic
B. Developer
C. Business
D. Enterprise
2. B. Using the public documentation available through the Basic plan won’t be enough to
address your specific needs. The Business and Enterprise plans are not necessary as you
don’t yet have production deployments.
3.
Your corporate website was offline last week for more than two hours—which caused
serious consequences, including the early retirement of your CTO. Your engineers have
been having a lot of trouble tracking down the source of the outage and admit that they
need outside help. Which of the following will most likely meet that need?
A. Basic
B. Developer
C. Business
D. Enterprise
3. D. The lower three support tiers provide limited access to only lower-level support
professionals, while the Enterprise plan provides full access to senior engineers and
dedicates a technical account manager (TAM) as your resource for all your AWS needs.
4.
For which of the following will AWS provide direct 24/7 support to all users—even those
on the Basic Support plan?
A. Help with infrastructure under a massive denial-of-service (DoS) attack
B. Help with failed and unavailable infrastructure
C. Help with making a bill payment to AWS
D. Help with accessing your infrastructure via the AWS CLI
4. C. Basic plan customers are given customer support access only for account management
issues and not for technical support or security breaches.
5.
The primary purpose of an AWS technical account manager is to:
A. Provide 24/7 customer service for your AWS account
B.
Provide deployment guidance and advocacy for Enterprise Support customers
C. Provide deployment guidance and advocacy for Business Support customers
D. Provide strategic cost estimates for Enterprise Support customers
5. B. The TAM is available only for Enterprise Support customers. The primary function is
one of guidance and advocacy.

6.
Your Linux-based EC2 instance requires a patch to a Linux kernel module. The problem
is that patching the module will, for some reason, break the connection between your
instance and data in an S3 bucket. Your team doesn’t know if it’s possible to work
around this problem. Which is the most cost-effective AWS plan through which support
professionals will try to help you?
A. Developer.
B. Business.
C. Enterprise.
D. No plan covers this kind of support.
6. B. Only the Business and Enterprise plans include help with troubleshooting
interoperability between AWS resources and third-party software and operating systems.
The Business plan is the least expensive that will get you this level of support.
7.
Your company enrolled in the Developer Support plan and, through the course of one
month, consumed $4,000 USD of AWS services. How much will the support plan cost the
company for the month?
A. $120
B. $29
C. $100
D. $480
7. A. The Developer plan costs the greater of $29 or 3 percent of the monthly usage. In this
case, 3 percent of the month’s usage is $120.

8.
Your company enrolled in the Business Support plan and, through the course of three
months, consumed $33,000 of AWS services (the consumption was equally divided
across the months). How much will the support plan cost the company for the full
three months?
A. $4,000
B. $100
C. $1,100
D. $2,310
8. D. The Business plan—when monthly consumption falls between $10,000 and $80,000—
costs the greater of $100 or 7 percent of the monthly usage. In this case, 7 percent of a
single month’s usage ($11,000) is $770. The three month total would, therefore, be $2,310.
9.
Which of the following AWS support services does not offer free documentation of
some sort?
A. AWS Professional Services
B. The Basic Support plan
C. AWS Partner Network
D. The Knowledge Center
9. C. The AWS Professional Services site includes tech talk webinars, white papers, and blog
posts. The Basic Support plan includes AWS documentation resources. The Knowledge
Center consists of FAQ documentation.

10. What is the key difference between the roles of AWS Professional Services and a technical
account manager (TAM)?
A. The Professional Services product helps AWS Partner Network cloud professionals
work alongside your own team to help you administrate your cloud infrastructure. The
TAM is a cloud professional employed by AWS to guide you through the planning and
execution of your infrastructure.
B. The TAM is a cloud professional employed by AWS to guide you through the planning
and execution of your infrastructure. The Professional Services product provides cloud
professionals to work alongside your own team to help you administrate your cloud
infrastructure.
C. The TAM is a member of your team designated as the point person for all AWS
projects. The Professional Services product provides consultants to work alongside
your own team to help you administrate your cloud infrastructure.
D. The Professional Services product is a network appliance that AWS installs in your
data center to test cloud-bound workloads for compliance with best practices. The
TAM is a cloud professional employed by AWS to guide you through the planning and
execution of your infrastructure.
10. A. The TAM is an AWS employee dedicated to guiding your developer and admin teams.
There is no such thing as a network appliance for workload testing.
11. AWS documentation is available in a number of formats, including which of the following?
(Select TWO.)
A. Microsoft Word (DOC)
B. Kindle
C. HTML
D. DocBook
11. B, C. Although DOC and DocBook are both popular and useful formats, neither is used by
AWS for its documentation.
12. Which of the following documentation sites are most likely to contain code snippets for you
to cut and (after making sure you understand exactly what they’ll do) paste into your AWS
operations? (Select TWO.)
A. https://aws.amazon.com/premiumsupport/knowledge-center
B. https://aws.amazon.com/premiumsupport/compare-plans
C. https://docs.aws.amazon.com
D. https://aws.amazon.com/professional-services
12. A, C. The compare-plans page provides general information about support plans, and the
professional-services site describes accessing that particular resource. Neither directly
includes technical guides.
13. What is the primary function of the content linked from the Knowledge Center?
A. To introduce new users to the functionality of the core AWS services
B. To explain how AWS deployments can be more efficient and secure than on-premises
C. To provide a public forum where AWS users can ask their technical questions
D. To present solutions to commonly encountered technical problems using AWS
infrastructure
13. D. The Knowledge Center is a FAQ for technical problems and their solutions. The main
documentation site is much better suited to introduction-level guides. The https://
forums.aws.amazon.com site is the discussion forum for AWS users.
14. On which of the following sites are you most likely to find information about encrypting
your AWS resources?
A. https://aws.amazon.com/premiumsupport/knowledge-center
B.
https://aws.amazon.com/security/security-resources
C. https://docs.aws.amazon.com
D. https://aws.amazon.com/security/encryption
14. B. The Knowledge Center is a general FAQ for technical problems and their solutions.
The docs.aws.amazon.com site is for general documentation. There is no https://aws
.amazon.com/security/encryption page.
15. When using AWS documentation pages, what is the best way to be sure the information
you’re reading is up-to-date?
A. The page URL will include the word latest.
B. The page URL will include the version number (i.e., 3.2).
C. The page will have the word Current at the top right.
D. There is no easy way to tell.
15. A. Version numbers are not publicly available, and the word Current isn’t used in this context.
16. Which of the following is not a Trusted Advisor category?
A. Performance
B. Service Limits
C. Replication
D. Fault Tolerance
16. C. Replication is, effectively, a subset of Fault Tolerance and therefore would not require its
own category.
17. “Data volumes that aren’t properly backed up” is an example of which of these Trusted
Advisor categories?
A. Fault Tolerance
B. Performance
C. Security
D. Cost Optimization
17. A. Performance identifies configuration settings that might be blocking performance
improvements. Security identifies any failures to use security best-practice configurations.
Cost Optimization identifies any resources that are running and unnecessarily costing
you money.
18. Instances that are running (mostly) idle should be identified by which of these Trusted
Advisor categories?
A. Performance
B. Cost Optimization
C. Service Limits
D. Replication
18. B. Performance identifies configuration settings that might be blocking performance
improvements. Service Limits identifies resource usage that’s approaching AWS Region or
service limits. There is no Replication category.
19. Within the context of Trusted Advisor, what is a false positive?
A. An alert for a service state that was actually intentional
B. A green OK icon for a service state that is failed or failing
C. A single status icon indicating that your account is completely compliant
D. Textual indication of a failed state
19. A. An OK status for a failed state is a false negative. There is no single status icon
indicating that your account is completely compliant in Trusted Advisor.
20. Which of the following Trusted Advisor alerts is available only for accounts on the Business
or Enterprise Support plan? (Select TWO.)
A. MFA on Root Account
B. Load Balancer Optimization
C. Service Limits
D. IAM Access Key Rotation
20. B, D. Both the MFA and Service Limits checks are available for all accounts.




- Chapter4. Summary
An AWS Region connects at least two Availability Zones located within a single geographic
area into a low-latency network. Because of the default isolation of their underlying hard-
ware, building secure, access-controlled regional environments is eminently possible.
An Availability Zone is a group of one or more independent (and fault-protected) data
centers located within a single geographic region.
It’s important to be aware of the region that’s currently selected by your interface (either
the AWS Management Console or a command-line terminal), as any operations you execute
will launch specifically within the context of that region.
The design structure of Amazon’s global system of regions allows you to build your
infrastructure in ways that provide the best possible user experience while meeting your
security and regulatory needs.
AWS offers some global resources whose use isn’t restricted to any one region. Those
include IAM, CloudFront, and S3.
You can connect to AWS service instances using their endpoint addresses, which will
(generally) incorporate the host region’s designation.
EC2 virtual machine instances are launched with an IP address issued from a network
subnet that’s associated with a single Availability Zone.
The principle of high availability can be used to make your infrastructure more resilient
and reliable by launching parallel redundant instances in multiple Availability Zones.
AWS edge locations are globally distributed data servers that can store cached copies of
AWS-based data from which—on behalf of the CloudFront service—they can be efficiently
served to end users.
The elements of the AWS platform that you’re expected to secure and maintain and
those whose administration is managed by Amazon are defined by the AWS Shared
Responsibility Model.

- Chapter4. Exam Essentials
  - Understand the importance of resource isolation for cloud deployments. Properly placing your
cloud resources within the right region and Availability Zone—along with carefully setting
appropriate access controls—can improve both application security and performance.
  - Understand the role of autoscaling in a highly available deployment. The scalability of
AWS resources means you can automate the process of increasing or decreasing the scale of
a deployment based on need. This can automate application recovery after a crash.
  - Understand the role of load balancing in a highly available deployment. The ability to
automatically redirect incoming requests away from a nonfunctioning instance and to a
backup replacement is managed by a load balancer.
  - Understand the principles of the AWS Shared Responsibility Model. AWS handles secu-
rity and administration for its underlying physical infrastructure and for the full stack of
all its managed services, while customers are responsible for everything else.
  - Understand the principles of the AWS Acceptable Use Policy. Using AWS resources to
commit crimes or launch attacks against any individual or organization will result in
account suspension or termination.

- Chapter4. Review Questions

1.
Which of the following designations would refer to the AWS US West (Oregon) region?
A. us-east-1
B. us-west-2
C. us-west-2a
D. us-west-2b
1. B. The letter (a, b…) at the end of a designation indicates an Availability Zone. us-east-1
would never be used for a Region in the western part of the United States.
2.
Which of the following is an AWS Region for which customer access is restricted?
A. AWS Admin
B. US-DOD
C. Asia Pacific (Tokyo)
D. AWS GovCloud
2. D. The AWS GovCloud Region is restricted to authorized customers only. Asia Pacific
(Tokyo) is a normal Region. AWS Admin and US-DOD don’t exist (as far as we know, at
any rate).
3.
When you request a new virtual machine instance in EC2, your instance will automatically
launch into the currently selected value of which of the following?
A. Service
B. Subnet
C. Availability Zone
D. Region
3. D. EC2 instances will automatically launch into the Region you currently have selected.
You can manually select the subnet that’s associated with a particular Availability Zone for
your new EC2 instance, but there’s no default choice.
4.
Which of the following are not globally based AWS services? (Select TWO.)
A. RDS
B. Route 53
C. EC2
D. CloudFront
4. B, D. Relational Database Service (RDS) and EC2 both use resources that can exist in only
one Region. Route 53 and CloudFront are truly global services in that they’re not located in
or restricted to any single AWS Region.
5.
Which of the following would be a valid endpoint your developers could use to access a
particular Relational Database Service instance you’re running in the Northern Virginia
region?
A. us-east-1.amazonaws.com.rds
B. ecs.eu-west-3.amazonaws.com
C. rds.us-east-1.amazonaws.com
D. rds.amazonaws.com.us-east-1
5. C. The correct syntax for an endpoint is <service-designation>.<region-
designation> .amazonaws.com—meaning, in this case, rds.us-east-1.amazonaws.com.
6.
What are the most significant architectural benefits of the way AWS designed its regions?
(Select TWO.)
A. It can make infrastructure more fault tolerant.
B. It can make applications available to end users with lower latency.
C. It can make applications more compliant with local regulations.
D. It can bring down the price of running.
6. B, C. For most uses, distributing your application infrastructure between multiple AZs
within a single Region gives them sufficient fault tolerance. While AWS services do enjoy
a significant economy of scale—bring prices down—little of that is due to the structure of
their Regions. Lower latency and compliance are the biggest benefits from this list.
7.
Why is it that most AWS resources are tied to a single region?
A. Because those resources are run on a physical device, and that device must live
somewhere
B. Because security considerations are best served by restricting access to a single physical
location
C. Because access to any one digital resource must always occur through a single physical
gateway
D. Because spreading them too far afield would introduce latency issues
7. A. Sharing a single resource among Regions wouldn’t cause any particular security,
networking, or latency problems. It’s a simple matter of finding a single physical host device
to run on.
8.
You want to improve the resilience of your EC2 web server. Which of the following is the
most effective and efficient approach?
A. Launch parallel, load-balanced instances in multiple AWS Regions.
B. Launch parallel, load-balanced instances in multiple Availability Zones within a single
AWS Region.
C. Launch parallel, autoscaled instances in multiple AWS Regions.
D. Launch parallel, autoscaled instances in multiple Availability Zones within a single
AWS Region.
8. B. Auto Scaling is an important working element of application high availability, but it’s
not what most directly drives it (that’s load balancing). The most effective and efficient
way to get the job done is through parallel, load-balanced instances in multiple Availability
Zones, not Regions.
9.
Which of the following is the most accurate description of an AWS Availability Zone?
A. One or more independently powered data centers running a wide range of hardware
host types
B. One or more independently powered data centers running a uniform hardware host
type
C. All the data centers located within a broad geographic area
D. The infrastructure running within a single physical data center
9. A. “Data centers running uniform host types” would describe an edge location. The data
centers within a “broad geographic area” would more closely describe an AWS Region. AZs
aren’t restricted to a single data center.
10. Which of the following most accurately describes a subnet within the AWS ecosystem?
A. The virtual limits imposed on the network access permitted to a resource instance
B. The block of IP addresses assigned for use within a single region
C. The block of IP addresses assigned for use within a single Availability Zone
D. The networking hardware used within a single Availability Zone
10. C. Imposing virtual networking limits on an instance would be the job of a security group
or access control list. IP address blocks are not assigned at the Region level. Customers have
no access to or control over AWS networking hardware.
11. What determines the order by which subnets/AZ options are displayed in EC2 configura-
tion dialogs?
A. Alphabetical order
B. They (appear) to be displayed in random order.
C. Numerical order
D. By order of capacity, with largest capacity first
11. B. AWS displays AZs in (apparently) random order to prevent too many resources from
being launched in too few zones.

12. What is the primary goal of autoscaling?
A. To ensure the long-term reliability of a particular physical resource
B. To ensure the long-term reliability of a particular virtual resource
C. To orchestrate the use of multiple parallel resources to direct incoming user requests
D. To ensure that a predefined service level is maintained regardless of external demand
or instance failures
12. D. Auto Scaling doesn’t focus on any one resource (physical or virtual) because it’s
interested only in the appropriate availability and quality of the overall service. The job of
orchestration is for load balancers, not autoscalers.
13. Which of the following design strategies is most effective for maintaining the reliability of a
cloud application?
A. Resource isolation
B. Resource automation
C. Resource redundancy
D. Resource geolocation
13. C. Resource isolation can play an important role in security, but not reliability. Automation
can improve administration processes, but neither it, nor geolocation, is the most effective
reliability strategy.
14. Which of the following AWS services are not likely to benefit from Amazon edge locations?
(Select TWO.)
A. RDS
B. EC2 load balancers
C. Elastic Block Store (EBS)
D. CloudFront
14. A, C. RDS database instances and Lambda functions are not qualified CloudFront origins.
EC2 load balancers can be used as CloudFront origins.
15. Which of the following is the primary benefit of using CloudFront distributions?
A. Automated protection from mass email campaigns
B. Greater availability through redundancy
C. Greater security through data encryption
D. Reduced latency access to your content no matter where your end users live
15. D. CloudFront can’t protect against spam and, while it can complement your application’s
existing redundancy and encryption, those aren’t its primary purpose.
16. What is the main purpose of Amazon Route 53?
A. Countering the threat of distributed denial-of-service (DDoS) attacks
B. Managing domain name registration and traffic routing
C. Protecting web applications from web-based threats
D. Using the serverless power of Lambda to customize CloudFront behavior
16. B. Countering the threat of DDoS attacks is the job of AWS Shield. Protecting web
applications from web-based threats is done by AWS Web Application Firewall. Using
Lambda to customize CloudFront behavior is for Lambda Edge.
17. According to the AWS Shared Responsibility Model, which of the following are responsi-
bilities of AWS? (Select TWO.)
A. The security of the cloud
B. Patching underlying virtualization software running in AWS data centers
C. Security of what’s in the cloud
D. Patching OSs running on EC2 instances
17. A, B. What’s in the cloud is your responsibility—it includes the administration of
EC2-based operating systems.
18. According to the AWS Shared Responsibility Model, what’s the best way to define the sta-
tus of the software driving an AWS managed service?
A. Everything associated with an AWS managed service is the responsibility of AWS.
B. Whatever is added by the customer (like application code) is the customer’s
responsibility.
C. Whatever the customer can control (application code and/or configuration settings) is
the customer’s responsibility.
D. Everything associated with an AWS managed service is the responsibility of the
customer.
18. C. There’s no one easy answer, as some managed services are pretty much entirely within
Amazon’s sphere, and others leave lots of responsibility with the customer. Remember, “if
you can edit it, you own it.”

19. Which of the following is one of the first places you should look when troubleshooting a
failing application?
A. AWS Acceptable Use Monitor
B. Service Status Dashboard
C. AWS Billing Dashboard
D. Service Health Dashboard
19. D. The AWS Billing Dashboard is focused on your account billing issues. Neither the AWS
Acceptable Use Monitor nor the Service Status Dashboard actually exists. But nice try.
20. Where will you find information on the limits AWS imposes on the ways you can use your
account resources?
A. AWS User Agreement Policy
B. AWS Acceptable Use Policy
C. AWS Acceptable Use Monitor
D. AWS Acceptable Use Dashboard
20. B. The correct document (and web page https://aws.amazon.com/aup/) for this
information is the AWS Acceptable Use Policy.
