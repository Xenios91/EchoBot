package com.echobot.echobot.service;

import java.util.HashMap;
import java.util.Map;
import java.util.Random;

import org.springframework.stereotype.Service;

import com.echobot.echobot.model.EchoRequest;

@Service
public class EchoRequestIndexer {
	private HashMap<String, EchoRequest> echoRequestIndex = new HashMap<>();
	private final Random random = new Random();

	private String generateToken() {
		final int leftLimit = 48;
		final int rightLimit = 122;
		final int targetStringLength = 24;

		String generatedString = this.random.ints(leftLimit, rightLimit + 1)
				.filter(i -> (i <= 57 || i >= 65) && (i <= 90 || i >= 97)).limit(targetStringLength)
				.collect(StringBuilder::new, StringBuilder::appendCodePoint, StringBuilder::append).toString();

		if (this.echoRequestIndex.containsKey(generatedString)) {
			generatedString = generateToken();
		}
		return generatedString;
	}

	public String addToQueue(final EchoRequest echoRequest) {
		final String token = generateToken();
		this.echoRequestIndex.put(token, echoRequest);
		return token;
	}

	public Map<String, EchoRequest> getEchoRequestIndex() {
		return this.echoRequestIndex;
	}

	public void setEchoRequestIndex(final Map<String, EchoRequest> echoRequestIndex) {
		this.echoRequestIndex = (HashMap<String, EchoRequest>) echoRequestIndex;
	}

}
