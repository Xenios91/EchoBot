package com.echobot.echobot.rest;

import java.util.concurrent.TimeUnit;

import javax.servlet.http.HttpServletRequest;

import com.echobot.echobot.model.EchoRequest;
import com.echobot.echobot.service.EchoRequestIndexer;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class EndPoints {

    @Autowired
    private EchoRequestIndexer echoRequestIndexer;
    @Autowired
    private HttpServletRequest request;

    @GetMapping("/echo")
    public String getEchoRequest(@RequestParam final String token) {
        final EchoRequest echoRequest = this.echoRequestIndexer.getEchoRequestIndex().get(token);
        if (echoRequest == null) {
            return "Invalid token, please try again";
        }

        int delay = echoRequest.getDelay();
        if (delay > 0) {
            try {
                TimeUnit.SECONDS.sleep(delay);
            } catch (InterruptedException e) {
                e.printStackTrace();
                Thread.currentThread().interrupt();
            }
        }
        return echoRequest.getRequestBody();
    }

    @PostMapping("/createEchoRequest")
    public ResponseEntity<String> setEchoRequest(@RequestBody final String echoBody,
            @RequestParam final String delayType) {
        final String requestAddress = this.request.getRemoteAddr();
        final EchoRequest echoRequest = new EchoRequest(requestAddress, echoBody, delayType);
        final String token = this.echoRequestIndexer.addToQueue(echoRequest);
        return ResponseEntity.ok(token);
    }

}
