package com.echobot.echobot.model;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class EchoRequest {

    private List<String> ips = new ArrayList<>();
    private String requestBody;
    private LocalDateTime terminateTime;
    private String delayType;

    public EchoRequest(final String ip, final String requestBody, String delayType) {
        this.ips.add(ip);
        this.requestBody = requestBody;
        this.terminateTime = LocalDateTime.now().plusHours(12);

        if (delayType == null || delayType.isEmpty()) {
            delayType = "normal";
        } else {
            this.delayType = delayType;
        }
    }

    public int getDelay() {
        switch (this.delayType) {
            case "normal":
                return 0;
            case "low":
                return 2;
            case "medium":
                return 5;
            case "high":
                return 10;
            case "insane":
                return 20;
            default:
                return 0;
        }
    }

    public String getDelayType() {
        return this.delayType;
    }

    public void setDelayType(final String delayType) {
        this.delayType = delayType;
    }

    public List<String> getIps() {
        return this.ips;
    }

    public void setIps(final List<String> ips) {
        this.ips = ips;
    }

    public String getRequestBody() {
        return this.requestBody;
    }

    public void setRequestBody(final String requestBody) {
        this.requestBody = requestBody;
    }

    public LocalDateTime getTerminateTime() {
        return this.terminateTime;
    }

    public void setTerminateTime(final LocalDateTime terminateTime) {
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
        return Objects.equals(ips, echoRequest.ips) && Objects.equals(requestBody, echoRequest.requestBody)
                && Objects.equals(terminateTime, echoRequest.terminateTime)
                && Objects.equals(delayType, echoRequest.delayType);
    }

    @Override
    public int hashCode() {
        return Objects.hash(ips, requestBody, terminateTime, delayType);
    }

    @Override
    public String toString() {
        return "{" + " ips='" + getIps() + "'" + ", requestBody='" + getRequestBody() + "'" + ", terminateTime='"
                + getTerminateTime() + "'" + ", delayType='" + getDelayType() + "'" + "}";
    }

}
