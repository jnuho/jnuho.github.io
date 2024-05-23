# 애플로그인 구현

## 애플 개발자센터

[애플 개발자센터](https://developer.apple.com/account) 로그인 후


1. ```Certificates, Identifiers & Profiles``` 클릭<br>
2. ```Keys``` 탭 클릭 하고 + 버튼 눌러 키 생성<br>
3. Sign in with Apple 선택하여 생성. 키 다운로드하면 재 다운로드안되서 주의

위의 1,2,3과정은 신매니저님이 이미 진행 하셔서 .p8파일 메일로 보내주심-로웸메일로 p8파일 방금 공유드렸습니다.

밑에 해당 코드작성 참고링크

https://stackoverflow.com/a/58252090

https://developer.apple.com/documentation/sign_in_with_apple/sign_in_with_apple_rest_api

https://spiralmoon.tistory.com/entry/Apple-%EC%95%A0%ED%94%8C-%EB%A1%9C%EA%B7%B8%EC%9D%B8-%EC%84%A4%EC%A0%95%ED%95%98%EA%B8%B0-Sign-In-with-Apple

``` java
package com.rowem.passicon.apis.join.controller;

import java.io.FileReader;
import java.security.PrivateKey;
import java.util.Date;

import org.bouncycastle.asn1.pkcs.PrivateKeyInfo;
import org.bouncycastle.openssl.PEMParser;
import org.bouncycastle.openssl.jcajce.JcaPEMKeyConverter;
import org.junit.Test;
import org.springframework.core.io.ClassPathResource;

import com.google.gson.Gson;
import com.mashape.unirest.http.HttpResponse;
import com.mashape.unirest.http.Unirest;
import com.rowem.passicon.apis.join.vo.IdTokenPayload;
import com.rowem.passicon.apis.join.vo.TokenResponse;

import io.jsonwebtoken.JwsHeader;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.io.Decoders;

public class AppleLoginTest {

	
	private static String APPLE_AUTH_URL = "https://appleid.apple.com/auth/token"; 
	private static String KEY_ID = "7MV3R7D2JW";
	private static String TEAM_ID = "Y67B6HMB9J";
	private static String CLIENT_ID = "com.rowem.PASSIKEY";

	private static PrivateKey pKey;

	@Test
	public void appleLoginTest() {
		System.out.println("asasa");
		try {
			appleAuth("0000");
			
		} catch(Exception e) {
			e.printStackTrace();
		}
	}

	private static PrivateKey getPrivateKey() throws Exception {
		// read your key
		String path = new ClassPathResource("apple/AuthKey_7MV3R7D2JW.p8").getFile().getAbsolutePath();

		final PEMParser pemParser = new PEMParser(new FileReader(path));
		final JcaPEMKeyConverter converter = new JcaPEMKeyConverter();
		final PrivateKeyInfo object = (PrivateKeyInfo) pemParser.readObject();
		final PrivateKey pKey = converter.getPrivateKey(object);

		pemParser.close();

		return pKey;
	}

	private static String generateJWT() throws Exception {
		if (pKey == null) {
			pKey = getPrivateKey();
		}

		String token = Jwts.builder()
				.setHeaderParam(JwsHeader.KEY_ID, KEY_ID)
				.setIssuer(TEAM_ID)
				.setAudience("https://appleid.apple.com")
				.setSubject(CLIENT_ID)
				.setExpiration(new Date(System.currentTimeMillis() + (1000 * 60 * 5)))
				.setIssuedAt(new Date(System.currentTimeMillis()))
				.signWith(pKey, SignatureAlgorithm.ES256)
				.compact();

		return token;
	}

	/*
	* Returns unique user id from apple
	* */
	public static String appleAuth(String authorizationCode) throws Exception {

		String token = generateJWT();

		HttpResponse<String> response = Unirest.post(APPLE_AUTH_URL)
				.header("Content-Type", "application/x-www-form-urlencoded")
				.field("client_id", CLIENT_ID)
				.field("client_secret", token)
				.field("grant_type", "authorization_code")
				.field("code", authorizationCode)
				.asString();

		TokenResponse tokenResponse = new Gson().fromJson(response.getBody(),TokenResponse.class);


		String idToken = tokenResponse.getId_token();
		String payload = idToken.split("\\.")[1];//0 is header we ignore it for now
		String decoded = new String(Decoders.BASE64.decode(payload));

		IdTokenPayload idTokenPayload = new Gson().fromJson(decoded,IdTokenPayload.class);

		return idTokenPayload.getSub();
	}
}

```