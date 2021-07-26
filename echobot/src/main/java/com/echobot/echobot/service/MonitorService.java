package com.echobot.echobot.service;

import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.Timer;
import java.util.TimerTask;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.echobot.echobot.model.EchoRequest;

@Service
public class MonitorService {
	@Autowired
	private EchoRequestIndexer echoRequestIndexer;

	private ExecutorService executorService = Executors.newSingleThreadExecutor();

	public MonitorService() {
		Runnable runnable = () -> new Timer().scheduleAtFixedRate(new TimerTask() {
			@Override
			public void run() {
				HashMap<String, EchoRequest> echoRequestMap = (HashMap<String, EchoRequest>) MonitorService.this.echoRequestIndexer
						.getEchoRequestIndex();
				echoRequestMap.forEach((key, value) -> {
					if (value.getTerminateTime().isBefore(LocalDateTime.now())) {
						echoRequestMap.remove(key);
					}
				});
			}
		}, TimeUnit.MINUTES.toMillis(1), TimeUnit.MINUTES.toMillis(15));

		this.executorService.submit(runnable);
	}

	public ExecutorService getExecutorService() {
		return this.executorService;
	}

	public void setExecutorService(final ExecutorService executorService) {
		this.executorService = executorService;
	}

}
