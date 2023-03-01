

- [Network layer](https://www.cloudflare.com/learning/network-layer/what-is-the-network-layer/)
  - [How does the internet works?](https://www.cloudflare.com/learning/network-layer/how-does-the-internet-work/)
  - [Go network](https://github.com/gyuho/learn/tree/main/doc/go_network)
  - [Router](https://www.cloudflare.com/learning/network-layer/what-is-a-router/)
    - A network switch forwards data packets between groups of devices in the same network,
    - whereas a router forwards data between different networks.

### OSI Reference Model

- The seven layers of OSI reference model are as follows:

```
L7 : 응용(Application) 계층 -> HTTP/FTP 프로토콜 통신
...
L2 : Data Link 계층
L1 : 물리(Physical) 계층

- 낮은 계층 일수록 물리적, 높은 계층 일수록 논리적

```

- 7. Application
  - HTTP, FTP, IRC, SSH, DNS
  - End user layer
  - 네트워크 어플리케이션 및 라이브러리들이 관계하는 레이어. 최종목적지로, 응용프로세스와 직접 관계하여 일반적인 응용 서비스를 수행한다
  - 사용자 인터페이스, 전자우편, 데이터베이스 관리 등의 서비스를 제공한다.
- 6. Presentation - SSL, SSH, IMAP, FTP, MPEG, JPEG
  - 데이터 표현에 대한 독립성을 제공하고, 암호화하는 역할. 파일인코딩, 명령어를 포장, 압축, 암호화 한다.
- 5. Session
  - Synch & send to port
  - API's, Sockets, WinSock
  - 네트워크의 노드들 간 커넥션 라이프사이클 관리
- 4. Transport
  - End-to-end connections
  - TCP (신뢰성, 연결지향적), UDP (비신뢰성, 비연결성, 실시간)
  - 데이터 전송의 안전성을 유지하면서 노드들간에 데이터전송을 제어하고 조율함
  - 데이터 전송의 안정성 유지: 에러 바로잡기, 데이터 전송 속도 제어, 데이터 조각내기, 누락된 데이터 재전송, 전달받은 데이터 확인 등이 있다.
- 3. Network
  - IP, ICMP, IPSec, IGMP, 라우터
  - Packets
  - 데이터를 목적지 까지 가장 안전하고 빠르게 전달
  - 라우터를 통해 이동할 경로르 선택하여 IP주소를 지정하고, 해당경로에 따라 Packet을 전달
  - 라우팅, 흐름제어, 오류제어, 세그먼테이션 등 수행
- 2. Data Link
  - Ethernet, PPP, Switch, Bridge
  - 물리계층으로 송수신되는 정보를 관리하여 안전하게 전달되도록 도움. Mac주소를 통해 통신
  - 프레임에 Mac 주소를 부여하고 에러검출, 재전송, 흐름제어
- 1. Physical
  - Coax, Fiber, Wireless, Hubs, Repeaters
  - 데이터 전기적인 신호로 변환해서 주고받는 기능하는 공간. 데이터 전송하는 역할만 진행



### TCP/IP 4-layer

- 애플리케이션 계층
  - 유저에게 제공되는 애플리케이션에서 사용하는 통신의 움직임을 결정
  - FTP, DNS, HTTP와 같은 공통 애플리케이션이 준비되어 있음
- 트랜스포트 계층
  - 애플리케이션 계층에 네트워크로 접속되어 있는 2대의 컴퓨터 사이 데이터 흐름을 제공
  - 서로다른 성질을 가진 TCP(Transmission Control Protocol), UDP(User Data Protocol) 2가지 프로토콜이 있음
- 네트워크 계층(혹은 인터넷 계층)
  - 네트워크 상에서 패킷의 이동을 다룸. 패킷의 이동 경로 결정
- 링크 계층 (혹은 데이터링크 계층, 네트워크 인터페이스 계층)
  - 네트워크에 접속하는 하는 하드웨어적인 면을 다룸
  - 디바이스 드라이버, 네트워크 인터페이스 카드 (NIC), 케이블 및 여러 전송매체 포함
  - 하드웨어적 측면은 모두 링크 계층의 역할 임




```
The OSI (Open Systems Interconnection) 7-layer model and the TCP/IP (Transmission Control Protocol/Internet Protocol) model
are both networking models that describe how data is transmitted between devices over a network.

The OSI model consists of seven layers:

1. Physical Layer - Deals with the physical aspects of network communication, such as the electrical and mechanical connections between devices.
2. Data Link Layer - Handles the transmission of data between devices on the same network.
3. Network Layer - Provides routing and addressing services for data packets as they travel across different networks.
4. Transport Layer - Ensures reliable data transmission between end points and manages flow control and error recovery.
5. Session Layer - Establishes, manages, and terminates sessions between applications on different devices.
6. Presentation Layer - Handles the formatting and translation of data between different formats and data structures.
7. Application Layer - Provides network services directly to applications and end users.

On the other hand, the TCP/IP model consists of four layers:

1. Network Access Layer - Defines the protocols and hardware necessary for communication on a specific network.
2. Internet Layer - Handles the addressing and routing of data across multiple networks.
3. Transport Layer - Provides reliable data transmission between applications on different devices.
4. Application Layer - Provides network services directly to applications and end users, including protocols such as HTTP, FTP, and DNS.

In summary, both models provide a framework for understanding how network communication occurs,
with the OSI model being more comprehensive and the TCP/IP model being more widely used in practice.
It's important to note that while these models provide a useful conceptual understanding,
actual network implementations often involve more complex interactions between layers and protocols.
```

