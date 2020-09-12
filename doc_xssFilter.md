
# XSS FILTER 가이드 (EDIT BY KYUNGJOO)

FORM 데이터 형식일때 / GET 방식일때 GETPARAMETER 등등의 메소드가 실행되는 로직이어서 지금 PASSIKEY 소스들과는 맞지 않음 그래서 네이버 LUCY 로 재도전

## XSS 란?

XSS(Cross Site Script)란 웹 어플리케이션에서 사용자 입력 값에 대한 필터링이 제대로 이루어지지 않을 경우, 공격자가 입력이 가능한 폼에 악의적인 스크립트를 삽입하여 해당 스크립트가 희생자 측에서 동작하도록 하여 악의적인 행위를 수행하는 취약점이다. 공격자는 취약점을 이용하여 사용자의 개인정보 및 쿠키정보 탈취, 악성코드 감염, 웹 페이지 변조 등의 공격을 수행한다.

### 확인 방법
```SH
게시판 입력창 아무대서나 
<script> alert("script ok"); </script>
 라고 입력하고 저장한 후
저장한 게시글 상세화면으로 들어갔을때 해당 코드(alert 창)가 실행되는지 확인하면 간단하다.
그래서 < , > 등을 입력 받지 못하도록 막아야 한다고 대형 사이트의 보안 권고사항이 나온다.
```



# Lucy를 이용한 Request Body 크로스사이트스크립팅 제어(JSON)

## 참조
https://github.com/naver/lucy-xss-filter

## pom.xml
lucy xss filter 의존성 주입
```sh
<!-- lucy xss filter -->
		<dependency>
			<groupId>com.navercorp.lucy</groupId>
			<artifactId>lucy-xss</artifactId>
			<version>1.6.3</version>
		</dependency>
	</dependencies>
```

## web.xml
현재 프로젝트에서 filter 는 encodingFilter / servletWrappingFilter / CroossDomainFilter 이렇게 3가지가 있다.
처음에는 제일 마지막에 xssfilter를 등록해주었는데 api 호출시 자꾸 에러가 나서 확인을 해보니
servletWrapping 을 먼저 태우고 나서 xssfilter 에 와서 inputstream 부분에서 토큰이 없다? 그런식의 에러가 나서 endcodingFilter 뒤에 xssFilter의 순서를 옮겨줌
```sh
	<!-- xss filter -->
	<filter>
		<filter-name>RequestBodyXSSFilter</filter-name>
		<filter-class>com.rowem.key.common.filter.RequestBodyXSSFilter</filter-class>
	</filter>
	<filter-mapping>
		<filter-name>RequestBodyXSSFilter</filter-name>
		<url-pattern>/*</url-pattern>
	</filter-mapping>
```

## RequestBodyXSSFilter
web.xml 에서 필터를 통해 RequestBodyXSSFilter 클래스에 온다 여기서 xss를 필터릴 해주는 XssFilterRequestWrapper를 호출
```sh
import java.io.IOException;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class RequestBodyXSSFilter implements Filter {
	@Override
	public void doFilter(ServletRequest req, ServletResponse res, FilterChain chain)
			throws IOException, ServletException {
		HttpServletRequest request = (HttpServletRequest)req;
 		HttpServletResponse response = (HttpServletResponse)res;
 		XssFilterRequestWrapper requestWrapper = null;
 		try{
 			requestWrapper = new XssFilterRequestWrapper(request);
 		}catch(Exception e){
 			e.printStackTrace();
 		}
 		chain.doFilter(requestWrapper, response);
	}
	
	@Override
	public void init(FilterConfig filterConfig) throws ServletException {}
	
	@Override
	public void destroy() {}

}
```

## XssFilterRequestWrapper

