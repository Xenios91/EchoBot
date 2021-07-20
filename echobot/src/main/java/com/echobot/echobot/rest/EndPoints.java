package com.echobot.echobot.rest;

import javax.servlet.http.HttpServletRequest;

import com.echobot.echobot.service.EchoRequestIndexer;
import com.echobot.model.EchoRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class EndPoints {

    private EchoRequestIndexer echoRequestIndexer;
    private HttpServletRequest request;

    @Autowired
    public void setRequest(HttpServletRequest request) {
        this.request = request;
    }

    @Autowired
    void setEchoRequestIndexer(EchoRequestIndexer echoRequestIndexer) {
        this.echoRequestIndexer = echoRequestIndexer;
    }

    @GetMapping("/echo")
    public String getEchoRequest(@RequestParam String token) {
        EchoRequest echoRequest = new EchoRequest();
        echoRequest.setRequestBody("test test test");
        return echoRequest.getRequestBody();
    }

    @PostMapping("/createEchoRequest")
    public void setEchoRequest() {
        String requestAddress = this.request.getRemoteAddr();
        EchoRequest echoRequest = new EchoRequest();
        echoRequest.setIp(requestAddress);
        this.echoRequestIndexer.addToQueue(echoRequest);
    }

}
