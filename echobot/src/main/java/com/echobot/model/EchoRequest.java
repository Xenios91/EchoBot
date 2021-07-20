package com.echobot.model;

import java.time.LocalDateTime;
import java.util.Objects;

public class EchoRequest {

    private String ip;
    private String requestBody;
    private String token;
    private LocalDateTime terminateTime;

    public String getIp() {
        return this.ip;
    }

    public void setIp(String ip) {
        this.ip = ip;
    }

    public String getRequestBody() {
        return this.requestBody;
    }

    public void setRequestBody(String requestBody) {
        this.requestBody = requestBody;
    }

    public String getToken() {
        return this.token;
    }

    public void setToken(String token) {
        this.token = token;
    }

    public LocalDateTime getTerminateTime() {
        return this.terminateTime;
    }

    public void setTerminateTime(LocalDateTime terminateTime) {
        this.terminateTime = terminateTime;
    }

    @Override
    public boolean equals(Object o) {
        if (o == this)
            return true;
        if (!(o instanceof EchoRequest)) {
            return false;
        }
        EchoRequest echoRequest = (EchoRequest) o;
        return Objects.equals(ip, echoRequest.ip) && Objects.equals(requestBody, echoRequest.requestBody)
                && Objects.equals(token, echoRequest.token);
    }

    @Override
    public int hashCode() {
        return Objects.hash(ip, requestBody, token);
    }

    @Override
    public String toString() {
        return "{" + " ip='" + getIp() + "'" + ", requestBody='" + getRequestBody() + "'" + ", token='" + getToken()
                + "'" + ", terminateTime='" + getTerminateTime() + "'" + "}";
    }

}
