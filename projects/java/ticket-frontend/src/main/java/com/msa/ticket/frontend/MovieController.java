package com.msa.ticket.frontend;

import com.msa.ticket.frontend.pojo.MovieData;
import com.msa.ticket.frontend.pojo.MovieDatas;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.http.*;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.ui.Model;
import org.springframework.web.client.RestClientException;
import org.springframework.web.client.RestTemplate;

import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;
import java.util.Map;
import java.util.logging.Logger;

@Slf4j
@Controller
public class MovieController {
    private static final Logger logger = Logger.getLogger(MovieController.class.getName());

    @Autowired
    private Environment env;
    private RestTemplate restTemplate ;
    private String ticketHost = "http://localhost:8080";

    @GetMapping("/")
    public String Home(@RequestParam(name="count", required = false, defaultValue = "3") Integer count, Model model) {
        logger.info(env.getProperty("TICKET_MANAGER_HOST"));
        if (env.getProperty("TICKET_MANAGER_HOST") != null) {
            ticketHost = env.getProperty("TICKET_MANAGER_HOST");
        }
        String url = ticketHost+"/api/movies?count="+count;
        RestTemplate restTemplate = new RestTemplate();
        MovieDatas movies = restTemplate.getForObject(url, MovieDatas.class);
        System.out.println(movies);

        assert movies != null;
        model.addAttribute("movies", movies.getData());
        return "home";
    }


    @GetMapping("/movies")
    public String GetMovies(@RequestParam(name="name", required = false, defaultValue = "ethan") String name, Model model) {
        logger.info(env.getProperty("TICKET_MANAGER_HOST"));
        if (env.getProperty("TICKET_MANAGER_HOST") != null) {
            ticketHost = env.getProperty("TICKET_MANAGER_HOST");
        }
        String url = ticketHost+"/api/movies";
        RestTemplate restTemplate = new RestTemplate();
        MovieDatas movies = restTemplate.getForObject(url, MovieDatas.class);
        System.out.println(movies);

        assert movies != null;
        model.addAttribute("movies", movies.getData());
        return "movies";
    }

    @GetMapping("/detail/{id}")
    public String GetDetail(@PathVariable Long id, Model model) {
        if (env.getProperty("TICKET_MANAGER_HOST") != null) {
            ticketHost = env.getProperty("TICKET_MANAGER_HOST");
        }
        String url = ticketHost+"/api/movie/"+id;
        RestTemplate restTemplate = new RestTemplate();
        MovieData movie = restTemplate.getForObject(url, MovieData.class);
        System.out.println(movie);

        assert movie != null;
        model.addAttribute("movie", movie.getData());
        return "detail";
    }


    public <T> Map postForMapResult(String baseUrl, String uri, T body) {
        Map map = null;
        try {
            HttpHeaders httpHeaders = new HttpHeaders();
            httpHeaders.setContentType(MediaType.APPLICATION_JSON);
            HttpEntity<T> requestEntity = new HttpEntity<>(body, httpHeaders);
            ResponseEntity<Map> exchange = restTemplate.exchange(baseUrl + uri, HttpMethod.POST, requestEntity, Map.class);
            map = exchange.getBody();
        } catch (RestClientException e) {
            System.out.println("postForMapResult异常");
        }
        return map;
    }

    public void  getMovies(HttpServletResponse response, Integer userId) throws Exception{
        String url = ticketHost+"/api/movies";
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        Map<String, Object> map = new HashMap<String, Object>();
        map.put("userId", userId);
        HttpEntity<Map<String, Object>> request = new HttpEntity<Map<String, Object>>(map, headers);
        ResponseEntity<String> entity = restTemplate.getForEntity(url, String.class);
        String body = entity.getBody();
        System.out.println(body);
    }

}