```sh
import java.io.BufferedReader;
import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;

import javax.servlet.ReadListener;
import javax.servlet.ServletInputStream;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletRequestWrapper;

import com.nhncorp.lucy.security.xss.XssFilter;

public class XssFilterRequestWrapper extends HttpServletRequestWrapper {
	private byte[] b;
	public XssFilterRequestWrapper(HttpServletRequest request) throws IOException {
		super(request);
 		XssFilter filter = XssFilter.getInstance("lucy-xss-superset-sax.xml" , true);			 //kyungjoo <!-- Not Allowed Tag Filtered --> 주석을 제거하기위해 getInstance 에 true 추가
 		b = new String(filter.doFilter(getBody(request))).getBytes("UTF-8");					 //kyungjoo .getBytes() -> 한글로 입력시 엔코딩 오류로 인해 UTF-8 추가
	}

	public ServletInputStream getInputStream() throws IOException {
 		final ByteArrayInputStream bis = new ByteArrayInputStream(b);

 		return new ServletInputStreamImpl(bis);
 	}

 	class ServletInputStreamImpl extends ServletInputStream{
 		private InputStream is;

 		public ServletInputStreamImpl(InputStream bis){
 			is = bis;
 		}


 		public int read() throws IOException {
 			return is.read();
 		}


 		public int read(byte[] b) throws IOException {
 			return is.read(b);
 		}

		public boolean isFinished() {
			// TODO Auto-generated method stub
			return false;
		}

		public boolean isReady() {
			// TODO Auto-generated method stub
			return false;
		}

		public void setReadListener(ReadListener readListener) {
			// TODO Auto-generated method stub
		}
 	}

 	public static String getBody(HttpServletRequest request) throws IOException {
 	    String body = null;
 	    StringBuilder stringBuilder = new StringBuilder();
 	    BufferedReader bufferedReader = null;

 	    try {
 	        InputStream inputStream = request.getInputStream();
 	        if (inputStream != null) {
 	            bufferedReader = new BufferedReader(new InputStreamReader(inputStream , "UTF-8"));		//kyungjoo .getBytes(UTF-8) 로 변경시 (寃쎌＜�뀒�뒪�듃) 이런식으로 나오는 부분 한글화를 위해 UTF-8 추가 
 	            char[] charBuffer = new char[128];
 	            int bytesRead = -1;
 	            while ((bytesRead = bufferedReader.read(charBuffer)) > 0) {
 	                stringBuilder.append(charBuffer, 0, bytesRead);
 	            }
 	        } else {
 	            stringBuilder.append("");
 	        }
 	    } catch (IOException ex) {
 	        throw ex;
 	    } finally {
 	        if (bufferedReader != null) {
 	            try {
 	                bufferedReader.close();
 	            } catch (IOException ex) {
 	                throw ex;
 	            }
 	        }
 	    }

 	    body = stringBuilder.toString();
 	    return body;
 	}
 	
```
### 수정한 부분
```sh
public XssFilterRequestWrapper(HttpServletRequest request) throws IOException {
		super(request);
 		XssFilter filter = XssFilter.getInstance("lucy-xss-sax.xml");			 
 		b = new String(filter.doFilter(getBody(request))).getBytes();					 
	}

처음에 소스를 가져왔을때 이런식이었다.
lucy git 에서 "lucy-xss-superset-sax.xml" 파일을 가져와 src/resource 파일 밑에 넣고 
.getInstance("lucy-xss-superset-sax.xml") 로 변경
.getBytes() 였을때 한글이 안들어가는 현상때문에 -> .getBytes("UTF-8") 로 변경
이러고 API 를 실행했을때 <!-- Not Allowed Tag Filtered --> 라는 주석이 계속 따라오는데 이를 제거하기 위해
.getInstance("lucy-xss-superset-sax.xml" , true) 를 추가 default = false 

.getBytes("UTF-8") 를 해주면 한글은 들어가지만 막상 파싱되었을때 한글이 깨져서 나온다 이를 방지하기위해
 bufferedReader = new BufferedReader(new InputStreamReader(inputStream)) 를
 bufferedReader = new BufferedReader(new InputStreamReader(inputStream , "UTF-8")) 로 변경 

```




