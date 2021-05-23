package com.msa.ticket.fronted;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.ui.Model;
import org.springframework.web.client.RestTemplate;

import java.util.Collections;
import java.util.Map;

@Slf4j
@Controller
public class TicketController {
    @Autowired
    private RestTemplate restTemplate ;

    @GetMapping("/home")
    public String Home() {
        return "home";
    }

    @GetMapping("/ticket")
    public String GetVersion(@RequestParam(name="name", required = false, defaultValue = "ethan") String name, Model model) {
        model.addAttribute("name", name);
        return "ticket";
    }

    @GetMapping("/detail")
    public String GetDetail(@RequestParam(name="id", required = false, defaultValue = "1") String name, Model model) {
        //请求地址
//        String url = "https://quoters.apps.pcfone.io/api/random";
        String url = "http://127.0.0.1:8081/api/v2/getMovieInfo";
        Movie result = restTemplate.getForObject(url, Movie.class,"42", "21");
/*        Map<String, String> vars = Collections.singletonMap("hotel", "42");
        String result = restTemplate.getForObject("http://example.com/hotels/{hotel}/rooms/{hotel}", String.class, vars);*/
//        log.info(result.toString());
        assert result != null;
        model.addAttribute("movie", result);
        return "detail";
    }
}
