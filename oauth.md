


OAuth 2.0
- 네이버, 카카오 로그인
- 해당 서비스의 회원임을 검증하기 위해 access_token 이용
- 토큰을 이용하여 고객이 안전하게 다른 서비스의 정보를 우리 서비스에 전달


Access Token
- 임의의 문자열 값: 토큰 발급 해 준 서비스만 알 수 있음
- JWT의 경우 Base64 인코딩 되어 있어 정보 볼 수 있음
- 이 토큰을 이용하여 이 토큰값과 관련된 고객의 정보를 해당 서비스에 요청 가능
- 서비스에서 access token 넘겨받는 방법
  - 네이버 Oauth 서버에서 redirect_uri에 access_token을 담아 (query string) 리다이렉트
  - redirect_uri 개발자센터에 등록
  - 헤더 Authorization

- 네이버에서 Access Token을 통해 정보조회 하는 API 제공

```
# 네이버 회원 프로필 조회 API
# GET방식 https://openapi.naver.com/v1/nid/me
# 요청헤더 Authorization: {토큰 타입} {접근 토큰}
```