# lucy-xss-servlet-filter 를 이용한 방식
- form-data에 대해서만 적용되고 Request Raw Body로 넘어가는 JSON에 대해서는 처리해주지 않는다는 단점이 있다. 그래서 JSON을 주고 받는 API 서버의 경우에는 직접 처리를 해줘야 한다.
 즉, Lucy는 RequestParameter관련한 지원만 해준다.
- 그래서 프로젝트에 반영 X


### 참조
- https://lee-mandu.tistory.com/450 

# JAVA FILTER 를 이용한 XSS 방식
- 이방식은  form-data에 대해서만 적용되고 Request Raw Body로 넘어가는 JSON에 대해서는 처리해주지 않는다.
- 그래서 프로젝트에 반영 X


### WEB.XML 에 FILTER 설정 추가
``` SH
<filter>
    <filter-name>XSS</filter-name>
    <filter-class>패키지명.CrossScriptingFilter</filter-class>
</filter>
<filter-mapping>
    <filter-name>XSS</filter-name>
    <url-pattern>/*</url-pattern>
</filter-mapping>
```

### CrossScriptingFilter 필터 파일
WEB.XML 에서 설정값이 바라보고 있는 필터 파일

```SH
 
import java.io.IOException;
import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;
 
 
public class CrossScriptingFilter implements Filter {

public FilterConfig filterConfig;

    public void init(FilterConfig filterConfig) throws ServletException {
        this.filterConfig = filterConfig;
    }
 
    public void destroy() {
        this.filterConfig = null;
    }
 
    public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
        throws IOException, ServletException {
 
        chain.doFilter(new RequestWrapper((HttpServletRequest) request), response);
 
    }
}
```

### RequestWrapper  필터링을 실행할 파일
CrossScriptingFilter 파일의 doFilter메소드 실행 시 바라보는 파일

```SH

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletRequestWrapper;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public final class RequestWrapper extends HttpServletRequestWrapper {

	public RequestWrapper(HttpServletRequest servletRequest) {
		super(servletRequest);
	}

	public String[] getParameterValues(String parameter) {

		String[] values = super.getParameterValues(parameter);
		if (values == null) {
			return null;
		}
		int count = values.length;
		String[] encodedValues = new String[count];
		for (int i = 0; i < count; i++) {
			encodedValues[i] = cleanXSS(values[i]);
		}
		return encodedValues;
	}

	public String getParameter(String parameter) {
		String value = super.getParameter(parameter);
		if (value == null) {
			return null;
		}
		return cleanXSS(value);
	}

	public String getHeader(String name) {
		String value = super.getHeader(name);
		if (value == null)
			return null;
		return cleanXSS(value);

	}

    private String cleanXSS(String value) {
    	StringBuffer sb = null;
    	String[] checkStr_arr = {
    			"<script>","</script>",
    			"<javascript>","</javascript>",
    			"<iframe>","</iframe>",
    			"<vbscript>","</vbscript>",
    			"<object","</object>",
    			"<img","</img>",
    			"<marquee","</marquee>",
    			"onerror","onclick","onload","onmouseover","onstart"
    	};
    	for(String checkStr:checkStr_arr) {
    		while(value.indexOf(checkStr)!=-1) {
    			value = value.replaceAll(checkStr, "");
    		}
    		while(value.toLowerCase().indexOf(checkStr)!=-1) {
    			sb = new StringBuffer(value);
    			sb = sb.replace(value.toLowerCase().indexOf(checkStr), value.toLowerCase().indexOf(checkStr)+checkStr.length(), "");
    			value = sb.toString();
    		}
    	}
                //You'll need to remove the spaces from the html entities below
        value = value.replaceAll("<", "& lt;").replaceAll(">", "& gt;");
        value = value.replaceAll("\\(", "& #40;").replaceAll("\\)", "& #41;");
        value = value.replaceAll("'", "& #39;");
        value = value.replaceAll("eval\\((.*)\\)", "");
        value = value.replaceAll("[\\\"\\\'][\\s]*javascript:(.*)[\\\"\\\']", "\"\"");
        value = value.replaceAll("script", "");
        return value;
    }
}
```
