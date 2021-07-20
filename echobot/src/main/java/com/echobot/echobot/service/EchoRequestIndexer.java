package com.echobot.echobot.service;

import java.util.HashMap;
import java.util.Random;

import com.echobot.model.EchoRequest;

import org.springframework.stereotype.Service;

@Service
public class EchoRequestIndexer {
    private HashMap<String, EchoRequest> echoRequestIndex = new HashMap<>();

    private String generateToken() {
        int leftLimit = 48;
        int rightLimit = 122;
        int targetStringLength = 24;
        Random random = new Random();

        String generatedString = random.ints(leftLimit, rightLimit + 1)
                .filter(i -> (i <= 57 || i >= 65) && (i <= 90 || i >= 97)).limit(targetStringLength)
                .collect(StringBuilder::new, StringBuilder::appendCodePoint, StringBuilder::append).toString();

        if (this.echoRequestIndex.containsKey(generatedString)) {
            generatedString = generateToken();
        }
        return generatedString;
    }

    public void addToQueue(EchoRequest echoRequest) {
        this.echoRequestIndex.put(generateToken(), echoRequest);
    }

    public HashMap<String, EchoRequest> getEchoRequestIndex() {
        return this.echoRequestIndex;
    }

    public void setEchoRequestIndex(HashMap<String, EchoRequest> echoRequestIndex) {
        this.echoRequestIndex = echoRequestIndex;
    }

}
