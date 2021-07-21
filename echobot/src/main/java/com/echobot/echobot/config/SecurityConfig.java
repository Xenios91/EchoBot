package com.echobot.echobot.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.provisioning.InMemoryUserDetailsManager;

@Configuration
@EnableWebSecurity
public class SecurityConfig extends WebSecurityConfigurerAdapter {
    @Value("${system.deployment}")
    private String systemDeployment;

    @Override
    protected void configure(final HttpSecurity http) throws Exception {
        if (systemDeployment.equals("prod")) {
            http.authorizeRequests().antMatchers("/").permitAll().anyRequest().authenticated().and().formLogin()
                    .loginPage("/login").permitAll().and().logout().permitAll();
        } else {
            http.cors().and().csrf().disable();
        }

    }

    @Bean
    @Override
    public UserDetailsService userDetailsService() {
        final UserDetails user = User.withDefaultPasswordEncoder().username("user").password("password").roles("USER")
                .build();

        return new InMemoryUserDetailsManager(user);
    }
}