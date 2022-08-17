package java

import "os"

func generateJavaException(path string) {
	APIException := path + "/APIException.java"
	fe, _ := os.Create(APIException)
	_, _ = fe.WriteString(generateJavaAPIException())

	APIExceptionHandler := path + "/APIExceptionHandler.java"
	fe, _ = os.Create(APIExceptionHandler)
	_, _ = fe.WriteString(generateJavaAPIExceptionHandler())

	ExceptionHandler := path + "/ExceptionHandler.java"
	fe, _ = os.Create(ExceptionHandler)
	_, _ = fe.WriteString(generateJavaExceptionHandler())
}

func generateJavaAPIException() string {
	return `package com.ne.exception;

import org.springframework.http.HttpStatus;

import java.time.ZonedDateTime;

public class APIException {
    private final String message;
    private final Throwable throwable;
    private final HttpStatus httpStatus;
    private final ZonedDateTime timeStamp;

    public APIException(String message, Throwable throwable, HttpStatus httpStatus, ZonedDateTime timeStamp) {
        this.message = message;
        this.throwable = throwable;
        this.httpStatus = httpStatus;
        this.timeStamp = timeStamp;
    }

    public String getMessage() {
        return message;
    }

    public Throwable getThrowable() {
        return throwable;
    }

    public HttpStatus getHttpStatus() {
        return httpStatus;
    }

    public ZonedDateTime getTimeStamp() {
        return timeStamp;
    }
}`
}

func generateJavaAPIExceptionHandler() string {
	return `package com.ne.exception;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;

import java.time.ZoneId;
import java.time.ZonedDateTime;

@ControllerAdvice
public class APIExceptionHandler {

    @org.springframework.web.bind.annotation.ExceptionHandler(value = {ExceptionHandler.class})
    public ResponseEntity<Object> handleAPIRequestException(ExceptionHandler e) {
        APIException exception = new APIException(
                e.getMessage(),
                e,
                HttpStatus.BAD_REQUEST,
                ZonedDateTime.now(ZoneId.of("Z"))
        );

        return new ResponseEntity<>(exception, HttpStatus.BAD_REQUEST);
    }
}`
}

func generateJavaExceptionHandler() string {
	return `package com.ne.exception;

public class ExceptionHandler extends RuntimeException{

    public ExceptionHandler(String message) {
        super(message);
    }

    public ExceptionHandler(String message, Throwable cause) {
        super(message, cause);
    }
}
`
}
