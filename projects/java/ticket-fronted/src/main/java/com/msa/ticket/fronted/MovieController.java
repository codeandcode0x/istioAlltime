package com.msa.ticket.fronted;

import com.msa.ticket.fronted.model.Movie;
import com.msa.ticket.fronted.pojo.MovieData;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.*;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.ui.Model;
import org.springframework.web.client.RestClientException;
import org.springframework.web.client.RestTemplate;

import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;
import java.util.Map;

@Slf4j
@Controller
public class MovieController {
    @Autowired
    private RestTemplate restTemplate ;

    @GetMapping("/")
    public String Home(@RequestParam(name="count", required = false, defaultValue = "3") Integer count, Model model) {
        var url = "http://127.0.0.1:8081/api/movies?count="+count;

        RestTemplate restTemplate = new RestTemplate();
        MovieData movies = restTemplate.getForObject(url, MovieData.class);
        System.out.println(movies);

        assert movies != null;
        model.addAttribute("movies", movies.getData());
        return "home";
    }


    @GetMapping("/movies")
    public String GetMovies(@RequestParam(name="name", required = false, defaultValue = "ethan") String name, Model model) {
        String url = "http://127.0.0.1:8081/api/movies";

        RestTemplate restTemplate = new RestTemplate();
        MovieData movies = restTemplate.getForObject(url, MovieData.class);
        System.out.println(movies);

        assert movies != null;
        model.addAttribute("movies", movies.getData());
        return "movies";
    }

    @GetMapping("/detail")
    public String GetDetail(@RequestParam(name="id", required = false, defaultValue = "1") Integer id, Model model) {
        var url = "http://127.0.0.1:8081/api/movie/="+id;

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
        String url = "http://127.0.0.1:8081/api/movies";
        HttpHeaders headers = new HttpHeaders();
        //定义请求参数类型，这里用json所以是MediaType.APPLICATION_JSON
        headers.setContentType(MediaType.APPLICATION_JSON);
        //RestTemplate带参传的时候要用HttpEntity<?>对象传递
        Map<String, Object> map = new HashMap<String, Object>();
        map.put("userId", userId);
        HttpEntity<Map<String, Object>> request = new HttpEntity<Map<String, Object>>(map, headers);

        ResponseEntity<String> entity = restTemplate.getForEntity(url, String.class);
        //获取3方接口返回的数据通过entity.getBody();它返回的是一个字符串；
        String body = entity.getBody();
        //然后把str转换成JSON再通过getJSONObject()方法获取到里面的result对象，因为我想要的数据都在result里面

        System.out.println(body);
    }

}
