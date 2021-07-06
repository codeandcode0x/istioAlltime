package com.msa.ticket.frontend;

import com.msa.ticket.frontend.pojo.MovieData;
import com.msa.ticket.frontend.pojo.MovieDatas;
import com.msa.ticket.frontend.pojo.ShowDatas;
import com.msa.ticket.frontend.pojo.InfoDatas;
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
public class HomeController {
    private static final Logger logger = Logger.getLogger(HomeController.class.getName());

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
        String movieUrl = ticketHost+"/api/movies?count="+count;
        RestTemplate restTemplate = new RestTemplate();
        MovieDatas movies = restTemplate.getForObject(movieUrl, MovieDatas.class);
        System.out.println(movies);

        String showMusicUrl = ticketHost+"/api/shows?mtype=演唱会&count="+count;
        ShowDatas showMusics = restTemplate.getForObject(showMusicUrl, ShowDatas.class);
        System.out.println(showMusics);

        String showDramaUrl = ticketHost+"/api/shows?mtype=歌舞剧&count="+count;
        ShowDatas showDramas = restTemplate.getForObject(showDramaUrl, ShowDatas.class);
        System.out.println(showDramas);

        String infoUrl = ticketHost+"/api/infos?count="+count;
        InfoDatas infos = restTemplate.getForObject(infoUrl, InfoDatas.class);
        System.out.println(infos);

        assert movies != null;
        assert showMusics != null;
        assert showDramas != null;
        assert infos != null;
        model.addAttribute("movies", movies.getData());
        model.addAttribute("showMusics", showMusics.getData());
        model.addAttribute("showDramas", showDramas.getData());
        model.addAttribute("infos", infos.getData());
        return "home";
    }


    @GetMapping("/movies")
    public String GetMovies(@RequestParam(name="name", required = false, defaultValue = "ethan") String name, Model model) {

        // Configuration config = new Configuration("Hello Jaeger");

        // Configuration.SenderConfiguration sender = new Configuration.SenderConfiguration();
        // sender.withAgentHost("localhost");
        // sender.withAgentPort(6831);
        // config.withReporter(new Configuration.ReporterConfiguration().withSender(sender).withFlushInterval(100).withLogSpans(false));

        // config.withSampler(new Configuration.SamplerConfiguration().withType("const").withParam(1));

        // io.opentracing.Tracer tracer = config.getTracer();
        // System.out.println(tracer.toString());
        // // GlobalTracer.register(tracer);
        // Tracer.SpanBuilder spanBuilder = GlobalTracer.get().buildSpan("hello");
        // Span parent = spanBuilder.start();
        
        // logger.info(env.getProperty("TICKET_MANAGER_HOST"));
        // if (env.getProperty("TICKET_MANAGER_HOST") != null) {
        //     ticketHost = env.getProperty("TICKET_MANAGER_HOST");
        // }

        String url = ticketHost+"/api/movies";
        RestTemplate restTemplate = new RestTemplate();
        HttpHeaders headers = new HttpHeaders();
        // headers.set("Uber-Trace-Id", parent.context().toString());
        headers.setContentType(MediaType.APPLICATION_JSON);
        Map<String, Object> map = new HashMap<String, Object>();
        HttpEntity<Map<String, Object>> request = new HttpEntity<Map<String, Object>>(map, headers);
        ResponseEntity<MovieDatas> response = restTemplate.exchange(url, HttpMethod.GET,request, MovieDatas.class);
        System.out.println(response.getBody());
        MovieDatas  movies = response.getBody();

        System.out.println("---------" + headers);
        // MovieDatas movies = restTemplate.getForObject(url ,MovieDatas.class);
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
