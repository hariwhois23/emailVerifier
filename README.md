# Mail Domain DNS Checker 

A simple tool written in Go that inspects email domain configurations by checking for:

* MX (Mail Exchange) records
* SPF (Sender Policy Framework) records
* DMARC (Domain-based Message Authentication, Reporting & Conformance) records

It helps quickly validate whether a domain is properly configured for sending and receiving emails.

###  The program uses standard `net` package:

* `net.LookupMX()` → fetches mail servers
* `net.LookupTXT()` → fetches TXT records (used for SPF & DMARC)


## What it does ? 

* Reads domain input from standard input (stdin)
* Checks:

  * MX records (mail servers)
  * SPF records (email authorization policy)
  * DMARC records (email authentication & policy)
* Outputs structured results in terminal

##  DNS Records Explained

### MX (Mail Exchange) Records

* Define which mail servers receive emails for a domain

  ```shell
  gmail.com → alt1.gmail-smtp-in.l.google.com
  ```
* Without MX records → domain **cannot receive emails**

### SPF (Sender Policy Framework)

* A TXT record that specifies **which servers can send emails** on behalf of a domain
* Helps prevent spoofing

```shell
v=spf1 include:_spf.google.com ~all
```

### DMARC (Domain-based Message Authentication)

* Builds on SPF & DKIM
* Defines what to do if authentication fails

```shell
v=DMARC1; p=reject; rua=mailto:admin@domain.com
```

## TO KNOW

* Input should be a **domain name**, not a full email address
  > NO `user@gmail.com`
  > YES `gmail.com`
* `ctrl + c` to terminate the running program.
* Network/DNS errors may occur if:
  * Domain does not exist
  * DNS records are missing

 ## Example Output

``` shell
example.com
--------------------------------------------------------------
mail |  hasMX |  hasSPF |  spfRecord | hasDMARC |  dmarcRecord
--------------------------------------------------------------
| MX Record    | Status: true  | Records: [...]
| SPF Record   | Status: true  | Value: v=spf1 include:_spf.google.com ~all
| DMARC Record | Status: true  | Value: v=DMARC1; p=none; ...
```



