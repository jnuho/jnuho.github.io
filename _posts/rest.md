
### Restful API

- Representational State Transfer (REST):
    - REST is an architectural style for designing networked applications, particularly web services and APIs.
    - Key points about REST:
        - Definition: REST is an acronym for REpresentational State Transfer.
        - Origin: It was introduced by computer scientist Roy Fielding in his dissertation in the year 2000.
        - Purpose: REST provides a set of principles and constraints for building distributed hypermedia systems.
        - Widely Used: It has become one of the most popular approaches for creating web-based APIs.
        - Not a Protocol: REST is not a protocol or a standard; rather, it's an architectural style.
        - Uniform Interface: REST defines a consistent and uniform interface for interactions between clients and servers.
        - HTTP Methods: HTTP-based REST APIs use standard HTTP methods (such as GET, POST, PUT, DELETE) and URIs (Uniform Resource Identifiers) to identify resources.
        - Guiding Principles: RESTful services adhere to specific guiding principles, which must be satisfied for a service to be considered RESTful.
- Statelessness in REST:
    - Statelessness is one of the key architectural constraints of REST.
    - It recommends that all interactions between clients and servers should be stateless.
    - Here's what it means:
        - Client Responsibility: In a stateless REST API, the server does not store any information about the client's previous requests.
        - Each Request Is New: Every HTTP request to a RESTful service is treated as a new request, regardless of any prior interactions.
        - Client Provides Context: Clients are responsible for providing all necessary information in each request. This includes authentication tokens, credentials, or any other context data.
        - No Server-Side Sessions: The server does not establish or maintain client sessions. There is no session affinity or sticky session.
        - Advantages:
            - Scalability: Stateless APIs are easier to scale because servers don't need to manage client-specific state.
            - Caching: Stateless design allows for efficient caching of responses.
            - Simplicity: Stateless interactions simplify the overall architecture.
- Application State vs. Resource State:
    - It's essential to differentiate between these two concepts:
        - Application State: This is server-side data that servers store to identify incoming client requests, their previous interactions, and current context information.
        - Resource State: Resource state refers to the current state of a resource on the server at any given time. It's what the server provides as a response (resource representation) to API requests.
    - REST statelessness specifically refers to being free from the application state, allowing each HTTP request to happen in complete isolation.
In summary, RESTful APIs follow the principles of statelessness, making them efficient, scalable, and straightforward. Clients provide all necessary information, and servers do not store any client-specific session data.

1. When a RESTful API receives a request (whether it’s a GET, POST, PUT, or DELETE request), the server processes it.

2. If the server intends for the response to be cached, it includes appropriate Cache-Control directives in the response headers.

3.The client (browser) then adheres to these directives:
  - If the response is cacheable (based on Cache-Control), the client stores a local copy in its cache.
  - For subsequent requests to the same resource, the client checks its cache first. If the cached copy is still valid (not expired), it uses it.
  - If the cached copy is stale or nonexistent, the client fetches the resource from the server.


### Cache-Control Headers:

Cache-Control is an HTTP header used to specify caching policies for both client requests and server responses.

These headers dictate how resources should be cached by browsers, proxies, and other intermediaries.

The Cache-Control header allows you to control caching behavior, including where resources are cached and their maximum age before expiring.