package org.foundee.api.configuration;

import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;

public class WebMvcConfiguration extends WebSecurityConfigurerAdapter {

	@Override
    protected void configure(HttpSecurity http) throws Exception {

        http.
            authorizeRequests()
            .antMatchers("/").permitAll()
            .antMatchers("/admin/**").authenticated()
        ;
    }
}
