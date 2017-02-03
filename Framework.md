# The C4 framework
C4 the Cinema Content Creation Cloud is an open source framework for content
creation that makes working with cloud, substantially easer.

C4 solves very hard problems with media production using widely distributed computing and storage resources that are constantly changing. These resources are often connected unreliably or even exist entirely off-line exchanging data only via physical storage.

C4 consists of 3 major pieces: An **identification system**, and **domain language**, and a public key cryptography based **security model**.

Key benefits include:

- Dead simple.
- Hard to break.
- Fully asynchronous.
- Supports incremental adoption.

## C4 ID - The C4 Identification System
[C4 ID Whitepaper](http://www.cccc.io/downloads/c4id_latest.pdf)

The C4 ID system provides an unambiguous, universally unique id for any file or block of data. It is not only universally unique it is also universally consistent. The means the same data has the same C4 ID regardless of who, where or when the data is identified.

C4 does this without needing centralized coordination. C4 does not require any network connect to produce correct and consistent C4 IDs of any data.  Nevertheless, users in different locations without knowledge of each other will always produce the same C4 ID for a given piece of data.

This allows organizations to communicate with others about assets consistently and unambiguously, while managing assets internally in anyway they choose.  

This 'agreement without communication' is an essential feature of C4 IDs, and enables the rest of the framework's powerful tools. 

C4 IDs have a structure that is easy for humans to recognize and for machines to read.

C4 IDS: 

- **start with `c4`**
- **are 90 characters long**
- **are safe in urls, filenames, and database records**
- **have no special, i.e. non-alphanumeric**
- **are easily and unambiguously selected by double clicking**

Here is an example of a C4 ID:

```
c44jVTEz8y7wCiJcXvsX66BHhZEUdmtf7TNcZPy1jdM6S14qqrzsiLyoZRSvRGcAMLnKn4zVBvAFimNg14NFKp46cC
```

C4 IDs are needed because there are no existing standards for identification of data that is a pure non-assigned representation of identity.

There are no universal standard encodings for common cryptographic hashes.  Labeled hex representations (i.e. sha512-cf83...) seem to be a tad more popular at the moment, so those are used for comparison below.  Note that c4 is a sha-512, yet it's only 19 characters longer than a hex encoded sha-256, and also faster to compute (on 64 bit hardware).

```yaml
# Comparison
sha-256: sha256-e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
c4:      c44jVTEz8y7wCiJcXvsX66BHhZEUdmtf7TNcZPy1jdM6S14qqrzsiLyoZRSvRGcAMLnKn4zVBvAFimNg14NFKp46cC
sha-512: sha512-cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e
```

It might seem like Base64 would make a shorter string, and therefore be a better encoding, but Base64 only saves 4 characters, and only if we cheat and remove the label.

```yaml
# Base64 vs C4

c4:      c45XyDwWmrPQwJPdULBhma6LGNaLghKtN7R9vLn2tFrepZJ9jJFSDzpCKei11EgA5r1veenBu3Q8qfvWeDuPc7fJK2
Sha-512: 9/u6bgY2+JDlb7vzKD5STG+jIErimDgtYkdB0NxmODJuKCxBvl5CVNiCB3LFUYosWowMf37aGVlKfrU5RT4e1w
```

The C4 ID, includes a label and is easily selectable by double clicking.  If you double click on the above base64 encoding you'll see that you don't select the entire string.

C4 IDs also lead to the ability to construct reliable deterministic dependency graphs for computing.

## C4 Lang - Dependency Oriented Domain Language

*(Whitepaper coming soon)*

The C4 Domain Specific Language (DSL) is a declarative language that is designed to represent a dependency graph of operations that are repeatable and verifiable. C4lang can describe processes that span any number of physical domains, making it much easer to design and reason about distributed workflows.

All data in c4lang is immutable, operations are idempotent.  With these constraints, and the cashing of results of any non-deterministic processes, a given c4lang graph node will eventually be reduced to a fixed and immutable result.  This provides "strong eventual consistency" and abstracts compute vs storage making them identical as far as the language is concerned. This fungibility of computation and storage is a very powerful property of the language and reduces much of the complexity of distributed computation for media production.


## C4 PKI - Public Key Cryptography Security Model

*(Whitepaper in the works)*

Under the C4 Public Key Infrastructure model there are no logins (other than a user on their own device). Identity is automatically federated without the need for an "Identity Provider" provider (i.e. the OAuth model).  Instead a standard x.509 certificate chain is used to validate *both* sides of all communications.  This system works automatically via strong cryptography. It even works in off-line environments such as productions in remote locations.

x.509 has a much longer history than OAuth, and is a well vetted component of standard secure web traffic. OAuth is a system designed around the idea that some identity providers want to 'own' a user's account. In media production, however, a more robust system that does not require a trusted 3rd party is required.


### History
Originally designed in 2008, and subsequently refined over a number of years, it has proven to be a generally useful set of tools for any production work flow.  It 2013 Studio Pyxis begin publishing portions of the framework as open source.

Since 2013 In collaboration with research conducted on behave of the major studios and vendors in Hollywood C4 has gone through and is continuing to go through a number of standards processes  

