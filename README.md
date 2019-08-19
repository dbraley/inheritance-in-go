# inheritance-in-go

## The Problem

Imagine the following database:

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

###### Ip_Credentials Table
| credential_id | ip              |
|---------------|-----------------|
| 101           | 192.168.0.27    |

Now, lets imagine we want a service (possibly an endpoint) that can list all the info we know about an account based on some query. Perhaps lets imagin a client that is going to try to print out all associated usernames and ip addresses of admin accounts. SECURITY AUDIT FTW!

## Potential Solutions

(Here will be ~~Dragons~~ links to branches with solutions and notes)
