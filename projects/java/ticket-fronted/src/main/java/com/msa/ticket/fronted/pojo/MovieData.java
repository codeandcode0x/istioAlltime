package com.msa.ticket.fronted.pojo;

import com.msa.ticket.fronted.model.Movie;

import java.util.List;

public class MovieData {
    private Integer code;
//    private Object message;
    private List<Movie> data;


    public Integer getCode() {
        return code;
    }

    public void setCode(Integer code) {
        this.code = code;
    }

    public List<Movie> getData() {
        return data;
    }

    public void setData(List<Movie> data) {
        this.data = data;
    }

    @Override
    public String toString() {
        return "Message{" +
                "code='" + code + '\'' +
                ", data=" + data +
                '}';
    }
}



