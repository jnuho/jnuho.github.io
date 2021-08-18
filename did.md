
### 디지털 신원인증 모델
- Centralized
  - e.g. Account-based website
  - [You] &rarr; [Org]

- Federated
  - service/identity provider (IDP) in the middle
  - The collection of all the sites that use the same IDP (or group of IDPs) is called a federation.
  - [You] &rarr; [IDP] &rarr; [Org]

- **Decentralized**
  - A new model, inspired by blockchain technology, since 2015
  - peer-to-peer
  - _```blockchain```_ is used based on _```public/private key cryptography```_ which secures direct and private connections
  - 블록체인 기술을 암호화폐가 아닌, DPKI (Decentralized PKI)에 적용
    - 공개키를 직접 교환 하여 거래 당사자(peer) 간에 private and secure 연결 생성 가능
    - 공개키를 공개 블록체인에 저장하여 디지털 신원 자격(VC)에 있는 서명 증명
      - VC(verifiable credentials): 실생활에서 신원증명 제공을 위해 교환 가능한 자격 증명
      - e.g. DID의 개인키는 디지털지갑, 공개키는 블록체인에 저장

<div>
<img src="https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_1-4.png"
height="20%" width="20%" alt="Figure1.4">
</div>

### Basic cryptography techniques for SSI
- Asymmetric-key cryptography
  - e.g. DID의 개인키는 디지털지갑, 공개키는 블록체인에 저장
- Digital signatures
  - rely on _public-key cryptography_
  - created using the private key of a key pair, and verified using the associated public key. 


### Decentralized identifiers (DIDs)

- The IP address itself doesn’t tell you anything about the identity holder.
- _public/private key cryptography_ provided digital proof that is verifiable
  - message is signed with private key and verified using public key
  - verifier must know the correct public key of the owner : ```PKI``` as a solution

- ***PKI*** 공개 키 기반 구조 - 중앙 집권형
  - proves _```ownership of their public keys.```_
  - obtain public key certificates from CAs around the world (_very centralized._)
  - 개인들이 여러개의 cryptographic 키set을 가지고 있는 SSI환경에서 쓰이기에는 한계 있음

- 분산 신원인증 DIDs
  - permanent
  - resolvable
  - cryptographically verifiable (identity holder can cryptographically prove using pk)
  - decentralized (decentralized network such as blockchain instead of relying on CAs)

<div>
<img src="https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_2-11.png"
height="35%" width="35%" alt="Figure2.11">

A DID functions as the address of a public key on a blockchain.
DID can also be used to locate an agent for the DID subject (the entity identified by the DID).
It is designed to take advantage of blockchain, DLT, etc, via a DID method
The DID method defines four atomic operations on any DID: CRUD 
</div>

<!-- ### 8 Decentralized Identifiers -->

- DIDs are the cryptographic counterpart to verifiable credentials (VCs)
- A new type of globally unique identifier

<div>
<img src="https://drek4537l1klr.cloudfront.net/preukschat/HighResolutionFigures/figure_8-2.png"
height="30%" width="30%" alt="Figure8.2">
</div>

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
