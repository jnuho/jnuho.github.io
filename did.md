
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
  - The peer-to-peer relationship enabled by the decentralized identity model—returning people to direct, private connections secured by public/private key cryptography
  - Neither of you has an 'account'; both share a 'connection'
<div>
<img src="https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_1-4.png" height="20%" width="20%" alt="Figure1.4">
</div>

- **Why does it need Blockchain technology?**
  - Public/private key cryptography is a way of securing data via mathematical
algorithms based on cyptographic keys held by each party.
  - Instead of using blockchain technology for cryptocurrency, identity management uses it for decentralized public key infrastructure (DPKI).

  - Blockchain technology and other decentralized network technologies can give us a strong, decentralized solution for
    - Exchanging public keys directly to form private, secure connections between any two peers
    - Storing some of these public keys on public blockchains to verify the signatures on digital identity credentials (aka verifiable credentials) that peers can exchange to provide proof of real-world identity

## 2 The basic building blocks of SSI

SSI is a set of technologies that build on core concepts in identity management, distributed computing, blockchain or distributed ledger technology (DLT), and cryptography.

- 2.1 Verifiable credentials (aka digital credentials)
- 2.2 The trust triangle: issuers, holders, and verifiers
- 2.3 Digital wallets
- 2.4 Digital agents

- 2.5 Decentralized identifiers (DIDs)

Knowing the IP address of a machine on the internet doesn’t tell you anything about the identity of the person, organization, or thing controlling that machine.
To do that, the controller (the identity holder) needs to be able to provide proof about their identity, attributes, relationships, or entitlements. And that proof has to be verifiable in some way.

We’ve had technology for creating digital proofs for decades now : **Public/private key cryptography**
```
The owner of a private key uses it to sign messages, and anyone else can verify this signature
using the owner’s corresponding public key.
The signature verification shows that the signature was created by the owner of
the private key and the message has not been tampered with since.
```

However, to rely on this verification, the verifier must know the correct public key for the owner.
So, for decentralized messaging between digital agents and wallets to be secure
—and for agents to be able to send cryptographically verifiable proofs of VCs to each other
—we need a strong, secure, scalable way for identity holders and their agents to prove ownership of their public keys.

The solution has been public key infrastructure (PKI)
: The system of obtaining public key certificates from a small set of certification authorities (CAs) around the world, which is also very centralized.
=> DIDs, a new identifier (permanent, resolvable, cryptographically verifiable, decentralized)

An example of a decentralized identifier (DID) and the associated public and private keys.
A DID functions as the address of a public key on a blockchain (or other decentralized network.)
In most cases, a DID can also be used to locate an agent for the DID subject (the entity identified by the DID).

<div>
<img src="https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_2-11.png"
height="50%" width="50%" alt="Figure2.11">
</div>



## 8 Decentralized Identifiers

- A new type of globally unique identifier

![figure8.2](https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_8-2.png)

A DID is a URI that can be either a URL or a URN and that can be looked up (resolved)
to get a standardized set of information (metadata) about the resource identified by the DID.
If the identified resource has one or more representations on the web, the metadata can include one or more of those URLs.
Four properties of DID:
1. Permanent
2. Resolvable
3. Cryptographically verifiable Identifier : You can prove control using cryptography
```
  cryptography is used to generate the DID.
  since the DID is now associated with exactly one public/private key pair,
  the controller of the private key can prove that they are also the controller of the DID.

```
4. Decentralized : No centralized registration autority is required
```
  cryptography eliminates the need for centralized registration authorities
  cryptographic algorithms for public/private key pairs
  are based on random number generators, large prime numbers, elliptic curves,
  or other cryptographic techniques for producing globally unique values
  that do not require a central registry to effectively guarantee uniqueness
```

As a result, anyone with the proper software can generate a DID according to a particular DID method and begin using it immediately without requiring the authorization or involvement of any centralized registration authority.

This is the same process used to create public addresses on the Bitcoin or Ethereum (or other popular) blockchains—it is the essence of what makes a DID decentralized.


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

<div>
<img src="https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_8-6.png"
height="70%" width="70%" alt="Figure8.6">
</div>

DID document contains :
- Public keys that can be used to authenticate the DID subject during an interaction; the essence of the DPKI
- Services associated with the DID subject that can be used for concrete interaction via protocols
- Certain additional metadata such as timestamps, digital signatures, and other cryptographic proofs, or metadata related to delegation and authorization.

```json
{
  // The first line is the JSON-LD context statement,
  // required in JSON-LD documents (but not in other DID document representations).
  "@context": "https://www.w3.org/ns/did/v1",
  // DID being described
  "id": "did:example:123456789abcdefghi",
  // public key for authenticating the DID subject.
  "authentication": [{
    "id": "did:example:123456789abcdefghi#keys-1",
    "type": "Ed25519VerificationKey2018",
    "controller": "did:example:123456789abcdefghi",
    "publicKeyBase58" : "H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"
  }],
  // service endpoint for exchanging verifiable credentials.
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

<div>
<img src="https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_8-7.png" height="70%" width="70%" alt="Figure8.7">
</div>

It is difficult to make generic statements about the four DID operations since DID methods can be designed in very different ways.
For example, some DID methods are based on blockchains or other distributed ledgers.
In this case, creating or updating a DID typically involves writing a transaction to that ledger.
Other DID methods do not use a blockchain; they implement the four DID operations in other ways (see section 8.2.7).

“Rubric” document to help adopters evaluate how well a particular DID method will meet the needs of a particular user community: https://w3c.github.io/did-rubric

8.2.3 DID resolution
The process of obtaining the DID document associated with a DID

<div>
<img
src="https://drek4537l1klr.cloudfront.net/preukschat/Figures/CH08_F08_Preukschat.png"
height="70%" width="70%" alt="Figure8.8">
</div>

<br>
<div>
<img
src="https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_8-9.png"
height="50%" width="50%" alt="Figure8.9">
</div>



8.2.4 DID URLs
