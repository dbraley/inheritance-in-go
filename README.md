# inheritance-in-go

## The Problem

Imagine we have some persisting data stor in which we have user accounts, as well as some sets of credentails which give our users access to the accounts. We know that each credential belongs to exactly one account, and that any account could have zero or more credentails of any type. We could model it with something like (but not necessarily) the following database:

###### Accounts Table
| account_id | name | is_admin |
|------------|------|----------|
| 1          | Dave | true     |
| 2          | Bob  | false    |

###### Credentials Table
| credential_id | account_id | type      |
|---------------|------------|-----------|
| 100           | 1          | USER_PASS |
| 101           | 1          | IP        |
| 102           | 2          | USER_PASS |

###### User_Pass_Credentials Table
| credential_id | username | password |
|---------------|----------|----------|
| 100           | Dave     | p@zzwVrd |
| 102           | Bob      | !!!y0!!! |

__Please do not store passwords in plaintext. This is a stupid example, do NOT do this__

###### Ip_Credentials Table
| credential_id | ip              |
|---------------|-----------------|
| 101           | 192.168.0.27    |

Now, lets imagine we want a service (possibly an endpoint) that can list all the information we know about accounts based on some query or another. Perhaps lets imagine a client that is going to try to print out all associated usernames and IP addresses of admin accounts. SECURITY AUDIT FTW!

## Goal
Understand some of the patterns we can use to model this, and how they compare to each other.

## Potential Solutions

(Here will be ~~Dragons~~ links to branches with solutions and notes)
