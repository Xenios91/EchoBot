package com.echobot.echobot.config;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.ViewControllerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class TemplateConfig implements WebMvcConfigurer {

	@Override
	public void addViewControllers(final ViewControllerRegistry registry) {
		registry.addViewController("/").setViewName("home");
		registry.addViewController("/greeting").setViewName("greeting");
		registry.addViewController("/login").setViewName("login");
		registry.addRedirectViewController("/home", "/");

	}

}