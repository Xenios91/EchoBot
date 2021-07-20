package com.echobot.echobot.service;

import java.net.http.HttpClient;
import java.net.http.HttpRequest;

import org.springframework.stereotype.Service;
import java.net.URI;
import java.net.http.HttpResponse.BodyHandlers;

import com.echobot.model.EchoRequest;

import java.net.http.HttpResponse;

@Service
public class Echo {

    public boolean sendEcho(EchoRequest echoRequest) {
        HttpClient client = HttpClient.newHttpClient();

        HttpRequest request = HttpRequest.newBuilder().uri(URI.create(echoRequest.getIp())).build();
        client.sendAsync(request, BodyHandlers.ofString()).thenApply(HttpResponse::body).thenAccept(System.out::println)
                .join();
        return true;
    }
}
