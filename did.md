
## 1 Three models of digital identity
- The centralized identity model
  - e.g. Account-based website
  - [You] &ndash;account&rarr; [Org]


- The federated identity model
  - The basic idea is to insert a service provider called an identity provider (IDP) in the middle
  - e.g. SAML, OAuth, OpenID Connect, SSO
  - [You] &ndash;account&rarr; [IDP] &rarr; [Org]


- **The decentralized identity model**
  - A new model, inspired by blockchain technology, surfaced in 2015
  - DIDs as a new decentralized identity standard.
  - Not account-based
  - peer-to-peer relationship, neither of you has an 'account'; both share a 'connection'

![Figure1.4](https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_1-4.png)


- Blockchain for decentralized identity model

```
Public/private key cryptography is a way of securing data via mathematical
algorithms based on cyptographic keys held by each party.

Instead of using blockchain technology for cryptocurrency,
identity management uses it for decentralized public key infrastructure (DPKI).

In essence, blockchain technology and other decentralized network technologies can give us a strong,
decentralized solution for
1. Exchanging public keys directly to form private, secure connections between any two peers
2. Storing some of these public keys on public blockchains to verify the signatures on digital identity credentials (aka verifiable credentials) that peers can exchange to provide proof of real-world identity
```

## 2 The basic building blocks of SSI
ighResolutionFigures/figure_2-11.png)



## 8 Decentralized Identifiers

A new type of globally unique identifier
![figure8.2](https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_8-2.png)

A DID is a URI that can be either a URL or a URN and that can be looked up (resolved)
to get a standardized set of information (metadata) about the resource identified by the DID.
If the identified resource has one or more representations on the web, the metadata can include one or more of those URLs.
Four properties of DID:
1. Permanent
2. Resolvable
3. Cryptographically verifiable : 
  cryptography is used to generate the DID.
  since the DID is now associated with exactly one public/private key pair,
  the controller of the private key can prove that they are also the controller of the DID.

4. Decentralized : 
  cryptography eliminates the need for centralized registration authorities
  By contrast, cryptographic algorithms for public/private key pairs
  are based on random number generators, large prime numbers, elliptic curves,
  or other cryptographic techniques for producing globally unique values
  that do not require a central registry to effectively guarantee uniqueness

```
As a result, anyone with the proper software can generate a DID
according to a particular DID method and begin using it immediately
without requiring the authorization or involvement of any centralized registration authority.
This is the same process used to create public addresses on the Bitcoin or Ethereum
(or other popular) blockchains—it is the essence of what makes a DID decentralized.
```

8.2 The functional level: How DIDs work

8.2.1 DID documents

Although it is not yet possible to type a DID into a web browser and have it do anything meaningful,
you can give DID to a specialized piece of software (or hardware) called a "DID resolver"
that will use it to retrieve a standardized data structure called a "DID document"
it is a machine-readable document designed to be consumed by digital identity applications
or services such as digital wallets, agents, or encrypted data stores,
all of which use DIDs as fundamental building blocks.

Every DID has exactly one associated DID document.
The DID document contains metadata about the DID subject,
which is the term for the resource identified by the DID and described by the DID document
For example, a DID for a person (the DID subject) has an associated DID document
that typically contains cryptographic keys, authentication methods,
and other metadata describing how to engage in trusted interactions with that person. 

![Figure8.6](https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_8-6.png)

DID document contains :
- Public keys that can be used to authenticate the DID subject during an interaction; the essence of the DPKI
- Services associated with the DID subject that can be used for concrete interaction via protocols
- Certain additional metadata

```json
{
  "@context": "https://www.w3.org/ns/did/v1",
  "id": "did:example:123456789abcdefghi",
  "authentication": [{
    "id": "did:example:123456789abcdefghi#keys-1",
    "type": "Ed25519VerificationKey2018",
    "controller": "did:example:123456789abcdefghi",
    "publicKeyBase58" : "H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"
  }],
  "service": [{
    "id":"did:example:123456789abcdefghi#vcs",
    "type": "VerifiableCredentialService",
    "serviceEndpoint": "https://example.com/vc/"
  }]
}
```

8.2.2 DID methods

Each DID method is required to have its own technical specification,
which must define the following aspects of the DID method:
- Method-specific identifier (sov,btcr,v1,ethr,jolo,...)
- Four basic operations can be executed on a DID: CRUD
- Security and privacy considerations specific to the DID method

![Figure8.7](https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_8-7.png)

It is difficult to make generic statements about the four DID operations
since DID methods can be designed in very different ways.
For example, some DID methods are based on blockchains or other distributed ledgers.
In this case, creating or updating a DID typically involves writing a transaction to that ledger.
Other DID methods do not use a blockchain; they implement the four DID operations in other ways (see section 8.2.7).


8.2.3 DID resolution
The process of obtaining the DID document associated with a DID


![Figure8.8](https://drek4537l1klr.cloudfront.net/preukschat/Figures/CH08_F08_Preukschat.png)
![Figure8.9](https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_8-9.png)


8.2.4 DID URLs
